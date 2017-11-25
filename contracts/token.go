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
const DDContractListABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_deleteIndex\",\"type\":\"uint256\"}],\"name\":\"deleteByIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContracList\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DDContractListBin is the compiled bytecode used for deploying new contracts.
const DDContractListBin = `0x6060604052341561000f57600080fd5b6102f28061001e6000396000f3006060604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631e69607e81146100665780637473aa771461007e57806389b09de7146100e4578063be1c766b14610110575b600080fd5b341561007157600080fd5b61007c600435610135565b005b341561008957600080fd5b610091610199565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100d05780820151838201526020016100b8565b505050509050019250505060405180910390f35b34156100ef57600080fd5b61007c73ffffffffffffffffffffffffffffffffffffffff6004351661020f565b341561011b57600080fd5b610123610267565b60405190815260200160405180910390f35b600054811061014357600080fd5b6000808281548110151561015357fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b6101a161026d565b600080548060200260200160405190810160405280929190818152602001828054801561020457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116101d9575b505050505090505b90565b6000805460018101610221838261027f565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005490565b60206040519081016040526000815290565b8154818355818115116102a3576000838152602090206102a39181019083016102a8565b505050565b61020c91905b808211156102c257600081556001016102ae565b50905600a165627a7a72305820701e2aa725bc016edab9f9a103b6aa9bc27e6688690e47354262f52f24ce5c920029`

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

// GetContracList is a free data retrieval call binding the contract method 0x7473aa77.
//
// Solidity: function getContracList() constant returns(address[])
func (_DDContractList *DDContractListCaller) GetContracList(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DDContractList.contract.Call(opts, out, "getContracList")
	return *ret0, err
}

// GetContracList is a free data retrieval call binding the contract method 0x7473aa77.
//
// Solidity: function getContracList() constant returns(address[])
func (_DDContractList *DDContractListSession) GetContracList() ([]common.Address, error) {
	return _DDContractList.Contract.GetContracList(&_DDContractList.CallOpts)
}

// GetContracList is a free data retrieval call binding the contract method 0x7473aa77.
//
// Solidity: function getContracList() constant returns(address[])
func (_DDContractList *DDContractListCallerSession) GetContracList() ([]common.Address, error) {
	return _DDContractList.Contract.GetContracList(&_DDContractList.CallOpts)
}

// GetLength is a free data retrieval call binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() constant returns(uint256)
func (_DDContractList *DDContractListCaller) GetLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DDContractList.contract.Call(opts, out, "getLength")
	return *ret0, err
}

// GetLength is a free data retrieval call binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() constant returns(uint256)
func (_DDContractList *DDContractListSession) GetLength() (*big.Int, error) {
	return _DDContractList.Contract.GetLength(&_DDContractList.CallOpts)
}

// GetLength is a free data retrieval call binding the contract method 0xbe1c766b.
//
// Solidity: function getLength() constant returns(uint256)
func (_DDContractList *DDContractListCallerSession) GetLength() (*big.Int, error) {
	return _DDContractList.Contract.GetLength(&_DDContractList.CallOpts)
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_someone\",\"type\":\"address\"}],\"name\":\"getAmountThisToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getIssuerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_someone\",\"type\":\"address\"},{\"name\":\"_newAmount\",\"type\":\"uint256\"}],\"name\":\"setAmoutThisToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenShops\",\"outputs\":[{\"name\":\"isActived\",\"type\":\"bool\"},{\"name\":\"amountTokens\",\"type\":\"uint256\"},{\"name\":\"tokenPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenIssuerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountTokens\",\"type\":\"uint256\"},{\"name\":\"_tokenPrice\",\"type\":\"uint256\"}],\"name\":\"sellTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNameThisToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenIssuer\",\"type\":\"address\"},{\"name\":\"_tokenName\",\"type\":\"bytes32\"},{\"name\":\"_amountTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenIssuerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenName\",\"type\":\"bytes32\"}],\"name\":\"SellToken_Ether\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_totalTokenBuy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenIssuerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenName\",\"type\":\"bytes32\"}],\"name\":\"BuyToken_Ehter\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x6060604052341561000f57600080fd5b60405160608061057383398101604052808051919060200180519190602001805160008054600160a060020a031916600160a060020a0396871617808255909516855260036020526040852055505060015561050290819061007190396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632e6cc053811461009d57806332e71ead146100ce57806357304815146100fd578063593fa37414610133578063b6d7855a14610178578063c25d01141461018b578063ec8ac4d81461019e578063ed9772b6146101b4578063fe3939f7146101cd575b600080fd5b34156100a857600080fd5b6100bc600160a060020a03600435166101e0565b60405190815260200160405180910390f35b34156100d957600080fd5b6100e16101fb565b604051600160a060020a03909116815260200160405180910390f35b341561010857600080fd5b61011f600160a060020a036004351660243561020a565b604051901515815260200160405180910390f35b341561013e57600080fd5b610152600160a060020a036004351661022a565b604051921515835260208301919091526040808301919091526060909101905180910390f35b341561018357600080fd5b6100bc610250565b341561019657600080fd5b6100e1610256565b6101b2600160a060020a0360043516610265565b005b34156101bf57600080fd5b6101b26004356024356103ce565b34156101d857600080fd5b6100bc6104d0565b600160a060020a031660009081526003602052604090205490565b600054600160a060020a031690565b600160a060020a0391909116600090815260036020526040902055600190565b600260208190526000918252604090912080546001820154919092015460ff9092169183565b60015481565b600054600160a060020a031681565b600160a060020a038116600090815260026020526040812054819060ff16151560011461029157600080fd5b600160a060020a03831660009081526002602081905260409091200154915081348115156102bb57fe5b600160a060020a03851660009081526002602052604090206001015491900491508190116102e857600080fd5b600160a060020a03808416600081815260026020908152604080832060010180548790039055600390915280822080548690039055339093168152829020805484019055903480156108fc029151600060405180830381858888f19350505050151561035357600080fd5b7fed12fd386ef6331096fee82eb252bcd8e3f36a469072400a221154fe7fd2c76e8333836000809054906101000a9004600160a060020a0316600154604051600160a060020a039586168152938516602085015260408085019390935293166060830152608082019290925260a001905180910390a1505050565b600160a060020a033316600090815260036020526040812054116103f157600080fd5b600160a060020a03331660009081526003602052604090205482901161041657600080fd5b6000821161042357600080fd5b6000811161043057600080fd5b600160a060020a033381811660009081526002602081905260408083206001808201899055928101879055805460ff191683179055915490547f1015d42154ec3496e8f271fd08fe4d0d30cd90c5001acf81efae08bc9784dfae9487938793909116919051600160a060020a039586168152602081019490945260408085019390935293166060830152608082019290925260a001905180910390a15050565b600154905600a165627a7a723058203ec8815f522fb2e5cf684e2f64cc4657e80b3ab8a95f3ff7c96c4a048c8191f10029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenIssuer common.Address, _tokenName [32]byte, _amountTokens *big.Int) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, _tokenIssuer, _tokenName, _amountTokens)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// TokenIssuerAddress is a free data retrieval call binding the contract method 0xc25d0114.
//
// Solidity: function TokenIssuerAddress() constant returns(address)
func (_Token *TokenCaller) TokenIssuerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TokenIssuerAddress")
	return *ret0, err
}

// TokenIssuerAddress is a free data retrieval call binding the contract method 0xc25d0114.
//
// Solidity: function TokenIssuerAddress() constant returns(address)
func (_Token *TokenSession) TokenIssuerAddress() (common.Address, error) {
	return _Token.Contract.TokenIssuerAddress(&_Token.CallOpts)
}

// TokenIssuerAddress is a free data retrieval call binding the contract method 0xc25d0114.
//
// Solidity: function TokenIssuerAddress() constant returns(address)
func (_Token *TokenCallerSession) TokenIssuerAddress() (common.Address, error) {
	return _Token.Contract.TokenIssuerAddress(&_Token.CallOpts)
}

// TokenName is a free data retrieval call binding the contract method 0xb6d7855a.
//
// Solidity: function TokenName() constant returns(bytes32)
func (_Token *TokenCaller) TokenName(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "TokenName")
	return *ret0, err
}

// TokenName is a free data retrieval call binding the contract method 0xb6d7855a.
//
// Solidity: function TokenName() constant returns(bytes32)
func (_Token *TokenSession) TokenName() ([32]byte, error) {
	return _Token.Contract.TokenName(&_Token.CallOpts)
}

// TokenName is a free data retrieval call binding the contract method 0xb6d7855a.
//
// Solidity: function TokenName() constant returns(bytes32)
func (_Token *TokenCallerSession) TokenName() ([32]byte, error) {
	return _Token.Contract.TokenName(&_Token.CallOpts)
}

// GetAmountThisToken is a free data retrieval call binding the contract method 0x2e6cc053.
//
// Solidity: function getAmountThisToken(_someone address) constant returns(uint256)
func (_Token *TokenCaller) GetAmountThisToken(opts *bind.CallOpts, _someone common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getAmountThisToken", _someone)
	return *ret0, err
}

// GetAmountThisToken is a free data retrieval call binding the contract method 0x2e6cc053.
//
// Solidity: function getAmountThisToken(_someone address) constant returns(uint256)
func (_Token *TokenSession) GetAmountThisToken(_someone common.Address) (*big.Int, error) {
	return _Token.Contract.GetAmountThisToken(&_Token.CallOpts, _someone)
}

// GetAmountThisToken is a free data retrieval call binding the contract method 0x2e6cc053.
//
// Solidity: function getAmountThisToken(_someone address) constant returns(uint256)
func (_Token *TokenCallerSession) GetAmountThisToken(_someone common.Address) (*big.Int, error) {
	return _Token.Contract.GetAmountThisToken(&_Token.CallOpts, _someone)
}

// GetNameThisToken is a free data retrieval call binding the contract method 0xfe3939f7.
//
// Solidity: function getNameThisToken() constant returns(bytes32)
func (_Token *TokenCaller) GetNameThisToken(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "getNameThisToken")
	return *ret0, err
}

// GetNameThisToken is a free data retrieval call binding the contract method 0xfe3939f7.
//
// Solidity: function getNameThisToken() constant returns(bytes32)
func (_Token *TokenSession) GetNameThisToken() ([32]byte, error) {
	return _Token.Contract.GetNameThisToken(&_Token.CallOpts)
}

// GetNameThisToken is a free data retrieval call binding the contract method 0xfe3939f7.
//
// Solidity: function getNameThisToken() constant returns(bytes32)
func (_Token *TokenCallerSession) GetNameThisToken() ([32]byte, error) {
	return _Token.Contract.GetNameThisToken(&_Token.CallOpts)
}

// TokenShops is a free data retrieval call binding the contract method 0x593fa374.
//
// Solidity: function tokenShops( address) constant returns(isActived bool, amountTokens uint256, tokenPrice uint256)
func (_Token *TokenCaller) TokenShops(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsActived    bool
	AmountTokens *big.Int
	TokenPrice   *big.Int
}, error) {
	ret := new(struct {
		IsActived    bool
		AmountTokens *big.Int
		TokenPrice   *big.Int
	})
	out := ret
	err := _Token.contract.Call(opts, out, "tokenShops", arg0)
	return *ret, err
}

// TokenShops is a free data retrieval call binding the contract method 0x593fa374.
//
// Solidity: function tokenShops( address) constant returns(isActived bool, amountTokens uint256, tokenPrice uint256)
func (_Token *TokenSession) TokenShops(arg0 common.Address) (struct {
	IsActived    bool
	AmountTokens *big.Int
	TokenPrice   *big.Int
}, error) {
	return _Token.Contract.TokenShops(&_Token.CallOpts, arg0)
}

// TokenShops is a free data retrieval call binding the contract method 0x593fa374.
//
// Solidity: function tokenShops( address) constant returns(isActived bool, amountTokens uint256, tokenPrice uint256)
func (_Token *TokenCallerSession) TokenShops(arg0 common.Address) (struct {
	IsActived    bool
	AmountTokens *big.Int
	TokenPrice   *big.Int
}, error) {
	return _Token.Contract.TokenShops(&_Token.CallOpts, arg0)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_from address) returns()
func (_Token *TokenTransactor) BuyTokens(opts *bind.TransactOpts, _from common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "buyTokens", _from)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_from address) returns()
func (_Token *TokenSession) BuyTokens(_from common.Address) (*types.Transaction, error) {
	return _Token.Contract.BuyTokens(&_Token.TransactOpts, _from)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_from address) returns()
func (_Token *TokenTransactorSession) BuyTokens(_from common.Address) (*types.Transaction, error) {
	return _Token.Contract.BuyTokens(&_Token.TransactOpts, _from)
}

// GetIssuerAddress is a paid mutator transaction binding the contract method 0x32e71ead.
//
// Solidity: function getIssuerAddress() returns(address)
func (_Token *TokenTransactor) GetIssuerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "getIssuerAddress")
}

// GetIssuerAddress is a paid mutator transaction binding the contract method 0x32e71ead.
//
// Solidity: function getIssuerAddress() returns(address)
func (_Token *TokenSession) GetIssuerAddress() (*types.Transaction, error) {
	return _Token.Contract.GetIssuerAddress(&_Token.TransactOpts)
}

// GetIssuerAddress is a paid mutator transaction binding the contract method 0x32e71ead.
//
// Solidity: function getIssuerAddress() returns(address)
func (_Token *TokenTransactorSession) GetIssuerAddress() (*types.Transaction, error) {
	return _Token.Contract.GetIssuerAddress(&_Token.TransactOpts)
}

// SellTokens is a paid mutator transaction binding the contract method 0xed9772b6.
//
// Solidity: function sellTokens(_amountTokens uint256, _tokenPrice uint256) returns()
func (_Token *TokenTransactor) SellTokens(opts *bind.TransactOpts, _amountTokens *big.Int, _tokenPrice *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "sellTokens", _amountTokens, _tokenPrice)
}

// SellTokens is a paid mutator transaction binding the contract method 0xed9772b6.
//
// Solidity: function sellTokens(_amountTokens uint256, _tokenPrice uint256) returns()
func (_Token *TokenSession) SellTokens(_amountTokens *big.Int, _tokenPrice *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SellTokens(&_Token.TransactOpts, _amountTokens, _tokenPrice)
}

// SellTokens is a paid mutator transaction binding the contract method 0xed9772b6.
//
// Solidity: function sellTokens(_amountTokens uint256, _tokenPrice uint256) returns()
func (_Token *TokenTransactorSession) SellTokens(_amountTokens *big.Int, _tokenPrice *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SellTokens(&_Token.TransactOpts, _amountTokens, _tokenPrice)
}

// SetAmoutThisToken is a paid mutator transaction binding the contract method 0x57304815.
//
// Solidity: function setAmoutThisToken(_someone address, _newAmount uint256) returns(bool)
func (_Token *TokenTransactor) SetAmoutThisToken(opts *bind.TransactOpts, _someone common.Address, _newAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAmoutThisToken", _someone, _newAmount)
}

// SetAmoutThisToken is a paid mutator transaction binding the contract method 0x57304815.
//
// Solidity: function setAmoutThisToken(_someone address, _newAmount uint256) returns(bool)
func (_Token *TokenSession) SetAmoutThisToken(_someone common.Address, _newAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetAmoutThisToken(&_Token.TransactOpts, _someone, _newAmount)
}

// SetAmoutThisToken is a paid mutator transaction binding the contract method 0x57304815.
//
// Solidity: function setAmoutThisToken(_someone address, _newAmount uint256) returns(bool)
func (_Token *TokenTransactorSession) SetAmoutThisToken(_someone common.Address, _newAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetAmoutThisToken(&_Token.TransactOpts, _someone, _newAmount)
}

// TradeTokenABI is the input ABI used to generate the binding from.
const TradeTokenABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getPairToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"someOneAccept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Tok\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAType\",\"type\":\"address\"},{\"name\":\"_amountA\",\"type\":\"uint256\"},{\"name\":\"_tokenBType\",\"type\":\"address\"},{\"name\":\"_amountB\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_A_Actor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amountA\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_B_Actor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenBAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amountB\",\"type\":\"uint256\"}],\"name\":\"exChangeToken\",\"type\":\"event\"}]"

// TradeTokenBin is the compiled bytecode used for deploying new contracts.
const TradeTokenBin = `0x6060604052341561000f57600080fd5b60405160808061074683398101604052808051919060200180519190602001805191906020018051915083905061005385336401000000006104136100ab82021704565b1161005d57600080fd5b60068054600160a060020a03338116600160a060020a031992831617909255600180549683169682169690961790955560028054939091169290941691909117909255600355600455610152565b60008054600160a060020a031916600160a060020a038481169190911780835516632e6cc0538383604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561013157600080fd5b6102c65a03f1151561014257600080fd5b5050506040518051949350505050565b6105e5806101616000396000f30060606040526004361061003d5763ffffffff60e060020a6000350416634323491581146100425780635a667e7f146100a85780636a45750c146100bd575b600080fd5b341561004d57600080fd5b6100556100ec565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561009457808201518382015260200161007c565b505050509050019250505060405180910390f35b34156100b357600080fd5b6100bb6101f0565b005b34156100c857600080fd5b6100d0610404565b604051600160a060020a03909116815260200160405180910390f35b6100f4610561565b600080548190600181016101088382610573565b50600091825260209091206001805492909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909316929092179091558154829181016101568382610573565b5060009182526020918290206002549101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390921691909117905581548291818102016040519081016040528092919081815260200182805480156101e457602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116101c6575b505050505091505b5090565b600080600080600060035411151561020757600080fd5b6004546000901161021757600080fd5b6007805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a0390811691909117909155600254600654610259929182169116610413565b60015460075491955061027891600160a060020a039182169116610413565b60015460065491945061029791600160a060020a039182169116610413565b6002546007549193506102b691600160a060020a039182169116610413565b6003549091508210156102c857600080fd5b6004548110156102d757600080fd5b600380546004805460009384905592905560015460075496830196958201959190940393919092039161031791600160a060020a039081169116856104ae565b5060025460075461033591600160a060020a039081169116836104ae565b5060015460065461035391600160a060020a039081169116846104ae565b5060025460065461037191600160a060020a039081169116866104ae565b506006546001546003546007546002546004547f6e0638be69a0949d215aa3013d8dae1f6533e6eb8526fb98ed393c314310b99f95600160a060020a0390811695811694938116921690604051600160a060020a03968716815294861660208601526040808601949094529185166060850152909316608083015260a082019290925260c001905180910390a150505050565b600054600160a060020a031681565b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038481169190911780835516632e6cc05383836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561048d57600080fd5b6102c65a03f1151561049e57600080fd5b5050506040518051949350505050565b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03858116919091178083551663573048158484846040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561052e57600080fd5b6102c65a03f1151561053f57600080fd5b50505060405180519050156105565750600161055a565b5060005b9392505050565b60206040519081016040526000815290565b8154818355818115116105975760008381526020902061059791810190830161059c565b505050565b6105b691905b808211156101ec57600081556001016105a2565b905600a165627a7a723058207ae35c129a366f3dadf3463423d5439a0e1939f8ec78c359610f619eb38feec00029`

// DeployTradeToken deploys a new Ethereum contract, binding an instance of TradeToken to it.
func DeployTradeToken(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenAType common.Address, _amountA *big.Int, _tokenBType common.Address, _amountB *big.Int) (common.Address, *types.Transaction, *TradeToken, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TradeTokenBin), backend, _tokenAType, _amountA, _tokenBType, _amountB)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TradeToken{TradeTokenCaller: TradeTokenCaller{contract: contract}, TradeTokenTransactor: TradeTokenTransactor{contract: contract}}, nil
}

// TradeToken is an auto generated Go binding around an Ethereum contract.
type TradeToken struct {
	TradeTokenCaller     // Read-only binding to the contract
	TradeTokenTransactor // Write-only binding to the contract
}

// TradeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeTokenSession struct {
	Contract     *TradeToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeTokenCallerSession struct {
	Contract *TradeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TradeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTokenTransactorSession struct {
	Contract     *TradeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TradeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeTokenRaw struct {
	Contract *TradeToken // Generic contract binding to access the raw methods on
}

// TradeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeTokenCallerRaw struct {
	Contract *TradeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTokenTransactorRaw struct {
	Contract *TradeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradeToken creates a new instance of TradeToken, bound to a specific deployed contract.
func NewTradeToken(address common.Address, backend bind.ContractBackend) (*TradeToken, error) {
	contract, err := bindTradeToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradeToken{TradeTokenCaller: TradeTokenCaller{contract: contract}, TradeTokenTransactor: TradeTokenTransactor{contract: contract}}, nil
}

// NewTradeTokenCaller creates a new read-only instance of TradeToken, bound to a specific deployed contract.
func NewTradeTokenCaller(address common.Address, caller bind.ContractCaller) (*TradeTokenCaller, error) {
	contract, err := bindTradeToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTokenCaller{contract: contract}, nil
}

// NewTradeTokenTransactor creates a new write-only instance of TradeToken, bound to a specific deployed contract.
func NewTradeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTokenTransactor, error) {
	contract, err := bindTradeToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TradeTokenTransactor{contract: contract}, nil
}

// bindTradeToken binds a generic wrapper to an already deployed contract.
func bindTradeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeToken *TradeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TradeToken.Contract.TradeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeToken *TradeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeToken.Contract.TradeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeToken *TradeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeToken.Contract.TradeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeToken *TradeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TradeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeToken *TradeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeToken *TradeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeToken.Contract.contract.Transact(opts, method, params...)
}

// Tok is a free data retrieval call binding the contract method 0x6a45750c.
//
// Solidity: function Tok() constant returns(address)
func (_TradeToken *TradeTokenCaller) Tok(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TradeToken.contract.Call(opts, out, "Tok")
	return *ret0, err
}

// Tok is a free data retrieval call binding the contract method 0x6a45750c.
//
// Solidity: function Tok() constant returns(address)
func (_TradeToken *TradeTokenSession) Tok() (common.Address, error) {
	return _TradeToken.Contract.Tok(&_TradeToken.CallOpts)
}

// Tok is a free data retrieval call binding the contract method 0x6a45750c.
//
// Solidity: function Tok() constant returns(address)
func (_TradeToken *TradeTokenCallerSession) Tok() (common.Address, error) {
	return _TradeToken.Contract.Tok(&_TradeToken.CallOpts)
}

// GetPairToken is a paid mutator transaction binding the contract method 0x43234915.
//
// Solidity: function getPairToken() returns(address[])
func (_TradeToken *TradeTokenTransactor) GetPairToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeToken.contract.Transact(opts, "getPairToken")
}

// GetPairToken is a paid mutator transaction binding the contract method 0x43234915.
//
// Solidity: function getPairToken() returns(address[])
func (_TradeToken *TradeTokenSession) GetPairToken() (*types.Transaction, error) {
	return _TradeToken.Contract.GetPairToken(&_TradeToken.TransactOpts)
}

// GetPairToken is a paid mutator transaction binding the contract method 0x43234915.
//
// Solidity: function getPairToken() returns(address[])
func (_TradeToken *TradeTokenTransactorSession) GetPairToken() (*types.Transaction, error) {
	return _TradeToken.Contract.GetPairToken(&_TradeToken.TransactOpts)
}

// SomeOneAccept is a paid mutator transaction binding the contract method 0x5a667e7f.
//
// Solidity: function someOneAccept() returns()
func (_TradeToken *TradeTokenTransactor) SomeOneAccept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeToken.contract.Transact(opts, "someOneAccept")
}

// SomeOneAccept is a paid mutator transaction binding the contract method 0x5a667e7f.
//
// Solidity: function someOneAccept() returns()
func (_TradeToken *TradeTokenSession) SomeOneAccept() (*types.Transaction, error) {
	return _TradeToken.Contract.SomeOneAccept(&_TradeToken.TransactOpts)
}

// SomeOneAccept is a paid mutator transaction binding the contract method 0x5a667e7f.
//
// Solidity: function someOneAccept() returns()
func (_TradeToken *TradeTokenTransactorSession) SomeOneAccept() (*types.Transaction, error) {
	return _TradeToken.Contract.SomeOneAccept(&_TradeToken.TransactOpts)
}
