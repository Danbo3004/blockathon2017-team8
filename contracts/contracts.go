package contracts

import (
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/k0kubun/pp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ddContractAddressFile = "blocker_ddcontract.txt"

var client *EthClient

// test addresses generated in https://www.myetherwallet.com
// network: rinkeby (etherscan.io)
// password: 1234567890
const (
	testIssuerA   = "0x42aa941997a340aa34c69f1be4f8772c5b33816f"
	testIssuerB   = "0xb8ef151495a467bbd7bf8e56ce6e83d991856e1e"
	testCustomerX = "0xd2434dcdff28a29c7dc23a9951c9c4c55c196dba"
	testCustomerY = "0x92c1aa640c694572e357569a7e91cad8a0b15275"
)

const szabo = 10 ^ 16

type EthClient struct {
	URL               string
	ContractOwnerAddr common.Address
	Conn              *ethclient.Client
	Auth              *bind.TransactOpts
	DDContractAddr    common.Address
}

func InitClient(url, ownerKey, ownerPassowrd string) (*EthClient, error) {
	conn, err := ethclient.Dial(url)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	auth, err := bind.NewTransactor(strings.NewReader(ownerKey), ownerPassowrd)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	var ddContractAddr common.Address
	if _, err := os.Stat(ddContractAddressFile); os.IsNotExist(err) {
		pp.Printf("%s doesn't exist. Let's deploy DDContractList", ddContractAddressFile)
		_address, transaction, _, sErr := DeployDDContractList(auth, conn)
		if sErr != nil {
			logrus.Errorln(sErr)
			return nil, sErr
		}

		logrus.Debugf("transaction id: %s", transaction.Hash().Hex())
		err := ioutil.WriteFile(ddContractAddressFile, []byte(_address.Hex()), 0644)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		logrus.Debugln("Stored DDContractList address to file")

		ddContractAddr = _address
	} else {
		logrus.Debugf("File %s exists, just read to get ddContractAddress", ddContractAddressFile)
		data, err := ioutil.ReadFile(ddContractAddressFile)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}

		ddContractAddr = common.HexToAddress(string(data))
		logrus.Debugf("ddContractAddress: %s", ddContractAddr.Hex())
	}

	client = &EthClient{
		URL:               url,
		ContractOwnerAddr: common.HexToAddress(ownerKey),
		Conn:              conn,
		Auth:              auth,
		DDContractAddr:    ddContractAddr,
	}

	// Just test some function :)
	_newContractAddr, err := client.DeployNewContract(testIssuerA, "A-Token", 10)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	// Check ...
	length, err := client.GetLengthOfDDContractList()
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	logrus.Infof("Total contract on blockchain: %d", length)

	list, err := client.GetListExistingContracts()
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	logrus.Infof("%v", list)

	// Sell
	logrus.Infof("NewContractAddr %s", _newContractAddr)
	_tran, err := client.SellTokens(_newContractAddr, testIssuerA, 7, 1*szabo)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	pp.Println(string(_tran.Data()))

	// Buy
	_tran, err = client.BuyTokens(_newContractAddr, testCustomerX, testIssuerA, 5)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	pp.Println(string(_tran.Data()))

	// Buy again
	_tran, err = client.BuyTokens(_newContractAddr, testCustomerX, testIssuerA, 2)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	pp.Println(string(_tran.Data()))

	return client, nil
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

func (c *EthClient) DeployNewContract(issuerAddress, tokenName string, amount int64) (string, error) {

	addrTokenIssuer := common.HexToAddress(issuerAddress)

	address, tx, contractor, err := DeployToken(c.Auth, c.Conn, addrTokenIssuer, to32Byte(tokenName), big.NewInt(amount))
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}

	logrus.Debugf("Contract pending deploy: 0x%x", address)
	logrus.Debugf("Transaction waiting to be mined: 0x%x", tx.Hash().Hex())

	for i := 0; i < 60; i++ {
		name, err := contractor.TokenName(&bind.CallOpts{Pending: true})
		if err != nil {
			logrus.Warnf("Failed to retrieve pending name: %v", err)
			time.Sleep(time.Second)
			continue
		}
		logrus.Debugf("Pending name: %s", name)

		return address.Hex(), c.CheckAndPushContractAddress(address)
	}

	return "", errors.New("Local node may be busy...")
}

func (c *EthClient) BuyTokens(contractAddress, buyerAddress, sellerAddress string, amount int64) (*types.Transaction, error) {
	contractor, err := c.GetContract(contractAddress)
	if err != nil {
		logrus.Errorln(err)
		return nil, err
	}

	_from := common.HexToAddress(sellerAddress)
	return contractor.BuyTokens(&bind.TransactOpts{
		From:     common.HexToAddress(buyerAddress), // Ethereum account to send the transaction from
		Value:    big.NewInt(amount),                // Funds to transfer along along the transaction (nil = 0 = no funds)
		GasPrice: nil,                               // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: nil,                               // Gas limit to set for the transaction execution (nil = estimate + 10%)
		// Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
		// Signer: c.Auth.Signer, // Method to use for signing the transaction (mandatory)
		// Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
	}, _from)
}

func (c *EthClient) SellTokens(contractAddress, sellerAddress string, amount, price int64) (*types.Transaction, error) {
	contractInstance, err := c.GetContract(contractAddress)
	if err != nil {
		logrus.Errorf("Get contract failed: %v", err)
		return nil, err
	}

	_trans, err := contractInstance.SellTokens(&bind.TransactOpts{
		From: common.HexToAddress(sellerAddress), // Ethereum account to send the transaction from
		// Value: big.NewInt(amount), // Funds to transfer along along the transaction (nil = 0 = no funds)
		GasPrice: nil, // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: nil, // Gas limit to set for the transaction execution (nil = estimate + 10%)
		// Nonce  *big.Int       // Nonce to use for the transaction execution (nil = use pending state)
		Signer: c.Auth.Signer, // Method to use for signing the transaction (mandatory)
		// Context context.Context // Network context to support cancellation and timeouts (nil = no timeout)
	}, big.NewInt(amount), big.NewInt(price))

	if err != nil {
		logrus.Errorf("SellTokens: %v", err)
		return nil, err
	}

	return _trans, nil
}

func (c *EthClient) CheckAndPushContractAddress(contractAddress common.Address) error {
	logrus.Infof("contract address: %s", contractAddress.Hex())
	addrList, err := c.GetListExistingContracts()
	if err != nil {
		logrus.Errorf("forGetListExistingContractsmat: %v", err)
		return err
	}

	for _, elem := range *addrList {
		if contractAddress.Hex() == elem {
			return nil
		}
	}

	ddContractList, err := NewDDContractList(c.DDContractAddr, c.Conn)
	if err != nil {
		logrus.Errorf("NewDDContractList: %v", err)
		return err
	}

	_, err = ddContractList.Push(c.Auth, contractAddress)

	if err != nil {
		logrus.Errorf("Push: %v", err)
		return err
	}
	return nil
}

func (c *EthClient) GetDDContractList() (*DDContractList, error) {
	return NewDDContractList(c.DDContractAddr, c.Conn)
}

func (c *EthClient) GetLengthOfDDContractList() (int64, error) {
	contractor, err := c.GetDDContractList()
	if err != nil {
		return 0, err
	}
	result, err := contractor.GetLength(&bind.CallOpts{Pending: true})
	return result.Int64(), err
}

func (c *EthClient) GetListExistingContracts() (*[]string, error) {
	contractor, err := c.GetDDContractList()
	if err != nil {
		logrus.Errorf("GetDDContractList: %v", err)
		return nil, err
	}

	list, err := contractor.GetContracList(&bind.CallOpts{Pending: true})
	if err != nil {
		logrus.Errorf("GetContracList: %v", err)
		return nil, err
	}

	result := []string{}
	for _, item := range list {
		if item.Hex() == "" {
			continue
		}
		result = append(result, item.Hex())
	}

	return &result, nil
}

func (c *EthClient) GetTokenPairOfContractAddress(contractAddress string) string {
	return "Need to dev"
}
