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
)

// EnsRegistryMetaData contains all meta data concerning the EnsRegistry contract.
var EnsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"ttl\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"NewTTL\",\"type\":\"event\"}]",
}

// EnsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EnsRegistryMetaData.ABI instead.
var EnsRegistryABI = EnsRegistryMetaData.ABI

// EnsRegistry is an auto generated Go binding around an Ethereum contract.
type EnsRegistry struct {
	EnsRegistryCaller     // Read-only binding to the contract
	EnsRegistryTransactor // Write-only binding to the contract
	EnsRegistryFilterer   // Log filterer for contract events
}

// EnsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnsRegistrySession struct {
	Contract     *EnsRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnsRegistryCallerSession struct {
	Contract *EnsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EnsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnsRegistryTransactorSession struct {
	Contract     *EnsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EnsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnsRegistryRaw struct {
	Contract *EnsRegistry // Generic contract binding to access the raw methods on
}

// EnsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnsRegistryCallerRaw struct {
	Contract *EnsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EnsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnsRegistryTransactorRaw struct {
	Contract *EnsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnsRegistry creates a new instance of EnsRegistry, bound to a specific deployed contract.
func NewEnsRegistry(address common.Address, backend bind.ContractBackend) (*EnsRegistry, error) {
	contract, err := bindEnsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnsRegistry{EnsRegistryCaller: EnsRegistryCaller{contract: contract}, EnsRegistryTransactor: EnsRegistryTransactor{contract: contract}, EnsRegistryFilterer: EnsRegistryFilterer{contract: contract}}, nil
}

// NewEnsRegistryCaller creates a new read-only instance of EnsRegistry, bound to a specific deployed contract.
func NewEnsRegistryCaller(address common.Address, caller bind.ContractCaller) (*EnsRegistryCaller, error) {
	contract, err := bindEnsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryCaller{contract: contract}, nil
}

// NewEnsRegistryTransactor creates a new write-only instance of EnsRegistry, bound to a specific deployed contract.
func NewEnsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EnsRegistryTransactor, error) {
	contract, err := bindEnsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryTransactor{contract: contract}, nil
}

// NewEnsRegistryFilterer creates a new log filterer instance of EnsRegistry, bound to a specific deployed contract.
func NewEnsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EnsRegistryFilterer, error) {
	contract, err := bindEnsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryFilterer{contract: contract}, nil
}

// bindEnsRegistry binds a generic wrapper to an already deployed contract.
func bindEnsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnsRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnsRegistry *EnsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnsRegistry.Contract.EnsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnsRegistry *EnsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnsRegistry.Contract.EnsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnsRegistry *EnsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnsRegistry.Contract.EnsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnsRegistry *EnsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnsRegistry *EnsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnsRegistry *EnsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnsRegistry.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) returns(address)
func (_EnsRegistry *EnsRegistryCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EnsRegistry.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) returns(address)
func (_EnsRegistry *EnsRegistrySession) Owner(node [32]byte) (common.Address, error) {
	return _EnsRegistry.Contract.Owner(&_EnsRegistry.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) returns(address)
func (_EnsRegistry *EnsRegistryCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _EnsRegistry.Contract.Owner(&_EnsRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) returns(address)
func (_EnsRegistry *EnsRegistryCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EnsRegistry.contract.Call(opts, &out, "resolver", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) returns(address)
func (_EnsRegistry *EnsRegistrySession) Resolver(node [32]byte) (common.Address, error) {
	return _EnsRegistry.Contract.Resolver(&_EnsRegistry.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(bytes32 node) returns(address)
func (_EnsRegistry *EnsRegistryCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _EnsRegistry.Contract.Resolver(&_EnsRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) returns(uint64)
func (_EnsRegistry *EnsRegistryCaller) Ttl(opts *bind.CallOpts, node [32]byte) (uint64, error) {
	var out []interface{}
	err := _EnsRegistry.contract.Call(opts, &out, "ttl", node)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) returns(uint64)
func (_EnsRegistry *EnsRegistrySession) Ttl(node [32]byte) (uint64, error) {
	return _EnsRegistry.Contract.Ttl(&_EnsRegistry.CallOpts, node)
}

// Ttl is a free data retrieval call binding the contract method 0x16a25cbd.
//
// Solidity: function ttl(bytes32 node) returns(uint64)
func (_EnsRegistry *EnsRegistryCallerSession) Ttl(node [32]byte) (uint64, error) {
	return _EnsRegistry.Contract.Ttl(&_EnsRegistry.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_EnsRegistry *EnsRegistryTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistry.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_EnsRegistry *EnsRegistrySession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetOwner(&_EnsRegistry.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(bytes32 node, address owner) returns()
func (_EnsRegistry *EnsRegistryTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetOwner(&_EnsRegistry.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_EnsRegistry *EnsRegistryTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _EnsRegistry.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_EnsRegistry *EnsRegistrySession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetResolver(&_EnsRegistry.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_EnsRegistry *EnsRegistryTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetResolver(&_EnsRegistry.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns()
func (_EnsRegistry *EnsRegistryTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistry.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns()
func (_EnsRegistry *EnsRegistrySession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetSubnodeOwner(&_EnsRegistry.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(bytes32 node, bytes32 label, address owner) returns()
func (_EnsRegistry *EnsRegistryTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetSubnodeOwner(&_EnsRegistry.TransactOpts, node, label, owner)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_EnsRegistry *EnsRegistryTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistry.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_EnsRegistry *EnsRegistrySession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetTTL(&_EnsRegistry.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_EnsRegistry *EnsRegistryTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _EnsRegistry.Contract.SetTTL(&_EnsRegistry.TransactOpts, node, ttl)
}

// EnsRegistryNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the EnsRegistry contract.
type EnsRegistryNewOwnerIterator struct {
	Event *EnsRegistryNewOwner // Event containing the contract specifics and raw log

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
func (it *EnsRegistryNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryNewOwner)
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
		it.Event = new(EnsRegistryNewOwner)
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
func (it *EnsRegistryNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryNewOwner represents a NewOwner event raised by the EnsRegistry contract.
type EnsRegistryNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_EnsRegistry *EnsRegistryFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*EnsRegistryNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _EnsRegistry.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryNewOwnerIterator{contract: _EnsRegistry.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_EnsRegistry *EnsRegistryFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *EnsRegistryNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _EnsRegistry.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryNewOwner)
				if err := _EnsRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
func (_EnsRegistry *EnsRegistryFilterer) ParseNewOwner(log types.Log) (*EnsRegistryNewOwner, error) {
	event := new(EnsRegistryNewOwner)
	if err := _EnsRegistry.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the EnsRegistry contract.
type EnsRegistryNewResolverIterator struct {
	Event *EnsRegistryNewResolver // Event containing the contract specifics and raw log

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
func (it *EnsRegistryNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryNewResolver)
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
		it.Event = new(EnsRegistryNewResolver)
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
func (it *EnsRegistryNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryNewResolver represents a NewResolver event raised by the EnsRegistry contract.
type EnsRegistryNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_EnsRegistry *EnsRegistryFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*EnsRegistryNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistry.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryNewResolverIterator{contract: _EnsRegistry.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_EnsRegistry *EnsRegistryFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *EnsRegistryNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistry.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryNewResolver)
				if err := _EnsRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ParseNewResolver is a log parse operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: event NewResolver(bytes32 indexed node, address resolver)
func (_EnsRegistry *EnsRegistryFilterer) ParseNewResolver(log types.Log) (*EnsRegistryNewResolver, error) {
	event := new(EnsRegistryNewResolver)
	if err := _EnsRegistry.contract.UnpackLog(event, "NewResolver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryNewTTLIterator is returned from FilterNewTTL and is used to iterate over the raw logs and unpacked data for NewTTL events raised by the EnsRegistry contract.
type EnsRegistryNewTTLIterator struct {
	Event *EnsRegistryNewTTL // Event containing the contract specifics and raw log

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
func (it *EnsRegistryNewTTLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryNewTTL)
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
		it.Event = new(EnsRegistryNewTTL)
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
func (it *EnsRegistryNewTTLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryNewTTLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryNewTTL represents a NewTTL event raised by the EnsRegistry contract.
type EnsRegistryNewTTL struct {
	Node [32]byte
	Ttl  uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTTL is a free log retrieval operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_EnsRegistry *EnsRegistryFilterer) FilterNewTTL(opts *bind.FilterOpts, node [][32]byte) (*EnsRegistryNewTTLIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistry.contract.FilterLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryNewTTLIterator{contract: _EnsRegistry.contract, event: "NewTTL", logs: logs, sub: sub}, nil
}

// WatchNewTTL is a free log subscription operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_EnsRegistry *EnsRegistryFilterer) WatchNewTTL(opts *bind.WatchOpts, sink chan<- *EnsRegistryNewTTL, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistry.contract.WatchLogs(opts, "NewTTL", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryNewTTL)
				if err := _EnsRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
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

// ParseNewTTL is a log parse operation binding the contract event 0x1d4f9bbfc9cab89d66e1a1562f2233ccbf1308cb4f63de2ead5787adddb8fa68.
//
// Solidity: event NewTTL(bytes32 indexed node, uint64 ttl)
func (_EnsRegistry *EnsRegistryFilterer) ParseNewTTL(log types.Log) (*EnsRegistryNewTTL, error) {
	event := new(EnsRegistryNewTTL)
	if err := _EnsRegistry.contract.UnpackLog(event, "NewTTL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnsRegistryTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EnsRegistry contract.
type EnsRegistryTransferIterator struct {
	Event *EnsRegistryTransfer // Event containing the contract specifics and raw log

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
func (it *EnsRegistryTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnsRegistryTransfer)
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
		it.Event = new(EnsRegistryTransfer)
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
func (it *EnsRegistryTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnsRegistryTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnsRegistryTransfer represents a Transfer event raised by the EnsRegistry contract.
type EnsRegistryTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_EnsRegistry *EnsRegistryFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*EnsRegistryTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistry.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &EnsRegistryTransferIterator{contract: _EnsRegistry.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_EnsRegistry *EnsRegistryFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EnsRegistryTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _EnsRegistry.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnsRegistryTransfer)
				if err := _EnsRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: event Transfer(bytes32 indexed node, address owner)
func (_EnsRegistry *EnsRegistryFilterer) ParseTransfer(log types.Log) (*EnsRegistryTransfer, error) {
	event := new(EnsRegistryTransfer)
	if err := _EnsRegistry.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
