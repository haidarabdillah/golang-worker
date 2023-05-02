package eth
import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"os"
)
func init() {
	// Get  settings from environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getLastBlock()(uint64,error){
	rpc = os.Getenv("RPC")
	// Connect to the Ethereum network
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}

	// Get the latest block number
	blockNumber, err := client.HeaderByNumber(context.Background(),nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(blockNumber.Number.Uint64())

	return blockNumber.Number.Uint64(), nil

}
