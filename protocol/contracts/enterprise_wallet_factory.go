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

// IEnterpriseWalletFactoryInitParams is an auto generated low-level Go binding around an user-defined struct.
type IEnterpriseWalletFactoryInitParams struct {
	Methods    [][4]byte
	Configs    []IEnterpriseWalletMethodConfig
	SuperAdmin common.Address
}

// EnterpriseWalletFactoryMetaData contains all meta data concerning the EnterpriseWalletFactory contract.
var EnterpriseWalletFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIEnterpriseWalletFactory.InitParams\",\"components\":[{\"name\":\"methods\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"},{\"name\":\"configs\",\"type\":\"tuple[]\",\"internalType\":\"structIEnterpriseWallet.MethodConfig[]\",\"components\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"superAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getWhitelistedImplementations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isImplementationWhitelisted\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"predictWalletAddress\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ImplementationAdded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationRemoved\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WalletCreated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ImplementationAlreadyExists\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ImplementationNotWhitelisted\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidImplementation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WalletAlreadyExists\",\"inputs\":[]}]",
}

// EnterpriseWalletFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseWalletFactoryMetaData.ABI instead.
var EnterpriseWalletFactoryABI = EnterpriseWalletFactoryMetaData.ABI

// EnterpriseWalletFactory is an auto generated Go binding around an Ethereum contract.
type EnterpriseWalletFactory struct {
	EnterpriseWalletFactoryCaller     // Read-only binding to the contract
	EnterpriseWalletFactoryTransactor // Write-only binding to the contract
	EnterpriseWalletFactoryFilterer   // Log filterer for contract events
}

// EnterpriseWalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseWalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseWalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseWalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseWalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseWalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseWalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseWalletFactorySession struct {
	Contract     *EnterpriseWalletFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnterpriseWalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseWalletFactoryCallerSession struct {
	Contract *EnterpriseWalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// EnterpriseWalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseWalletFactoryTransactorSession struct {
	Contract     *EnterpriseWalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// EnterpriseWalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseWalletFactoryRaw struct {
	Contract *EnterpriseWalletFactory // Generic contract binding to access the raw methods on
}

// EnterpriseWalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseWalletFactoryCallerRaw struct {
	Contract *EnterpriseWalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseWalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseWalletFactoryTransactorRaw struct {
	Contract *EnterpriseWalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterpriseWalletFactory creates a new instance of EnterpriseWalletFactory, bound to a specific deployed contract.
func NewEnterpriseWalletFactory(address common.Address, backend bind.ContractBackend) (*EnterpriseWalletFactory, error) {
	contract, err := bindEnterpriseWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactory{EnterpriseWalletFactoryCaller: EnterpriseWalletFactoryCaller{contract: contract}, EnterpriseWalletFactoryTransactor: EnterpriseWalletFactoryTransactor{contract: contract}, EnterpriseWalletFactoryFilterer: EnterpriseWalletFactoryFilterer{contract: contract}}, nil
}

// NewEnterpriseWalletFactoryCaller creates a new read-only instance of EnterpriseWalletFactory, bound to a specific deployed contract.
func NewEnterpriseWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseWalletFactoryCaller, error) {
	contract, err := bindEnterpriseWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryCaller{contract: contract}, nil
}

// NewEnterpriseWalletFactoryTransactor creates a new write-only instance of EnterpriseWalletFactory, bound to a specific deployed contract.
func NewEnterpriseWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseWalletFactoryTransactor, error) {
	contract, err := bindEnterpriseWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryTransactor{contract: contract}, nil
}

// NewEnterpriseWalletFactoryFilterer creates a new log filterer instance of EnterpriseWalletFactory, bound to a specific deployed contract.
func NewEnterpriseWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseWalletFactoryFilterer, error) {
	contract, err := bindEnterpriseWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryFilterer{contract: contract}, nil
}

// bindEnterpriseWalletFactory binds a generic wrapper to an already deployed contract.
func bindEnterpriseWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnterpriseWalletFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseWalletFactory.Contract.EnterpriseWalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.EnterpriseWalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.EnterpriseWalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseWalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.contract.Transact(opts, method, params...)
}

// GetWhitelistedImplementations is a free data retrieval call binding the contract method 0x19bc1a0c.
//
// Solidity: function getWhitelistedImplementations() view returns(address[])
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCaller) GetWhitelistedImplementations(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EnterpriseWalletFactory.contract.Call(opts, &out, "getWhitelistedImplementations")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedImplementations is a free data retrieval call binding the contract method 0x19bc1a0c.
//
// Solidity: function getWhitelistedImplementations() view returns(address[])
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) GetWhitelistedImplementations() ([]common.Address, error) {
	return _EnterpriseWalletFactory.Contract.GetWhitelistedImplementations(&_EnterpriseWalletFactory.CallOpts)
}

// GetWhitelistedImplementations is a free data retrieval call binding the contract method 0x19bc1a0c.
//
// Solidity: function getWhitelistedImplementations() view returns(address[])
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCallerSession) GetWhitelistedImplementations() ([]common.Address, error) {
	return _EnterpriseWalletFactory.Contract.GetWhitelistedImplementations(&_EnterpriseWalletFactory.CallOpts)
}

// IsImplementationWhitelisted is a free data retrieval call binding the contract method 0x8876c2eb.
//
// Solidity: function isImplementationWhitelisted(address implementation) view returns(bool)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCaller) IsImplementationWhitelisted(opts *bind.CallOpts, implementation common.Address) (bool, error) {
	var out []interface{}
	err := _EnterpriseWalletFactory.contract.Call(opts, &out, "isImplementationWhitelisted", implementation)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsImplementationWhitelisted is a free data retrieval call binding the contract method 0x8876c2eb.
//
// Solidity: function isImplementationWhitelisted(address implementation) view returns(bool)
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) IsImplementationWhitelisted(implementation common.Address) (bool, error) {
	return _EnterpriseWalletFactory.Contract.IsImplementationWhitelisted(&_EnterpriseWalletFactory.CallOpts, implementation)
}

// IsImplementationWhitelisted is a free data retrieval call binding the contract method 0x8876c2eb.
//
// Solidity: function isImplementationWhitelisted(address implementation) view returns(bool)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCallerSession) IsImplementationWhitelisted(implementation common.Address) (bool, error) {
	return _EnterpriseWalletFactory.Contract.IsImplementationWhitelisted(&_EnterpriseWalletFactory.CallOpts, implementation)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseWalletFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) Owner() (common.Address, error) {
	return _EnterpriseWalletFactory.Contract.Owner(&_EnterpriseWalletFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCallerSession) Owner() (common.Address, error) {
	return _EnterpriseWalletFactory.Contract.Owner(&_EnterpriseWalletFactory.CallOpts)
}

// PredictWalletAddress is a free data retrieval call binding the contract method 0x65205ee0.
//
// Solidity: function predictWalletAddress(address implementation, bytes32 salt, address deployer) view returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCaller) PredictWalletAddress(opts *bind.CallOpts, implementation common.Address, salt [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseWalletFactory.contract.Call(opts, &out, "predictWalletAddress", implementation, salt, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictWalletAddress is a free data retrieval call binding the contract method 0x65205ee0.
//
// Solidity: function predictWalletAddress(address implementation, bytes32 salt, address deployer) view returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) PredictWalletAddress(implementation common.Address, salt [32]byte, deployer common.Address) (common.Address, error) {
	return _EnterpriseWalletFactory.Contract.PredictWalletAddress(&_EnterpriseWalletFactory.CallOpts, implementation, salt, deployer)
}

// PredictWalletAddress is a free data retrieval call binding the contract method 0x65205ee0.
//
// Solidity: function predictWalletAddress(address implementation, bytes32 salt, address deployer) view returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryCallerSession) PredictWalletAddress(implementation common.Address, salt [32]byte, deployer common.Address) (common.Address, error) {
	return _EnterpriseWalletFactory.Contract.PredictWalletAddress(&_EnterpriseWalletFactory.CallOpts, implementation, salt, deployer)
}

// AddImplementation is a paid mutator transaction binding the contract method 0xc6e2a400.
//
// Solidity: function addImplementation(address implementation) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactor) AddImplementation(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.contract.Transact(opts, "addImplementation", implementation)
}

// AddImplementation is a paid mutator transaction binding the contract method 0xc6e2a400.
//
// Solidity: function addImplementation(address implementation) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) AddImplementation(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.AddImplementation(&_EnterpriseWalletFactory.TransactOpts, implementation)
}

// AddImplementation is a paid mutator transaction binding the contract method 0xc6e2a400.
//
// Solidity: function addImplementation(address implementation) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorSession) AddImplementation(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.AddImplementation(&_EnterpriseWalletFactory.TransactOpts, implementation)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xdec4d1af.
//
// Solidity: function createWallet(address implementation, bytes32 salt, (bytes4[],(address)[],address) params) returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactor) CreateWallet(opts *bind.TransactOpts, implementation common.Address, salt [32]byte, params IEnterpriseWalletFactoryInitParams) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.contract.Transact(opts, "createWallet", implementation, salt, params)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xdec4d1af.
//
// Solidity: function createWallet(address implementation, bytes32 salt, (bytes4[],(address)[],address) params) returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) CreateWallet(implementation common.Address, salt [32]byte, params IEnterpriseWalletFactoryInitParams) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.CreateWallet(&_EnterpriseWalletFactory.TransactOpts, implementation, salt, params)
}

// CreateWallet is a paid mutator transaction binding the contract method 0xdec4d1af.
//
// Solidity: function createWallet(address implementation, bytes32 salt, (bytes4[],(address)[],address) params) returns(address)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorSession) CreateWallet(implementation common.Address, salt [32]byte, params IEnterpriseWalletFactoryInitParams) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.CreateWallet(&_EnterpriseWalletFactory.TransactOpts, implementation, salt, params)
}

// RemoveImplementation is a paid mutator transaction binding the contract method 0x22175a32.
//
// Solidity: function removeImplementation(address implementation) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactor) RemoveImplementation(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.contract.Transact(opts, "removeImplementation", implementation)
}

// RemoveImplementation is a paid mutator transaction binding the contract method 0x22175a32.
//
// Solidity: function removeImplementation(address implementation) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) RemoveImplementation(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.RemoveImplementation(&_EnterpriseWalletFactory.TransactOpts, implementation)
}

// RemoveImplementation is a paid mutator transaction binding the contract method 0x22175a32.
//
// Solidity: function removeImplementation(address implementation) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorSession) RemoveImplementation(implementation common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.RemoveImplementation(&_EnterpriseWalletFactory.TransactOpts, implementation)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.RenounceOwnership(&_EnterpriseWalletFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.RenounceOwnership(&_EnterpriseWalletFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.TransferOwnership(&_EnterpriseWalletFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EnterpriseWalletFactory.Contract.TransferOwnership(&_EnterpriseWalletFactory.TransactOpts, newOwner)
}

// EnterpriseWalletFactoryImplementationAddedIterator is returned from FilterImplementationAdded and is used to iterate over the raw logs and unpacked data for ImplementationAdded events raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryImplementationAddedIterator struct {
	Event *EnterpriseWalletFactoryImplementationAdded // Event containing the contract specifics and raw log

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
func (it *EnterpriseWalletFactoryImplementationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletFactoryImplementationAdded)
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
		it.Event = new(EnterpriseWalletFactoryImplementationAdded)
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
func (it *EnterpriseWalletFactoryImplementationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletFactoryImplementationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletFactoryImplementationAdded represents a ImplementationAdded event raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryImplementationAdded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterImplementationAdded is a free log retrieval operation binding the contract event 0x331cedc71f28c46d467691770675b586e8aa77a0d4fe09f257d01ef00bc15458.
//
// Solidity: event ImplementationAdded(address indexed implementation)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) FilterImplementationAdded(opts *bind.FilterOpts, implementation []common.Address) (*EnterpriseWalletFactoryImplementationAddedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.FilterLogs(opts, "ImplementationAdded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryImplementationAddedIterator{contract: _EnterpriseWalletFactory.contract, event: "ImplementationAdded", logs: logs, sub: sub}, nil
}

// WatchImplementationAdded is a free log subscription operation binding the contract event 0x331cedc71f28c46d467691770675b586e8aa77a0d4fe09f257d01ef00bc15458.
//
// Solidity: event ImplementationAdded(address indexed implementation)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) WatchImplementationAdded(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletFactoryImplementationAdded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.WatchLogs(opts, "ImplementationAdded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletFactoryImplementationAdded)
				if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "ImplementationAdded", log); err != nil {
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

// ParseImplementationAdded is a log parse operation binding the contract event 0x331cedc71f28c46d467691770675b586e8aa77a0d4fe09f257d01ef00bc15458.
//
// Solidity: event ImplementationAdded(address indexed implementation)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) ParseImplementationAdded(log types.Log) (*EnterpriseWalletFactoryImplementationAdded, error) {
	event := new(EnterpriseWalletFactoryImplementationAdded)
	if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "ImplementationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletFactoryImplementationRemovedIterator is returned from FilterImplementationRemoved and is used to iterate over the raw logs and unpacked data for ImplementationRemoved events raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryImplementationRemovedIterator struct {
	Event *EnterpriseWalletFactoryImplementationRemoved // Event containing the contract specifics and raw log

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
func (it *EnterpriseWalletFactoryImplementationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletFactoryImplementationRemoved)
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
		it.Event = new(EnterpriseWalletFactoryImplementationRemoved)
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
func (it *EnterpriseWalletFactoryImplementationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletFactoryImplementationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletFactoryImplementationRemoved represents a ImplementationRemoved event raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryImplementationRemoved struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterImplementationRemoved is a free log retrieval operation binding the contract event 0xaf23121e2402485071dadf421078b368d7b67e54cabcc81540563c5d6bf1a4c3.
//
// Solidity: event ImplementationRemoved(address indexed implementation)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) FilterImplementationRemoved(opts *bind.FilterOpts, implementation []common.Address) (*EnterpriseWalletFactoryImplementationRemovedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.FilterLogs(opts, "ImplementationRemoved", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryImplementationRemovedIterator{contract: _EnterpriseWalletFactory.contract, event: "ImplementationRemoved", logs: logs, sub: sub}, nil
}

// WatchImplementationRemoved is a free log subscription operation binding the contract event 0xaf23121e2402485071dadf421078b368d7b67e54cabcc81540563c5d6bf1a4c3.
//
// Solidity: event ImplementationRemoved(address indexed implementation)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) WatchImplementationRemoved(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletFactoryImplementationRemoved, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.WatchLogs(opts, "ImplementationRemoved", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletFactoryImplementationRemoved)
				if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "ImplementationRemoved", log); err != nil {
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

// ParseImplementationRemoved is a log parse operation binding the contract event 0xaf23121e2402485071dadf421078b368d7b67e54cabcc81540563c5d6bf1a4c3.
//
// Solidity: event ImplementationRemoved(address indexed implementation)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) ParseImplementationRemoved(log types.Log) (*EnterpriseWalletFactoryImplementationRemoved, error) {
	event := new(EnterpriseWalletFactoryImplementationRemoved)
	if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "ImplementationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryOwnershipTransferredIterator struct {
	Event *EnterpriseWalletFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EnterpriseWalletFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletFactoryOwnershipTransferred)
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
		it.Event = new(EnterpriseWalletFactoryOwnershipTransferred)
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
func (it *EnterpriseWalletFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EnterpriseWalletFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryOwnershipTransferredIterator{contract: _EnterpriseWalletFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletFactoryOwnershipTransferred)
				if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*EnterpriseWalletFactoryOwnershipTransferred, error) {
	event := new(EnterpriseWalletFactoryOwnershipTransferred)
	if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnterpriseWalletFactoryWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryWalletCreatedIterator struct {
	Event *EnterpriseWalletFactoryWalletCreated // Event containing the contract specifics and raw log

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
func (it *EnterpriseWalletFactoryWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseWalletFactoryWalletCreated)
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
		it.Event = new(EnterpriseWalletFactoryWalletCreated)
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
func (it *EnterpriseWalletFactoryWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseWalletFactoryWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseWalletFactoryWalletCreated represents a WalletCreated event raised by the EnterpriseWalletFactory contract.
type EnterpriseWalletFactoryWalletCreated struct {
	Wallet         common.Address
	Owner          common.Address
	Implementation common.Address
	Salt           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xdf0ba70dab079d6aae8592d45293b0f7ffbbaaf7b10676c06b702d67a3c2cf90.
//
// Solidity: event WalletCreated(address indexed wallet, address indexed owner, address indexed implementation, bytes32 salt)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) FilterWalletCreated(opts *bind.FilterOpts, wallet []common.Address, owner []common.Address, implementation []common.Address) (*EnterpriseWalletFactoryWalletCreatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.FilterLogs(opts, "WalletCreated", walletRule, ownerRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseWalletFactoryWalletCreatedIterator{contract: _EnterpriseWalletFactory.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xdf0ba70dab079d6aae8592d45293b0f7ffbbaaf7b10676c06b702d67a3c2cf90.
//
// Solidity: event WalletCreated(address indexed wallet, address indexed owner, address indexed implementation, bytes32 salt)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *EnterpriseWalletFactoryWalletCreated, wallet []common.Address, owner []common.Address, implementation []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EnterpriseWalletFactory.contract.WatchLogs(opts, "WalletCreated", walletRule, ownerRule, implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseWalletFactoryWalletCreated)
				if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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

// ParseWalletCreated is a log parse operation binding the contract event 0xdf0ba70dab079d6aae8592d45293b0f7ffbbaaf7b10676c06b702d67a3c2cf90.
//
// Solidity: event WalletCreated(address indexed wallet, address indexed owner, address indexed implementation, bytes32 salt)
func (_EnterpriseWalletFactory *EnterpriseWalletFactoryFilterer) ParseWalletCreated(log types.Log) (*EnterpriseWalletFactoryWalletCreated, error) {
	event := new(EnterpriseWalletFactoryWalletCreated)
	if err := _EnterpriseWalletFactory.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
