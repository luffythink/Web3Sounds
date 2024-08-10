package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	// 直接导入本地包

	// 直接导入本地包
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 全局变量
var (
	client     *ethclient.Client
	instance   *FileStorageContract
	auth       *bind.TransactOpts
	privateKey *ecdsa.PrivateKey
)

// 初始化函数，连接到以太坊网络并创建合约实例
func init() {
	var err error

	// 连接到以太坊网络
	client, err = ethclient.Dial("https://galileo.web3q.io:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 设置合约地址
	contractAddress := common.HexToAddress("0x27F34A93A30B86b3Af43E958259CD78d1688FC7f")

	// 创建合约实例
	instance, err = fsc.NewFileStorageContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// 设置私钥（注意：在生产环境中应该安全地管理私钥）
	privateKey, err = crypto.HexToECDSA("your_private_key_here")
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	// 创建交易选项
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
}

// 获取所有用户文件
func getAllUserFiles() {
	files, err := instance.GetAllUserFiles(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to get all user files: %v", err)
	}
	fmt.Println("All user files:", files)
}

// 上传文件
func uploadFile(name string, size *big.Int, hash [32]byte) {
	tx, err := instance.UploadFile(auth, name, size, hash)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
	fmt.Printf("File upload transaction sent: %s\n", tx.Hash().Hex())
}

// 修改文件
func modifyFile(hash [32]byte, newName string, newSize *big.Int) {
	tx, err := instance.ModifyFile(auth, hash, newName, newSize)
	if err != nil {
		log.Fatalf("Failed to modify file: %v", err)
	}
	fmt.Printf("File modification transaction sent: %s\n", tx.Hash().Hex())
}

// 删除文件
func deleteFile(hash [32]byte) {
	tx, err := instance.DeleteFile(auth, hash)
	if err != nil {
		log.Fatalf("Failed to delete file: %v", err)
	}
	fmt.Printf("File deletion transaction sent: %s\n", tx.Hash().Hex())
}

// 监听文件上传事件
func watchFileUploaded() {
	filterer, err := fsc.NewFileStorageContractFilterer(instance.Address(), client)
	if err != nil {
		log.Fatalf("Failed to create event filterer: %v", err)
	}

	sink := make(chan *fsc.FileStorageContractFileUploaded)
	sub, err := filterer.WatchFileUploaded(nil, sink, nil, nil)
	if err != nil {
		log.Fatalf("Failed to watch for FileUploaded events: %v", err)
	}

	for {
		select {
		case event := <-sink:
			fmt.Printf("File uploaded: Name=%s, Hash=%x, User=%s\n", event.Name, event.FileHash, event.User.Hex())
		case err := <-sub.Err():
			log.Fatalf("Error in event subscription: %v", err)
		}
	}
}

func main() {
	// 示例：获取所有用户文件
	getAllUserFiles()

	// 示例：上传文件
	uploadFile("example.txt", big.NewInt(1024), [32]byte{1, 2, 3, 4})

	// 示例：修改文件
	modifyFile([32]byte{1, 2, 3, 4}, "new_example.txt", big.NewInt(2048))

	// 示例：删除文件
	deleteFile([32]byte{1, 2, 3, 4})

	// 监听文件上传事件
	watchFileUploaded()
}
