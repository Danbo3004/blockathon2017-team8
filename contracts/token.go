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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_someone\",\"type\":\"address\"}],\"name\":\"getAmountThisToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_someone\",\"type\":\"address\"},{\"name\":\"_newAmount\",\"type\":\"uint256\"}],\"name\":\"setAmoutThisToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenShops\",\"outputs\":[{\"name\":\"isActived\",\"type\":\"bool\"},{\"name\":\"amountTokens\",\"type\":\"uint256\"},{\"name\":\"tokenPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenName\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TokenIssuerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amountTokens\",\"type\":\"uint256\"},{\"name\":\"_tokenPrice\",\"type\":\"uint256\"}],\"name\":\"sellTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNameThisToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenIssuer\",\"type\":\"address\"},{\"name\":\"_tokenName\",\"type\":\"bytes32\"},{\"name\":\"_amountTokens\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenIssuerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenName\",\"type\":\"bytes32\"}],\"name\":\"SellToken_Ether\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_totalTokenBuy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenIssuerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenName\",\"type\":\"bytes32\"}],\"name\":\"BuyToken_Ehter\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x6060604052341561000f57600080fd5b60405160608061054683398101604052808051919060200180519190602001805160008054600160a060020a031916600160a060020a039687161780825590951685526003602052604085205550506001556104d590819061007190396000f30060606040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632e6cc053811461009257806357304815146100c3578063593fa374146100f9578063b6d7855a1461013e578063c25d011414610151578063ec8ac4d814610180578063ed9772b614610196578063fe3939f7146101af575b600080fd5b341561009d57600080fd5b6100b1600160a060020a03600435166101c2565b60405190815260200160405180910390f35b34156100ce57600080fd5b6100e5600160a060020a03600435166024356101dd565b604051901515815260200160405180910390f35b341561010457600080fd5b610118600160a060020a03600435166101fd565b604051921515835260208301919091526040808301919091526060909101905180910390f35b341561014957600080fd5b6100b1610223565b341561015c57600080fd5b610164610229565b604051600160a060020a03909116815260200160405180910390f35b610194600160a060020a0360043516610238565b005b34156101a157600080fd5b6101946004356024356103a1565b34156101ba57600080fd5b6100b16104a3565b600160a060020a031660009081526003602052604090205490565b600160a060020a0391909116600090815260036020526040902055600190565b600260208190526000918252604090912080546001820154919092015460ff9092169183565b60015481565b600054600160a060020a031681565b600160a060020a038116600090815260026020526040812054819060ff16151560011461026457600080fd5b600160a060020a038316600090815260026020819052604090912001549150813481151561028e57fe5b600160a060020a03851660009081526002602052604090206001015491900491508190116102bb57600080fd5b600160a060020a03808416600081815260026020908152604080832060010180548790039055600390915280822080548690039055339093168152829020805484019055903480156108fc029151600060405180830381858888f19350505050151561032657600080fd5b7fed12fd386ef6331096fee82eb252bcd8e3f36a469072400a221154fe7fd2c76e8333836000809054906101000a9004600160a060020a0316600154604051600160a060020a039586168152938516602085015260408085019390935293166060830152608082019290925260a001905180910390a1505050565b600160a060020a033316600090815260036020526040812054116103c457600080fd5b600160a060020a0333166000908152600360205260409020548290116103e957600080fd5b600082116103f657600080fd5b6000811161040357600080fd5b600160a060020a033381811660009081526002602081905260408083206001808201899055928101879055805460ff191683179055915490547f1015d42154ec3496e8f271fd08fe4d0d30cd90c5001acf81efae08bc9784dfae9487938793909116919051600160a060020a039586168152602081019490945260408085019390935293166060830152608082019290925260a001905180910390a15050565b600154905600a165627a7a723058209fd8d85c1bc637882e1f9c47c5e0e4ab962c17e66e1c857d957e53bbafce55980029`

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
