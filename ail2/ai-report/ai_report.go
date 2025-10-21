// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aireport

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AireportMetaData contains all meta data concerning the Aireport contract.
var AireportMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"AddAuthorizedReporter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"canUpgradeAddress\",\"type\":\"address\"}],\"name\":\"AuthorizedUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toBeNotifiedMachineStateUpdateContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toReportStakingStatusContractAddress\",\"type\":\"address\"}],\"name\":\"ContractRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dbcContractAddr\",\"type\":\"address\"}],\"name\":\"DBCContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"MachineInfoContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumAI.NotifyType\",\"name\":\"tp\",\"type\":\"uint8\"}],\"name\":\"MachineStateUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"targetContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumAI.NotifyType\",\"name\":\"tp\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"NotifiedTargetContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"RemoveAuthorizedReporter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAI.NotifyType\",\"name\":\"tp\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ReportFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumAI.StakingType\",\"name\":\"tp\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gpuNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isStake\",\"type\":\"bool\"}],\"name\":\"reportedStakingStatus\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"addAuthorizedReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedReporters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canUpgradeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dbcContract\",\"outputs\":[{\"internalType\":\"contractDBCStakingContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"}],\"name\":\"deleteRegisteredProjectStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"freeGpuAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isDeepLink\",\"type\":\"bool\"}],\"name\":\"getMachineInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"calcPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"mem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getMachineRegion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"}],\"name\":\"getMachineState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"machineInfoContractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dbcContractAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_authorizedReporters\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"machineInProject2States\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"updateAtTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"machineInfoContract\",\"outputs\":[{\"internalType\":\"contractMachineInfoContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"enumAI.StakingType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"projectName2StakingContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"toBeNotifiedMachineStateUpdateContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toReportStakingStatusContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"toBeNotifiedMachineStateUpdateContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toReportStakingStatusContractAddress\",\"type\":\"address\"}],\"name\":\"registerProjectStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"}],\"name\":\"removeAuthorizedReporter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumAI.NotifyType\",\"name\":\"tp\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"projectName\",\"type\":\"string\"},{\"internalType\":\"enumAI.StakingType\",\"name\":\"stakingType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuNum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isStake\",\"type\":\"bool\"}],\"name\":\"reportStakingStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_canUpgradeAddress\",\"type\":\"address\"}],\"name\":\"requestSetUpgradePermission\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dbcContractAddr\",\"type\":\"address\"}],\"name\":\"setDBCContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setMachineInfoContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_canUpgradeAddress\",\"type\":\"address\"}],\"name\":\"setUpgradePermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AireportABI is the input ABI used to generate the binding from.
// Deprecated: Use AireportMetaData.ABI instead.
var AireportABI = AireportMetaData.ABI

// Aireport is an auto generated Go binding around an Ethereum contract.
type Aireport struct {
	AireportCaller     // Read-only binding to the contract
	AireportTransactor // Write-only binding to the contract
	AireportFilterer   // Log filterer for contract events
}

// AireportCaller is an auto generated read-only Go binding around an Ethereum contract.
type AireportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AireportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AireportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AireportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AireportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AireportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AireportSession struct {
	Contract     *Aireport         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AireportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AireportCallerSession struct {
	Contract *AireportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AireportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AireportTransactorSession struct {
	Contract     *AireportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AireportRaw is an auto generated low-level Go binding around an Ethereum contract.
type AireportRaw struct {
	Contract *Aireport // Generic contract binding to access the raw methods on
}

// AireportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AireportCallerRaw struct {
	Contract *AireportCaller // Generic read-only contract binding to access the raw methods on
}

// AireportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AireportTransactorRaw struct {
	Contract *AireportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAireport creates a new instance of Aireport, bound to a specific deployed contract.
func NewAireport(address common.Address, backend bind.ContractBackend) (*Aireport, error) {
	contract, err := bindAireport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aireport{AireportCaller: AireportCaller{contract: contract}, AireportTransactor: AireportTransactor{contract: contract}, AireportFilterer: AireportFilterer{contract: contract}}, nil
}

// NewAireportCaller creates a new read-only instance of Aireport, bound to a specific deployed contract.
func NewAireportCaller(address common.Address, caller bind.ContractCaller) (*AireportCaller, error) {
	contract, err := bindAireport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AireportCaller{contract: contract}, nil
}

// NewAireportTransactor creates a new write-only instance of Aireport, bound to a specific deployed contract.
func NewAireportTransactor(address common.Address, transactor bind.ContractTransactor) (*AireportTransactor, error) {
	contract, err := bindAireport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AireportTransactor{contract: contract}, nil
}

// NewAireportFilterer creates a new log filterer instance of Aireport, bound to a specific deployed contract.
func NewAireportFilterer(address common.Address, filterer bind.ContractFilterer) (*AireportFilterer, error) {
	contract, err := bindAireport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AireportFilterer{contract: contract}, nil
}

// bindAireport binds a generic wrapper to an already deployed contract.
func bindAireport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AireportMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aireport *AireportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aireport.Contract.AireportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aireport *AireportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aireport.Contract.AireportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aireport *AireportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aireport.Contract.AireportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aireport *AireportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aireport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aireport *AireportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aireport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aireport *AireportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aireport.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedReporters is a free data retrieval call binding the contract method 0x3a41429a.
//
// Solidity: function authorizedReporters(address ) view returns(bool)
func (_Aireport *AireportCaller) AuthorizedReporters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "authorizedReporters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedReporters is a free data retrieval call binding the contract method 0x3a41429a.
//
// Solidity: function authorizedReporters(address ) view returns(bool)
func (_Aireport *AireportSession) AuthorizedReporters(arg0 common.Address) (bool, error) {
	return _Aireport.Contract.AuthorizedReporters(&_Aireport.CallOpts, arg0)
}

// AuthorizedReporters is a free data retrieval call binding the contract method 0x3a41429a.
//
// Solidity: function authorizedReporters(address ) view returns(bool)
func (_Aireport *AireportCallerSession) AuthorizedReporters(arg0 common.Address) (bool, error) {
	return _Aireport.Contract.AuthorizedReporters(&_Aireport.CallOpts, arg0)
}

// CanUpgradeAddress is a free data retrieval call binding the contract method 0x75dfe221.
//
// Solidity: function canUpgradeAddress() view returns(address)
func (_Aireport *AireportCaller) CanUpgradeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "canUpgradeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanUpgradeAddress is a free data retrieval call binding the contract method 0x75dfe221.
//
// Solidity: function canUpgradeAddress() view returns(address)
func (_Aireport *AireportSession) CanUpgradeAddress() (common.Address, error) {
	return _Aireport.Contract.CanUpgradeAddress(&_Aireport.CallOpts)
}

// CanUpgradeAddress is a free data retrieval call binding the contract method 0x75dfe221.
//
// Solidity: function canUpgradeAddress() view returns(address)
func (_Aireport *AireportCallerSession) CanUpgradeAddress() (common.Address, error) {
	return _Aireport.Contract.CanUpgradeAddress(&_Aireport.CallOpts)
}

// DbcContract is a free data retrieval call binding the contract method 0xde1dc9bc.
//
// Solidity: function dbcContract() view returns(address)
func (_Aireport *AireportCaller) DbcContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "dbcContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DbcContract is a free data retrieval call binding the contract method 0xde1dc9bc.
//
// Solidity: function dbcContract() view returns(address)
func (_Aireport *AireportSession) DbcContract() (common.Address, error) {
	return _Aireport.Contract.DbcContract(&_Aireport.CallOpts)
}

// DbcContract is a free data retrieval call binding the contract method 0xde1dc9bc.
//
// Solidity: function dbcContract() view returns(address)
func (_Aireport *AireportCallerSession) DbcContract() (common.Address, error) {
	return _Aireport.Contract.DbcContract(&_Aireport.CallOpts)
}

// FreeGpuAmount is a free data retrieval call binding the contract method 0xaa41d9ba.
//
// Solidity: function freeGpuAmount(string _id) view returns(uint256)
func (_Aireport *AireportCaller) FreeGpuAmount(opts *bind.CallOpts, _id string) (*big.Int, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "freeGpuAmount", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeGpuAmount is a free data retrieval call binding the contract method 0xaa41d9ba.
//
// Solidity: function freeGpuAmount(string _id) view returns(uint256)
func (_Aireport *AireportSession) FreeGpuAmount(_id string) (*big.Int, error) {
	return _Aireport.Contract.FreeGpuAmount(&_Aireport.CallOpts, _id)
}

// FreeGpuAmount is a free data retrieval call binding the contract method 0xaa41d9ba.
//
// Solidity: function freeGpuAmount(string _id) view returns(uint256)
func (_Aireport *AireportCallerSession) FreeGpuAmount(_id string) (*big.Int, error) {
	return _Aireport.Contract.FreeGpuAmount(&_Aireport.CallOpts, _id)
}

// GetMachineInfo is a free data retrieval call binding the contract method 0xc8ffe871.
//
// Solidity: function getMachineInfo(string id, bool isDeepLink) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, uint256 mem)
func (_Aireport *AireportCaller) GetMachineInfo(opts *bind.CallOpts, id string, isDeepLink bool) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Mem          *big.Int
}, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "getMachineInfo", id, isDeepLink)

	outstruct := new(struct {
		MachineOwner common.Address
		CalcPoint    *big.Int
		CpuRate      *big.Int
		GpuType      string
		GpuMem       *big.Int
		CpuType      string
		GpuCount     *big.Int
		MachineId    string
		Mem          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MachineOwner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CalcPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CpuRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.GpuType = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.GpuMem = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CpuType = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.GpuCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MachineId = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Mem = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMachineInfo is a free data retrieval call binding the contract method 0xc8ffe871.
//
// Solidity: function getMachineInfo(string id, bool isDeepLink) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, uint256 mem)
func (_Aireport *AireportSession) GetMachineInfo(id string, isDeepLink bool) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Mem          *big.Int
}, error) {
	return _Aireport.Contract.GetMachineInfo(&_Aireport.CallOpts, id, isDeepLink)
}

// GetMachineInfo is a free data retrieval call binding the contract method 0xc8ffe871.
//
// Solidity: function getMachineInfo(string id, bool isDeepLink) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, uint256 mem)
func (_Aireport *AireportCallerSession) GetMachineInfo(id string, isDeepLink bool) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Mem          *big.Int
}, error) {
	return _Aireport.Contract.GetMachineInfo(&_Aireport.CallOpts, id, isDeepLink)
}

// GetMachineRegion is a free data retrieval call binding the contract method 0x1b97cc77.
//
// Solidity: function getMachineRegion(string _id) view returns(string)
func (_Aireport *AireportCaller) GetMachineRegion(opts *bind.CallOpts, _id string) (string, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "getMachineRegion", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMachineRegion is a free data retrieval call binding the contract method 0x1b97cc77.
//
// Solidity: function getMachineRegion(string _id) view returns(string)
func (_Aireport *AireportSession) GetMachineRegion(_id string) (string, error) {
	return _Aireport.Contract.GetMachineRegion(&_Aireport.CallOpts, _id)
}

// GetMachineRegion is a free data retrieval call binding the contract method 0x1b97cc77.
//
// Solidity: function getMachineRegion(string _id) view returns(string)
func (_Aireport *AireportCallerSession) GetMachineRegion(_id string) (string, error) {
	return _Aireport.Contract.GetMachineRegion(&_Aireport.CallOpts, _id)
}

// GetMachineState is a free data retrieval call binding the contract method 0x49e5950b.
//
// Solidity: function getMachineState(string machineId, string projectName, uint8 stakingType) view returns(bool isOnline, bool isRegistered)
func (_Aireport *AireportCaller) GetMachineState(opts *bind.CallOpts, machineId string, projectName string, stakingType uint8) (struct {
	IsOnline     bool
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "getMachineState", machineId, projectName, stakingType)

	outstruct := new(struct {
		IsOnline     bool
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsOnline = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsRegistered = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetMachineState is a free data retrieval call binding the contract method 0x49e5950b.
//
// Solidity: function getMachineState(string machineId, string projectName, uint8 stakingType) view returns(bool isOnline, bool isRegistered)
func (_Aireport *AireportSession) GetMachineState(machineId string, projectName string, stakingType uint8) (struct {
	IsOnline     bool
	IsRegistered bool
}, error) {
	return _Aireport.Contract.GetMachineState(&_Aireport.CallOpts, machineId, projectName, stakingType)
}

// GetMachineState is a free data retrieval call binding the contract method 0x49e5950b.
//
// Solidity: function getMachineState(string machineId, string projectName, uint8 stakingType) view returns(bool isOnline, bool isRegistered)
func (_Aireport *AireportCallerSession) GetMachineState(machineId string, projectName string, stakingType uint8) (struct {
	IsOnline     bool
	IsRegistered bool
}, error) {
	return _Aireport.Contract.GetMachineState(&_Aireport.CallOpts, machineId, projectName, stakingType)
}

// MachineInProject2States is a free data retrieval call binding the contract method 0x365af396.
//
// Solidity: function machineInProject2States(string ) view returns(bool isOnline, bool isRegistered, uint256 updateAtTimestamp)
func (_Aireport *AireportCaller) MachineInProject2States(opts *bind.CallOpts, arg0 string) (struct {
	IsOnline          bool
	IsRegistered      bool
	UpdateAtTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "machineInProject2States", arg0)

	outstruct := new(struct {
		IsOnline          bool
		IsRegistered      bool
		UpdateAtTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsOnline = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsRegistered = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.UpdateAtTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MachineInProject2States is a free data retrieval call binding the contract method 0x365af396.
//
// Solidity: function machineInProject2States(string ) view returns(bool isOnline, bool isRegistered, uint256 updateAtTimestamp)
func (_Aireport *AireportSession) MachineInProject2States(arg0 string) (struct {
	IsOnline          bool
	IsRegistered      bool
	UpdateAtTimestamp *big.Int
}, error) {
	return _Aireport.Contract.MachineInProject2States(&_Aireport.CallOpts, arg0)
}

// MachineInProject2States is a free data retrieval call binding the contract method 0x365af396.
//
// Solidity: function machineInProject2States(string ) view returns(bool isOnline, bool isRegistered, uint256 updateAtTimestamp)
func (_Aireport *AireportCallerSession) MachineInProject2States(arg0 string) (struct {
	IsOnline          bool
	IsRegistered      bool
	UpdateAtTimestamp *big.Int
}, error) {
	return _Aireport.Contract.MachineInProject2States(&_Aireport.CallOpts, arg0)
}

// MachineInfoContract is a free data retrieval call binding the contract method 0x183d0f0c.
//
// Solidity: function machineInfoContract() view returns(address)
func (_Aireport *AireportCaller) MachineInfoContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "machineInfoContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MachineInfoContract is a free data retrieval call binding the contract method 0x183d0f0c.
//
// Solidity: function machineInfoContract() view returns(address)
func (_Aireport *AireportSession) MachineInfoContract() (common.Address, error) {
	return _Aireport.Contract.MachineInfoContract(&_Aireport.CallOpts)
}

// MachineInfoContract is a free data retrieval call binding the contract method 0x183d0f0c.
//
// Solidity: function machineInfoContract() view returns(address)
func (_Aireport *AireportCallerSession) MachineInfoContract() (common.Address, error) {
	return _Aireport.Contract.MachineInfoContract(&_Aireport.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aireport *AireportCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aireport *AireportSession) Owner() (common.Address, error) {
	return _Aireport.Contract.Owner(&_Aireport.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aireport *AireportCallerSession) Owner() (common.Address, error) {
	return _Aireport.Contract.Owner(&_Aireport.CallOpts)
}

// ProjectName2StakingContractAddress is a free data retrieval call binding the contract method 0xc4146a23.
//
// Solidity: function projectName2StakingContractAddress(string , uint8 ) view returns(address toBeNotifiedMachineStateUpdateContractAddress, address toReportStakingStatusContractAddress)
func (_Aireport *AireportCaller) ProjectName2StakingContractAddress(opts *bind.CallOpts, arg0 string, arg1 uint8) (struct {
	ToBeNotifiedMachineStateUpdateContractAddress common.Address
	ToReportStakingStatusContractAddress          common.Address
}, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "projectName2StakingContractAddress", arg0, arg1)

	outstruct := new(struct {
		ToBeNotifiedMachineStateUpdateContractAddress common.Address
		ToReportStakingStatusContractAddress          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ToBeNotifiedMachineStateUpdateContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ToReportStakingStatusContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ProjectName2StakingContractAddress is a free data retrieval call binding the contract method 0xc4146a23.
//
// Solidity: function projectName2StakingContractAddress(string , uint8 ) view returns(address toBeNotifiedMachineStateUpdateContractAddress, address toReportStakingStatusContractAddress)
func (_Aireport *AireportSession) ProjectName2StakingContractAddress(arg0 string, arg1 uint8) (struct {
	ToBeNotifiedMachineStateUpdateContractAddress common.Address
	ToReportStakingStatusContractAddress          common.Address
}, error) {
	return _Aireport.Contract.ProjectName2StakingContractAddress(&_Aireport.CallOpts, arg0, arg1)
}

// ProjectName2StakingContractAddress is a free data retrieval call binding the contract method 0xc4146a23.
//
// Solidity: function projectName2StakingContractAddress(string , uint8 ) view returns(address toBeNotifiedMachineStateUpdateContractAddress, address toReportStakingStatusContractAddress)
func (_Aireport *AireportCallerSession) ProjectName2StakingContractAddress(arg0 string, arg1 uint8) (struct {
	ToBeNotifiedMachineStateUpdateContractAddress common.Address
	ToReportStakingStatusContractAddress          common.Address
}, error) {
	return _Aireport.Contract.ProjectName2StakingContractAddress(&_Aireport.CallOpts, arg0, arg1)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Aireport *AireportCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Aireport *AireportSession) ProxiableUUID() ([32]byte, error) {
	return _Aireport.Contract.ProxiableUUID(&_Aireport.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Aireport *AireportCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Aireport.Contract.ProxiableUUID(&_Aireport.CallOpts)
}

// RequestSetUpgradePermission is a free data retrieval call binding the contract method 0x95b5d975.
//
// Solidity: function requestSetUpgradePermission(address _canUpgradeAddress) pure returns(bytes)
func (_Aireport *AireportCaller) RequestSetUpgradePermission(opts *bind.CallOpts, _canUpgradeAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _Aireport.contract.Call(opts, &out, "requestSetUpgradePermission", _canUpgradeAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RequestSetUpgradePermission is a free data retrieval call binding the contract method 0x95b5d975.
//
// Solidity: function requestSetUpgradePermission(address _canUpgradeAddress) pure returns(bytes)
func (_Aireport *AireportSession) RequestSetUpgradePermission(_canUpgradeAddress common.Address) ([]byte, error) {
	return _Aireport.Contract.RequestSetUpgradePermission(&_Aireport.CallOpts, _canUpgradeAddress)
}

// RequestSetUpgradePermission is a free data retrieval call binding the contract method 0x95b5d975.
//
// Solidity: function requestSetUpgradePermission(address _canUpgradeAddress) pure returns(bytes)
func (_Aireport *AireportCallerSession) RequestSetUpgradePermission(_canUpgradeAddress common.Address) ([]byte, error) {
	return _Aireport.Contract.RequestSetUpgradePermission(&_Aireport.CallOpts, _canUpgradeAddress)
}

// AddAuthorizedReporter is a paid mutator transaction binding the contract method 0xe63fb88f.
//
// Solidity: function addAuthorizedReporter(address reporter) returns()
func (_Aireport *AireportTransactor) AddAuthorizedReporter(opts *bind.TransactOpts, reporter common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "addAuthorizedReporter", reporter)
}

// AddAuthorizedReporter is a paid mutator transaction binding the contract method 0xe63fb88f.
//
// Solidity: function addAuthorizedReporter(address reporter) returns()
func (_Aireport *AireportSession) AddAuthorizedReporter(reporter common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.AddAuthorizedReporter(&_Aireport.TransactOpts, reporter)
}

// AddAuthorizedReporter is a paid mutator transaction binding the contract method 0xe63fb88f.
//
// Solidity: function addAuthorizedReporter(address reporter) returns()
func (_Aireport *AireportTransactorSession) AddAuthorizedReporter(reporter common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.AddAuthorizedReporter(&_Aireport.TransactOpts, reporter)
}

// DeleteRegisteredProjectStakingContract is a paid mutator transaction binding the contract method 0x77051437.
//
// Solidity: function deleteRegisteredProjectStakingContract(string projectName, uint8 stakingType) returns()
func (_Aireport *AireportTransactor) DeleteRegisteredProjectStakingContract(opts *bind.TransactOpts, projectName string, stakingType uint8) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "deleteRegisteredProjectStakingContract", projectName, stakingType)
}

// DeleteRegisteredProjectStakingContract is a paid mutator transaction binding the contract method 0x77051437.
//
// Solidity: function deleteRegisteredProjectStakingContract(string projectName, uint8 stakingType) returns()
func (_Aireport *AireportSession) DeleteRegisteredProjectStakingContract(projectName string, stakingType uint8) (*types.Transaction, error) {
	return _Aireport.Contract.DeleteRegisteredProjectStakingContract(&_Aireport.TransactOpts, projectName, stakingType)
}

// DeleteRegisteredProjectStakingContract is a paid mutator transaction binding the contract method 0x77051437.
//
// Solidity: function deleteRegisteredProjectStakingContract(string projectName, uint8 stakingType) returns()
func (_Aireport *AireportTransactorSession) DeleteRegisteredProjectStakingContract(projectName string, stakingType uint8) (*types.Transaction, error) {
	return _Aireport.Contract.DeleteRegisteredProjectStakingContract(&_Aireport.TransactOpts, projectName, stakingType)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address machineInfoContractAddr, address dbcContractAddr, address[] _authorizedReporters) returns()
func (_Aireport *AireportTransactor) Initialize(opts *bind.TransactOpts, machineInfoContractAddr common.Address, dbcContractAddr common.Address, _authorizedReporters []common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "initialize", machineInfoContractAddr, dbcContractAddr, _authorizedReporters)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address machineInfoContractAddr, address dbcContractAddr, address[] _authorizedReporters) returns()
func (_Aireport *AireportSession) Initialize(machineInfoContractAddr common.Address, dbcContractAddr common.Address, _authorizedReporters []common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.Initialize(&_Aireport.TransactOpts, machineInfoContractAddr, dbcContractAddr, _authorizedReporters)
}

// Initialize is a paid mutator transaction binding the contract method 0x77a24f36.
//
// Solidity: function initialize(address machineInfoContractAddr, address dbcContractAddr, address[] _authorizedReporters) returns()
func (_Aireport *AireportTransactorSession) Initialize(machineInfoContractAddr common.Address, dbcContractAddr common.Address, _authorizedReporters []common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.Initialize(&_Aireport.TransactOpts, machineInfoContractAddr, dbcContractAddr, _authorizedReporters)
}

// RegisterProjectStakingContract is a paid mutator transaction binding the contract method 0xb5cf512f.
//
// Solidity: function registerProjectStakingContract(string projectName, uint8 stakingType, address toBeNotifiedMachineStateUpdateContractAddress, address toReportStakingStatusContractAddress) returns()
func (_Aireport *AireportTransactor) RegisterProjectStakingContract(opts *bind.TransactOpts, projectName string, stakingType uint8, toBeNotifiedMachineStateUpdateContractAddress common.Address, toReportStakingStatusContractAddress common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "registerProjectStakingContract", projectName, stakingType, toBeNotifiedMachineStateUpdateContractAddress, toReportStakingStatusContractAddress)
}

// RegisterProjectStakingContract is a paid mutator transaction binding the contract method 0xb5cf512f.
//
// Solidity: function registerProjectStakingContract(string projectName, uint8 stakingType, address toBeNotifiedMachineStateUpdateContractAddress, address toReportStakingStatusContractAddress) returns()
func (_Aireport *AireportSession) RegisterProjectStakingContract(projectName string, stakingType uint8, toBeNotifiedMachineStateUpdateContractAddress common.Address, toReportStakingStatusContractAddress common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.RegisterProjectStakingContract(&_Aireport.TransactOpts, projectName, stakingType, toBeNotifiedMachineStateUpdateContractAddress, toReportStakingStatusContractAddress)
}

// RegisterProjectStakingContract is a paid mutator transaction binding the contract method 0xb5cf512f.
//
// Solidity: function registerProjectStakingContract(string projectName, uint8 stakingType, address toBeNotifiedMachineStateUpdateContractAddress, address toReportStakingStatusContractAddress) returns()
func (_Aireport *AireportTransactorSession) RegisterProjectStakingContract(projectName string, stakingType uint8, toBeNotifiedMachineStateUpdateContractAddress common.Address, toReportStakingStatusContractAddress common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.RegisterProjectStakingContract(&_Aireport.TransactOpts, projectName, stakingType, toBeNotifiedMachineStateUpdateContractAddress, toReportStakingStatusContractAddress)
}

// RemoveAuthorizedReporter is a paid mutator transaction binding the contract method 0xa6b2f395.
//
// Solidity: function removeAuthorizedReporter(address reporter) returns()
func (_Aireport *AireportTransactor) RemoveAuthorizedReporter(opts *bind.TransactOpts, reporter common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "removeAuthorizedReporter", reporter)
}

// RemoveAuthorizedReporter is a paid mutator transaction binding the contract method 0xa6b2f395.
//
// Solidity: function removeAuthorizedReporter(address reporter) returns()
func (_Aireport *AireportSession) RemoveAuthorizedReporter(reporter common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.RemoveAuthorizedReporter(&_Aireport.TransactOpts, reporter)
}

// RemoveAuthorizedReporter is a paid mutator transaction binding the contract method 0xa6b2f395.
//
// Solidity: function removeAuthorizedReporter(address reporter) returns()
func (_Aireport *AireportTransactorSession) RemoveAuthorizedReporter(reporter common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.RemoveAuthorizedReporter(&_Aireport.TransactOpts, reporter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aireport *AireportTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aireport *AireportSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aireport.Contract.RenounceOwnership(&_Aireport.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aireport *AireportTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aireport.Contract.RenounceOwnership(&_Aireport.TransactOpts)
}

// Report is a paid mutator transaction binding the contract method 0xcee5a655.
//
// Solidity: function report(uint8 tp, string projectName, uint8 stakingType, string machineId) returns()
func (_Aireport *AireportTransactor) Report(opts *bind.TransactOpts, tp uint8, projectName string, stakingType uint8, machineId string) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "report", tp, projectName, stakingType, machineId)
}

// Report is a paid mutator transaction binding the contract method 0xcee5a655.
//
// Solidity: function report(uint8 tp, string projectName, uint8 stakingType, string machineId) returns()
func (_Aireport *AireportSession) Report(tp uint8, projectName string, stakingType uint8, machineId string) (*types.Transaction, error) {
	return _Aireport.Contract.Report(&_Aireport.TransactOpts, tp, projectName, stakingType, machineId)
}

// Report is a paid mutator transaction binding the contract method 0xcee5a655.
//
// Solidity: function report(uint8 tp, string projectName, uint8 stakingType, string machineId) returns()
func (_Aireport *AireportTransactorSession) Report(tp uint8, projectName string, stakingType uint8, machineId string) (*types.Transaction, error) {
	return _Aireport.Contract.Report(&_Aireport.TransactOpts, tp, projectName, stakingType, machineId)
}

// ReportStakingStatus is a paid mutator transaction binding the contract method 0x9359f17a.
//
// Solidity: function reportStakingStatus(string projectName, uint8 stakingType, string id, uint256 gpuNum, bool isStake) returns()
func (_Aireport *AireportTransactor) ReportStakingStatus(opts *bind.TransactOpts, projectName string, stakingType uint8, id string, gpuNum *big.Int, isStake bool) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "reportStakingStatus", projectName, stakingType, id, gpuNum, isStake)
}

// ReportStakingStatus is a paid mutator transaction binding the contract method 0x9359f17a.
//
// Solidity: function reportStakingStatus(string projectName, uint8 stakingType, string id, uint256 gpuNum, bool isStake) returns()
func (_Aireport *AireportSession) ReportStakingStatus(projectName string, stakingType uint8, id string, gpuNum *big.Int, isStake bool) (*types.Transaction, error) {
	return _Aireport.Contract.ReportStakingStatus(&_Aireport.TransactOpts, projectName, stakingType, id, gpuNum, isStake)
}

// ReportStakingStatus is a paid mutator transaction binding the contract method 0x9359f17a.
//
// Solidity: function reportStakingStatus(string projectName, uint8 stakingType, string id, uint256 gpuNum, bool isStake) returns()
func (_Aireport *AireportTransactorSession) ReportStakingStatus(projectName string, stakingType uint8, id string, gpuNum *big.Int, isStake bool) (*types.Transaction, error) {
	return _Aireport.Contract.ReportStakingStatus(&_Aireport.TransactOpts, projectName, stakingType, id, gpuNum, isStake)
}

// SetDBCContract is a paid mutator transaction binding the contract method 0xc80dd5bb.
//
// Solidity: function setDBCContract(address dbcContractAddr) returns()
func (_Aireport *AireportTransactor) SetDBCContract(opts *bind.TransactOpts, dbcContractAddr common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "setDBCContract", dbcContractAddr)
}

// SetDBCContract is a paid mutator transaction binding the contract method 0xc80dd5bb.
//
// Solidity: function setDBCContract(address dbcContractAddr) returns()
func (_Aireport *AireportSession) SetDBCContract(dbcContractAddr common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.SetDBCContract(&_Aireport.TransactOpts, dbcContractAddr)
}

// SetDBCContract is a paid mutator transaction binding the contract method 0xc80dd5bb.
//
// Solidity: function setDBCContract(address dbcContractAddr) returns()
func (_Aireport *AireportTransactorSession) SetDBCContract(dbcContractAddr common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.SetDBCContract(&_Aireport.TransactOpts, dbcContractAddr)
}

// SetMachineInfoContract is a paid mutator transaction binding the contract method 0x8694541f.
//
// Solidity: function setMachineInfoContract(address addr) returns()
func (_Aireport *AireportTransactor) SetMachineInfoContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "setMachineInfoContract", addr)
}

// SetMachineInfoContract is a paid mutator transaction binding the contract method 0x8694541f.
//
// Solidity: function setMachineInfoContract(address addr) returns()
func (_Aireport *AireportSession) SetMachineInfoContract(addr common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.SetMachineInfoContract(&_Aireport.TransactOpts, addr)
}

// SetMachineInfoContract is a paid mutator transaction binding the contract method 0x8694541f.
//
// Solidity: function setMachineInfoContract(address addr) returns()
func (_Aireport *AireportTransactorSession) SetMachineInfoContract(addr common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.SetMachineInfoContract(&_Aireport.TransactOpts, addr)
}

// SetUpgradePermission is a paid mutator transaction binding the contract method 0x4a59e7bf.
//
// Solidity: function setUpgradePermission(address _canUpgradeAddress) returns()
func (_Aireport *AireportTransactor) SetUpgradePermission(opts *bind.TransactOpts, _canUpgradeAddress common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "setUpgradePermission", _canUpgradeAddress)
}

// SetUpgradePermission is a paid mutator transaction binding the contract method 0x4a59e7bf.
//
// Solidity: function setUpgradePermission(address _canUpgradeAddress) returns()
func (_Aireport *AireportSession) SetUpgradePermission(_canUpgradeAddress common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.SetUpgradePermission(&_Aireport.TransactOpts, _canUpgradeAddress)
}

// SetUpgradePermission is a paid mutator transaction binding the contract method 0x4a59e7bf.
//
// Solidity: function setUpgradePermission(address _canUpgradeAddress) returns()
func (_Aireport *AireportTransactorSession) SetUpgradePermission(_canUpgradeAddress common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.SetUpgradePermission(&_Aireport.TransactOpts, _canUpgradeAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aireport *AireportTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aireport *AireportSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.TransferOwnership(&_Aireport.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aireport *AireportTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.TransferOwnership(&_Aireport.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Aireport *AireportTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Aireport *AireportSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.UpgradeTo(&_Aireport.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Aireport *AireportTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Aireport.Contract.UpgradeTo(&_Aireport.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Aireport *AireportTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Aireport.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Aireport *AireportSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Aireport.Contract.UpgradeToAndCall(&_Aireport.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Aireport *AireportTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Aireport.Contract.UpgradeToAndCall(&_Aireport.TransactOpts, newImplementation, data)
}

// AireportAddAuthorizedReporterIterator is returned from FilterAddAuthorizedReporter and is used to iterate over the raw logs and unpacked data for AddAuthorizedReporter events raised by the Aireport contract.
type AireportAddAuthorizedReporterIterator struct {
	Event *AireportAddAuthorizedReporter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportAddAuthorizedReporterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportAddAuthorizedReporter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportAddAuthorizedReporter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportAddAuthorizedReporterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportAddAuthorizedReporterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportAddAuthorizedReporter represents a AddAuthorizedReporter event raised by the Aireport contract.
type AireportAddAuthorizedReporter struct {
	Reporter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddAuthorizedReporter is a free log retrieval operation binding the contract event 0xa0aa2a8303c134fa8985c6e53af53dccad59dd37ed022720d4d082ed36e1a9d9.
//
// Solidity: event AddAuthorizedReporter(address indexed reporter)
func (_Aireport *AireportFilterer) FilterAddAuthorizedReporter(opts *bind.FilterOpts, reporter []common.Address) (*AireportAddAuthorizedReporterIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "AddAuthorizedReporter", reporterRule)
	if err != nil {
		return nil, err
	}
	return &AireportAddAuthorizedReporterIterator{contract: _Aireport.contract, event: "AddAuthorizedReporter", logs: logs, sub: sub}, nil
}

// WatchAddAuthorizedReporter is a free log subscription operation binding the contract event 0xa0aa2a8303c134fa8985c6e53af53dccad59dd37ed022720d4d082ed36e1a9d9.
//
// Solidity: event AddAuthorizedReporter(address indexed reporter)
func (_Aireport *AireportFilterer) WatchAddAuthorizedReporter(opts *bind.WatchOpts, sink chan<- *AireportAddAuthorizedReporter, reporter []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "AddAuthorizedReporter", reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportAddAuthorizedReporter)
				if err := _Aireport.contract.UnpackLog(event, "AddAuthorizedReporter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddAuthorizedReporter is a log parse operation binding the contract event 0xa0aa2a8303c134fa8985c6e53af53dccad59dd37ed022720d4d082ed36e1a9d9.
//
// Solidity: event AddAuthorizedReporter(address indexed reporter)
func (_Aireport *AireportFilterer) ParseAddAuthorizedReporter(log types.Log) (*AireportAddAuthorizedReporter, error) {
	event := new(AireportAddAuthorizedReporter)
	if err := _Aireport.contract.UnpackLog(event, "AddAuthorizedReporter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Aireport contract.
type AireportAdminChangedIterator struct {
	Event *AireportAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportAdminChanged represents a AdminChanged event raised by the Aireport contract.
type AireportAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Aireport *AireportFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AireportAdminChangedIterator, error) {

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AireportAdminChangedIterator{contract: _Aireport.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Aireport *AireportFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AireportAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportAdminChanged)
				if err := _Aireport.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Aireport *AireportFilterer) ParseAdminChanged(log types.Log) (*AireportAdminChanged, error) {
	event := new(AireportAdminChanged)
	if err := _Aireport.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportAuthorizedUpgradeIterator is returned from FilterAuthorizedUpgrade and is used to iterate over the raw logs and unpacked data for AuthorizedUpgrade events raised by the Aireport contract.
type AireportAuthorizedUpgradeIterator struct {
	Event *AireportAuthorizedUpgrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportAuthorizedUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportAuthorizedUpgrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportAuthorizedUpgrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportAuthorizedUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportAuthorizedUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportAuthorizedUpgrade represents a AuthorizedUpgrade event raised by the Aireport contract.
type AireportAuthorizedUpgrade struct {
	CanUpgradeAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedUpgrade is a free log retrieval operation binding the contract event 0x8d93ff81d7be78d4fd10818ecc85b3a529925b2149689913097f4039345cb62c.
//
// Solidity: event AuthorizedUpgrade(address indexed canUpgradeAddress)
func (_Aireport *AireportFilterer) FilterAuthorizedUpgrade(opts *bind.FilterOpts, canUpgradeAddress []common.Address) (*AireportAuthorizedUpgradeIterator, error) {

	var canUpgradeAddressRule []interface{}
	for _, canUpgradeAddressItem := range canUpgradeAddress {
		canUpgradeAddressRule = append(canUpgradeAddressRule, canUpgradeAddressItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "AuthorizedUpgrade", canUpgradeAddressRule)
	if err != nil {
		return nil, err
	}
	return &AireportAuthorizedUpgradeIterator{contract: _Aireport.contract, event: "AuthorizedUpgrade", logs: logs, sub: sub}, nil
}

// WatchAuthorizedUpgrade is a free log subscription operation binding the contract event 0x8d93ff81d7be78d4fd10818ecc85b3a529925b2149689913097f4039345cb62c.
//
// Solidity: event AuthorizedUpgrade(address indexed canUpgradeAddress)
func (_Aireport *AireportFilterer) WatchAuthorizedUpgrade(opts *bind.WatchOpts, sink chan<- *AireportAuthorizedUpgrade, canUpgradeAddress []common.Address) (event.Subscription, error) {

	var canUpgradeAddressRule []interface{}
	for _, canUpgradeAddressItem := range canUpgradeAddress {
		canUpgradeAddressRule = append(canUpgradeAddressRule, canUpgradeAddressItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "AuthorizedUpgrade", canUpgradeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportAuthorizedUpgrade)
				if err := _Aireport.contract.UnpackLog(event, "AuthorizedUpgrade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAuthorizedUpgrade is a log parse operation binding the contract event 0x8d93ff81d7be78d4fd10818ecc85b3a529925b2149689913097f4039345cb62c.
//
// Solidity: event AuthorizedUpgrade(address indexed canUpgradeAddress)
func (_Aireport *AireportFilterer) ParseAuthorizedUpgrade(log types.Log) (*AireportAuthorizedUpgrade, error) {
	event := new(AireportAuthorizedUpgrade)
	if err := _Aireport.contract.UnpackLog(event, "AuthorizedUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Aireport contract.
type AireportBeaconUpgradedIterator struct {
	Event *AireportBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportBeaconUpgraded represents a BeaconUpgraded event raised by the Aireport contract.
type AireportBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Aireport *AireportFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AireportBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AireportBeaconUpgradedIterator{contract: _Aireport.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Aireport *AireportFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AireportBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportBeaconUpgraded)
				if err := _Aireport.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Aireport *AireportFilterer) ParseBeaconUpgraded(log types.Log) (*AireportBeaconUpgraded, error) {
	event := new(AireportBeaconUpgraded)
	if err := _Aireport.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportContractRegisterIterator is returned from FilterContractRegister and is used to iterate over the raw logs and unpacked data for ContractRegister events raised by the Aireport contract.
type AireportContractRegisterIterator struct {
	Event *AireportContractRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportContractRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportContractRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportContractRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportContractRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportContractRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportContractRegister represents a ContractRegister event raised by the Aireport contract.
type AireportContractRegister struct {
	Caller                                        common.Address
	ProjectName                                   string
	ToBeNotifiedMachineStateUpdateContractAddress common.Address
	ToReportStakingStatusContractAddress          common.Address
	Raw                                           types.Log // Blockchain specific contextual infos
}

// FilterContractRegister is a free log retrieval operation binding the contract event 0x3bb0c9c4e86902fc67f92deb14cc8e8775454751bb8d7d5019d59daa06b9010d.
//
// Solidity: event ContractRegister(address indexed caller, string projectName, address indexed toBeNotifiedMachineStateUpdateContractAddress, address indexed toReportStakingStatusContractAddress)
func (_Aireport *AireportFilterer) FilterContractRegister(opts *bind.FilterOpts, caller []common.Address, toBeNotifiedMachineStateUpdateContractAddress []common.Address, toReportStakingStatusContractAddress []common.Address) (*AireportContractRegisterIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var toBeNotifiedMachineStateUpdateContractAddressRule []interface{}
	for _, toBeNotifiedMachineStateUpdateContractAddressItem := range toBeNotifiedMachineStateUpdateContractAddress {
		toBeNotifiedMachineStateUpdateContractAddressRule = append(toBeNotifiedMachineStateUpdateContractAddressRule, toBeNotifiedMachineStateUpdateContractAddressItem)
	}
	var toReportStakingStatusContractAddressRule []interface{}
	for _, toReportStakingStatusContractAddressItem := range toReportStakingStatusContractAddress {
		toReportStakingStatusContractAddressRule = append(toReportStakingStatusContractAddressRule, toReportStakingStatusContractAddressItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "ContractRegister", callerRule, toBeNotifiedMachineStateUpdateContractAddressRule, toReportStakingStatusContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &AireportContractRegisterIterator{contract: _Aireport.contract, event: "ContractRegister", logs: logs, sub: sub}, nil
}

// WatchContractRegister is a free log subscription operation binding the contract event 0x3bb0c9c4e86902fc67f92deb14cc8e8775454751bb8d7d5019d59daa06b9010d.
//
// Solidity: event ContractRegister(address indexed caller, string projectName, address indexed toBeNotifiedMachineStateUpdateContractAddress, address indexed toReportStakingStatusContractAddress)
func (_Aireport *AireportFilterer) WatchContractRegister(opts *bind.WatchOpts, sink chan<- *AireportContractRegister, caller []common.Address, toBeNotifiedMachineStateUpdateContractAddress []common.Address, toReportStakingStatusContractAddress []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var toBeNotifiedMachineStateUpdateContractAddressRule []interface{}
	for _, toBeNotifiedMachineStateUpdateContractAddressItem := range toBeNotifiedMachineStateUpdateContractAddress {
		toBeNotifiedMachineStateUpdateContractAddressRule = append(toBeNotifiedMachineStateUpdateContractAddressRule, toBeNotifiedMachineStateUpdateContractAddressItem)
	}
	var toReportStakingStatusContractAddressRule []interface{}
	for _, toReportStakingStatusContractAddressItem := range toReportStakingStatusContractAddress {
		toReportStakingStatusContractAddressRule = append(toReportStakingStatusContractAddressRule, toReportStakingStatusContractAddressItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "ContractRegister", callerRule, toBeNotifiedMachineStateUpdateContractAddressRule, toReportStakingStatusContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportContractRegister)
				if err := _Aireport.contract.UnpackLog(event, "ContractRegister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractRegister is a log parse operation binding the contract event 0x3bb0c9c4e86902fc67f92deb14cc8e8775454751bb8d7d5019d59daa06b9010d.
//
// Solidity: event ContractRegister(address indexed caller, string projectName, address indexed toBeNotifiedMachineStateUpdateContractAddress, address indexed toReportStakingStatusContractAddress)
func (_Aireport *AireportFilterer) ParseContractRegister(log types.Log) (*AireportContractRegister, error) {
	event := new(AireportContractRegister)
	if err := _Aireport.contract.UnpackLog(event, "ContractRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportDBCContractChangedIterator is returned from FilterDBCContractChanged and is used to iterate over the raw logs and unpacked data for DBCContractChanged events raised by the Aireport contract.
type AireportDBCContractChangedIterator struct {
	Event *AireportDBCContractChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportDBCContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportDBCContractChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportDBCContractChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportDBCContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportDBCContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportDBCContractChanged represents a DBCContractChanged event raised by the Aireport contract.
type AireportDBCContractChanged struct {
	DbcContractAddr common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDBCContractChanged is a free log retrieval operation binding the contract event 0x430b5699f8f42827478afbe097e5f2d9a9cf40af6fda93f449b533a9fa92fe4a.
//
// Solidity: event DBCContractChanged(address indexed dbcContractAddr)
func (_Aireport *AireportFilterer) FilterDBCContractChanged(opts *bind.FilterOpts, dbcContractAddr []common.Address) (*AireportDBCContractChangedIterator, error) {

	var dbcContractAddrRule []interface{}
	for _, dbcContractAddrItem := range dbcContractAddr {
		dbcContractAddrRule = append(dbcContractAddrRule, dbcContractAddrItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "DBCContractChanged", dbcContractAddrRule)
	if err != nil {
		return nil, err
	}
	return &AireportDBCContractChangedIterator{contract: _Aireport.contract, event: "DBCContractChanged", logs: logs, sub: sub}, nil
}

// WatchDBCContractChanged is a free log subscription operation binding the contract event 0x430b5699f8f42827478afbe097e5f2d9a9cf40af6fda93f449b533a9fa92fe4a.
//
// Solidity: event DBCContractChanged(address indexed dbcContractAddr)
func (_Aireport *AireportFilterer) WatchDBCContractChanged(opts *bind.WatchOpts, sink chan<- *AireportDBCContractChanged, dbcContractAddr []common.Address) (event.Subscription, error) {

	var dbcContractAddrRule []interface{}
	for _, dbcContractAddrItem := range dbcContractAddr {
		dbcContractAddrRule = append(dbcContractAddrRule, dbcContractAddrItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "DBCContractChanged", dbcContractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportDBCContractChanged)
				if err := _Aireport.contract.UnpackLog(event, "DBCContractChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDBCContractChanged is a log parse operation binding the contract event 0x430b5699f8f42827478afbe097e5f2d9a9cf40af6fda93f449b533a9fa92fe4a.
//
// Solidity: event DBCContractChanged(address indexed dbcContractAddr)
func (_Aireport *AireportFilterer) ParseDBCContractChanged(log types.Log) (*AireportDBCContractChanged, error) {
	event := new(AireportDBCContractChanged)
	if err := _Aireport.contract.UnpackLog(event, "DBCContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aireport contract.
type AireportInitializedIterator struct {
	Event *AireportInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportInitialized represents a Initialized event raised by the Aireport contract.
type AireportInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aireport *AireportFilterer) FilterInitialized(opts *bind.FilterOpts) (*AireportInitializedIterator, error) {

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AireportInitializedIterator{contract: _Aireport.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aireport *AireportFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AireportInitialized) (event.Subscription, error) {

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportInitialized)
				if err := _Aireport.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aireport *AireportFilterer) ParseInitialized(log types.Log) (*AireportInitialized, error) {
	event := new(AireportInitialized)
	if err := _Aireport.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportMachineInfoContractChangedIterator is returned from FilterMachineInfoContractChanged and is used to iterate over the raw logs and unpacked data for MachineInfoContractChanged events raised by the Aireport contract.
type AireportMachineInfoContractChangedIterator struct {
	Event *AireportMachineInfoContractChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportMachineInfoContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportMachineInfoContractChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportMachineInfoContractChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportMachineInfoContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportMachineInfoContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportMachineInfoContractChanged represents a MachineInfoContractChanged event raised by the Aireport contract.
type AireportMachineInfoContractChanged struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMachineInfoContractChanged is a free log retrieval operation binding the contract event 0x064453794ffce71f07f5cdf4581f62542de70cd615ccd48480dce390fe996d6c.
//
// Solidity: event MachineInfoContractChanged(address indexed addr)
func (_Aireport *AireportFilterer) FilterMachineInfoContractChanged(opts *bind.FilterOpts, addr []common.Address) (*AireportMachineInfoContractChangedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "MachineInfoContractChanged", addrRule)
	if err != nil {
		return nil, err
	}
	return &AireportMachineInfoContractChangedIterator{contract: _Aireport.contract, event: "MachineInfoContractChanged", logs: logs, sub: sub}, nil
}

// WatchMachineInfoContractChanged is a free log subscription operation binding the contract event 0x064453794ffce71f07f5cdf4581f62542de70cd615ccd48480dce390fe996d6c.
//
// Solidity: event MachineInfoContractChanged(address indexed addr)
func (_Aireport *AireportFilterer) WatchMachineInfoContractChanged(opts *bind.WatchOpts, sink chan<- *AireportMachineInfoContractChanged, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "MachineInfoContractChanged", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportMachineInfoContractChanged)
				if err := _Aireport.contract.UnpackLog(event, "MachineInfoContractChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMachineInfoContractChanged is a log parse operation binding the contract event 0x064453794ffce71f07f5cdf4581f62542de70cd615ccd48480dce390fe996d6c.
//
// Solidity: event MachineInfoContractChanged(address indexed addr)
func (_Aireport *AireportFilterer) ParseMachineInfoContractChanged(log types.Log) (*AireportMachineInfoContractChanged, error) {
	event := new(AireportMachineInfoContractChanged)
	if err := _Aireport.contract.UnpackLog(event, "MachineInfoContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportMachineStateUpdateIterator is returned from FilterMachineStateUpdate and is used to iterate over the raw logs and unpacked data for MachineStateUpdate events raised by the Aireport contract.
type AireportMachineStateUpdateIterator struct {
	Event *AireportMachineStateUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportMachineStateUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportMachineStateUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportMachineStateUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportMachineStateUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportMachineStateUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportMachineStateUpdate represents a MachineStateUpdate event raised by the Aireport contract.
type AireportMachineStateUpdate struct {
	MachineId   string
	ProjectName string
	StakingType uint8
	Tp          uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMachineStateUpdate is a free log retrieval operation binding the contract event 0x8ee03bdb4deed82a4bf4a9269f3bbbd591cc6a8b075b8b03a7cf8c19c5ebcdf2.
//
// Solidity: event MachineStateUpdate(string machineId, string projectName, uint8 stakingType, uint8 tp)
func (_Aireport *AireportFilterer) FilterMachineStateUpdate(opts *bind.FilterOpts) (*AireportMachineStateUpdateIterator, error) {

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "MachineStateUpdate")
	if err != nil {
		return nil, err
	}
	return &AireportMachineStateUpdateIterator{contract: _Aireport.contract, event: "MachineStateUpdate", logs: logs, sub: sub}, nil
}

// WatchMachineStateUpdate is a free log subscription operation binding the contract event 0x8ee03bdb4deed82a4bf4a9269f3bbbd591cc6a8b075b8b03a7cf8c19c5ebcdf2.
//
// Solidity: event MachineStateUpdate(string machineId, string projectName, uint8 stakingType, uint8 tp)
func (_Aireport *AireportFilterer) WatchMachineStateUpdate(opts *bind.WatchOpts, sink chan<- *AireportMachineStateUpdate) (event.Subscription, error) {

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "MachineStateUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportMachineStateUpdate)
				if err := _Aireport.contract.UnpackLog(event, "MachineStateUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMachineStateUpdate is a log parse operation binding the contract event 0x8ee03bdb4deed82a4bf4a9269f3bbbd591cc6a8b075b8b03a7cf8c19c5ebcdf2.
//
// Solidity: event MachineStateUpdate(string machineId, string projectName, uint8 stakingType, uint8 tp)
func (_Aireport *AireportFilterer) ParseMachineStateUpdate(log types.Log) (*AireportMachineStateUpdate, error) {
	event := new(AireportMachineStateUpdate)
	if err := _Aireport.contract.UnpackLog(event, "MachineStateUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportNotifiedTargetContractIterator is returned from FilterNotifiedTargetContract and is used to iterate over the raw logs and unpacked data for NotifiedTargetContract events raised by the Aireport contract.
type AireportNotifiedTargetContractIterator struct {
	Event *AireportNotifiedTargetContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportNotifiedTargetContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportNotifiedTargetContract)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportNotifiedTargetContract)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportNotifiedTargetContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportNotifiedTargetContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportNotifiedTargetContract represents a NotifiedTargetContract event raised by the Aireport contract.
type AireportNotifiedTargetContract struct {
	TargetContractAddress common.Address
	Tp                    uint8
	MachineId             string
	Result                bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNotifiedTargetContract is a free log retrieval operation binding the contract event 0x9b3cd1355dfd8f1f59422e3e7618dcc9d2e48e504d02a569310afe9a1b87f4da.
//
// Solidity: event NotifiedTargetContract(address indexed targetContractAddress, uint8 tp, string machineId, bool result)
func (_Aireport *AireportFilterer) FilterNotifiedTargetContract(opts *bind.FilterOpts, targetContractAddress []common.Address) (*AireportNotifiedTargetContractIterator, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "NotifiedTargetContract", targetContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &AireportNotifiedTargetContractIterator{contract: _Aireport.contract, event: "NotifiedTargetContract", logs: logs, sub: sub}, nil
}

// WatchNotifiedTargetContract is a free log subscription operation binding the contract event 0x9b3cd1355dfd8f1f59422e3e7618dcc9d2e48e504d02a569310afe9a1b87f4da.
//
// Solidity: event NotifiedTargetContract(address indexed targetContractAddress, uint8 tp, string machineId, bool result)
func (_Aireport *AireportFilterer) WatchNotifiedTargetContract(opts *bind.WatchOpts, sink chan<- *AireportNotifiedTargetContract, targetContractAddress []common.Address) (event.Subscription, error) {

	var targetContractAddressRule []interface{}
	for _, targetContractAddressItem := range targetContractAddress {
		targetContractAddressRule = append(targetContractAddressRule, targetContractAddressItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "NotifiedTargetContract", targetContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportNotifiedTargetContract)
				if err := _Aireport.contract.UnpackLog(event, "NotifiedTargetContract", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotifiedTargetContract is a log parse operation binding the contract event 0x9b3cd1355dfd8f1f59422e3e7618dcc9d2e48e504d02a569310afe9a1b87f4da.
//
// Solidity: event NotifiedTargetContract(address indexed targetContractAddress, uint8 tp, string machineId, bool result)
func (_Aireport *AireportFilterer) ParseNotifiedTargetContract(log types.Log) (*AireportNotifiedTargetContract, error) {
	event := new(AireportNotifiedTargetContract)
	if err := _Aireport.contract.UnpackLog(event, "NotifiedTargetContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aireport contract.
type AireportOwnershipTransferredIterator struct {
	Event *AireportOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportOwnershipTransferred represents a OwnershipTransferred event raised by the Aireport contract.
type AireportOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aireport *AireportFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AireportOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AireportOwnershipTransferredIterator{contract: _Aireport.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aireport *AireportFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AireportOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportOwnershipTransferred)
				if err := _Aireport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aireport *AireportFilterer) ParseOwnershipTransferred(log types.Log) (*AireportOwnershipTransferred, error) {
	event := new(AireportOwnershipTransferred)
	if err := _Aireport.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportRemoveAuthorizedReporterIterator is returned from FilterRemoveAuthorizedReporter and is used to iterate over the raw logs and unpacked data for RemoveAuthorizedReporter events raised by the Aireport contract.
type AireportRemoveAuthorizedReporterIterator struct {
	Event *AireportRemoveAuthorizedReporter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportRemoveAuthorizedReporterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportRemoveAuthorizedReporter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportRemoveAuthorizedReporter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportRemoveAuthorizedReporterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportRemoveAuthorizedReporterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportRemoveAuthorizedReporter represents a RemoveAuthorizedReporter event raised by the Aireport contract.
type AireportRemoveAuthorizedReporter struct {
	Reporter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuthorizedReporter is a free log retrieval operation binding the contract event 0x9746786bebd40aafb5cdd5e71c600accffc705724efc342a313d7c11c8751457.
//
// Solidity: event RemoveAuthorizedReporter(address indexed reporter)
func (_Aireport *AireportFilterer) FilterRemoveAuthorizedReporter(opts *bind.FilterOpts, reporter []common.Address) (*AireportRemoveAuthorizedReporterIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "RemoveAuthorizedReporter", reporterRule)
	if err != nil {
		return nil, err
	}
	return &AireportRemoveAuthorizedReporterIterator{contract: _Aireport.contract, event: "RemoveAuthorizedReporter", logs: logs, sub: sub}, nil
}

// WatchRemoveAuthorizedReporter is a free log subscription operation binding the contract event 0x9746786bebd40aafb5cdd5e71c600accffc705724efc342a313d7c11c8751457.
//
// Solidity: event RemoveAuthorizedReporter(address indexed reporter)
func (_Aireport *AireportFilterer) WatchRemoveAuthorizedReporter(opts *bind.WatchOpts, sink chan<- *AireportRemoveAuthorizedReporter, reporter []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "RemoveAuthorizedReporter", reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportRemoveAuthorizedReporter)
				if err := _Aireport.contract.UnpackLog(event, "RemoveAuthorizedReporter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveAuthorizedReporter is a log parse operation binding the contract event 0x9746786bebd40aafb5cdd5e71c600accffc705724efc342a313d7c11c8751457.
//
// Solidity: event RemoveAuthorizedReporter(address indexed reporter)
func (_Aireport *AireportFilterer) ParseRemoveAuthorizedReporter(log types.Log) (*AireportRemoveAuthorizedReporter, error) {
	event := new(AireportRemoveAuthorizedReporter)
	if err := _Aireport.contract.UnpackLog(event, "RemoveAuthorizedReporter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportReportFailedIterator is returned from FilterReportFailed and is used to iterate over the raw logs and unpacked data for ReportFailed events raised by the Aireport contract.
type AireportReportFailedIterator struct {
	Event *AireportReportFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportReportFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportReportFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportReportFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportReportFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportReportFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportReportFailed represents a ReportFailed event raised by the Aireport contract.
type AireportReportFailed struct {
	Tp          uint8
	ProjectName string
	StakingType uint8
	MachineId   string
	Reason      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReportFailed is a free log retrieval operation binding the contract event 0xb9431bf2a1a48f37e3f5570e5c0d6919111de411f7b5d9d335865416338cd5ff.
//
// Solidity: event ReportFailed(uint8 tp, string projectName, uint8 stakingType, string machineId, string reason)
func (_Aireport *AireportFilterer) FilterReportFailed(opts *bind.FilterOpts) (*AireportReportFailedIterator, error) {

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "ReportFailed")
	if err != nil {
		return nil, err
	}
	return &AireportReportFailedIterator{contract: _Aireport.contract, event: "ReportFailed", logs: logs, sub: sub}, nil
}

// WatchReportFailed is a free log subscription operation binding the contract event 0xb9431bf2a1a48f37e3f5570e5c0d6919111de411f7b5d9d335865416338cd5ff.
//
// Solidity: event ReportFailed(uint8 tp, string projectName, uint8 stakingType, string machineId, string reason)
func (_Aireport *AireportFilterer) WatchReportFailed(opts *bind.WatchOpts, sink chan<- *AireportReportFailed) (event.Subscription, error) {

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "ReportFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportReportFailed)
				if err := _Aireport.contract.UnpackLog(event, "ReportFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportFailed is a log parse operation binding the contract event 0xb9431bf2a1a48f37e3f5570e5c0d6919111de411f7b5d9d335865416338cd5ff.
//
// Solidity: event ReportFailed(uint8 tp, string projectName, uint8 stakingType, string machineId, string reason)
func (_Aireport *AireportFilterer) ParseReportFailed(log types.Log) (*AireportReportFailed, error) {
	event := new(AireportReportFailed)
	if err := _Aireport.contract.UnpackLog(event, "ReportFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Aireport contract.
type AireportUpgradedIterator struct {
	Event *AireportUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportUpgraded represents a Upgraded event raised by the Aireport contract.
type AireportUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Aireport *AireportFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AireportUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AireportUpgradedIterator{contract: _Aireport.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Aireport *AireportFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AireportUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportUpgraded)
				if err := _Aireport.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Aireport *AireportFilterer) ParseUpgraded(log types.Log) (*AireportUpgraded, error) {
	event := new(AireportUpgraded)
	if err := _Aireport.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AireportReportedStakingStatusIterator is returned from FilterReportedStakingStatus and is used to iterate over the raw logs and unpacked data for ReportedStakingStatus events raised by the Aireport contract.
type AireportReportedStakingStatusIterator struct {
	Event *AireportReportedStakingStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AireportReportedStakingStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AireportReportedStakingStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AireportReportedStakingStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AireportReportedStakingStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AireportReportedStakingStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AireportReportedStakingStatus represents a ReportedStakingStatus event raised by the Aireport contract.
type AireportReportedStakingStatus struct {
	ProjectName string
	Tp          uint8
	MachineId   string
	GpuNum      *big.Int
	IsStake     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReportedStakingStatus is a free log retrieval operation binding the contract event 0xe04797838f35c6500e42eb74d5a345f7b66ffe2d60a1807521e96031c6730502.
//
// Solidity: event reportedStakingStatus(string projectName, uint8 tp, string machineId, uint256 gpuNum, bool isStake)
func (_Aireport *AireportFilterer) FilterReportedStakingStatus(opts *bind.FilterOpts) (*AireportReportedStakingStatusIterator, error) {

	logs, sub, err := _Aireport.contract.FilterLogs(opts, "reportedStakingStatus")
	if err != nil {
		return nil, err
	}
	return &AireportReportedStakingStatusIterator{contract: _Aireport.contract, event: "reportedStakingStatus", logs: logs, sub: sub}, nil
}

// WatchReportedStakingStatus is a free log subscription operation binding the contract event 0xe04797838f35c6500e42eb74d5a345f7b66ffe2d60a1807521e96031c6730502.
//
// Solidity: event reportedStakingStatus(string projectName, uint8 tp, string machineId, uint256 gpuNum, bool isStake)
func (_Aireport *AireportFilterer) WatchReportedStakingStatus(opts *bind.WatchOpts, sink chan<- *AireportReportedStakingStatus) (event.Subscription, error) {

	logs, sub, err := _Aireport.contract.WatchLogs(opts, "reportedStakingStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AireportReportedStakingStatus)
				if err := _Aireport.contract.UnpackLog(event, "reportedStakingStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportedStakingStatus is a log parse operation binding the contract event 0xe04797838f35c6500e42eb74d5a345f7b66ffe2d60a1807521e96031c6730502.
//
// Solidity: event reportedStakingStatus(string projectName, uint8 tp, string machineId, uint256 gpuNum, bool isStake)
func (_Aireport *AireportFilterer) ParseReportedStakingStatus(log types.Log) (*AireportReportedStakingStatus, error) {
	event := new(AireportReportedStakingStatus)
	if err := _Aireport.contract.UnpackLog(event, "reportedStakingStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
