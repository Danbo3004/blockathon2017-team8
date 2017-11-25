package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/infinityblockchainlabs/blockathon2017-team8/cmd"
	"github.com/infinityblockchainlabs/blockathon2017-team8/contracts"
)

const (
	url = "http://www.blockathon.asia:8545"
	key = `{"version":3,"id":"05f3b680-e728-4088-8b47-ce6e71d8cda9","address":"d0039884b06163d17a118c890786960d557593f8","Crypto":{"ciphertext":"85fc901a5efc330cde99ee4ecf39ca589d295b8195815265df50151a2db15d7e","cipherparams":{"iv":"f009d427e91e28be84643205bd069ed6"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"78fe3b426ec49fefc22ff99587be0da5cb74b79b294d7de73fb3d8c0414c45d7","n":1024,"r":8,"p":1},"mac":"2d7fa1c7823c062b17ad6dc7a0a818cd1349353ada6093aea24982aaadc6c56c"}}`
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	_, err := contracts.InitClient(url, key, "innerpeace")
	if err != nil {
		panic("wtf")
	}

	router := cmd.NewRouter()
	logrus.Info("Listen in :8080")
	err = http.ListenAndServe(":8080", router)
	logrus.Fatal("Error occur while serving requests", err)
}
