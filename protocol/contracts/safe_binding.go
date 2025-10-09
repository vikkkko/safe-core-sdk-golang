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

// SafeBindingMetaData contains all meta data concerning the SafeBinding contract.
var SafeBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SafeBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeBindingMetaData.ABI instead.
var SafeBindingABI = SafeBindingMetaData.ABI

// SafeBinding is an auto generated Go binding around an Ethereum contract.
type SafeBinding struct {
	SafeBindingCaller     // Read-only binding to the contract
	SafeBindingTransactor // Write-only binding to the contract
	SafeBindingFilterer   // Log filterer for contract events
}

// SafeBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeBindingSession struct {
	Contract     *SafeBinding      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeBindingCallerSession struct {
	Contract *SafeBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SafeBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeBindingTransactorSession struct {
	Contract     *SafeBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SafeBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeBindingRaw struct {
	Contract *SafeBinding // Generic contract binding to access the raw methods on
}

// SafeBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeBindingCallerRaw struct {
	Contract *SafeBindingCaller // Generic read-only contract binding to access the raw methods on
}

// SafeBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeBindingTransactorRaw struct {
	Contract *SafeBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeBinding creates a new instance of SafeBinding, bound to a specific deployed contract.
func NewSafeBinding(address common.Address, backend bind.ContractBackend) (*SafeBinding, error) {
	contract, err := bindSafeBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeBinding{SafeBindingCaller: SafeBindingCaller{contract: contract}, SafeBindingTransactor: SafeBindingTransactor{contract: contract}, SafeBindingFilterer: SafeBindingFilterer{contract: contract}}, nil
}

// NewSafeBindingCaller creates a new read-only instance of SafeBinding, bound to a specific deployed contract.
func NewSafeBindingCaller(address common.Address, caller bind.ContractCaller) (*SafeBindingCaller, error) {
	contract, err := bindSafeBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeBindingCaller{contract: contract}, nil
}

// NewSafeBindingTransactor creates a new write-only instance of SafeBinding, bound to a specific deployed contract.
func NewSafeBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeBindingTransactor, error) {
	contract, err := bindSafeBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeBindingTransactor{contract: contract}, nil
}

// NewSafeBindingFilterer creates a new log filterer instance of SafeBinding, bound to a specific deployed contract.
func NewSafeBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeBindingFilterer, error) {
	contract, err := bindSafeBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeBindingFilterer{contract: contract}, nil
}

// bindSafeBinding binds a generic wrapper to an already deployed contract.
func bindSafeBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeBinding *SafeBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeBinding.Contract.SafeBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeBinding *SafeBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeBinding.Contract.SafeBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeBinding *SafeBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeBinding.Contract.SafeBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeBinding *SafeBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeBinding *SafeBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeBinding *SafeBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeBinding.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeBinding *SafeBindingCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeBinding *SafeBindingSession) VERSION() (string, error) {
	return _SafeBinding.Contract.VERSION(&_SafeBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeBinding *SafeBindingCallerSession) VERSION() (string, error) {
	return _SafeBinding.Contract.VERSION(&_SafeBinding.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeBinding *SafeBindingCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeBinding *SafeBindingSession) DomainSeparator() ([32]byte, error) {
	return _SafeBinding.Contract.DomainSeparator(&_SafeBinding.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeBinding *SafeBindingCallerSession) DomainSeparator() ([32]byte, error) {
	return _SafeBinding.Contract.DomainSeparator(&_SafeBinding.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeBinding *SafeBindingCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeBinding *SafeBindingSession) GetChainId() (*big.Int, error) {
	return _SafeBinding.Contract.GetChainId(&_SafeBinding.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeBinding *SafeBindingCallerSession) GetChainId() (*big.Int, error) {
	return _SafeBinding.Contract.GetChainId(&_SafeBinding.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0xc9106389.
//
// Solidity: function getGuard() view returns(address)
func (_SafeBinding *SafeBindingCaller) GetGuard(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "getGuard")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0xc9106389.
//
// Solidity: function getGuard() view returns(address)
func (_SafeBinding *SafeBindingSession) GetGuard() (common.Address, error) {
	return _SafeBinding.Contract.GetGuard(&_SafeBinding.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0xc9106389.
//
// Solidity: function getGuard() view returns(address)
func (_SafeBinding *SafeBindingCallerSession) GetGuard() (common.Address, error) {
	return _SafeBinding.Contract.GetGuard(&_SafeBinding.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeBinding *SafeBindingCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeBinding *SafeBindingSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SafeBinding.Contract.GetModulesPaginated(&_SafeBinding.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeBinding *SafeBindingCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SafeBinding.Contract.GetModulesPaginated(&_SafeBinding.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeBinding *SafeBindingCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeBinding *SafeBindingSession) GetOwners() ([]common.Address, error) {
	return _SafeBinding.Contract.GetOwners(&_SafeBinding.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeBinding *SafeBindingCallerSession) GetOwners() ([]common.Address, error) {
	return _SafeBinding.Contract.GetOwners(&_SafeBinding.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeBinding *SafeBindingCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeBinding *SafeBindingSession) GetThreshold() (*big.Int, error) {
	return _SafeBinding.Contract.GetThreshold(&_SafeBinding.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeBinding *SafeBindingCallerSession) GetThreshold() (*big.Int, error) {
	return _SafeBinding.Contract.GetThreshold(&_SafeBinding.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_SafeBinding *SafeBindingCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_SafeBinding *SafeBindingSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _SafeBinding.Contract.GetTransactionHash(&_SafeBinding.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_SafeBinding *SafeBindingCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _SafeBinding.Contract.GetTransactionHash(&_SafeBinding.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeBinding *SafeBindingCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeBinding *SafeBindingSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SafeBinding.Contract.IsModuleEnabled(&_SafeBinding.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeBinding *SafeBindingCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SafeBinding.Contract.IsModuleEnabled(&_SafeBinding.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeBinding *SafeBindingCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeBinding *SafeBindingSession) IsOwner(owner common.Address) (bool, error) {
	return _SafeBinding.Contract.IsOwner(&_SafeBinding.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeBinding *SafeBindingCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _SafeBinding.Contract.IsOwner(&_SafeBinding.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeBinding *SafeBindingCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeBinding.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeBinding *SafeBindingSession) Nonce() (*big.Int, error) {
	return _SafeBinding.Contract.Nonce(&_SafeBinding.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeBinding *SafeBindingCallerSession) Nonce() (*big.Int, error) {
	return _SafeBinding.Contract.Nonce(&_SafeBinding.CallOpts)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeBinding *SafeBindingTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeBinding *SafeBindingSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.Contract.AddOwnerWithThreshold(&_SafeBinding.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeBinding *SafeBindingTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.Contract.AddOwnerWithThreshold(&_SafeBinding.TransactOpts, owner, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeBinding *SafeBindingTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeBinding *SafeBindingSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.Contract.ChangeThreshold(&_SafeBinding.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeBinding *SafeBindingTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.Contract.ChangeThreshold(&_SafeBinding.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeBinding *SafeBindingTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeBinding *SafeBindingSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.DisableModule(&_SafeBinding.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeBinding *SafeBindingTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.DisableModule(&_SafeBinding.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeBinding *SafeBindingTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeBinding *SafeBindingSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.EnableModule(&_SafeBinding.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeBinding *SafeBindingTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.EnableModule(&_SafeBinding.TransactOpts, module)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeBinding *SafeBindingTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeBinding *SafeBindingSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.Contract.RemoveOwner(&_SafeBinding.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeBinding *SafeBindingTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeBinding.Contract.RemoveOwner(&_SafeBinding.TransactOpts, prevOwner, owner, _threshold)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeBinding *SafeBindingTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeBinding *SafeBindingSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.SetFallbackHandler(&_SafeBinding.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeBinding *SafeBindingTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.SetFallbackHandler(&_SafeBinding.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeBinding *SafeBindingTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeBinding *SafeBindingSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.SetGuard(&_SafeBinding.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeBinding *SafeBindingTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.SetGuard(&_SafeBinding.TransactOpts, guard)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeBinding *SafeBindingTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeBinding.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeBinding *SafeBindingSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.SwapOwner(&_SafeBinding.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeBinding *SafeBindingTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeBinding.Contract.SwapOwner(&_SafeBinding.TransactOpts, prevOwner, oldOwner, newOwner)
}
