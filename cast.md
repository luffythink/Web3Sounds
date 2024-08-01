以下是针对 FileStorage 合约的所有函数的详细 `cast` 调用命令：

1. 上传文件 (uploadFile):
```bash
cast send $CONTRACT_ADDRESS "uploadFile(string,uint256,bytes32)" "test.txt" 1024 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef --from $ADDRESS --private-key $PRIVATE_KEY

blockHash               0x9caeee8a4af356837af51b197564db542caf7718750ffde6bbc90cca0c091ae5
blockNumber             2
contractAddress         
cumulativeGasUsed       203765
effectiveGasPrice       883817335
from                    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
gasUsed                 203765
logs                    [{"address":"0x5fbdb2315678afecb367f032d93f642f64180aa3","topics":["0xd404a946e4879a6393311b452f8f3945c811eca1ce782109276b7487f148c5aa","0x000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266","0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"],"data":"0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000008746573742e747874000000000000000000000000000000000000000000000000","blockHash":"0x9caeee8a4af356837af51b197564db542caf7718750ffde6bbc90cca0c091ae5","blockNumber":"0x2","blockTimestamp":"0x66ab638d","transactionHash":"0xfd2045bc173d42dee2df263daa24ad14208d19e9aea6c39e000dea8a128fc68b","transactionIndex":"0x0","logIndex":"0x0","removed":false}]
logsBloom               0x00000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000100000000000000000000000000040000000001000000000000000000000000000000000000000000000000000000000200000000000000080000000000000000000000000000040000000000000000000000000000000040000000200000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
root                    0x68e88719c1d1f54e5a833b175d860ba70f16e39e9085a223aa2d2cece354b742
status                  1 (success)
transactionHash         0xfd2045bc173d42dee2df263daa24ad14208d19e9aea6c39e000dea8a128fc68b
transactionIndex        0
type                    2
blobGasPrice            1
blobGasUsed             
authorizationList       
to                      0x5FbDB2315678afecb367f032d93F642f64180aa3
```

2. 修改文件 (modifyFile):
```bash
cast send $CONTRACT_ADDRESS "modifyFile(bytes32,string,uint256)" 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef "updated_test.txt" 2048 --from $ADDRESS --private-key $PRIVATE_KEY
```

3. 删除文件 (deleteFile):
```bash
cast send $CONTRACT_ADDRESS "deleteFile(bytes32)" 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef --from $ADDRESS --private-key $PRIVATE_KEY
```

4. 获取用户文件 (getUserFiles):
```bash
cast call $CONTRACT_ADDRESS "getUserFiles()" --from $ADDRESS
```

5. 获取所有用户文件 (getAllUserFiles) - 仅合约所有者可调用:
```bash
cast call $CONTRACT_ADDRESS "getAllUserFiles()" --from $ADDRESS
```

6. 查询合约所有者 (owner):
```bash
cast call $CONTRACT_ADDRESS "owner()"
```

7. 转移所有权 (transferOwnership) - 仅当前所有者可调用:
```bash
cast send $CONTRACT_ADDRESS "transferOwnership(address)" <NEW_OWNER_ADDRESS> --from $ADDRESS --private-key $PRIVATE_KEY
```

8. 放弃所有权 (renounceOwnership) - 仅当前所有者可调用:
```bash
cast send $CONTRACT_ADDRESS "renounceOwnership()" --from $ADDRESS --private-key $PRIVATE_KEY
```

注意事项：

1. 对于 `bytes32` 类型的参数（如文件哈希），确保使用 32 字节的十六进制字符串。

2. 对于 `send` 交易，你可能需要等待几秒钟让交易被确认，然后才能查询结果。

3. 对于 `call` 查询，结果会直接返回。

4. 如果你想查看更详细的交易信息，可以在命令后添加 `--json` 参数。

5. 如果你想指定 gas 限制或 gas 价格，可以使用 `--gas-limit` 和 `--gas-price` 参数。

6. 对于 Anvil，你可能不需要指定 `--private-key`，因为它默认使用预定义的账户。

7. 如果你想监听事件，可以使用 `cast logs` 命令。例如：
   ```bash
   cast logs $CONTRACT_ADDRESS "FileUploaded(address,bytes32,string)" --from-block 0
   ```
