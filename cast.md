当然，我很乐意为您提供一系列的 cast 命令来测试修改后的 FileStorage 合约。这些命令将覆盖合约的主要功能。请确保在运行这些命令之前，您已经部署了合约并设置了必要的环境变量。

假设您已经设置了以下环境变量：
- `$CONTRACT_ADDRESS`: 部署的合约地址
- `$OWNER_ADDRESS`: 合约所有者的地址
- `$ADDRESS`: 测试用户1的地址
- `$ADDRESS2`: 测试用户2的地址
- `$OWNER_PRIVATE_KEY`: 合约所有者的私钥
- `$PRIVATE_KEY`: 测试用户1的私钥
- `$PRIVATE_KEY2`: 测试用户2的私钥

以下是测试命令：

1. 上传文件（用户1）：
```bash
cast send $CONTRACT_ADDRESS "uploadFile(string,uint256,bytes32)" "test1.txt" 1024 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef --from $ADDRESS --private-key $PRIVATE_KEY
```

2. 上传另一个文件（用户2）：
```bash
cast send $CONTRACT_ADDRESS "uploadFile(string,uint256,bytes32)" "test2.txt" 2048 0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890 --from $ADDRESS2 --private-key $PRIVATE_KEY2
```

3. 获取用户1的文件：
```bash
cast call $CONTRACT_ADDRESS "getUserFiles()" --from $ADDRESS
```

4. 获取用户2的文件：
```bash
cast call $CONTRACT_ADDRESS "getUserFiles()" --from $ADDRESS2
```

5. 修改文件（用户1）：
```bash
cast send $CONTRACT_ADDRESS "modifyFile(bytes32,string,uint256)" 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef "updated_test1.txt" 1500 --from $ADDRESS --private-key $PRIVATE_KEY
```

6. 尝试修改不属于自己的文件（用户2尝试修改用户1的文件，应该失败）：
```bash
cast send $CONTRACT_ADDRESS "modifyFile(bytes32,string,uint256)" 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef "hacked.txt" 1000 --from $ADDRESS2 --private-key $PRIVATE_KEY2
```

7. 删除文件（用户1）：
```bash
cast send $CONTRACT_ADDRESS "deleteFile(bytes32)" 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef --from $ADDRESS --private-key $PRIVATE_KEY
```

8. 尝试删除不存在的文件（应该失败）：
```bash
cast send $CONTRACT_ADDRESS "deleteFile(bytes32)" 0x0000000000000000000000000000000000000000000000000000000000000000 --from $ADDRESS --private-key $PRIVATE_KEY
```

9. 获取所有用户的所有文件（只有合约所有者可以调用）：
```bash
cast call $CONTRACT_ADDRESS "getAllUserFiles()" --from $OWNER_ADDRESS
```

10. 尝试以非所有者身份获取所有文件（应该失败）：
```bash
cast call $CONTRACT_ADDRESS "getAllUserFiles()" --from $ADDRESS
```

这些命令将帮助您测试合约的主要功能，包括上传、修改、删除文件，以及获取用户文件和所有文件。请注意，某些命令预期会失败（如非所有者调用 getAllUserFiles 或修改/删除不属于自己的文件），这是为了测试合约的权限控制。

在运行这些命令时，请确保您连接到了正确的网络（如果是在测试网络上），并且账户中有足够的 ETH 来支付 gas 费用。如果遇到任何问题，请随时告诉我，我会很乐意帮助您解决。