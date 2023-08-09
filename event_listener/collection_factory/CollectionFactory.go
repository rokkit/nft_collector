// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package collection_factory

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

// CollectionFactoryMetaData contains all meta data concerning the CollectionFactory contract.
var CollectionFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"CollectionCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"createCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CollectionFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CollectionFactoryMetaData.ABI instead.
var CollectionFactoryABI = CollectionFactoryMetaData.ABI

// CollectionFactory is an auto generated Go binding around an Ethereum contract.
type CollectionFactory struct {
	CollectionFactoryCaller     // Read-only binding to the contract
	CollectionFactoryTransactor // Write-only binding to the contract
	CollectionFactoryFilterer   // Log filterer for contract events
}

// CollectionFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionFactorySession struct {
	Contract     *CollectionFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CollectionFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionFactoryCallerSession struct {
	Contract *CollectionFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CollectionFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionFactoryTransactorSession struct {
	Contract     *CollectionFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CollectionFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionFactoryRaw struct {
	Contract *CollectionFactory // Generic contract binding to access the raw methods on
}

// CollectionFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionFactoryCallerRaw struct {
	Contract *CollectionFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionFactoryTransactorRaw struct {
	Contract *CollectionFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionFactory creates a new instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactory(address common.Address, backend bind.ContractBackend) (*CollectionFactory, error) {
	contract, err := bindCollectionFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionFactory{CollectionFactoryCaller: CollectionFactoryCaller{contract: contract}, CollectionFactoryTransactor: CollectionFactoryTransactor{contract: contract}, CollectionFactoryFilterer: CollectionFactoryFilterer{contract: contract}}, nil
}

// NewCollectionFactoryCaller creates a new read-only instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactoryCaller(address common.Address, caller bind.ContractCaller) (*CollectionFactoryCaller, error) {
	contract, err := bindCollectionFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryCaller{contract: contract}, nil
}

// NewCollectionFactoryTransactor creates a new write-only instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionFactoryTransactor, error) {
	contract, err := bindCollectionFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryTransactor{contract: contract}, nil
}

// NewCollectionFactoryFilterer creates a new log filterer instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionFactoryFilterer, error) {
	contract, err := bindCollectionFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryFilterer{contract: contract}, nil
}

// bindCollectionFactory binds a generic wrapper to an already deployed contract.
func bindCollectionFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollectionFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionFactory *CollectionFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionFactory.Contract.CollectionFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionFactory *CollectionFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CollectionFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionFactory *CollectionFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CollectionFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionFactory *CollectionFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionFactory *CollectionFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionFactory *CollectionFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x176b9b75.
//
// Solidity: function createCollection(string name, string symbol) returns(address)
func (_CollectionFactory *CollectionFactoryTransactor) CreateCollection(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _CollectionFactory.contract.Transact(opts, "createCollection", name, symbol)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x176b9b75.
//
// Solidity: function createCollection(string name, string symbol) returns(address)
func (_CollectionFactory *CollectionFactorySession) CreateCollection(name string, symbol string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CreateCollection(&_CollectionFactory.TransactOpts, name, symbol)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x176b9b75.
//
// Solidity: function createCollection(string name, string symbol) returns(address)
func (_CollectionFactory *CollectionFactoryTransactorSession) CreateCollection(name string, symbol string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CreateCollection(&_CollectionFactory.TransactOpts, name, symbol)
}

// CollectionFactoryCollectionCreatedIterator is returned from FilterCollectionCreated and is used to iterate over the raw logs and unpacked data for CollectionCreated events raised by the CollectionFactory contract.
type CollectionFactoryCollectionCreatedIterator struct {
	Event *CollectionFactoryCollectionCreated // Event containing the contract specifics and raw log

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
func (it *CollectionFactoryCollectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionFactoryCollectionCreated)
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
		it.Event = new(CollectionFactoryCollectionCreated)
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
func (it *CollectionFactoryCollectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionFactoryCollectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionFactoryCollectionCreated represents a CollectionCreated event raised by the CollectionFactory contract.
type CollectionFactoryCollectionCreated struct {
	CollectionAddress common.Address
	Name              common.Hash
	Symbol            common.Hash
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollectionCreated is a free log retrieval operation binding the contract event 0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617.
//
// Solidity: event CollectionCreated(address indexed collectionAddress, string indexed name, string indexed symbol)
func (_CollectionFactory *CollectionFactoryFilterer) FilterCollectionCreated(opts *bind.FilterOpts, collectionAddress []common.Address, name []string, symbol []string) (*CollectionFactoryCollectionCreatedIterator, error) {

	var collectionAddressRule []interface{}
	for _, collectionAddressItem := range collectionAddress {
		collectionAddressRule = append(collectionAddressRule, collectionAddressItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _CollectionFactory.contract.FilterLogs(opts, "CollectionCreated", collectionAddressRule, nameRule, symbolRule)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryCollectionCreatedIterator{contract: _CollectionFactory.contract, event: "CollectionCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionCreated is a free log subscription operation binding the contract event 0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617.
//
// Solidity: event CollectionCreated(address indexed collectionAddress, string indexed name, string indexed symbol)
func (_CollectionFactory *CollectionFactoryFilterer) WatchCollectionCreated(opts *bind.WatchOpts, sink chan<- *CollectionFactoryCollectionCreated, collectionAddress []common.Address, name []string, symbol []string) (event.Subscription, error) {

	var collectionAddressRule []interface{}
	for _, collectionAddressItem := range collectionAddress {
		collectionAddressRule = append(collectionAddressRule, collectionAddressItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _CollectionFactory.contract.WatchLogs(opts, "CollectionCreated", collectionAddressRule, nameRule, symbolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionFactoryCollectionCreated)
				if err := _CollectionFactory.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
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

// ParseCollectionCreated is a log parse operation binding the contract event 0x3454b57f2dca4f5a54e8358d096ac9d1a0d2dab98991ddb89ff9ea1746260617.
//
// Solidity: event CollectionCreated(address indexed collectionAddress, string indexed name, string indexed symbol)
func (_CollectionFactory *CollectionFactoryFilterer) ParseCollectionCreated(log types.Log) (*CollectionFactoryCollectionCreated, error) {
	event := new(CollectionFactoryCollectionCreated)
	if err := _CollectionFactory.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
