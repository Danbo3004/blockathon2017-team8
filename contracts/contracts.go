package contracts

import (
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ddContractAddressFile = "blocker_ddcontract.txt"

var client *EthClient

type EthClient struct {
	URL               string
	ContractOwnerAddr common.Address
	Conn              *ethclient.Client
	Auth              *bind.TransactOpts
	DDContractAddress string
}

func InitClient(url, ownerKey, ownerPassowrd string) (*EthClient, error) {
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewTransactor(strings.NewReader(ownerKey), ownerPassowrd)
	if err != nil {
		return nil, err
	}

	var ddContractAddress string
	if _, err := os.Stat(ddContractAddressFile); os.IsNotExist(err) {
		pp.Printf("%s doesn't exist. Let deploy DDContractList", ddContractAddressFile)
		_address, transaction, _, sErr := DeployDDContractList(auth, conn)
		if sErr != nil {
			return nil, err
		}

		pp.Printf("transaction hash: %s\n", transaction.Hash())
		pp.Println("Let's save DDContractList address")
		err := ioutil.WriteFile(ddContractAddressFile, []byte(_address.Hex()), 0644)
		if err != nil {
			return nil, err
		}

		ddContractAddress = _address.Hex()
	} else {
		data, err := ioutil.ReadFile("/tmp/dat")
		if err != nil {
			return nil, err
		}

		ddContractAddress = string(data)
	}

	return &EthClient{
		URL:               url,
		ContractOwnerAddr: common.HexToAddress(ownerKey),
		Conn:              conn,
		Auth:              auth,
		DDContractAddress: ddContractAddress,
	}, nil
}

func GetClient() *EthClient {
	if client == nil {
		panic("You must init client first")
	}

	return client
}

func (c *EthClient) GetContract(contractAddress string) (*Token, error) {
	return NewToken(common.HexToAddress(contractAddress), c.Conn)
}

func to32Byte(str string) [32]byte {
	arr := [32]byte{}
	_ = copy(arr[:], str)
	return arr
}

func (c *EthClient) DeployNewContract(issuerAddress, tokenName string, amount int64) error {

	addrTokenIssuer := common.HexToAddress(issuerAddress)

	address, tx, contractor, err := DeployToken(c.Auth, c.Conn, addrTokenIssuer, to32Byte(tokenName), big.NewInt(amount))
	if err != nil {
		return err
	}

	pp.Printf("Contract pending deploy: 0x%x\n", address)
	pp.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	for i := 0; i < 60; i++ {
		name, err := contractor.TokenName(&bind.CallOpts{Pending: true})
		if err != nil {
			pp.Printf("Failed to retrieve pending name: %v", err)
			time.Sleep(time.Second)
			continue
		}
		pp.Println("Pending name:", name)

		return c.CheckAndPushContractAddress(address.Hex())
	}

	return errors.New("Local node may be busy...")
}

func (c *EthClient) BuyTokens(contractAddress, buyerAddress, sellerAddress string, amount int64) (*types.Transaction, error) {
	contractor, err := c.GetContract(contractAddress)
	if err != nil {
		return nil, err
	}

	_from := common.HexToAddress(sellerAddress)
	return contractor.BuyTokens(&bind.TransactOpts{
		From:     common.HexToAddress(buyerAddress), // Ethereum account to send the transaction from
		Value:    big.NewInt(amount),                // Funds to transfer along along the transaction (nil = 0 = no funds)
		GasPrice: nil,                               // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: nil,                               // Gas limit to set for the transaction execution (nil = estimate + 10%)
		// Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
		// Signer SignerFn       // Method to use for signing the transaction (mandatory)
		// Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
	}, _from)
}

func (c *EthClient) SellTokens(contractAddress, buyerAddress string, amount, price int64) (*types.Transaction, error) {
	contractor, err := c.GetContract(contractAddress)
	if err != nil {
		return nil, err
	}
	return contractor.SellTokens(&bind.TransactOpts{
		From: common.HexToAddress(buyerAddress), // Ethereum account to send the transaction from
		// Value: big.NewInt(amount), // Funds to transfer along along the transaction (nil = 0 = no funds)
		GasPrice: nil, // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: nil, // Gas limit to set for the transaction execution (nil = estimate + 10%)
		// Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
		// Signer SignerFn       // Method to use for signing the transaction (mandatory)
		// Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
	}, big.NewInt(amount), big.NewInt(price))
}

func (c *EthClient) CheckAndPushContractAddress(contractAddress string) error {
	addrList, err := c.GetListOfContracts()
	if err != nil {
		return err
	}

	for _, elem := range *addrList {
		if contractAddress == elem {
			return nil
		}
	}

	ddContractor, err := NewDDContractList(c.ContractOwnerAddr, c.Conn)
	if err != nil {
		return err
	}

	_, err = ddContractor.Push(&bind.TransactOpts{
		From:     c.ContractOwnerAddr,
		GasPrice: nil,
		GasLimit: nil,
	}, common.HexToAddress(contractAddress))

	return err
}

func (c *EthClient) GetListOfContracts() (*[]string, error) {
	ddContractor, err := NewDDContractList(c.ContractOwnerAddr, c.Conn)
	if err != nil {
		return nil, err
	}

	transaction, err := ddContractor.GetContracList(nil)
	if err != nil {
		return nil, err
	}

	pp.Println(transaction)
	panic("TODO here pls, return list of contract addresses")
	return nil, nil
}

func (c *EthClient) GetTokenPairOfContractAddress(contractAddress string) string {
	return "Need to dev"
}
