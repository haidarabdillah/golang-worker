
// check transaction list
// connect database
// check if address receiver (to:) === Payment address
// send to rabbit-MQ to be process in que transaction
package main

import (
	"deposit_collector/services"
	"fmt"
)

func main() {
	// connect RPC
	// check block height
	currentHeight,err := services.GetLastBlock()
	if err != nil {
		log.Fatalf("Error connecting to rpc: %v", err)
	}
	dbHeight,err := services.DBHeight()
}

