package helpers

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
)

func getLastBlock()(uint64,error){
	// Connect to the Ethereum network
	client, err := ethclient.Dial("https://rpctestnet.spectachains.org/")
	if err != nil {
		panic(err)
	}

	// Get the latest block number
	blockNumber, err := client.HeaderByNumber(context.Background(),nil)
	if err != nil {
		panic(err)
	}

	return blockNumber.Number.Uint64(),err

}
