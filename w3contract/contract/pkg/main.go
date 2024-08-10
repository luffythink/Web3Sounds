package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"your_project_name/contracts" // 确保这里的路径正确
)

func main() {
	// (int64)ChainID  :=3334
	// 连接到以太坊网络
	client, err := ethclient.Dial("https://galileo.web3q.io:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 设置私钥（请替换为您的私钥）
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to get public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 设置合约地址
	contractAddress := common.HexToAddress("0x27F34A93A30B86b3Af43E958259CD78d1688FC7f")

	// 创建合约实例
	instance, err := contracts.NewFileStorageContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3334)) // 512 是 Galileo 测试网的 Chain ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前gas平均费用:", gasPrice)

	auth.GasPrice = new(big.Int).Mul(gasPrice, big.NewInt(2))
	auth.GasLimit = uint64(300000)

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account balance: %s\n", balance.String())

	// 1. 获取所有用户文件
	// Get user-specific files
	files, err := instance.GetUserFiles(&bind.CallOpts{From: fromAddress})
	if err != nil {
		log.Fatalf("Failed to get user files: %v", err)
	}
	fmt.Printf("Files for address %s: %v\n", fromAddress.Hex(), files)

	fmt.Println("All user files:", files)

	// 2. 上传文件
	fileName := "example.txt"
	fileSize := big.NewInt(1024)     // 1KB
	fileHash := [32]byte{1, 2, 3, 4} // 示例哈希值
	tx, err := instance.UploadFile(auth, fileName, fileSize, fileHash)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}
	fmt.Printf("File upload transaction sent: %s\n", tx.Hash().Hex())

	// 3. 修改文件
	newFileName := "new_example.txt"
	newFileSize := big.NewInt(2048) // 2KB
	tx, err = instance.ModifyFile(auth, fileHash, newFileName, newFileSize)
	if err != nil {
		log.Fatalf("Failed to modify file: %v", err)
	}
	fmt.Printf("File modification transaction sent: %s\n", tx.Hash().Hex())

	userFiles, err := instance.GetUserFiles(&bind.CallOpts{From: fromAddress})
	if err != nil {
		log.Fatalf("Failed to get user files: %v", err)
	}
	fmt.Println("User files:", userFiles)

	// 4. 删除文件
	tx, err = instance.DeleteFile(auth, fileHash)
	if err != nil {
		log.Fatalf("Failed to delete file: %v", err)
	}
	fmt.Printf("File deletion transaction sent: %s\n", tx.Hash().Hex())

	// 5. 获取用户文件
	userFiles1, err := instance.GetUserFiles(&bind.CallOpts{From: fromAddress})
	if err != nil {
		log.Fatalf("Failed to get user files: %v", err)
	}
	fmt.Println("User files:", userFiles1)

	// 6. 获取合约所有者
	owner, err := instance.Owner(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to get owner: %v", err)
	}
	fmt.Println("Contract owner:", owner.Hex())

	go watchFileUploadedEvents(instance)

	// 7. 监听文件上传事件（示例）
	// filterer, err := contracts.NewFileStorageContractFilterer(contractAddress, client)
	// if err != nil {
	// 	log.Fatalf("Failed to create event filterer: %v", err)
	// }

	// sink := make(chan *contracts.FileStorageContractFileUploaded)
	// sub, err := filterer.WatchFileUploaded(nil, sink, nil, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to watch for FileUploaded events: %v", err)
	// }

	// go func() {
	// 	for {
	// 		select {
	// 		case event := <-sink:
	// 			fmt.Printf("File uploaded: User=%s, Hash=%x, Name=%s\n", event.User.Hex(), event.FileHash, event.Name)
	// 		case err := <-sub.Err():
	// 			log.Fatalf("Error in event subscription: %v", err)
	// 		}
	// 	}
	// }()

	// 等待一段时间以观察事件
	select {}
}

func watchFileUploadedEvents(contract *contracts.FileStorageContract) {
	filterOpts := &bind.FilterOpts{
		Start:   0, // You might want to start from a specific block
		Context: context.Background(),
		End:     nil,
	}

	for {
		events, err := contract.FilterFileUploaded(filterOpts, nil, nil)
		if err != nil {
			log.Printf("Failed to filter FileUploaded events: %v", err)
			time.Sleep(time.Second * 10) // Wait before retrying
			continue
		}

		for events.Next() {
			event := events.Event
			log.Printf("File uploaded: User=%s, Hash=%x, Name=%s\n", event.User.Hex(), event.FileHash, event.Name)
		}

		if events.Error() != nil {
			log.Printf("Error iterating events: %v", events.Error())
		}

		// Update the start block for the next iteration
		filterOpts.Start = *filterOpts.End + 1

		// Wait before the next poll
		time.Sleep(time.Second * 15)
	}
}
