package dbc

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	machineinfos "AIL2-ProveNode/ail2/machine-infos"
	mt "AIL2-ProveNode/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Contract address of the deployed AI contract
const reportContractAddressOnTestnet = "0x5d72d4f8be9055f519cF49a7B5ED3De07FDDDa39"
const machineinfosContractAddressOnTestnet = "0xF9335c71583132d58E5320f73713beEf6da5257D"

const dbcTestNetChainID = 19850818

const dbcTestNetRPC = "https://rpc-testnet.dbcwallet.io"

// go test -v -timeout 30s -count=1 -run TestDbcContract AIL2-ProveNode/ail2
func TestDbcContract(t *testing.T) {
	// Load ABI from file
	abiFile := "ai_abi.json"
	abiData, err := os.ReadFile(abiFile)
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Connect to Ethereum node
	client, err := ethclient.Dial(dbcTestNetRPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	cid, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain id: %v", err)
	}
	log.Printf("chain id: %v", cid.Int64())

	// Private key of the account that will sign the transaction
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(dbcTestNetChainID))
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Define the inputs for the report function
	notifyType := uint8(1) // Example: NotifyType.MachineRegister
	projectName := "ExampleProject"
	stakingType := uint8(0) // Example: StakingType.ShortTerm
	machineId := "example-machine-id"

	// Encode the function call
	data, err := parsedABI.Pack("report", notifyType, projectName, stakingType, machineId)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Define the target contract address
	toAddress := common.HexToAddress(reportContractAddressOnTestnet)

	// Estimate gas limit
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &toAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Fatalf("Failed to estimate gas limit: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Fetch the nonce for the public address
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s", signedTx.Hash().Hex())
}

// go test -v -timeout 120s -count=1 -run TestContractReport AIL2-ProveNode/ail2
func TestContractReport(t *testing.T) {
	ctx := context.Background()
	var cancel context.CancelFunc
	if deadline, ok := t.Deadline(); ok {
		ctx, cancel = context.WithDeadline(ctx, deadline.Add(-time.Second))
		defer cancel()
	}
	chainConfig := mt.ChainConfig{
		Rpc:        dbcTestNetRPC,
		ChainId:    dbcTestNetChainID,
		PrivateKey: os.Getenv("PRIVATE_KEY"),
		ReportContract: mt.ContractConfig{
			AbiFile:         "ai_abi.json",
			ContractAddress: reportContractAddressOnTestnet,
		},
		MachineInfoContract: mt.ContractConfig{
			AbiFile:         "machineinfos.json",
			ContractAddress: machineinfosContractAddressOnTestnet,
		},
	}
	if err := InitDbcChain(ctx, chainConfig); err != nil {
		log.Fatalf("Failed to init dbc chain: %v", err)
	}

	isOnline, isRegistered, err := DbcChain.GetMachineState(ctx, "deeplink", "123456789", mt.ShortTerm)
	log.Printf("Get machine state {Online: %v, Registered: %v} %v", isOnline, isRegistered, err)

	// hash, err := chain.Report(ctx, mt.MachineUnregister, mt.ShortTerm, "deeplink", "123456789")
	// if err != nil {
	// 	log.Fatalf("Unregister machine failed: %v with hash %v", err, hash)
	// }
	// log.Printf("Unregister machine success with hash %v", hash)

	hash, err := DbcChain.Report(ctx, mt.MachineRegister, mt.ShortTerm, "deeplink", "123456789")
	if err != nil {
		log.Fatalf("Register machine failed: %v with hash %v", err, hash)
	}
	log.Printf("Register machine success with hash %v", hash)

	isOnline, isRegistered, err = DbcChain.GetMachineState(ctx, "deeplink", "123456789", mt.ShortTerm)
	log.Printf("Get machine state {Online: %v, Registered: %v} %v", isOnline, isRegistered, err)

	// hash, err = chain.Report(ctx, mt.MachineRegister, mt.ShortTerm, "deeplink", "123456789")
	// if err != nil {
	// 	log.Fatalf("Register machine failed: %v with hash %v", err, hash)
	// }
	// log.Printf("Register machine success with hash %v", hash)

	hash, err = DbcChain.Report(ctx, mt.MachineOnline, mt.ShortTerm, "deeplink", "123456789")
	if err != nil {
		log.Fatalf("Online machine failed: %v with hash %v", err, hash)
	}
	log.Printf("Online machine success with hash %v", hash)

	isOnline, isRegistered, err = DbcChain.GetMachineState(ctx, "deeplink", "123456789", mt.ShortTerm)
	log.Printf("Get machine state {Online: %v, Registered: %v} %v", isOnline, isRegistered, err)

	hash, err = DbcChain.Report(ctx, mt.MachineOffline, mt.ShortTerm, "deeplink", "123456789")
	if err != nil {
		log.Fatalf("Offline machine failed: %v with hash %v", err, hash)
	}
	log.Printf("Offline machine success with hash %v", hash)

	isOnline, isRegistered, err = DbcChain.GetMachineState(ctx, "deeplink", "123456789", mt.ShortTerm)
	log.Printf("Get machine state {Online: %v, Registered: %v} %v", isOnline, isRegistered, err)

	hash, err = DbcChain.Report(ctx, mt.MachineUnregister, mt.ShortTerm, "deeplink", "123456789")
	if err != nil {
		log.Fatalf("Unregister machine failed: %v with hash %v", err, hash)
	}
	log.Printf("Unregister machine success with hash %v", hash)

	isOnline, isRegistered, err = DbcChain.GetMachineState(ctx, "deeplink", "123456789", mt.ShortTerm)
	log.Printf("Get machine state {Online: %v, Registered: %v} %v", isOnline, isRegistered, err)
}

// go test -v -timeout 30s -count=1 -run TestGetMachineInfo AIL2-ProveNode/ail2
func TestGetMachineInfo(t *testing.T) {
	// Define the target contract address
	toAddress := common.HexToAddress(machineinfosContractAddressOnTestnet)

	// Connect to Ethereum node
	client, err := ethclient.Dial(dbcTestNetRPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	cid, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain id: %v", err)
	}
	log.Printf("chain id: %v", cid.Int64())

	instance, err := machineinfos.NewMachineinfos(toAddress, client)
	if err != nil {
		log.Fatalf("Failed to new machineinfos instance: %v", err)
	}

	mi, err := instance.GetMachineInfo(nil, "98900e56cb4a2e0859bde5b1416797d110bd383603c18abcb967877b24a9dabe", false)
	if err != nil {
		log.Fatalf("Failed to get machine info: %v", err)
	}
	log.Printf("Machine owner: %v", mi.MachineOwner.String())
	log.Printf("CalcPoint: %v", mi.CalcPoint.Int64())
	log.Printf("CpuRate: %v", mi.CpuRate.Int64())
	log.Printf("GpuType: %v", mi.GpuType)
	log.Printf("GpuMem: %v", mi.GpuMem.Int64())
	log.Printf("CpuType: %v", mi.CpuType)
	log.Printf("GpuCount: %v", mi.GpuCount.Int64())
	log.Printf("MachineId: %v", mi.MachineId)
	log.Printf("Longitude: %v", mi.Longitude)
	log.Printf("Latitude: %v", mi.Latitude)
	log.Printf("MachineMem: %v", mi.MachineMem.Int64())
}

// go test -v -timeout 60s -count=1 -run TestSetMachineInfoWithAbi AIL2-ProveNode/ail2
func TestSetMachineInfoWithAbi(t *testing.T) {
	// Load ABI from file
	abiData, err := os.ReadFile("machineinfos.json")
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Connect to Ethereum node
	client, err := ethclient.Dial(dbcTestNetRPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	cid, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain id: %v", err)
	}
	log.Printf("chain id: %v", cid.Int64())

	// Private key of the account that will sign the transaction
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, cid)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Define the inputs for the report function
	// notifyType := uint8(1) // Example: NotifyType.MachineRegister
	// projectName := "ExampleProject"
	// stakingType := uint8(0) // Example: StakingType.ShortTerm
	// machineId := "example-machine-id"

	// 'gpu_type': 'GeForceRTX4070SUPER',
	// 'gpu_num': 1,
	// 'cuda_core': 7168,
	// 'gpu_mem':12,
	// 'calc_point':12689,
	// 'sys_disk':100,
	// 'data_disk':35000,
	// 'cpu_type':'13th Gen Intel(R) Core(TM) i7-13790F',
	// 'cpu_core_num':16,
	// 'cpu_rate':2100,
	// 'mem_num':32,
	mi := machineinfos.MachineInfosMachineInfo{
		// MachineOwner: common.HexToAddress("0xc759c00adf496ba19cf80bd878745c9a97a485f7"),
		MachineOwner: publicAddress,
		CalcPoint:    big.NewInt(931500),
		CpuRate:      big.NewInt(3.5e3),
		GpuType:      "NVIDIA GeForce RTX 4060 Ti",
		GpuMem:       big.NewInt(8),
		CpuType:      "13th Gen Intel(R) Core(TM) i7-13790F",
		GpuCount:     big.NewInt(1),
		MachineId:    "",
		Longitude:    "",
		Latitude:     "",
		MachineMem:   big.NewInt(32),
		Region:       "",
	}

	// Encode the function call
	data, err := parsedABI.Pack("setMachineInfo", "123456789", mi)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	// Define the target contract address
	toAddress := common.HexToAddress(machineinfosContractAddressOnTestnet)

	// Estimate gas limit
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &toAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Fatalf("Failed to estimate gas limit: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Fetch the nonce for the public address
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// Sign the transaction
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s", signedTx.Hash().Hex())
}

// go test -v -timeout 60s -count=1 -run TestSetMachineInfoWithoutAbi AIL2-ProveNode/ail2
func TestSetMachineInfoWithoutAbi(t *testing.T) {
	// Define the target contract address
	toAddress := common.HexToAddress(machineinfosContractAddressOnTestnet)

	// Private key of the account that will sign the transaction
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Connect to Ethereum node
	client, err := ethclient.Dial(dbcTestNetRPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	cid, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain id: %v", err)
	}
	log.Printf("chain id: %v", cid.Int64())

	// Load ABI from file
	abiData, err := os.ReadFile("machineinfos.json")
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(abiData)))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Define the inputs for the report function
	// 'gpu_type': 'GeForceRTX4070SUPER',
	// 'gpu_num': 1,
	// 'cuda_core': 7168,
	// 'gpu_mem':12,
	// 'calc_point':12689,
	// 'sys_disk':100,
	// 'data_disk':35000,
	// 'cpu_type':'13th Gen Intel(R) Core(TM) i7-13790F',
	// 'cpu_core_num':16,
	// 'cpu_rate':2100,
	// 'mem_num':32,
	mi := machineinfos.MachineInfosMachineInfo{
		MachineOwner: publicAddress,
		CalcPoint:    big.NewInt(120),
		CpuRate:      big.NewInt(2100),
		GpuType:      "GeForceRTX4070SUPER",
		GpuMem:       big.NewInt(12),
		CpuType:      "13th Gen Intel(R) Core(TM) i7-13790F",
		GpuCount:     big.NewInt(1),
		MachineId:    "123456789",
		Longitude:    "",
		Latitude:     "",
		MachineMem:   big.NewInt(32),
	}

	// Encode the function call
	data, err := parsedABI.Pack("setMachineInfo", "123456789", mi)
	if err != nil {
		log.Fatalf("Failed to pack data: %v", err)
	}

	// Estimate gas limit
	callMsg := ethereum.CallMsg{
		From:  publicAddress,
		To:    &toAddress,
		Gas:   0,
		Value: big.NewInt(0),
		Data:  data,
	}
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Fatalf("Failed to estimate gas limit: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	// Fetch the nonce for the public address
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Create an authorized transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, cid)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	instance, err := machineinfos.NewMachineinfos(toAddress, client)
	if err != nil {
		log.Fatalf("Failed to new machineinfos instance: %v", err)
	}

	tx, err := instance.SetMachineInfo(auth, "123456789", mi)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Printf("Transaction sent: %s", tx.Hash().Hex())
}
