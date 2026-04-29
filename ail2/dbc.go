package dbc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	aireport "AIL2-ProveNode/ail2/ai-report"
	machineinfos "AIL2-ProveNode/ail2/machine-infos"
	mt "AIL2-ProveNode/types"
)

var DbcChain *dbcChain = nil

type chainContract struct {
	abi             abi.ABI
	contractAddress common.Address
	chainId         *big.Int
}

type dbcChain struct {
	rpc          string
	privateKey   *ecdsa.PrivateKey
	report       *chainContract
	machineInfos *chainContract
	// txMutex serializes the "fetch nonce → sign → send → wait mined"
	// sequence across all transaction methods on this chain instance, so
	// two concurrent callers can't both read the same PendingNonceAt and
	// produce two transactions with the same nonce (one of which would be
	// rejected by the node). Held only for the duration of a single tx.
	txMutex sync.Mutex
}

func InitDbcChain(ctx context.Context, config mt.ChainConfig) error {
	reportContract, err := initChainContract(ctx, config.ReportContract, config.Rpc, config.ChainId)
	if err != nil {
		return err
	}
	machineInfoContract, err := initChainContract(ctx, config.MachineInfoContract, config.Rpc, config.ChainId)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}
	DbcChain = &dbcChain{
		rpc:          config.Rpc,
		privateKey:   privateKey,
		report:       reportContract,
		machineInfos: machineInfoContract,
	}
	return nil
}

func initChainContract(ctx context.Context, config mt.ContractConfig, rpc string, chainId int64) (*chainContract, error) {
	file, err := os.Open(config.AbiFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read ABI file: %v", err)
	}
	defer file.Close()
	abi, err := abi.JSON(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %v", err)
	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}
	cid, err := client.NetworkID(ctx)
	if err != nil {
		if chainId != 0 {
			cid = big.NewInt(chainId)
		} else {
			return nil, fmt.Errorf("failed to get chain id: %v", err)
		}
	}
	if !(common.IsHexAddress(config.ContractAddress)) {
		return nil, fmt.Errorf("invalid contract address: %v", config.ContractAddress)
	}
	addr := common.HexToAddress(config.ContractAddress)
	return &chainContract{
		abi:             abi,
		contractAddress: addr,
		chainId:         cid,
	}, nil
}

// sendTx submits a contract transaction and blocks until the on-chain
// receipt arrives. Three guarantees that the previous per-method copies of
// this code did not provide:
//
//  1. Connection cleanup: every successful Dial is matched by Close, so
//     long-running callers don't exhaust the local file-descriptor budget.
//  2. Nonce serialization: chain.txMutex prevents concurrent senders from
//     both reading the same PendingNonceAt value, which used to make one of
//     the transactions land in the mempool with a duplicate nonce and get
//     silently dropped.
//  3. Receipt confirmation: callers used to receive a tx hash and a nil
//     error even when the transaction reverted on chain. We now wait for
//     the receipt and return a non-nil error if status != 1, so the caller
//     can distinguish "submitted" from "confirmed".
//
// ctx controls how long we wait for the receipt; pass a bounded context if
// you don't want a stalled mempool to block the caller indefinitely.
func (chain *dbcChain) sendTx(
	ctx context.Context,
	contract *chainContract,
	data []byte,
) (string, error) {
	chain.txMutex.Lock()
	defer chain.txMutex.Unlock()

	client, err := ethclient.Dial(chain.rpc)
	if err != nil {
		return "", fmt.Errorf("dial ethereum client: %w", err)
	}
	defer client.Close()

	auth, err := bind.NewKeyedTransactorWithChainID(chain.privateKey, contract.chainId)
	if err != nil {
		return "", fmt.Errorf("create transactor: %w", err)
	}

	publicAddress := crypto.PubkeyToAddress(chain.privateKey.PublicKey)
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &contract.contractAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return "", fmt.Errorf("estimate gas: %w", err)
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("suggest gas price: %w", err)
	}
	nonce, err := client.PendingNonceAt(ctx, publicAddress)
	if err != nil {
		return "", fmt.Errorf("get nonce: %w", err)
	}

	tx := types.NewTransaction(nonce, contract.contractAddress, big.NewInt(0), gasLimit, gasPrice, data)
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", fmt.Errorf("sign transaction: %w", err)
	}
	txHash := signedTx.Hash().Hex()

	if err := client.SendTransaction(ctx, signedTx); err != nil {
		return txHash, fmt.Errorf("send transaction: %w", err)
	}

	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		return txHash, fmt.Errorf("wait for receipt: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return txHash, fmt.Errorf("transaction reverted on chain (status=%d)", receipt.Status)
	}
	return txHash, nil
}

func (chain *dbcChain) Report(
	ctx context.Context,
	notifyType mt.NotifyType,
	stakingType mt.StakingType,
	projectName, machineId string,
) (string, error) {
	data, err := chain.report.abi.Pack("report", notifyType, projectName, stakingType, machineId)
	if err != nil {
		return "", fmt.Errorf("pack report data: %w", err)
	}
	return chain.sendTx(ctx, chain.report, data)
}

func (chain *dbcChain) GetMachineState(
	ctx context.Context,
	projectName, machineId string,
	stakingType mt.StakingType,
) (bool, bool, error) {
	client, err := ethclient.Dial(chain.rpc)
	if err != nil {
		return false, false, fmt.Errorf("dial ethereum client: %w", err)
	}
	defer client.Close()

	instance, err := aireport.NewAireport(chain.report.contractAddress, client)
	if err != nil {
		return false, false, fmt.Errorf("new aireport instance: %w", err)
	}

	ms, err := instance.GetMachineState(nil, machineId, projectName, uint8(stakingType))
	return ms.IsOnline, ms.IsRegistered, err
}

func (chain *dbcChain) SetDeepLinkMachineInfoST(
	ctx context.Context,
	mk mt.MachineKey,
	mi mt.DeepLinkMachineInfoST,
	calcPoint int64,
	longitude, latitude float32,
	region string,
) (string, error) {
	info := machineinfos.MachineInfosMachineInfo{
		MachineOwner: common.HexToAddress(mi.Wallet),
		CalcPoint:    big.NewInt(calcPoint),
		CpuRate:      big.NewInt(int64(mi.CpuRate)),
		GpuType:      mi.GPUNames[0],
		GpuMem:       big.NewInt(int64(mi.GPUMemoryTotal[0])),
		CpuType:      mi.CpuType,
		GpuCount:     big.NewInt(int64(len(mi.GPUNames))),
		MachineId:    mk.ContainerId,
		Longitude:    fmt.Sprintf("%f", longitude),
		Latitude:     fmt.Sprintf("%f", latitude),
		MachineMem:   big.NewInt(mi.MemoryTotal),
		Region:       region,
		Model:        "",
	}
	data, err := chain.machineInfos.abi.Pack("setMachineInfo", mk.MachineId, info)
	if err != nil {
		return "", fmt.Errorf("pack setMachineInfo data: %w", err)
	}
	return chain.sendTx(ctx, chain.machineInfos, data)
}

func (chain *dbcChain) SetDeepLinkMachineInfoBandwidth(
	ctx context.Context,
	mk mt.MachineKey,
	mi mt.DeepLinkMachineInfoBandwidth,
	region string,
) (string, error) {
	info := machineinfos.MachineInfosBandWidthMintInfo{
		MachineOwner: common.HexToAddress(mi.Wallet),
		MachineId:    mk.ContainerId,
		CpuCores:     big.NewInt(int64(mi.CpuCores)),
		MachineMem:   big.NewInt(mi.MemoryTotal),
		Region:       region,
		Hdd:          big.NewInt(mi.Hdd),
		Bandwidth:    big.NewInt(int64(mi.Bandwidth)),
	}
	data, err := chain.machineInfos.abi.Pack("setBandWidthInfos", mk.MachineId, info)
	if err != nil {
		return "", fmt.Errorf("pack setBandWidthInfos data: %w", err)
	}
	return chain.sendTx(ctx, chain.machineInfos, data)
}
