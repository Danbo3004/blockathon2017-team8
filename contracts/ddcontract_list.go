// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DDContractListABI is the input ABI used to generate the binding from.
const DDContractListABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_deleteIndex\",\"type\":\"uint256\"}],\"name\":\"deleteByIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getContracList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DDContractListBin is the compiled bytecode used for deploying new contracts.
const DDContractListBin = `0x6060604052341561000f57600080fd5b6102f28061001e6000396000f3006060604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631e69607e81146100665780637473aa771461007e57806389b09de7146100e4578063be1c766b14610110575b600080fd5b341561007157600080fd5b61007c600435610135565b005b341561008957600080fd5b610091610199565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100d05780820151838201526020016100b8565b505050509050019250505060405180910390f35b34156100ef57600080fd5b61007c73ffffffffffffffffffffffffffffffffffffffff6004351661020f565b341561011b57600080fd5b610123610267565b60405190815260200160405180910390f35b600054811061014357600080fd5b6000808281548110151561015357fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b6101a161026d565b600080548060200260200160405190810160405280929190818152602001828054801561020457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116101d9575b505050505090505b90565b6000805460018101610221838261027f565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005490565b60206040519081016040526000815290565b8154818355818115116102a3576000838152602090206102a39181019083016102a8565b505050565b61020c91905b808211156102c257600081556001016102ae565b50905600a165627a7a72305820bb345dbcbe36d607442092e3cf4bee2ae1e8ed4557ebdf357f6e124e02976bbc0029`

// DeployDDContractList deploys a new Ethereum contract, binding an instance of DDContractList to it.
func DeployDDContractList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DDContractList, error) {
	parsed, err := abi.JSON(strings.NewReader(DDContractListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DDContractListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DDContractList{DDContractListCaller: DDContractListCaller{contract: contract}, DDContractListTransactor: DDContractListTransactor{contract: contract}}, nil
}

// DDContractList is an auto generated Go binding around an Ethereum contract.
type DDContractList struct {
	DDContractListCaller     // Read-only binding to the contract
	DDContractListTransactor // Write-only binding to the contract
}

// DDContractListCaller is an auto generated read-only Go binding around an Ethereum contract.
type DDContractListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDContractListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DDContractListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DDContractListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DDContractListSession struct {
	Contract     *DDContractList   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DDContractListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DDContractListCallerSession struct {
	Contract *DDContractListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DDContractListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DDContractListTransactorSession struct {
	Contract     *DDContractListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DDContractListRaw is an auto generated low-level Go binding around an Ethereum contract.
type DDContractListRaw struct {
	Contract *DDContractList // Generic contract binding to access the raw methods on
}

// DDContractListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DDContractListCallerRaw struct {
	Contract *DDContractListCaller // Generic read-only contract binding to access the raw methods on
}

// DDContractListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DDContractListTransactorRaw struct {
	Contract *DDContractListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDDContractList creates a new instance of DDContractList, bound to a specific deployed contract.
func NewDDContractList(address common.Address, backend bind.ContractBackend) (*DDContractList, error) {
	contract, err := bindDDContractList(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DDContractList{DDContractListCaller: DDContractListCaller{contract: contract}, DDContractListTransactor: DDContractListTransactor{contract: contract}}, nil
}

// NewDDContractListCaller creates a new read-only instance of DDContractList, bound to a specific deployed contract.
func NewDDContractListCaller(address common.Address, caller bind.ContractCaller) (*DDContractListCaller, error) {
	contract, err := bindDDContractList(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DDContractListCaller{contract: contract}, nil
}

// NewDDContractListTransactor creates a new write-only instance of DDContractList, bound to a specific deployed contract.
func NewDDContractListTransactor(address common.Address, transactor bind.ContractTransactor) (*DDContractListTransactor, error) {
	contract, err := bindDDContractList(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DDContractListTransactor{contract: contract}, nil
}

// bindDDContractList binds a generic wrapper to an already deployed contract.
func bindDDContractList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DDContractListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDContractList *DDContractListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DDContractList.Contract.DDContractListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDContractList *DDContractListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDContractList.Contract.DDContractListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDContractList *DDContractListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDContractList.Contract.DDContractListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DDContractList *DDContractListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DDContractList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DDContractList *DDContractListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDContractList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DDContractList *DDContractListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DDContractList.Contract.contract.Transact(opts, method, params...)
}

// DeleteByIndex is a paid mutator transaction binding the contract method 0x1e69607e.
//
// Solidity: function deleteByIndex(_deleteIndex uint256) returns()
func (_DDContractList *DDContractListTransactor) DeleteByIndex(opts *bind.TransactOpts, _deleteIndex *big.Int) (*types.Transaction, error) {
	return _DDContractList.contract.Transact(opts, "deleteByIndex", _deleteIndex)
}

// DeleteByIndex is a paid mutator transaction binding the contract method 0x1e69607e.
//
// Solidity: function deleteByIndex(_deleteIndex uint256) returns()
func (_DDContractList *DDContractListSession) DeleteByIndex(_deleteIndex *big.Int) (*types.Transaction, error) {
	return _DDContractList.Contract.DeleteByIndex(&_DDContractList.TransactOpts, _deleteIndex)
}

// DeleteByIndex is a paid mutator transaction binding the contract method 0x1e69607e.
//
// Solidity: function deleteByIndex(_deleteIndex uint256) returns()
func (_DDContractList *DDContractListTransactorSession) DeleteByIndex(_deleteIndex *big.Int) (*types.Transaction, error) {
	return _DDContractList.Contract.DeleteByIndex(&_DDContractList.TransactOpts, _deleteIndex)
}

// GetContracList is a paid mutator transaction binding the contract method 0x7473aa77.
//
// Solidity: function getContracList() returns(address[])
func (_DDContractList *DDContractListTransactor) GetContracList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDContractList.contract.Transact(opts, "getContracList")
}

// GetContracList is a paid mutator transaction binding the contract method 0x7473aa77.
//
// Solidity: function getContracList() returns(address[])
func (_DDContractList *DDContractListSession) GetContracList() (*types.Transaction, error) {
	return _DDContractList.Contract.GetContracList(&_DDContractList.TransactOpts)
}

// GetContracList is a paid mutator transaction binding the contract method 0x7473aa77.
//
// Solidity: function getContracList() returns(address[])
func (_DDContractList *DDContractListTransactorSession) GetContracList() (*types.Transaction, error) {
	return _DDContractList.Contract.GetContracList(&_DDContractList.TransactOpts)
}

// GetLength is a paid mutator transaction binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() returns(uint256)
func (_DDContractList *DDContractListTransactor) GetLength(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DDContractList.contract.Transact(opts, "getLength")
}

// GetLength is a paid mutator transaction binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() returns(uint256)
func (_DDContractList *DDContractListSession) GetLength() (*types.Transaction, error) {
	return _DDContractList.Contract.GetLength(&_DDContractList.TransactOpts)
}

// GetLength is a paid mutator transaction binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() returns(uint256)
func (_DDContractList *DDContractListTransactorSession) GetLength() (*types.Transaction, error) {
	return _DDContractList.Contract.GetLength(&_DDContractList.TransactOpts)
}

// Push is a paid mutator transaction binding the contract method 0x89b09de7.
//
// Solidity: function push(_address address) returns()
func (_DDContractList *DDContractListTransactor) Push(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DDContractList.contract.Transact(opts, "push", _address)
}

// Push is a paid mutator transaction binding the contract method 0x89b09de7.
//
// Solidity: function push(_address address) returns()
func (_DDContractList *DDContractListSession) Push(_address common.Address) (*types.Transaction, error) {
	return _DDContractList.Contract.Push(&_DDContractList.TransactOpts, _address)
}

// Push is a paid mutator transaction binding the contract method 0x89b09de7.
//
// Solidity: function push(_address address) returns()
func (_DDContractList *DDContractListTransactorSession) Push(_address common.Address) (*types.Transaction, error) {
	return _DDContractList.Contract.Push(&_DDContractList.TransactOpts, _address)
}
