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

// SimpleUniversalGuardFactoryMetaData contains all meta data concerning the SimpleUniversalGuardFactory contract.
var SimpleUniversalGuardFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createGuard\",\"inputs\":[{\"name\":\"requiredSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"guard\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getGuard\",\"inputs\":[{\"name\":\"requiredSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"predictGuardAddress\",\"inputs\":[{\"name\":\"requiredSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"predicted\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GuardCreated\",\"inputs\":[{\"name\":\"guard\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"requiredSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// SimpleUniversalGuardFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleUniversalGuardFactoryMetaData.ABI instead.
var SimpleUniversalGuardFactoryABI = SimpleUniversalGuardFactoryMetaData.ABI

// SimpleUniversalGuardFactory is an auto generated Go binding around an Ethereum contract.
type SimpleUniversalGuardFactory struct {
	SimpleUniversalGuardFactoryCaller     // Read-only binding to the contract
	SimpleUniversalGuardFactoryTransactor // Write-only binding to the contract
	SimpleUniversalGuardFactoryFilterer   // Log filterer for contract events
}

// SimpleUniversalGuardFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleUniversalGuardFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleUniversalGuardFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleUniversalGuardFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleUniversalGuardFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleUniversalGuardFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleUniversalGuardFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleUniversalGuardFactorySession struct {
	Contract     *SimpleUniversalGuardFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SimpleUniversalGuardFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleUniversalGuardFactoryCallerSession struct {
	Contract *SimpleUniversalGuardFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// SimpleUniversalGuardFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleUniversalGuardFactoryTransactorSession struct {
	Contract     *SimpleUniversalGuardFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// SimpleUniversalGuardFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleUniversalGuardFactoryRaw struct {
	Contract *SimpleUniversalGuardFactory // Generic contract binding to access the raw methods on
}

// SimpleUniversalGuardFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleUniversalGuardFactoryCallerRaw struct {
	Contract *SimpleUniversalGuardFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleUniversalGuardFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleUniversalGuardFactoryTransactorRaw struct {
	Contract *SimpleUniversalGuardFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleUniversalGuardFactory creates a new instance of SimpleUniversalGuardFactory, bound to a specific deployed contract.
func NewSimpleUniversalGuardFactory(address common.Address, backend bind.ContractBackend) (*SimpleUniversalGuardFactory, error) {
	contract, err := bindSimpleUniversalGuardFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleUniversalGuardFactory{SimpleUniversalGuardFactoryCaller: SimpleUniversalGuardFactoryCaller{contract: contract}, SimpleUniversalGuardFactoryTransactor: SimpleUniversalGuardFactoryTransactor{contract: contract}, SimpleUniversalGuardFactoryFilterer: SimpleUniversalGuardFactoryFilterer{contract: contract}}, nil
}

// NewSimpleUniversalGuardFactoryCaller creates a new read-only instance of SimpleUniversalGuardFactory, bound to a specific deployed contract.
func NewSimpleUniversalGuardFactoryCaller(address common.Address, caller bind.ContractCaller) (*SimpleUniversalGuardFactoryCaller, error) {
	contract, err := bindSimpleUniversalGuardFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleUniversalGuardFactoryCaller{contract: contract}, nil
}

// NewSimpleUniversalGuardFactoryTransactor creates a new write-only instance of SimpleUniversalGuardFactory, bound to a specific deployed contract.
func NewSimpleUniversalGuardFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleUniversalGuardFactoryTransactor, error) {
	contract, err := bindSimpleUniversalGuardFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleUniversalGuardFactoryTransactor{contract: contract}, nil
}

// NewSimpleUniversalGuardFactoryFilterer creates a new log filterer instance of SimpleUniversalGuardFactory, bound to a specific deployed contract.
func NewSimpleUniversalGuardFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleUniversalGuardFactoryFilterer, error) {
	contract, err := bindSimpleUniversalGuardFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleUniversalGuardFactoryFilterer{contract: contract}, nil
}

// bindSimpleUniversalGuardFactory binds a generic wrapper to an already deployed contract.
func bindSimpleUniversalGuardFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleUniversalGuardFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleUniversalGuardFactory.Contract.SimpleUniversalGuardFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.Contract.SimpleUniversalGuardFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.Contract.SimpleUniversalGuardFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleUniversalGuardFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.Contract.contract.Transact(opts, method, params...)
}

// GetGuard is a free data retrieval call binding the contract method 0x7027ea6c.
//
// Solidity: function getGuard(address requiredSigner) view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCaller) GetGuard(opts *bind.CallOpts, requiredSigner common.Address) (common.Address, error) {
	var out []interface{}
	err := _SimpleUniversalGuardFactory.contract.Call(opts, &out, "getGuard", requiredSigner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0x7027ea6c.
//
// Solidity: function getGuard(address requiredSigner) view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactorySession) GetGuard(requiredSigner common.Address) (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.GetGuard(&_SimpleUniversalGuardFactory.CallOpts, requiredSigner)
}

// GetGuard is a free data retrieval call binding the contract method 0x7027ea6c.
//
// Solidity: function getGuard(address requiredSigner) view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCallerSession) GetGuard(requiredSigner common.Address) (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.GetGuard(&_SimpleUniversalGuardFactory.CallOpts, requiredSigner)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleUniversalGuardFactory.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactorySession) GetImplementation() (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.GetImplementation(&_SimpleUniversalGuardFactory.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCallerSession) GetImplementation() (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.GetImplementation(&_SimpleUniversalGuardFactory.CallOpts)
}

// Guards is a free data retrieval call binding the contract method 0x6b462151.
//
// Solidity: function guards(address ) view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCaller) Guards(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SimpleUniversalGuardFactory.contract.Call(opts, &out, "guards", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guards is a free data retrieval call binding the contract method 0x6b462151.
//
// Solidity: function guards(address ) view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactorySession) Guards(arg0 common.Address) (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.Guards(&_SimpleUniversalGuardFactory.CallOpts, arg0)
}

// Guards is a free data retrieval call binding the contract method 0x6b462151.
//
// Solidity: function guards(address ) view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCallerSession) Guards(arg0 common.Address) (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.Guards(&_SimpleUniversalGuardFactory.CallOpts, arg0)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleUniversalGuardFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactorySession) Implementation() (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.Implementation(&_SimpleUniversalGuardFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCallerSession) Implementation() (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.Implementation(&_SimpleUniversalGuardFactory.CallOpts)
}

// PredictGuardAddress is a free data retrieval call binding the contract method 0x500b4f54.
//
// Solidity: function predictGuardAddress(address requiredSigner) view returns(address predicted)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCaller) PredictGuardAddress(opts *bind.CallOpts, requiredSigner common.Address) (common.Address, error) {
	var out []interface{}
	err := _SimpleUniversalGuardFactory.contract.Call(opts, &out, "predictGuardAddress", requiredSigner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictGuardAddress is a free data retrieval call binding the contract method 0x500b4f54.
//
// Solidity: function predictGuardAddress(address requiredSigner) view returns(address predicted)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactorySession) PredictGuardAddress(requiredSigner common.Address) (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.PredictGuardAddress(&_SimpleUniversalGuardFactory.CallOpts, requiredSigner)
}

// PredictGuardAddress is a free data retrieval call binding the contract method 0x500b4f54.
//
// Solidity: function predictGuardAddress(address requiredSigner) view returns(address predicted)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryCallerSession) PredictGuardAddress(requiredSigner common.Address) (common.Address, error) {
	return _SimpleUniversalGuardFactory.Contract.PredictGuardAddress(&_SimpleUniversalGuardFactory.CallOpts, requiredSigner)
}

// CreateGuard is a paid mutator transaction binding the contract method 0xc4c8a628.
//
// Solidity: function createGuard(address requiredSigner) returns(address guard)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryTransactor) CreateGuard(opts *bind.TransactOpts, requiredSigner common.Address) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.contract.Transact(opts, "createGuard", requiredSigner)
}

// CreateGuard is a paid mutator transaction binding the contract method 0xc4c8a628.
//
// Solidity: function createGuard(address requiredSigner) returns(address guard)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactorySession) CreateGuard(requiredSigner common.Address) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.Contract.CreateGuard(&_SimpleUniversalGuardFactory.TransactOpts, requiredSigner)
}

// CreateGuard is a paid mutator transaction binding the contract method 0xc4c8a628.
//
// Solidity: function createGuard(address requiredSigner) returns(address guard)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryTransactorSession) CreateGuard(requiredSigner common.Address) (*types.Transaction, error) {
	return _SimpleUniversalGuardFactory.Contract.CreateGuard(&_SimpleUniversalGuardFactory.TransactOpts, requiredSigner)
}

// SimpleUniversalGuardFactoryGuardCreatedIterator is returned from FilterGuardCreated and is used to iterate over the raw logs and unpacked data for GuardCreated events raised by the SimpleUniversalGuardFactory contract.
type SimpleUniversalGuardFactoryGuardCreatedIterator struct {
	Event *SimpleUniversalGuardFactoryGuardCreated // Event containing the contract specifics and raw log

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
func (it *SimpleUniversalGuardFactoryGuardCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleUniversalGuardFactoryGuardCreated)
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
		it.Event = new(SimpleUniversalGuardFactoryGuardCreated)
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
func (it *SimpleUniversalGuardFactoryGuardCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleUniversalGuardFactoryGuardCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleUniversalGuardFactoryGuardCreated represents a GuardCreated event raised by the SimpleUniversalGuardFactory contract.
type SimpleUniversalGuardFactoryGuardCreated struct {
	Guard          common.Address
	RequiredSigner common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGuardCreated is a free log retrieval operation binding the contract event 0x0a81e862badd25ab8f33c54bb33582c2625144e5f5bd92e61d0111895b6d3188.
//
// Solidity: event GuardCreated(address indexed guard, address indexed requiredSigner)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryFilterer) FilterGuardCreated(opts *bind.FilterOpts, guard []common.Address, requiredSigner []common.Address) (*SimpleUniversalGuardFactoryGuardCreatedIterator, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}
	var requiredSignerRule []interface{}
	for _, requiredSignerItem := range requiredSigner {
		requiredSignerRule = append(requiredSignerRule, requiredSignerItem)
	}

	logs, sub, err := _SimpleUniversalGuardFactory.contract.FilterLogs(opts, "GuardCreated", guardRule, requiredSignerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleUniversalGuardFactoryGuardCreatedIterator{contract: _SimpleUniversalGuardFactory.contract, event: "GuardCreated", logs: logs, sub: sub}, nil
}

// WatchGuardCreated is a free log subscription operation binding the contract event 0x0a81e862badd25ab8f33c54bb33582c2625144e5f5bd92e61d0111895b6d3188.
//
// Solidity: event GuardCreated(address indexed guard, address indexed requiredSigner)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryFilterer) WatchGuardCreated(opts *bind.WatchOpts, sink chan<- *SimpleUniversalGuardFactoryGuardCreated, guard []common.Address, requiredSigner []common.Address) (event.Subscription, error) {

	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}
	var requiredSignerRule []interface{}
	for _, requiredSignerItem := range requiredSigner {
		requiredSignerRule = append(requiredSignerRule, requiredSignerItem)
	}

	logs, sub, err := _SimpleUniversalGuardFactory.contract.WatchLogs(opts, "GuardCreated", guardRule, requiredSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleUniversalGuardFactoryGuardCreated)
				if err := _SimpleUniversalGuardFactory.contract.UnpackLog(event, "GuardCreated", log); err != nil {
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

// ParseGuardCreated is a log parse operation binding the contract event 0x0a81e862badd25ab8f33c54bb33582c2625144e5f5bd92e61d0111895b6d3188.
//
// Solidity: event GuardCreated(address indexed guard, address indexed requiredSigner)
func (_SimpleUniversalGuardFactory *SimpleUniversalGuardFactoryFilterer) ParseGuardCreated(log types.Log) (*SimpleUniversalGuardFactoryGuardCreated, error) {
	event := new(SimpleUniversalGuardFactoryGuardCreated)
	if err := _SimpleUniversalGuardFactory.contract.UnpackLog(event, "GuardCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
