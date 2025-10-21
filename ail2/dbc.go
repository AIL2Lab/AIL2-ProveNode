package dbc

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

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

func (chain *dbcChain) Report(
	ctx context.Context,
	notifyType mt.NotifyType,
	stakingType mt.StakingType,
	projectName, machineId string,
) (string, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(chain.rpc)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(chain.privateKey, chain.report.chainId)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	// Define the inputs for the report function
	// notifyType := uint8(1) // Example: NotifyType.MachineRegister
	// projectName := "ExampleProject"
	// stakingType := uint8(0) // Example: StakingType.ShortTerm
	// machineId := "example-machine-id"

	// Encode the function call
	data, err := chain.report.abi.Pack("report", notifyType, projectName, stakingType, machineId)
	if err != nil {
		return "", fmt.Errorf("failed to pack data: %v", err)
	}

	publicAddress := crypto.PubkeyToAddress(chain.privateKey.PublicKey)

	// Estimate gas limit
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &chain.report.contractAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas limit: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Fetch the nonce for the public address
	nonce, err := client.PendingNonceAt(ctx, publicAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, chain.report.contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return signedTx.Hash().Hex(), fmt.Errorf("failed to send transaction: %v", err)
	}
	return signedTx.Hash().Hex(), nil
}

func (chain *dbcChain) GetMachineState(
	ctx context.Context,
	projectName, machineId string,
	stakingType mt.StakingType,
) (bool, bool, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(chain.rpc)
	if err != nil {
		return false, false, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	instance, err := aireport.NewAireport(chain.report.contractAddress, client)
	if err != nil {
		return false, false, fmt.Errorf("failed to new aireport instance: %v", err)
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
	// Connect to Ethereum node
	client, err := ethclient.Dial(chain.rpc)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(chain.privateKey, chain.machineInfos.chainId)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	// Define the inputs for the report function
	// notifyType := uint8(1) // Example: NotifyType.MachineRegister
	// projectName := "ExampleProject"
	// stakingType := uint8(0) // Example: StakingType.ShortTerm
	// machineId := "example-machine-id"
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

	// Encode the function call
	data, err := chain.machineInfos.abi.Pack("setMachineInfo", mk.MachineId, info)
	if err != nil {
		return "", fmt.Errorf("failed to pack data: %v", err)
	}

	publicAddress := crypto.PubkeyToAddress(chain.privateKey.PublicKey)

	// Estimate gas limit
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &chain.machineInfos.contractAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas limit: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Fetch the nonce for the public address
	nonce, err := client.PendingNonceAt(ctx, publicAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, chain.machineInfos.contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return signedTx.Hash().Hex(), fmt.Errorf("failed to send transaction: %v", err)
	}
	return signedTx.Hash().Hex(), nil
}

func (chain *dbcChain) SetDeepLinkMachineInfoBandwidth(
	ctx context.Context,
	mk mt.MachineKey,
	mi mt.DeepLinkMachineInfoBandwidth,
	region string,
) (string, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(chain.rpc)
	if err != nil {
		return "", fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(chain.privateKey, chain.machineInfos.chainId)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %v", err)
	}

	// Define the inputs for the report function
	// notifyType := uint8(1) // Example: NotifyType.MachineRegister
	// projectName := "ExampleProject"
	// stakingType := uint8(0) // Example: StakingType.ShortTerm
	// machineId := "example-machine-id"
	info := machineinfos.MachineInfosBandWidthMintInfo{
		MachineOwner: common.HexToAddress(mi.Wallet),
		MachineId:    mk.ContainerId,
		CpuCores:     big.NewInt(int64(mi.CpuCores)),
		MachineMem:   big.NewInt(mi.MemoryTotal),
		Region:       region,
		Hdd:          big.NewInt(mi.Hdd),
		Bandwidth:    big.NewInt(int64(mi.Bandwidth)),
	}

	// Encode the function call
	data, err := chain.machineInfos.abi.Pack("setBandWidthInfos", mk.MachineId, info)
	if err != nil {
		return "", fmt.Errorf("failed to pack data: %v", err)
	}

	publicAddress := crypto.PubkeyToAddress(chain.privateKey.PublicKey)

	// Estimate gas limit
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &chain.machineInfos.contractAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(ctx, callMsg)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas limit: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Fetch the nonce for the public address
	nonce, err := client.PendingNonceAt(ctx, publicAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, chain.machineInfos.contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return signedTx.Hash().Hex(), fmt.Errorf("failed to send transaction: %v", err)
	}
	return signedTx.Hash().Hex(), nil
}
