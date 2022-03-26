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

// IEtherAlleyStoreTokenListing is an auto generated low-level Go binding around an user-defined struct.
type IEtherAlleyStoreTokenListing struct {
	Purchasable  bool
	Transferable bool
	Price        *big.Int
	SupplyLimit  *big.Int
	BalanceLimit *big.Int
	Supply       *big.Int
}

// EtherAlleyStoreMetaData contains all meta data concerning the EtherAlleyStore contract.
var EtherAlleyStoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"purchasable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"ListingChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"purchasable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"internalType\":\"structIEtherAlleyStore.TokenListing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"getListingBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"purchasable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"internalType\":\"structIEtherAlleyStore.TokenListing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"purchase\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"purchaseBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"purchasable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceLimit\",\"type\":\"uint256\"}],\"name\":\"setListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EtherAlleyStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherAlleyStoreMetaData.ABI instead.
var EtherAlleyStoreABI = EtherAlleyStoreMetaData.ABI

// EtherAlleyStore is an auto generated Go binding around an Ethereum contract.
type EtherAlleyStore struct {
	EtherAlleyStoreCaller     // Read-only binding to the contract
	EtherAlleyStoreTransactor // Write-only binding to the contract
	EtherAlleyStoreFilterer   // Log filterer for contract events
}

// EtherAlleyStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherAlleyStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherAlleyStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherAlleyStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherAlleyStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherAlleyStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherAlleyStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherAlleyStoreSession struct {
	Contract     *EtherAlleyStore  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherAlleyStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherAlleyStoreCallerSession struct {
	Contract *EtherAlleyStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EtherAlleyStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherAlleyStoreTransactorSession struct {
	Contract     *EtherAlleyStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EtherAlleyStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherAlleyStoreRaw struct {
	Contract *EtherAlleyStore // Generic contract binding to access the raw methods on
}

// EtherAlleyStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherAlleyStoreCallerRaw struct {
	Contract *EtherAlleyStoreCaller // Generic read-only contract binding to access the raw methods on
}

// EtherAlleyStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherAlleyStoreTransactorRaw struct {
	Contract *EtherAlleyStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherAlleyStore creates a new instance of EtherAlleyStore, bound to a specific deployed contract.
func NewEtherAlleyStore(address common.Address, backend bind.ContractBackend) (*EtherAlleyStore, error) {
	contract, err := bindEtherAlleyStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStore{EtherAlleyStoreCaller: EtherAlleyStoreCaller{contract: contract}, EtherAlleyStoreTransactor: EtherAlleyStoreTransactor{contract: contract}, EtherAlleyStoreFilterer: EtherAlleyStoreFilterer{contract: contract}}, nil
}

// NewEtherAlleyStoreCaller creates a new read-only instance of EtherAlleyStore, bound to a specific deployed contract.
func NewEtherAlleyStoreCaller(address common.Address, caller bind.ContractCaller) (*EtherAlleyStoreCaller, error) {
	contract, err := bindEtherAlleyStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreCaller{contract: contract}, nil
}

// NewEtherAlleyStoreTransactor creates a new write-only instance of EtherAlleyStore, bound to a specific deployed contract.
func NewEtherAlleyStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherAlleyStoreTransactor, error) {
	contract, err := bindEtherAlleyStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreTransactor{contract: contract}, nil
}

// NewEtherAlleyStoreFilterer creates a new log filterer instance of EtherAlleyStore, bound to a specific deployed contract.
func NewEtherAlleyStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherAlleyStoreFilterer, error) {
	contract, err := bindEtherAlleyStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreFilterer{contract: contract}, nil
}

// bindEtherAlleyStore binds a generic wrapper to an already deployed contract.
func bindEtherAlleyStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherAlleyStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherAlleyStore *EtherAlleyStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherAlleyStore.Contract.EtherAlleyStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherAlleyStore *EtherAlleyStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.EtherAlleyStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherAlleyStore *EtherAlleyStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.EtherAlleyStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherAlleyStore *EtherAlleyStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherAlleyStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherAlleyStore *EtherAlleyStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherAlleyStore *EtherAlleyStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_EtherAlleyStore *EtherAlleyStoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_EtherAlleyStore *EtherAlleyStoreSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _EtherAlleyStore.Contract.BalanceOf(&_EtherAlleyStore.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _EtherAlleyStore.Contract.BalanceOf(&_EtherAlleyStore.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_EtherAlleyStore *EtherAlleyStoreCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_EtherAlleyStore *EtherAlleyStoreSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _EtherAlleyStore.Contract.BalanceOfBatch(&_EtherAlleyStore.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _EtherAlleyStore.Contract.BalanceOfBatch(&_EtherAlleyStore.CallOpts, accounts, ids)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 id) view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_EtherAlleyStore *EtherAlleyStoreCaller) GetListing(opts *bind.CallOpts, id *big.Int) (IEtherAlleyStoreTokenListing, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "getListing", id)

	if err != nil {
		return *new(IEtherAlleyStoreTokenListing), err
	}

	out0 := *abi.ConvertType(out[0], new(IEtherAlleyStoreTokenListing)).(*IEtherAlleyStoreTokenListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 id) view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_EtherAlleyStore *EtherAlleyStoreSession) GetListing(id *big.Int) (IEtherAlleyStoreTokenListing, error) {
	return _EtherAlleyStore.Contract.GetListing(&_EtherAlleyStore.CallOpts, id)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 id) view returns((bool,bool,uint256,uint256,uint256,uint256))
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) GetListing(id *big.Int) (IEtherAlleyStoreTokenListing, error) {
	return _EtherAlleyStore.Contract.GetListing(&_EtherAlleyStore.CallOpts, id)
}

// GetListingBatch is a free data retrieval call binding the contract method 0x45384ec3.
//
// Solidity: function getListingBatch(uint256[] ids) view returns((bool,bool,uint256,uint256,uint256,uint256)[])
func (_EtherAlleyStore *EtherAlleyStoreCaller) GetListingBatch(opts *bind.CallOpts, ids []*big.Int) ([]IEtherAlleyStoreTokenListing, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "getListingBatch", ids)

	if err != nil {
		return *new([]IEtherAlleyStoreTokenListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]IEtherAlleyStoreTokenListing)).(*[]IEtherAlleyStoreTokenListing)

	return out0, err

}

// GetListingBatch is a free data retrieval call binding the contract method 0x45384ec3.
//
// Solidity: function getListingBatch(uint256[] ids) view returns((bool,bool,uint256,uint256,uint256,uint256)[])
func (_EtherAlleyStore *EtherAlleyStoreSession) GetListingBatch(ids []*big.Int) ([]IEtherAlleyStoreTokenListing, error) {
	return _EtherAlleyStore.Contract.GetListingBatch(&_EtherAlleyStore.CallOpts, ids)
}

// GetListingBatch is a free data retrieval call binding the contract method 0x45384ec3.
//
// Solidity: function getListingBatch(uint256[] ids) view returns((bool,bool,uint256,uint256,uint256,uint256)[])
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) GetListingBatch(ids []*big.Int) ([]IEtherAlleyStoreTokenListing, error) {
	return _EtherAlleyStore.Contract.GetListingBatch(&_EtherAlleyStore.CallOpts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_EtherAlleyStore *EtherAlleyStoreCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_EtherAlleyStore *EtherAlleyStoreSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _EtherAlleyStore.Contract.IsApprovedForAll(&_EtherAlleyStore.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _EtherAlleyStore.Contract.IsApprovedForAll(&_EtherAlleyStore.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EtherAlleyStore *EtherAlleyStoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EtherAlleyStore *EtherAlleyStoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherAlleyStore.Contract.SupportsInterface(&_EtherAlleyStore.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherAlleyStore.Contract.SupportsInterface(&_EtherAlleyStore.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_EtherAlleyStore *EtherAlleyStoreCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _EtherAlleyStore.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_EtherAlleyStore *EtherAlleyStoreSession) Uri(id *big.Int) (string, error) {
	return _EtherAlleyStore.Contract.Uri(&_EtherAlleyStore.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_EtherAlleyStore *EtherAlleyStoreCallerSession) Uri(id *big.Int) (string, error) {
	return _EtherAlleyStore.Contract.Uri(&_EtherAlleyStore.CallOpts, id)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) Pause() (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.Pause(&_EtherAlleyStore.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) Pause() (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.Pause(&_EtherAlleyStore.TransactOpts)
}

// Purchase is a paid mutator transaction binding the contract method 0xeea3ea3f.
//
// Solidity: function purchase(address account, uint256 id, uint256 amount, bytes data) payable returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) Purchase(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "purchase", account, id, amount, data)
}

// Purchase is a paid mutator transaction binding the contract method 0xeea3ea3f.
//
// Solidity: function purchase(address account, uint256 id, uint256 amount, bytes data) payable returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) Purchase(account common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.Purchase(&_EtherAlleyStore.TransactOpts, account, id, amount, data)
}

// Purchase is a paid mutator transaction binding the contract method 0xeea3ea3f.
//
// Solidity: function purchase(address account, uint256 id, uint256 amount, bytes data) payable returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) Purchase(account common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.Purchase(&_EtherAlleyStore.TransactOpts, account, id, amount, data)
}

// PurchaseBatch is a paid mutator transaction binding the contract method 0x8b2be628.
//
// Solidity: function purchaseBatch(address to, uint256[] ids, uint256[] amounts, bytes data) payable returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) PurchaseBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "purchaseBatch", to, ids, amounts, data)
}

// PurchaseBatch is a paid mutator transaction binding the contract method 0x8b2be628.
//
// Solidity: function purchaseBatch(address to, uint256[] ids, uint256[] amounts, bytes data) payable returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) PurchaseBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.PurchaseBatch(&_EtherAlleyStore.TransactOpts, to, ids, amounts, data)
}

// PurchaseBatch is a paid mutator transaction binding the contract method 0x8b2be628.
//
// Solidity: function purchaseBatch(address to, uint256[] ids, uint256[] amounts, bytes data) payable returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) PurchaseBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.PurchaseBatch(&_EtherAlleyStore.TransactOpts, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SafeBatchTransferFrom(&_EtherAlleyStore.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SafeBatchTransferFrom(&_EtherAlleyStore.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SafeTransferFrom(&_EtherAlleyStore.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SafeTransferFrom(&_EtherAlleyStore.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SetApprovalForAll(&_EtherAlleyStore.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SetApprovalForAll(&_EtherAlleyStore.TransactOpts, operator, approved)
}

// SetListing is a paid mutator transaction binding the contract method 0x2fe8a4c3.
//
// Solidity: function setListing(uint256 id, bool purchasable, bool transferable, uint256 price, uint256 supplyLimit, uint256 balanceLimit) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) SetListing(opts *bind.TransactOpts, id *big.Int, purchasable bool, transferable bool, price *big.Int, supplyLimit *big.Int, balanceLimit *big.Int) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "setListing", id, purchasable, transferable, price, supplyLimit, balanceLimit)
}

// SetListing is a paid mutator transaction binding the contract method 0x2fe8a4c3.
//
// Solidity: function setListing(uint256 id, bool purchasable, bool transferable, uint256 price, uint256 supplyLimit, uint256 balanceLimit) returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) SetListing(id *big.Int, purchasable bool, transferable bool, price *big.Int, supplyLimit *big.Int, balanceLimit *big.Int) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SetListing(&_EtherAlleyStore.TransactOpts, id, purchasable, transferable, price, supplyLimit, balanceLimit)
}

// SetListing is a paid mutator transaction binding the contract method 0x2fe8a4c3.
//
// Solidity: function setListing(uint256 id, bool purchasable, bool transferable, uint256 price, uint256 supplyLimit, uint256 balanceLimit) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) SetListing(id *big.Int, purchasable bool, transferable bool, price *big.Int, supplyLimit *big.Int, balanceLimit *big.Int) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SetListing(&_EtherAlleyStore.TransactOpts, id, purchasable, transferable, price, supplyLimit, balanceLimit)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) SetURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "setURI", newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) SetURI(newuri string) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SetURI(&_EtherAlleyStore.TransactOpts, newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) SetURI(newuri string) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.SetURI(&_EtherAlleyStore.TransactOpts, newuri)
}

// TransferBalance is a paid mutator transaction binding the contract method 0x56a6d9ef.
//
// Solidity: function transferBalance(address to, uint256 amount) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) TransferBalance(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "transferBalance", to, amount)
}

// TransferBalance is a paid mutator transaction binding the contract method 0x56a6d9ef.
//
// Solidity: function transferBalance(address to, uint256 amount) returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) TransferBalance(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.TransferBalance(&_EtherAlleyStore.TransactOpts, to, amount)
}

// TransferBalance is a paid mutator transaction binding the contract method 0x56a6d9ef.
//
// Solidity: function transferBalance(address to, uint256 amount) returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) TransferBalance(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.TransferBalance(&_EtherAlleyStore.TransactOpts, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherAlleyStore.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EtherAlleyStore *EtherAlleyStoreSession) Unpause() (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.Unpause(&_EtherAlleyStore.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EtherAlleyStore *EtherAlleyStoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _EtherAlleyStore.Contract.Unpause(&_EtherAlleyStore.TransactOpts)
}

// EtherAlleyStoreApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the EtherAlleyStore contract.
type EtherAlleyStoreApprovalForAllIterator struct {
	Event *EtherAlleyStoreApprovalForAll // Event containing the contract specifics and raw log

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
func (it *EtherAlleyStoreApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherAlleyStoreApprovalForAll)
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
		it.Event = new(EtherAlleyStoreApprovalForAll)
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
func (it *EtherAlleyStoreApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherAlleyStoreApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherAlleyStoreApprovalForAll represents a ApprovalForAll event raised by the EtherAlleyStore contract.
type EtherAlleyStoreApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*EtherAlleyStoreApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreApprovalForAllIterator{contract: _EtherAlleyStore.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *EtherAlleyStoreApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherAlleyStoreApprovalForAll)
				if err := _EtherAlleyStore.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) ParseApprovalForAll(log types.Log) (*EtherAlleyStoreApprovalForAll, error) {
	event := new(EtherAlleyStoreApprovalForAll)
	if err := _EtherAlleyStore.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherAlleyStoreListingChangeIterator is returned from FilterListingChange and is used to iterate over the raw logs and unpacked data for ListingChange events raised by the EtherAlleyStore contract.
type EtherAlleyStoreListingChangeIterator struct {
	Event *EtherAlleyStoreListingChange // Event containing the contract specifics and raw log

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
func (it *EtherAlleyStoreListingChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherAlleyStoreListingChange)
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
		it.Event = new(EtherAlleyStoreListingChange)
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
func (it *EtherAlleyStoreListingChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherAlleyStoreListingChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherAlleyStoreListingChange represents a ListingChange event raised by the EtherAlleyStore contract.
type EtherAlleyStoreListingChange struct {
	Id           *big.Int
	Purchasable  bool
	Transferable bool
	Price        *big.Int
	SupplyLimit  *big.Int
	BalanceLimit *big.Int
	Supply       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterListingChange is a free log retrieval operation binding the contract event 0xd62d82d72db3289f1f82f55480c2e382f45f54eb61a9b93034c426343bd8a55f.
//
// Solidity: event ListingChange(uint256 id, bool purchasable, bool transferable, uint256 price, uint256 supplyLimit, uint256 balanceLimit, uint256 supply)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) FilterListingChange(opts *bind.FilterOpts) (*EtherAlleyStoreListingChangeIterator, error) {

	logs, sub, err := _EtherAlleyStore.contract.FilterLogs(opts, "ListingChange")
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreListingChangeIterator{contract: _EtherAlleyStore.contract, event: "ListingChange", logs: logs, sub: sub}, nil
}

// WatchListingChange is a free log subscription operation binding the contract event 0xd62d82d72db3289f1f82f55480c2e382f45f54eb61a9b93034c426343bd8a55f.
//
// Solidity: event ListingChange(uint256 id, bool purchasable, bool transferable, uint256 price, uint256 supplyLimit, uint256 balanceLimit, uint256 supply)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) WatchListingChange(opts *bind.WatchOpts, sink chan<- *EtherAlleyStoreListingChange) (event.Subscription, error) {

	logs, sub, err := _EtherAlleyStore.contract.WatchLogs(opts, "ListingChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherAlleyStoreListingChange)
				if err := _EtherAlleyStore.contract.UnpackLog(event, "ListingChange", log); err != nil {
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

// ParseListingChange is a log parse operation binding the contract event 0xd62d82d72db3289f1f82f55480c2e382f45f54eb61a9b93034c426343bd8a55f.
//
// Solidity: event ListingChange(uint256 id, bool purchasable, bool transferable, uint256 price, uint256 supplyLimit, uint256 balanceLimit, uint256 supply)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) ParseListingChange(log types.Log) (*EtherAlleyStoreListingChange, error) {
	event := new(EtherAlleyStoreListingChange)
	if err := _EtherAlleyStore.contract.UnpackLog(event, "ListingChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherAlleyStoreTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the EtherAlleyStore contract.
type EtherAlleyStoreTransferBatchIterator struct {
	Event *EtherAlleyStoreTransferBatch // Event containing the contract specifics and raw log

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
func (it *EtherAlleyStoreTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherAlleyStoreTransferBatch)
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
		it.Event = new(EtherAlleyStoreTransferBatch)
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
func (it *EtherAlleyStoreTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherAlleyStoreTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherAlleyStoreTransferBatch represents a TransferBatch event raised by the EtherAlleyStore contract.
type EtherAlleyStoreTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*EtherAlleyStoreTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreTransferBatchIterator{contract: _EtherAlleyStore.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *EtherAlleyStoreTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherAlleyStoreTransferBatch)
				if err := _EtherAlleyStore.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) ParseTransferBatch(log types.Log) (*EtherAlleyStoreTransferBatch, error) {
	event := new(EtherAlleyStoreTransferBatch)
	if err := _EtherAlleyStore.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherAlleyStoreTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the EtherAlleyStore contract.
type EtherAlleyStoreTransferSingleIterator struct {
	Event *EtherAlleyStoreTransferSingle // Event containing the contract specifics and raw log

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
func (it *EtherAlleyStoreTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherAlleyStoreTransferSingle)
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
		it.Event = new(EtherAlleyStoreTransferSingle)
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
func (it *EtherAlleyStoreTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherAlleyStoreTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherAlleyStoreTransferSingle represents a TransferSingle event raised by the EtherAlleyStore contract.
type EtherAlleyStoreTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*EtherAlleyStoreTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreTransferSingleIterator{contract: _EtherAlleyStore.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *EtherAlleyStoreTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherAlleyStoreTransferSingle)
				if err := _EtherAlleyStore.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) ParseTransferSingle(log types.Log) (*EtherAlleyStoreTransferSingle, error) {
	event := new(EtherAlleyStoreTransferSingle)
	if err := _EtherAlleyStore.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherAlleyStoreURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the EtherAlleyStore contract.
type EtherAlleyStoreURIIterator struct {
	Event *EtherAlleyStoreURI // Event containing the contract specifics and raw log

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
func (it *EtherAlleyStoreURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherAlleyStoreURI)
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
		it.Event = new(EtherAlleyStoreURI)
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
func (it *EtherAlleyStoreURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherAlleyStoreURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherAlleyStoreURI represents a URI event raised by the EtherAlleyStore contract.
type EtherAlleyStoreURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*EtherAlleyStoreURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherAlleyStoreURIIterator{contract: _EtherAlleyStore.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *EtherAlleyStoreURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherAlleyStore.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherAlleyStoreURI)
				if err := _EtherAlleyStore.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_EtherAlleyStore *EtherAlleyStoreFilterer) ParseURI(log types.Log) (*EtherAlleyStoreURI, error) {
	event := new(EtherAlleyStoreURI)
	if err := _EtherAlleyStore.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
