package services
import (
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/joho/godotenv"
	"os"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type Setting struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Key   string `gorm:"column:key;unique"`
	Value uint64 `gorm:"column:value;not null"`
}

func init() {
	// Get  settings from environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetLastBlock()(uint64,error){
	// Connect to the Ethereum network
	client, err := ethclient.Dial("https://rpctestnet.spectachains.org/")
	if err != nil {
		log.Fatalf("Error connecting to rpc: %v", err)
	}

	// Get the latest block number
	blockNumber, err := client.HeaderByNumber(context.Background(),nil)
	if err != nil {
		log.Fatalf("Error connecting to rpc: %v", err)
	}

	return blockNumber.Number.Uint64(),err

}
func DbHeight() (uint64, error) {
	// Open a database connection using GORM
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// get block height value from db
	var setting Setting
	err = db.Where("ID = ?", 1).First(&setting).Error
	return setting.Value.Uint64(),err

}
