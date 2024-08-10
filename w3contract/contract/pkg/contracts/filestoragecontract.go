// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// FileStorageFile is an auto generated low-level Go binding around an user-defined struct.
type FileStorageFile struct {
	Name       string
	Size       *big.Int
	Hash       [32]byte
	UploadTime *big.Int
	Uploader   common.Address
	Exists     bool
}

// FileStorageContractMetaData contains all meta data concerning the FileStorageContract contract.
var FileStorageContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyDeniedAdminAccess\",\"inputs\":[]},{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllUserFiles\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structFileStorage.File[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"uploadTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uploader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserFiles\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structFileStorage.File[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"uploadTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uploader\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyFile\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_newName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_newSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"uploadFile\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FileDeleted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fileHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FileModified\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fileHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FileUploaded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"fileHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// FileStorageContractABI is the input ABI used to generate the binding from.
// Deprecated: Use FileStorageContractMetaData.ABI instead.
var FileStorageContractABI = FileStorageContractMetaData.ABI

// FileStorageContract is an auto generated Go binding around an Ethereum contract.
type FileStorageContract struct {
	FileStorageContractCaller     // Read-only binding to the contract
	FileStorageContractTransactor // Write-only binding to the contract
	FileStorageContractFilterer   // Log filterer for contract events
}

// FileStorageContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileStorageContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStorageContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileStorageContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStorageContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileStorageContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStorageContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileStorageContractSession struct {
	Contract     *FileStorageContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FileStorageContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileStorageContractCallerSession struct {
	Contract *FileStorageContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FileStorageContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileStorageContractTransactorSession struct {
	Contract     *FileStorageContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FileStorageContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileStorageContractRaw struct {
	Contract *FileStorageContract // Generic contract binding to access the raw methods on
}

// FileStorageContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileStorageContractCallerRaw struct {
	Contract *FileStorageContractCaller // Generic read-only contract binding to access the raw methods on
}

// FileStorageContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileStorageContractTransactorRaw struct {
	Contract *FileStorageContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileStorageContract creates a new instance of FileStorageContract, bound to a specific deployed contract.
func NewFileStorageContract(address common.Address, backend bind.ContractBackend) (*FileStorageContract, error) {
	contract, err := bindFileStorageContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileStorageContract{FileStorageContractCaller: FileStorageContractCaller{contract: contract}, FileStorageContractTransactor: FileStorageContractTransactor{contract: contract}, FileStorageContractFilterer: FileStorageContractFilterer{contract: contract}}, nil
}

// NewFileStorageContractCaller creates a new read-only instance of FileStorageContract, bound to a specific deployed contract.
func NewFileStorageContractCaller(address common.Address, caller bind.ContractCaller) (*FileStorageContractCaller, error) {
	contract, err := bindFileStorageContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractCaller{contract: contract}, nil
}

// NewFileStorageContractTransactor creates a new write-only instance of FileStorageContract, bound to a specific deployed contract.
func NewFileStorageContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FileStorageContractTransactor, error) {
	contract, err := bindFileStorageContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractTransactor{contract: contract}, nil
}

// NewFileStorageContractFilterer creates a new log filterer instance of FileStorageContract, bound to a specific deployed contract.
func NewFileStorageContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FileStorageContractFilterer, error) {
	contract, err := bindFileStorageContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractFilterer{contract: contract}, nil
}

// bindFileStorageContract binds a generic wrapper to an already deployed contract.
func bindFileStorageContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FileStorageContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileStorageContract *FileStorageContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileStorageContract.Contract.FileStorageContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileStorageContract *FileStorageContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStorageContract.Contract.FileStorageContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileStorageContract *FileStorageContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileStorageContract.Contract.FileStorageContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileStorageContract *FileStorageContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileStorageContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileStorageContract *FileStorageContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStorageContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileStorageContract *FileStorageContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileStorageContract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FileStorageContract *FileStorageContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FileStorageContract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FileStorageContract *FileStorageContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FileStorageContract.Contract.UPGRADEINTERFACEVERSION(&_FileStorageContract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FileStorageContract *FileStorageContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FileStorageContract.Contract.UPGRADEINTERFACEVERSION(&_FileStorageContract.CallOpts)
}

// GetAllUserFiles is a free data retrieval call binding the contract method 0xc3ba2cd9.
//
// Solidity: function getAllUserFiles() view returns((string,uint256,bytes32,uint256,address,bool)[])
func (_FileStorageContract *FileStorageContractCaller) GetAllUserFiles(opts *bind.CallOpts) ([]FileStorageFile, error) {
	var out []interface{}
	err := _FileStorageContract.contract.Call(opts, &out, "getAllUserFiles")

	if err != nil {
		return *new([]FileStorageFile), err
	}

	out0 := *abi.ConvertType(out[0], new([]FileStorageFile)).(*[]FileStorageFile)

	return out0, err

}

// GetAllUserFiles is a free data retrieval call binding the contract method 0xc3ba2cd9.
//
// Solidity: function getAllUserFiles() view returns((string,uint256,bytes32,uint256,address,bool)[])
func (_FileStorageContract *FileStorageContractSession) GetAllUserFiles() ([]FileStorageFile, error) {
	return _FileStorageContract.Contract.GetAllUserFiles(&_FileStorageContract.CallOpts)
}

// GetAllUserFiles is a free data retrieval call binding the contract method 0xc3ba2cd9.
//
// Solidity: function getAllUserFiles() view returns((string,uint256,bytes32,uint256,address,bool)[])
func (_FileStorageContract *FileStorageContractCallerSession) GetAllUserFiles() ([]FileStorageFile, error) {
	return _FileStorageContract.Contract.GetAllUserFiles(&_FileStorageContract.CallOpts)
}

// GetUserFiles is a free data retrieval call binding the contract method 0x11942989.
//
// Solidity: function getUserFiles() view returns((string,uint256,bytes32,uint256,address,bool)[])
func (_FileStorageContract *FileStorageContractCaller) GetUserFiles(opts *bind.CallOpts) ([]FileStorageFile, error) {
	var out []interface{}
	err := _FileStorageContract.contract.Call(opts, &out, "getUserFiles")

	if err != nil {
		return *new([]FileStorageFile), err
	}

	out0 := *abi.ConvertType(out[0], new([]FileStorageFile)).(*[]FileStorageFile)

	return out0, err

}

// GetUserFiles is a free data retrieval call binding the contract method 0x11942989.
//
// Solidity: function getUserFiles() view returns((string,uint256,bytes32,uint256,address,bool)[])
func (_FileStorageContract *FileStorageContractSession) GetUserFiles() ([]FileStorageFile, error) {
	return _FileStorageContract.Contract.GetUserFiles(&_FileStorageContract.CallOpts)
}

// GetUserFiles is a free data retrieval call binding the contract method 0x11942989.
//
// Solidity: function getUserFiles() view returns((string,uint256,bytes32,uint256,address,bool)[])
func (_FileStorageContract *FileStorageContractCallerSession) GetUserFiles() ([]FileStorageFile, error) {
	return _FileStorageContract.Contract.GetUserFiles(&_FileStorageContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FileStorageContract *FileStorageContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileStorageContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FileStorageContract *FileStorageContractSession) Owner() (common.Address, error) {
	return _FileStorageContract.Contract.Owner(&_FileStorageContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FileStorageContract *FileStorageContractCallerSession) Owner() (common.Address, error) {
	return _FileStorageContract.Contract.Owner(&_FileStorageContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FileStorageContract *FileStorageContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FileStorageContract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FileStorageContract *FileStorageContractSession) ProxiableUUID() ([32]byte, error) {
	return _FileStorageContract.Contract.ProxiableUUID(&_FileStorageContract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FileStorageContract *FileStorageContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FileStorageContract.Contract.ProxiableUUID(&_FileStorageContract.CallOpts)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x6ab799f1.
//
// Solidity: function deleteFile(bytes32 _hash) returns()
func (_FileStorageContract *FileStorageContractTransactor) DeleteFile(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "deleteFile", _hash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x6ab799f1.
//
// Solidity: function deleteFile(bytes32 _hash) returns()
func (_FileStorageContract *FileStorageContractSession) DeleteFile(_hash [32]byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.DeleteFile(&_FileStorageContract.TransactOpts, _hash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0x6ab799f1.
//
// Solidity: function deleteFile(bytes32 _hash) returns()
func (_FileStorageContract *FileStorageContractTransactorSession) DeleteFile(_hash [32]byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.DeleteFile(&_FileStorageContract.TransactOpts, _hash)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_FileStorageContract *FileStorageContractTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_FileStorageContract *FileStorageContractSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _FileStorageContract.Contract.Initialize(&_FileStorageContract.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_FileStorageContract *FileStorageContractTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _FileStorageContract.Contract.Initialize(&_FileStorageContract.TransactOpts, initialOwner)
}

// ModifyFile is a paid mutator transaction binding the contract method 0x9d5aca64.
//
// Solidity: function modifyFile(bytes32 _hash, string _newName, uint256 _newSize) returns()
func (_FileStorageContract *FileStorageContractTransactor) ModifyFile(opts *bind.TransactOpts, _hash [32]byte, _newName string, _newSize *big.Int) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "modifyFile", _hash, _newName, _newSize)
}

// ModifyFile is a paid mutator transaction binding the contract method 0x9d5aca64.
//
// Solidity: function modifyFile(bytes32 _hash, string _newName, uint256 _newSize) returns()
func (_FileStorageContract *FileStorageContractSession) ModifyFile(_hash [32]byte, _newName string, _newSize *big.Int) (*types.Transaction, error) {
	return _FileStorageContract.Contract.ModifyFile(&_FileStorageContract.TransactOpts, _hash, _newName, _newSize)
}

// ModifyFile is a paid mutator transaction binding the contract method 0x9d5aca64.
//
// Solidity: function modifyFile(bytes32 _hash, string _newName, uint256 _newSize) returns()
func (_FileStorageContract *FileStorageContractTransactorSession) ModifyFile(_hash [32]byte, _newName string, _newSize *big.Int) (*types.Transaction, error) {
	return _FileStorageContract.Contract.ModifyFile(&_FileStorageContract.TransactOpts, _hash, _newName, _newSize)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FileStorageContract *FileStorageContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FileStorageContract *FileStorageContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _FileStorageContract.Contract.RenounceOwnership(&_FileStorageContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FileStorageContract *FileStorageContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FileStorageContract.Contract.RenounceOwnership(&_FileStorageContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FileStorageContract *FileStorageContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FileStorageContract *FileStorageContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FileStorageContract.Contract.TransferOwnership(&_FileStorageContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FileStorageContract *FileStorageContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FileStorageContract.Contract.TransferOwnership(&_FileStorageContract.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FileStorageContract *FileStorageContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FileStorageContract *FileStorageContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.UpgradeToAndCall(&_FileStorageContract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FileStorageContract *FileStorageContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.UpgradeToAndCall(&_FileStorageContract.TransactOpts, newImplementation, data)
}

// UploadFile is a paid mutator transaction binding the contract method 0x854c932c.
//
// Solidity: function uploadFile(string _name, uint256 _size, bytes32 _hash) returns()
func (_FileStorageContract *FileStorageContractTransactor) UploadFile(opts *bind.TransactOpts, _name string, _size *big.Int, _hash [32]byte) (*types.Transaction, error) {
	return _FileStorageContract.contract.Transact(opts, "uploadFile", _name, _size, _hash)
}

// UploadFile is a paid mutator transaction binding the contract method 0x854c932c.
//
// Solidity: function uploadFile(string _name, uint256 _size, bytes32 _hash) returns()
func (_FileStorageContract *FileStorageContractSession) UploadFile(_name string, _size *big.Int, _hash [32]byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.UploadFile(&_FileStorageContract.TransactOpts, _name, _size, _hash)
}

// UploadFile is a paid mutator transaction binding the contract method 0x854c932c.
//
// Solidity: function uploadFile(string _name, uint256 _size, bytes32 _hash) returns()
func (_FileStorageContract *FileStorageContractTransactorSession) UploadFile(_name string, _size *big.Int, _hash [32]byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.UploadFile(&_FileStorageContract.TransactOpts, _name, _size, _hash)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FileStorageContract *FileStorageContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FileStorageContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FileStorageContract *FileStorageContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.Fallback(&_FileStorageContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FileStorageContract *FileStorageContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FileStorageContract.Contract.Fallback(&_FileStorageContract.TransactOpts, calldata)
}

// FileStorageContractAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the FileStorageContract contract.
type FileStorageContractAdminChangedIterator struct {
	Event *FileStorageContractAdminChanged // Event containing the contract specifics and raw log

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
func (it *FileStorageContractAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractAdminChanged)
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
		it.Event = new(FileStorageContractAdminChanged)
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
func (it *FileStorageContractAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractAdminChanged represents a AdminChanged event raised by the FileStorageContract contract.
type FileStorageContractAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FileStorageContract *FileStorageContractFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FileStorageContractAdminChangedIterator, error) {

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FileStorageContractAdminChangedIterator{contract: _FileStorageContract.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_FileStorageContract *FileStorageContractFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FileStorageContractAdminChanged) (event.Subscription, error) {

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractAdminChanged)
				if err := _FileStorageContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_FileStorageContract *FileStorageContractFilterer) ParseAdminChanged(log types.Log) (*FileStorageContractAdminChanged, error) {
	event := new(FileStorageContractAdminChanged)
	if err := _FileStorageContract.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractFileDeletedIterator is returned from FilterFileDeleted and is used to iterate over the raw logs and unpacked data for FileDeleted events raised by the FileStorageContract contract.
type FileStorageContractFileDeletedIterator struct {
	Event *FileStorageContractFileDeleted // Event containing the contract specifics and raw log

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
func (it *FileStorageContractFileDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractFileDeleted)
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
		it.Event = new(FileStorageContractFileDeleted)
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
func (it *FileStorageContractFileDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractFileDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractFileDeleted represents a FileDeleted event raised by the FileStorageContract contract.
type FileStorageContractFileDeleted struct {
	User     common.Address
	FileHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileDeleted is a free log retrieval operation binding the contract event 0x0b61407fda4ea870fbaa77e68ef9d3f6e3e78eaa5c4fa4674d5ac417f56055f9.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash)
func (_FileStorageContract *FileStorageContractFilterer) FilterFileDeleted(opts *bind.FilterOpts, user []common.Address, fileHash [][32]byte) (*FileStorageContractFileDeletedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fileHashRule []interface{}
	for _, fileHashItem := range fileHash {
		fileHashRule = append(fileHashRule, fileHashItem)
	}

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "FileDeleted", userRule, fileHashRule)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractFileDeletedIterator{contract: _FileStorageContract.contract, event: "FileDeleted", logs: logs, sub: sub}, nil
}

// WatchFileDeleted is a free log subscription operation binding the contract event 0x0b61407fda4ea870fbaa77e68ef9d3f6e3e78eaa5c4fa4674d5ac417f56055f9.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash)
func (_FileStorageContract *FileStorageContractFilterer) WatchFileDeleted(opts *bind.WatchOpts, sink chan<- *FileStorageContractFileDeleted, user []common.Address, fileHash [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fileHashRule []interface{}
	for _, fileHashItem := range fileHash {
		fileHashRule = append(fileHashRule, fileHashItem)
	}

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "FileDeleted", userRule, fileHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractFileDeleted)
				if err := _FileStorageContract.contract.UnpackLog(event, "FileDeleted", log); err != nil {
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

// ParseFileDeleted is a log parse operation binding the contract event 0x0b61407fda4ea870fbaa77e68ef9d3f6e3e78eaa5c4fa4674d5ac417f56055f9.
//
// Solidity: event FileDeleted(address indexed user, bytes32 indexed fileHash)
func (_FileStorageContract *FileStorageContractFilterer) ParseFileDeleted(log types.Log) (*FileStorageContractFileDeleted, error) {
	event := new(FileStorageContractFileDeleted)
	if err := _FileStorageContract.contract.UnpackLog(event, "FileDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractFileModifiedIterator is returned from FilterFileModified and is used to iterate over the raw logs and unpacked data for FileModified events raised by the FileStorageContract contract.
type FileStorageContractFileModifiedIterator struct {
	Event *FileStorageContractFileModified // Event containing the contract specifics and raw log

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
func (it *FileStorageContractFileModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractFileModified)
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
		it.Event = new(FileStorageContractFileModified)
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
func (it *FileStorageContractFileModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractFileModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractFileModified represents a FileModified event raised by the FileStorageContract contract.
type FileStorageContractFileModified struct {
	User     common.Address
	FileHash [32]byte
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileModified is a free log retrieval operation binding the contract event 0xd0ef5aa776e62fe721358136c990a215323daee3e913bf55795d3ada3de5b2ed.
//
// Solidity: event FileModified(address indexed user, bytes32 indexed fileHash, string name)
func (_FileStorageContract *FileStorageContractFilterer) FilterFileModified(opts *bind.FilterOpts, user []common.Address, fileHash [][32]byte) (*FileStorageContractFileModifiedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fileHashRule []interface{}
	for _, fileHashItem := range fileHash {
		fileHashRule = append(fileHashRule, fileHashItem)
	}

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "FileModified", userRule, fileHashRule)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractFileModifiedIterator{contract: _FileStorageContract.contract, event: "FileModified", logs: logs, sub: sub}, nil
}

// WatchFileModified is a free log subscription operation binding the contract event 0xd0ef5aa776e62fe721358136c990a215323daee3e913bf55795d3ada3de5b2ed.
//
// Solidity: event FileModified(address indexed user, bytes32 indexed fileHash, string name)
func (_FileStorageContract *FileStorageContractFilterer) WatchFileModified(opts *bind.WatchOpts, sink chan<- *FileStorageContractFileModified, user []common.Address, fileHash [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fileHashRule []interface{}
	for _, fileHashItem := range fileHash {
		fileHashRule = append(fileHashRule, fileHashItem)
	}

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "FileModified", userRule, fileHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractFileModified)
				if err := _FileStorageContract.contract.UnpackLog(event, "FileModified", log); err != nil {
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

// ParseFileModified is a log parse operation binding the contract event 0xd0ef5aa776e62fe721358136c990a215323daee3e913bf55795d3ada3de5b2ed.
//
// Solidity: event FileModified(address indexed user, bytes32 indexed fileHash, string name)
func (_FileStorageContract *FileStorageContractFilterer) ParseFileModified(log types.Log) (*FileStorageContractFileModified, error) {
	event := new(FileStorageContractFileModified)
	if err := _FileStorageContract.contract.UnpackLog(event, "FileModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractFileUploadedIterator is returned from FilterFileUploaded and is used to iterate over the raw logs and unpacked data for FileUploaded events raised by the FileStorageContract contract.
type FileStorageContractFileUploadedIterator struct {
	Event *FileStorageContractFileUploaded // Event containing the contract specifics and raw log

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
func (it *FileStorageContractFileUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractFileUploaded)
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
		it.Event = new(FileStorageContractFileUploaded)
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
func (it *FileStorageContractFileUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractFileUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractFileUploaded represents a FileUploaded event raised by the FileStorageContract contract.
type FileStorageContractFileUploaded struct {
	User     common.Address
	FileHash [32]byte
	Name     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileUploaded is a free log retrieval operation binding the contract event 0xd404a946e4879a6393311b452f8f3945c811eca1ce782109276b7487f148c5aa.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string name)
func (_FileStorageContract *FileStorageContractFilterer) FilterFileUploaded(opts *bind.FilterOpts, user []common.Address, fileHash [][32]byte) (*FileStorageContractFileUploadedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fileHashRule []interface{}
	for _, fileHashItem := range fileHash {
		fileHashRule = append(fileHashRule, fileHashItem)
	}

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "FileUploaded", userRule, fileHashRule)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractFileUploadedIterator{contract: _FileStorageContract.contract, event: "FileUploaded", logs: logs, sub: sub}, nil
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0xd404a946e4879a6393311b452f8f3945c811eca1ce782109276b7487f148c5aa.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string name)
func (_FileStorageContract *FileStorageContractFilterer) WatchFileUploaded(opts *bind.WatchOpts, sink chan<- *FileStorageContractFileUploaded, user []common.Address, fileHash [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var fileHashRule []interface{}
	for _, fileHashItem := range fileHash {
		fileHashRule = append(fileHashRule, fileHashItem)
	}

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "FileUploaded", userRule, fileHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractFileUploaded)
				if err := _FileStorageContract.contract.UnpackLog(event, "FileUploaded", log); err != nil {
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

// ParseFileUploaded is a log parse operation binding the contract event 0xd404a946e4879a6393311b452f8f3945c811eca1ce782109276b7487f148c5aa.
//
// Solidity: event FileUploaded(address indexed user, bytes32 indexed fileHash, string name)
func (_FileStorageContract *FileStorageContractFilterer) ParseFileUploaded(log types.Log) (*FileStorageContractFileUploaded, error) {
	event := new(FileStorageContractFileUploaded)
	if err := _FileStorageContract.contract.UnpackLog(event, "FileUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FileStorageContract contract.
type FileStorageContractInitializedIterator struct {
	Event *FileStorageContractInitialized // Event containing the contract specifics and raw log

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
func (it *FileStorageContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractInitialized)
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
		it.Event = new(FileStorageContractInitialized)
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
func (it *FileStorageContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractInitialized represents a Initialized event raised by the FileStorageContract contract.
type FileStorageContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FileStorageContract *FileStorageContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*FileStorageContractInitializedIterator, error) {

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FileStorageContractInitializedIterator{contract: _FileStorageContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FileStorageContract *FileStorageContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FileStorageContractInitialized) (event.Subscription, error) {

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractInitialized)
				if err := _FileStorageContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FileStorageContract *FileStorageContractFilterer) ParseInitialized(log types.Log) (*FileStorageContractInitialized, error) {
	event := new(FileStorageContractInitialized)
	if err := _FileStorageContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FileStorageContract contract.
type FileStorageContractOwnershipTransferredIterator struct {
	Event *FileStorageContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FileStorageContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractOwnershipTransferred)
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
		it.Event = new(FileStorageContractOwnershipTransferred)
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
func (it *FileStorageContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractOwnershipTransferred represents a OwnershipTransferred event raised by the FileStorageContract contract.
type FileStorageContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FileStorageContract *FileStorageContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FileStorageContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractOwnershipTransferredIterator{contract: _FileStorageContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FileStorageContract *FileStorageContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FileStorageContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractOwnershipTransferred)
				if err := _FileStorageContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FileStorageContract *FileStorageContractFilterer) ParseOwnershipTransferred(log types.Log) (*FileStorageContractOwnershipTransferred, error) {
	event := new(FileStorageContractOwnershipTransferred)
	if err := _FileStorageContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FileStorageContract contract.
type FileStorageContractUpgradedIterator struct {
	Event *FileStorageContractUpgraded // Event containing the contract specifics and raw log

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
func (it *FileStorageContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractUpgraded)
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
		it.Event = new(FileStorageContractUpgraded)
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
func (it *FileStorageContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractUpgraded represents a Upgraded event raised by the FileStorageContract contract.
type FileStorageContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileStorageContract *FileStorageContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FileStorageContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractUpgradedIterator{contract: _FileStorageContract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileStorageContract *FileStorageContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FileStorageContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractUpgraded)
				if err := _FileStorageContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FileStorageContract *FileStorageContractFilterer) ParseUpgraded(log types.Log) (*FileStorageContractUpgraded, error) {
	event := new(FileStorageContractUpgraded)
	if err := _FileStorageContract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileStorageContractUpgraded0Iterator is returned from FilterUpgraded0 and is used to iterate over the raw logs and unpacked data for Upgraded0 events raised by the FileStorageContract contract.
type FileStorageContractUpgraded0Iterator struct {
	Event *FileStorageContractUpgraded0 // Event containing the contract specifics and raw log

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
func (it *FileStorageContractUpgraded0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStorageContractUpgraded0)
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
		it.Event = new(FileStorageContractUpgraded0)
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
func (it *FileStorageContractUpgraded0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStorageContractUpgraded0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStorageContractUpgraded0 represents a Upgraded0 event raised by the FileStorageContract contract.
type FileStorageContractUpgraded0 struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded0 is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileStorageContract *FileStorageContractFilterer) FilterUpgraded0(opts *bind.FilterOpts, implementation []common.Address) (*FileStorageContractUpgraded0Iterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FileStorageContract.contract.FilterLogs(opts, "Upgraded0", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FileStorageContractUpgraded0Iterator{contract: _FileStorageContract.contract, event: "Upgraded0", logs: logs, sub: sub}, nil
}

// WatchUpgraded0 is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileStorageContract *FileStorageContractFilterer) WatchUpgraded0(opts *bind.WatchOpts, sink chan<- *FileStorageContractUpgraded0, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FileStorageContract.contract.WatchLogs(opts, "Upgraded0", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStorageContractUpgraded0)
				if err := _FileStorageContract.contract.UnpackLog(event, "Upgraded0", log); err != nil {
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

// ParseUpgraded0 is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileStorageContract *FileStorageContractFilterer) ParseUpgraded0(log types.Log) (*FileStorageContractUpgraded0, error) {
	event := new(FileStorageContractUpgraded0)
	if err := _FileStorageContract.contract.UnpackLog(event, "Upgraded0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
