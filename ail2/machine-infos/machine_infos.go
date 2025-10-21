// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machineinfos

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

// MachineInfosBandWidthMintInfo is an auto generated low-level Go binding around an user-defined struct.
type MachineInfosBandWidthMintInfo struct {
	MachineOwner common.Address
	MachineId    string
	CpuCores     *big.Int
	MachineMem   *big.Int
	Region       string
	Hdd          *big.Int
	Bandwidth    *big.Int
}

// MachineInfosMachineInfo is an auto generated low-level Go binding around an user-defined struct.
type MachineInfosMachineInfo struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
	Region       string
	Model        string
}

// MachineinfosMetaData contains all meta data concerning the Machineinfos contract.
var MachineinfosMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFY_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isDeepLink\",\"type\":\"bool\"}],\"name\":\"getMachineInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"calcPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"longitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"latitude\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"machineMem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isDeepLink\",\"type\":\"bool\"}],\"name\":\"getMachineInfoTotal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"calcPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"longitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"latitude\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"machineMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"}],\"internalType\":\"structMachineInfos.MachineInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"}],\"name\":\"getMachineRegion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"machineBandWidthInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cpuCores\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machineMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hdd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bandwidth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"machineInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"calcPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"longitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"latitude\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"machineMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cpuCores\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"machineMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"hdd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bandwidth\",\"type\":\"uint256\"}],\"internalType\":\"structMachineInfos.BandWidthMintInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"setBandWidthInfos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_Id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"machineOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"calcPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cpuRate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"cpuType\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"gpuCount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"machineId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"longitude\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"latitude\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"machineMem\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"region\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"model\",\"type\":\"string\"}],\"internalType\":\"structMachineInfos.MachineInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"setMachineInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ope\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_bool\",\"type\":\"bool\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MachineinfosABI is the input ABI used to generate the binding from.
// Deprecated: Use MachineinfosMetaData.ABI instead.
var MachineinfosABI = MachineinfosMetaData.ABI

// Machineinfos is an auto generated Go binding around an Ethereum contract.
type Machineinfos struct {
	MachineinfosCaller     // Read-only binding to the contract
	MachineinfosTransactor // Write-only binding to the contract
	MachineinfosFilterer   // Log filterer for contract events
}

// MachineinfosCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineinfosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineinfosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineinfosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineinfosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineinfosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineinfosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineinfosSession struct {
	Contract     *Machineinfos     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineinfosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineinfosCallerSession struct {
	Contract *MachineinfosCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MachineinfosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineinfosTransactorSession struct {
	Contract     *MachineinfosTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MachineinfosRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineinfosRaw struct {
	Contract *Machineinfos // Generic contract binding to access the raw methods on
}

// MachineinfosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineinfosCallerRaw struct {
	Contract *MachineinfosCaller // Generic read-only contract binding to access the raw methods on
}

// MachineinfosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineinfosTransactorRaw struct {
	Contract *MachineinfosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachineinfos creates a new instance of Machineinfos, bound to a specific deployed contract.
func NewMachineinfos(address common.Address, backend bind.ContractBackend) (*Machineinfos, error) {
	contract, err := bindMachineinfos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Machineinfos{MachineinfosCaller: MachineinfosCaller{contract: contract}, MachineinfosTransactor: MachineinfosTransactor{contract: contract}, MachineinfosFilterer: MachineinfosFilterer{contract: contract}}, nil
}

// NewMachineinfosCaller creates a new read-only instance of Machineinfos, bound to a specific deployed contract.
func NewMachineinfosCaller(address common.Address, caller bind.ContractCaller) (*MachineinfosCaller, error) {
	contract, err := bindMachineinfos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineinfosCaller{contract: contract}, nil
}

// NewMachineinfosTransactor creates a new write-only instance of Machineinfos, bound to a specific deployed contract.
func NewMachineinfosTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineinfosTransactor, error) {
	contract, err := bindMachineinfos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineinfosTransactor{contract: contract}, nil
}

// NewMachineinfosFilterer creates a new log filterer instance of Machineinfos, bound to a specific deployed contract.
func NewMachineinfosFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineinfosFilterer, error) {
	contract, err := bindMachineinfos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineinfosFilterer{contract: contract}, nil
}

// bindMachineinfos binds a generic wrapper to an already deployed contract.
func bindMachineinfos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MachineinfosMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machineinfos *MachineinfosRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Machineinfos.Contract.MachineinfosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machineinfos *MachineinfosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machineinfos.Contract.MachineinfosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machineinfos *MachineinfosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machineinfos.Contract.MachineinfosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Machineinfos *MachineinfosCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Machineinfos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Machineinfos *MachineinfosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machineinfos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Machineinfos *MachineinfosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Machineinfos.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Machineinfos *MachineinfosCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Machineinfos *MachineinfosSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Machineinfos.Contract.DOMAINTYPEHASH(&_Machineinfos.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Machineinfos *MachineinfosCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Machineinfos.Contract.DOMAINTYPEHASH(&_Machineinfos.CallOpts)
}

// VERIFYTYPEHASH is a free data retrieval call binding the contract method 0xcb6535b5.
//
// Solidity: function VERIFY_TYPEHASH() view returns(bytes32)
func (_Machineinfos *MachineinfosCaller) VERIFYTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "VERIFY_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFYTYPEHASH is a free data retrieval call binding the contract method 0xcb6535b5.
//
// Solidity: function VERIFY_TYPEHASH() view returns(bytes32)
func (_Machineinfos *MachineinfosSession) VERIFYTYPEHASH() ([32]byte, error) {
	return _Machineinfos.Contract.VERIFYTYPEHASH(&_Machineinfos.CallOpts)
}

// VERIFYTYPEHASH is a free data retrieval call binding the contract method 0xcb6535b5.
//
// Solidity: function VERIFY_TYPEHASH() view returns(bytes32)
func (_Machineinfos *MachineinfosCallerSession) VERIFYTYPEHASH() ([32]byte, error) {
	return _Machineinfos.Contract.VERIFYTYPEHASH(&_Machineinfos.CallOpts)
}

// GetMachineInfo is a free data retrieval call binding the contract method 0xc8ffe871.
//
// Solidity: function getMachineInfo(string _id, bool _isDeepLink) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, string longitude, string latitude, uint256 machineMem)
func (_Machineinfos *MachineinfosCaller) GetMachineInfo(opts *bind.CallOpts, _id string, _isDeepLink bool) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
}, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "getMachineInfo", _id, _isDeepLink)

	outstruct := new(struct {
		MachineOwner common.Address
		CalcPoint    *big.Int
		CpuRate      *big.Int
		GpuType      string
		GpuMem       *big.Int
		CpuType      string
		GpuCount     *big.Int
		MachineId    string
		Longitude    string
		Latitude     string
		MachineMem   *big.Int
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
	outstruct.Longitude = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.Latitude = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.MachineMem = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMachineInfo is a free data retrieval call binding the contract method 0xc8ffe871.
//
// Solidity: function getMachineInfo(string _id, bool _isDeepLink) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, string longitude, string latitude, uint256 machineMem)
func (_Machineinfos *MachineinfosSession) GetMachineInfo(_id string, _isDeepLink bool) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
}, error) {
	return _Machineinfos.Contract.GetMachineInfo(&_Machineinfos.CallOpts, _id, _isDeepLink)
}

// GetMachineInfo is a free data retrieval call binding the contract method 0xc8ffe871.
//
// Solidity: function getMachineInfo(string _id, bool _isDeepLink) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, string longitude, string latitude, uint256 machineMem)
func (_Machineinfos *MachineinfosCallerSession) GetMachineInfo(_id string, _isDeepLink bool) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
}, error) {
	return _Machineinfos.Contract.GetMachineInfo(&_Machineinfos.CallOpts, _id, _isDeepLink)
}

// GetMachineInfoTotal is a free data retrieval call binding the contract method 0xbde4e409.
//
// Solidity: function getMachineInfoTotal(string _id, bool _isDeepLink) view returns((address,uint256,uint256,string,uint256,string,uint256,string,string,string,uint256,string,string))
func (_Machineinfos *MachineinfosCaller) GetMachineInfoTotal(opts *bind.CallOpts, _id string, _isDeepLink bool) (MachineInfosMachineInfo, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "getMachineInfoTotal", _id, _isDeepLink)

	if err != nil {
		return *new(MachineInfosMachineInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(MachineInfosMachineInfo)).(*MachineInfosMachineInfo)

	return out0, err

}

// GetMachineInfoTotal is a free data retrieval call binding the contract method 0xbde4e409.
//
// Solidity: function getMachineInfoTotal(string _id, bool _isDeepLink) view returns((address,uint256,uint256,string,uint256,string,uint256,string,string,string,uint256,string,string))
func (_Machineinfos *MachineinfosSession) GetMachineInfoTotal(_id string, _isDeepLink bool) (MachineInfosMachineInfo, error) {
	return _Machineinfos.Contract.GetMachineInfoTotal(&_Machineinfos.CallOpts, _id, _isDeepLink)
}

// GetMachineInfoTotal is a free data retrieval call binding the contract method 0xbde4e409.
//
// Solidity: function getMachineInfoTotal(string _id, bool _isDeepLink) view returns((address,uint256,uint256,string,uint256,string,uint256,string,string,string,uint256,string,string))
func (_Machineinfos *MachineinfosCallerSession) GetMachineInfoTotal(_id string, _isDeepLink bool) (MachineInfosMachineInfo, error) {
	return _Machineinfos.Contract.GetMachineInfoTotal(&_Machineinfos.CallOpts, _id, _isDeepLink)
}

// GetMachineRegion is a free data retrieval call binding the contract method 0x1b97cc77.
//
// Solidity: function getMachineRegion(string _id) view returns(string)
func (_Machineinfos *MachineinfosCaller) GetMachineRegion(opts *bind.CallOpts, _id string) (string, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "getMachineRegion", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMachineRegion is a free data retrieval call binding the contract method 0x1b97cc77.
//
// Solidity: function getMachineRegion(string _id) view returns(string)
func (_Machineinfos *MachineinfosSession) GetMachineRegion(_id string) (string, error) {
	return _Machineinfos.Contract.GetMachineRegion(&_Machineinfos.CallOpts, _id)
}

// GetMachineRegion is a free data retrieval call binding the contract method 0x1b97cc77.
//
// Solidity: function getMachineRegion(string _id) view returns(string)
func (_Machineinfos *MachineinfosCallerSession) GetMachineRegion(_id string) (string, error) {
	return _Machineinfos.Contract.GetMachineRegion(&_Machineinfos.CallOpts, _id)
}

// MachineBandWidthInfos is a free data retrieval call binding the contract method 0x042f2acf.
//
// Solidity: function machineBandWidthInfos(string ) view returns(address machineOwner, string machineId, uint256 cpuCores, uint256 machineMem, string region, uint256 hdd, uint256 bandwidth)
func (_Machineinfos *MachineinfosCaller) MachineBandWidthInfos(opts *bind.CallOpts, arg0 string) (struct {
	MachineOwner common.Address
	MachineId    string
	CpuCores     *big.Int
	MachineMem   *big.Int
	Region       string
	Hdd          *big.Int
	Bandwidth    *big.Int
}, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "machineBandWidthInfos", arg0)

	outstruct := new(struct {
		MachineOwner common.Address
		MachineId    string
		CpuCores     *big.Int
		MachineMem   *big.Int
		Region       string
		Hdd          *big.Int
		Bandwidth    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MachineOwner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MachineId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.CpuCores = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MachineMem = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Region = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Hdd = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Bandwidth = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MachineBandWidthInfos is a free data retrieval call binding the contract method 0x042f2acf.
//
// Solidity: function machineBandWidthInfos(string ) view returns(address machineOwner, string machineId, uint256 cpuCores, uint256 machineMem, string region, uint256 hdd, uint256 bandwidth)
func (_Machineinfos *MachineinfosSession) MachineBandWidthInfos(arg0 string) (struct {
	MachineOwner common.Address
	MachineId    string
	CpuCores     *big.Int
	MachineMem   *big.Int
	Region       string
	Hdd          *big.Int
	Bandwidth    *big.Int
}, error) {
	return _Machineinfos.Contract.MachineBandWidthInfos(&_Machineinfos.CallOpts, arg0)
}

// MachineBandWidthInfos is a free data retrieval call binding the contract method 0x042f2acf.
//
// Solidity: function machineBandWidthInfos(string ) view returns(address machineOwner, string machineId, uint256 cpuCores, uint256 machineMem, string region, uint256 hdd, uint256 bandwidth)
func (_Machineinfos *MachineinfosCallerSession) MachineBandWidthInfos(arg0 string) (struct {
	MachineOwner common.Address
	MachineId    string
	CpuCores     *big.Int
	MachineMem   *big.Int
	Region       string
	Hdd          *big.Int
	Bandwidth    *big.Int
}, error) {
	return _Machineinfos.Contract.MachineBandWidthInfos(&_Machineinfos.CallOpts, arg0)
}

// MachineInfos is a free data retrieval call binding the contract method 0xacbab99e.
//
// Solidity: function machineInfos(string ) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, string longitude, string latitude, uint256 machineMem, string region, string model)
func (_Machineinfos *MachineinfosCaller) MachineInfos(opts *bind.CallOpts, arg0 string) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
	Region       string
	Model        string
}, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "machineInfos", arg0)

	outstruct := new(struct {
		MachineOwner common.Address
		CalcPoint    *big.Int
		CpuRate      *big.Int
		GpuType      string
		GpuMem       *big.Int
		CpuType      string
		GpuCount     *big.Int
		MachineId    string
		Longitude    string
		Latitude     string
		MachineMem   *big.Int
		Region       string
		Model        string
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
	outstruct.Longitude = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.Latitude = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.MachineMem = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Region = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.Model = *abi.ConvertType(out[12], new(string)).(*string)

	return *outstruct, err

}

// MachineInfos is a free data retrieval call binding the contract method 0xacbab99e.
//
// Solidity: function machineInfos(string ) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, string longitude, string latitude, uint256 machineMem, string region, string model)
func (_Machineinfos *MachineinfosSession) MachineInfos(arg0 string) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
	Region       string
	Model        string
}, error) {
	return _Machineinfos.Contract.MachineInfos(&_Machineinfos.CallOpts, arg0)
}

// MachineInfos is a free data retrieval call binding the contract method 0xacbab99e.
//
// Solidity: function machineInfos(string ) view returns(address machineOwner, uint256 calcPoint, uint256 cpuRate, string gpuType, uint256 gpuMem, string cpuType, uint256 gpuCount, string machineId, string longitude, string latitude, uint256 machineMem, string region, string model)
func (_Machineinfos *MachineinfosCallerSession) MachineInfos(arg0 string) (struct {
	MachineOwner common.Address
	CalcPoint    *big.Int
	CpuRate      *big.Int
	GpuType      string
	GpuMem       *big.Int
	CpuType      string
	GpuCount     *big.Int
	MachineId    string
	Longitude    string
	Latitude     string
	MachineMem   *big.Int
	Region       string
	Model        string
}, error) {
	return _Machineinfos.Contract.MachineInfos(&_Machineinfos.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Machineinfos *MachineinfosCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Machineinfos *MachineinfosSession) Name() (string, error) {
	return _Machineinfos.Contract.Name(&_Machineinfos.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Machineinfos *MachineinfosCallerSession) Name() (string, error) {
	return _Machineinfos.Contract.Name(&_Machineinfos.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Machineinfos *MachineinfosCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Machineinfos *MachineinfosSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Machineinfos.Contract.Nonces(&_Machineinfos.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Machineinfos *MachineinfosCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Machineinfos.Contract.Nonces(&_Machineinfos.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Machineinfos *MachineinfosCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Machineinfos *MachineinfosSession) Operators(arg0 common.Address) (bool, error) {
	return _Machineinfos.Contract.Operators(&_Machineinfos.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Machineinfos *MachineinfosCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Machineinfos.Contract.Operators(&_Machineinfos.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Machineinfos *MachineinfosCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Machineinfos *MachineinfosSession) Owner() (common.Address, error) {
	return _Machineinfos.Contract.Owner(&_Machineinfos.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Machineinfos *MachineinfosCallerSession) Owner() (common.Address, error) {
	return _Machineinfos.Contract.Owner(&_Machineinfos.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Machineinfos *MachineinfosCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Machineinfos.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Machineinfos *MachineinfosSession) Version() (string, error) {
	return _Machineinfos.Contract.Version(&_Machineinfos.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Machineinfos *MachineinfosCallerSession) Version() (string, error) {
	return _Machineinfos.Contract.Version(&_Machineinfos.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Machineinfos *MachineinfosTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Machineinfos *MachineinfosSession) Initialize() (*types.Transaction, error) {
	return _Machineinfos.Contract.Initialize(&_Machineinfos.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Machineinfos *MachineinfosTransactorSession) Initialize() (*types.Transaction, error) {
	return _Machineinfos.Contract.Initialize(&_Machineinfos.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Machineinfos *MachineinfosTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Machineinfos *MachineinfosSession) RenounceOwnership() (*types.Transaction, error) {
	return _Machineinfos.Contract.RenounceOwnership(&_Machineinfos.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Machineinfos *MachineinfosTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Machineinfos.Contract.RenounceOwnership(&_Machineinfos.TransactOpts)
}

// SetBandWidthInfos is a paid mutator transaction binding the contract method 0xca0070f8.
//
// Solidity: function setBandWidthInfos(string _id, (address,string,uint256,uint256,string,uint256,uint256) _info) returns()
func (_Machineinfos *MachineinfosTransactor) SetBandWidthInfos(opts *bind.TransactOpts, _id string, _info MachineInfosBandWidthMintInfo) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "setBandWidthInfos", _id, _info)
}

// SetBandWidthInfos is a paid mutator transaction binding the contract method 0xca0070f8.
//
// Solidity: function setBandWidthInfos(string _id, (address,string,uint256,uint256,string,uint256,uint256) _info) returns()
func (_Machineinfos *MachineinfosSession) SetBandWidthInfos(_id string, _info MachineInfosBandWidthMintInfo) (*types.Transaction, error) {
	return _Machineinfos.Contract.SetBandWidthInfos(&_Machineinfos.TransactOpts, _id, _info)
}

// SetBandWidthInfos is a paid mutator transaction binding the contract method 0xca0070f8.
//
// Solidity: function setBandWidthInfos(string _id, (address,string,uint256,uint256,string,uint256,uint256) _info) returns()
func (_Machineinfos *MachineinfosTransactorSession) SetBandWidthInfos(_id string, _info MachineInfosBandWidthMintInfo) (*types.Transaction, error) {
	return _Machineinfos.Contract.SetBandWidthInfos(&_Machineinfos.TransactOpts, _id, _info)
}

// SetMachineInfo is a paid mutator transaction binding the contract method 0xfc17c6f0.
//
// Solidity: function setMachineInfo(string _Id, (address,uint256,uint256,string,uint256,string,uint256,string,string,string,uint256,string,string) _info) returns()
func (_Machineinfos *MachineinfosTransactor) SetMachineInfo(opts *bind.TransactOpts, _Id string, _info MachineInfosMachineInfo) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "setMachineInfo", _Id, _info)
}

// SetMachineInfo is a paid mutator transaction binding the contract method 0xfc17c6f0.
//
// Solidity: function setMachineInfo(string _Id, (address,uint256,uint256,string,uint256,string,uint256,string,string,string,uint256,string,string) _info) returns()
func (_Machineinfos *MachineinfosSession) SetMachineInfo(_Id string, _info MachineInfosMachineInfo) (*types.Transaction, error) {
	return _Machineinfos.Contract.SetMachineInfo(&_Machineinfos.TransactOpts, _Id, _info)
}

// SetMachineInfo is a paid mutator transaction binding the contract method 0xfc17c6f0.
//
// Solidity: function setMachineInfo(string _Id, (address,uint256,uint256,string,uint256,string,uint256,string,string,string,uint256,string,string) _info) returns()
func (_Machineinfos *MachineinfosTransactorSession) SetMachineInfo(_Id string, _info MachineInfosMachineInfo) (*types.Transaction, error) {
	return _Machineinfos.Contract.SetMachineInfo(&_Machineinfos.TransactOpts, _Id, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Machineinfos *MachineinfosTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Machineinfos *MachineinfosSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Machineinfos.Contract.TransferOwnership(&_Machineinfos.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Machineinfos *MachineinfosTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Machineinfos.Contract.TransferOwnership(&_Machineinfos.TransactOpts, newOwner)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0x6d44a3b2.
//
// Solidity: function updateOperator(address _ope, bool _bool) returns()
func (_Machineinfos *MachineinfosTransactor) UpdateOperator(opts *bind.TransactOpts, _ope common.Address, _bool bool) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "updateOperator", _ope, _bool)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0x6d44a3b2.
//
// Solidity: function updateOperator(address _ope, bool _bool) returns()
func (_Machineinfos *MachineinfosSession) UpdateOperator(_ope common.Address, _bool bool) (*types.Transaction, error) {
	return _Machineinfos.Contract.UpdateOperator(&_Machineinfos.TransactOpts, _ope, _bool)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0x6d44a3b2.
//
// Solidity: function updateOperator(address _ope, bool _bool) returns()
func (_Machineinfos *MachineinfosTransactorSession) UpdateOperator(_ope common.Address, _bool bool) (*types.Transaction, error) {
	return _Machineinfos.Contract.UpdateOperator(&_Machineinfos.TransactOpts, _ope, _bool)
}

// Verify is a paid mutator transaction binding the contract method 0x5eec311a.
//
// Solidity: function verify(address _signer, uint8 v, bytes32 r, bytes32 s, uint256 deadline) returns()
func (_Machineinfos *MachineinfosTransactor) Verify(opts *bind.TransactOpts, _signer common.Address, v uint8, r [32]byte, s [32]byte, deadline *big.Int) (*types.Transaction, error) {
	return _Machineinfos.contract.Transact(opts, "verify", _signer, v, r, s, deadline)
}

// Verify is a paid mutator transaction binding the contract method 0x5eec311a.
//
// Solidity: function verify(address _signer, uint8 v, bytes32 r, bytes32 s, uint256 deadline) returns()
func (_Machineinfos *MachineinfosSession) Verify(_signer common.Address, v uint8, r [32]byte, s [32]byte, deadline *big.Int) (*types.Transaction, error) {
	return _Machineinfos.Contract.Verify(&_Machineinfos.TransactOpts, _signer, v, r, s, deadline)
}

// Verify is a paid mutator transaction binding the contract method 0x5eec311a.
//
// Solidity: function verify(address _signer, uint8 v, bytes32 r, bytes32 s, uint256 deadline) returns()
func (_Machineinfos *MachineinfosTransactorSession) Verify(_signer common.Address, v uint8, r [32]byte, s [32]byte, deadline *big.Int) (*types.Transaction, error) {
	return _Machineinfos.Contract.Verify(&_Machineinfos.TransactOpts, _signer, v, r, s, deadline)
}

// MachineinfosInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Machineinfos contract.
type MachineinfosInitializedIterator struct {
	Event *MachineinfosInitialized // Event containing the contract specifics and raw log

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
func (it *MachineinfosInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineinfosInitialized)
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
		it.Event = new(MachineinfosInitialized)
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
func (it *MachineinfosInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineinfosInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineinfosInitialized represents a Initialized event raised by the Machineinfos contract.
type MachineinfosInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Machineinfos *MachineinfosFilterer) FilterInitialized(opts *bind.FilterOpts) (*MachineinfosInitializedIterator, error) {

	logs, sub, err := _Machineinfos.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MachineinfosInitializedIterator{contract: _Machineinfos.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Machineinfos *MachineinfosFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MachineinfosInitialized) (event.Subscription, error) {

	logs, sub, err := _Machineinfos.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineinfosInitialized)
				if err := _Machineinfos.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Machineinfos *MachineinfosFilterer) ParseInitialized(log types.Log) (*MachineinfosInitialized, error) {
	event := new(MachineinfosInitialized)
	if err := _Machineinfos.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineinfosOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Machineinfos contract.
type MachineinfosOwnershipTransferredIterator struct {
	Event *MachineinfosOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MachineinfosOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineinfosOwnershipTransferred)
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
		it.Event = new(MachineinfosOwnershipTransferred)
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
func (it *MachineinfosOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineinfosOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineinfosOwnershipTransferred represents a OwnershipTransferred event raised by the Machineinfos contract.
type MachineinfosOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Machineinfos *MachineinfosFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MachineinfosOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Machineinfos.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MachineinfosOwnershipTransferredIterator{contract: _Machineinfos.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Machineinfos *MachineinfosFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MachineinfosOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Machineinfos.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineinfosOwnershipTransferred)
				if err := _Machineinfos.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Machineinfos *MachineinfosFilterer) ParseOwnershipTransferred(log types.Log) (*MachineinfosOwnershipTransferred, error) {
	event := new(MachineinfosOwnershipTransferred)
	if err := _Machineinfos.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
