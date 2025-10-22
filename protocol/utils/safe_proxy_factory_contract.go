// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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
	projectabi "github.com/vikkkko/safe-core-sdk-golang/abi"
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

// SafeProxyFactoryContractMetaData contains all meta data concerning the SafeProxyFactoryContract contract.
var SafeProxyFactoryContractMetaData = &bind.MetaData{
	ABI: string(projectabi.SafeProxyFactory),
}

// SafeProxyFactoryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeProxyFactoryContractMetaData.ABI instead.
var SafeProxyFactoryContractABI = SafeProxyFactoryContractMetaData.ABI

// SafeProxyFactoryContract is an auto generated Go binding around an Ethereum contract.
type SafeProxyFactoryContract struct {
	SafeProxyFactoryContractCaller     // Read-only binding to the contract
	SafeProxyFactoryContractTransactor // Write-only binding to the contract
	SafeProxyFactoryContractFilterer   // Log filterer for contract events
}

// SafeProxyFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeProxyFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeProxyFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeProxyFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeProxyFactoryContractSession struct {
	Contract     *SafeProxyFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SafeProxyFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeProxyFactoryContractCallerSession struct {
	Contract *SafeProxyFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SafeProxyFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeProxyFactoryContractTransactorSession struct {
	Contract     *SafeProxyFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SafeProxyFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeProxyFactoryContractRaw struct {
	Contract *SafeProxyFactoryContract // Generic contract binding to access the raw methods on
}

// SafeProxyFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeProxyFactoryContractCallerRaw struct {
	Contract *SafeProxyFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// SafeProxyFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeProxyFactoryContractTransactorRaw struct {
	Contract *SafeProxyFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeProxyFactoryContract creates a new instance of SafeProxyFactoryContract, bound to a specific deployed contract.
func NewSafeProxyFactoryContract(address common.Address, backend bind.ContractBackend) (*SafeProxyFactoryContract, error) {
	contract, err := bindSafeProxyFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContract{SafeProxyFactoryContractCaller: SafeProxyFactoryContractCaller{contract: contract}, SafeProxyFactoryContractTransactor: SafeProxyFactoryContractTransactor{contract: contract}, SafeProxyFactoryContractFilterer: SafeProxyFactoryContractFilterer{contract: contract}}, nil
}

// NewSafeProxyFactoryContractCaller creates a new read-only instance of SafeProxyFactoryContract, bound to a specific deployed contract.
func NewSafeProxyFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*SafeProxyFactoryContractCaller, error) {
	contract, err := bindSafeProxyFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContractCaller{contract: contract}, nil
}

// NewSafeProxyFactoryContractTransactor creates a new write-only instance of SafeProxyFactoryContract, bound to a specific deployed contract.
func NewSafeProxyFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeProxyFactoryContractTransactor, error) {
	contract, err := bindSafeProxyFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContractTransactor{contract: contract}, nil
}

// NewSafeProxyFactoryContractFilterer creates a new log filterer instance of SafeProxyFactoryContract, bound to a specific deployed contract.
func NewSafeProxyFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeProxyFactoryContractFilterer, error) {
	contract, err := bindSafeProxyFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContractFilterer{contract: contract}, nil
}

// bindSafeProxyFactoryContract binds a generic wrapper to an already deployed contract.
func bindSafeProxyFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeProxyFactoryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactoryContract *SafeProxyFactoryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactoryContract.Contract.SafeProxyFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactoryContract *SafeProxyFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.SafeProxyFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactoryContract *SafeProxyFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.SafeProxyFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeProxyFactoryContract.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactoryContract.Contract.GetChainId(&_SafeProxyFactoryContract.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCallerSession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactoryContract.Contract.GetChainId(&_SafeProxyFactoryContract.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SafeProxyFactoryContract.contract.Call(opts, &out, "proxyCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactoryContract.Contract.ProxyCreationCode(&_SafeProxyFactoryContract.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCallerSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactoryContract.Contract.ProxyCreationCode(&_SafeProxyFactoryContract.CallOpts)
}

// ProxyCreationCodehash is a free data retrieval call binding the contract method 0x1fdac6c5.
//
// Solidity: function proxyCreationCodehash(address singleton) pure returns(bytes32)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCaller) ProxyCreationCodehash(opts *bind.CallOpts, singleton common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SafeProxyFactoryContract.contract.Call(opts, &out, "proxyCreationCodehash", singleton)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxyCreationCodehash is a free data retrieval call binding the contract method 0x1fdac6c5.
//
// Solidity: function proxyCreationCodehash(address singleton) pure returns(bytes32)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) ProxyCreationCodehash(singleton common.Address) ([32]byte, error) {
	return _SafeProxyFactoryContract.Contract.ProxyCreationCodehash(&_SafeProxyFactoryContract.CallOpts, singleton)
}

// ProxyCreationCodehash is a free data retrieval call binding the contract method 0x1fdac6c5.
//
// Solidity: function proxyCreationCodehash(address singleton) pure returns(bytes32)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractCallerSession) ProxyCreationCodehash(singleton common.Address) ([32]byte, error) {
	return _SafeProxyFactoryContract.Contract.ProxyCreationCodehash(&_SafeProxyFactoryContract.CallOpts, singleton)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactor) CreateChainSpecificProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.contract.Transact(opts, "createChainSpecificProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactorSession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonceL2 is a paid mutator transaction binding the contract method 0x1a334b0b.
//
// Solidity: function createChainSpecificProxyWithNonceL2(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactor) CreateChainSpecificProxyWithNonceL2(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.contract.Transact(opts, "createChainSpecificProxyWithNonceL2", _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonceL2 is a paid mutator transaction binding the contract method 0x1a334b0b.
//
// Solidity: function createChainSpecificProxyWithNonceL2(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) CreateChainSpecificProxyWithNonceL2(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateChainSpecificProxyWithNonceL2(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonceL2 is a paid mutator transaction binding the contract method 0x1a334b0b.
//
// Solidity: function createChainSpecificProxyWithNonceL2(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactorSession) CreateChainSpecificProxyWithNonceL2(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateChainSpecificProxyWithNonceL2(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactor) CreateProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.contract.Transact(opts, "createProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateProxyWithNonce(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactorSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateProxyWithNonce(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonceL2 is a paid mutator transaction binding the contract method 0x4662f9d6.
//
// Solidity: function createProxyWithNonceL2(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactor) CreateProxyWithNonceL2(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.contract.Transact(opts, "createProxyWithNonceL2", _singleton, initializer, saltNonce)
}

// CreateProxyWithNonceL2 is a paid mutator transaction binding the contract method 0x4662f9d6.
//
// Solidity: function createProxyWithNonceL2(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractSession) CreateProxyWithNonceL2(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateProxyWithNonceL2(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonceL2 is a paid mutator transaction binding the contract method 0x4662f9d6.
//
// Solidity: function createProxyWithNonceL2(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractTransactorSession) CreateProxyWithNonceL2(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactoryContract.Contract.CreateProxyWithNonceL2(&_SafeProxyFactoryContract.TransactOpts, _singleton, initializer, saltNonce)
}

// SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator is returned from FilterChainSpecificProxyCreationL2 and is used to iterate over the raw logs and unpacked data for ChainSpecificProxyCreationL2 events raised by the SafeProxyFactoryContract contract.
type SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator struct {
	Event *SafeProxyFactoryContractChainSpecificProxyCreationL2 // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryContractChainSpecificProxyCreationL2)
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
		it.Event = new(SafeProxyFactoryContractChainSpecificProxyCreationL2)
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
func (it *SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryContractChainSpecificProxyCreationL2 represents a ChainSpecificProxyCreationL2 event raised by the SafeProxyFactoryContract contract.
type SafeProxyFactoryContractChainSpecificProxyCreationL2 struct {
	Proxy       common.Address
	Singleton   common.Address
	Initializer []byte
	SaltNonce   *big.Int
	ChainId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChainSpecificProxyCreationL2 is a free log retrieval operation binding the contract event 0x7a1c96f74273709828c0e67ab96189772005d3b4948bc6902693eb21e4e2cc6a.
//
// Solidity: event ChainSpecificProxyCreationL2(address indexed proxy, address singleton, bytes initializer, uint256 saltNonce, uint256 chainId)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) FilterChainSpecificProxyCreationL2(opts *bind.FilterOpts, proxy []common.Address) (*SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryContract.contract.FilterLogs(opts, "ChainSpecificProxyCreationL2", proxyRule)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContractChainSpecificProxyCreationL2Iterator{contract: _SafeProxyFactoryContract.contract, event: "ChainSpecificProxyCreationL2", logs: logs, sub: sub}, nil
}

// WatchChainSpecificProxyCreationL2 is a free log subscription operation binding the contract event 0x7a1c96f74273709828c0e67ab96189772005d3b4948bc6902693eb21e4e2cc6a.
//
// Solidity: event ChainSpecificProxyCreationL2(address indexed proxy, address singleton, bytes initializer, uint256 saltNonce, uint256 chainId)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) WatchChainSpecificProxyCreationL2(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryContractChainSpecificProxyCreationL2, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryContract.contract.WatchLogs(opts, "ChainSpecificProxyCreationL2", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryContractChainSpecificProxyCreationL2)
				if err := _SafeProxyFactoryContract.contract.UnpackLog(event, "ChainSpecificProxyCreationL2", log); err != nil {
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

// ParseChainSpecificProxyCreationL2 is a log parse operation binding the contract event 0x7a1c96f74273709828c0e67ab96189772005d3b4948bc6902693eb21e4e2cc6a.
//
// Solidity: event ChainSpecificProxyCreationL2(address indexed proxy, address singleton, bytes initializer, uint256 saltNonce, uint256 chainId)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) ParseChainSpecificProxyCreationL2(log types.Log) (*SafeProxyFactoryContractChainSpecificProxyCreationL2, error) {
	event := new(SafeProxyFactoryContractChainSpecificProxyCreationL2)
	if err := _SafeProxyFactoryContract.contract.UnpackLog(event, "ChainSpecificProxyCreationL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeProxyFactoryContractProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the SafeProxyFactoryContract contract.
type SafeProxyFactoryContractProxyCreationIterator struct {
	Event *SafeProxyFactoryContractProxyCreation // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryContractProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryContractProxyCreation)
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
		it.Event = new(SafeProxyFactoryContractProxyCreation)
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
func (it *SafeProxyFactoryContractProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryContractProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryContractProxyCreation represents a ProxyCreation event raised by the SafeProxyFactoryContract contract.
type SafeProxyFactoryContractProxyCreation struct {
	Proxy     common.Address
	Singleton common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) FilterProxyCreation(opts *bind.FilterOpts, proxy []common.Address) (*SafeProxyFactoryContractProxyCreationIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryContract.contract.FilterLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContractProxyCreationIterator{contract: _SafeProxyFactoryContract.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryContractProxyCreation, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryContract.contract.WatchLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryContractProxyCreation)
				if err := _SafeProxyFactoryContract.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) ParseProxyCreation(log types.Log) (*SafeProxyFactoryContractProxyCreation, error) {
	event := new(SafeProxyFactoryContractProxyCreation)
	if err := _SafeProxyFactoryContract.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeProxyFactoryContractProxyCreationL2Iterator is returned from FilterProxyCreationL2 and is used to iterate over the raw logs and unpacked data for ProxyCreationL2 events raised by the SafeProxyFactoryContract contract.
type SafeProxyFactoryContractProxyCreationL2Iterator struct {
	Event *SafeProxyFactoryContractProxyCreationL2 // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryContractProxyCreationL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryContractProxyCreationL2)
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
		it.Event = new(SafeProxyFactoryContractProxyCreationL2)
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
func (it *SafeProxyFactoryContractProxyCreationL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryContractProxyCreationL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryContractProxyCreationL2 represents a ProxyCreationL2 event raised by the SafeProxyFactoryContract contract.
type SafeProxyFactoryContractProxyCreationL2 struct {
	Proxy       common.Address
	Singleton   common.Address
	Initializer []byte
	SaltNonce   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProxyCreationL2 is a free log retrieval operation binding the contract event 0x6e78ae8c51534f0b801d7aa7cce5a2113f3c3368c61b06fa61415f361d400431.
//
// Solidity: event ProxyCreationL2(address indexed proxy, address singleton, bytes initializer, uint256 saltNonce)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) FilterProxyCreationL2(opts *bind.FilterOpts, proxy []common.Address) (*SafeProxyFactoryContractProxyCreationL2Iterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryContract.contract.FilterLogs(opts, "ProxyCreationL2", proxyRule)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryContractProxyCreationL2Iterator{contract: _SafeProxyFactoryContract.contract, event: "ProxyCreationL2", logs: logs, sub: sub}, nil
}

// WatchProxyCreationL2 is a free log subscription operation binding the contract event 0x6e78ae8c51534f0b801d7aa7cce5a2113f3c3368c61b06fa61415f361d400431.
//
// Solidity: event ProxyCreationL2(address indexed proxy, address singleton, bytes initializer, uint256 saltNonce)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) WatchProxyCreationL2(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryContractProxyCreationL2, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactoryContract.contract.WatchLogs(opts, "ProxyCreationL2", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryContractProxyCreationL2)
				if err := _SafeProxyFactoryContract.contract.UnpackLog(event, "ProxyCreationL2", log); err != nil {
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

// ParseProxyCreationL2 is a log parse operation binding the contract event 0x6e78ae8c51534f0b801d7aa7cce5a2113f3c3368c61b06fa61415f361d400431.
//
// Solidity: event ProxyCreationL2(address indexed proxy, address singleton, bytes initializer, uint256 saltNonce)
func (_SafeProxyFactoryContract *SafeProxyFactoryContractFilterer) ParseProxyCreationL2(log types.Log) (*SafeProxyFactoryContractProxyCreationL2, error) {
	event := new(SafeProxyFactoryContractProxyCreationL2)
	if err := _SafeProxyFactoryContract.contract.UnpackLog(event, "ProxyCreationL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
