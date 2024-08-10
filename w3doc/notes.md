FlatDirectory Address: 0xA4A53bFa26d0F5b8FAbf789D09bc19BEC6222B10
ethfs-cli upload -f "test.html"  -a 0xA4A53bFa26d0F5b8FAbf789D09bc19BEC6222B10  -p $PRIVATE_KEY  -r https://galileo.web3q.io:8545  -t 1 -c 3334

npx ethfs-uploader hello.txt 0xA4A53bFa26d0F5b8FAbf789D09bc19BEC6222B10 --privateKey $PRIVATE_KEY --chainId 3334
providerUrl = https://galileo.web3q.io:8545
chainId = 3334
address: 0xA4A53bFa26d0F5b8FAbf789D09bc19BEC6222B10

Start upload File.......
hello.txt, chunkId: 0
Transaction Id: 0x3489f40b4b833bcb412ad5a32f8d92fa22183d3fef33c56d4f5d49a3c589eb2f
File hello.txt chunkId: 0 uploaded!

Total Cost: 0 W3Q
Total File Count: 1
Total File Size: 0.0234375 KB

const provider = new JsonRpcProvider(process.env.NEXT_PUBLIC_RPC_URL);

Available Accounts
==================

(0) 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 (10000.000000000000000000 ETH)
(1) 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 (10000.000000000000000000 ETH)
(2) 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC (10000.000000000000000000 ETH)
(3) 0x90F79bf6EB2c4f870365E785982E1f101E93b906 (10000.000000000000000000 ETH)
(4) 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 (10000.000000000000000000 ETH)
(5) 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc (10000.000000000000000000 ETH)
(6) 0x976EA74026E726554dB657fA54763abd0C3a0aa9 (10000.000000000000000000 ETH)
(7) 0x14dC79964da2C08b23698B3D3cc7Ca32193d9955 (10000.000000000000000000 ETH)
(8) 0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f (10000.000000000000000000 ETH)
(9) 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720 (10000.000000000000000000 ETH)

Private Keys
==================

(0) 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
(1) 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
(2) 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
(3) 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
(4) 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
(5) 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
(6) 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
(7) 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
(8) 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
(9) 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
== Logs ==
  FileStorage deployed at: 0x5FbDB2315678afecb367f032d93F642f64180aa3
  Deployed By: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

## Setting up 1 EVM.

==========================

Chain 31337

Estimated gas price: 2.000000001 gwei

Estimated total gas used for script: 1375111

Estimated amount required: 0.002750222001375111 ETH

==========================

##### anvil-hardhat
✅  [Success]Hash: 0x4a10a8fc7e862d09ca3f690388616d378687301d57e09108e0956e575beae313
Contract Address: 0x5FbDB2315678afecb367f032d93F642f64180aa3
Block: 1
Paid: 0.00105808000105808 ETH (1058080 gas * 1.000000001 gwei)

✅ Sequence #1 on anvil-hardhat | Total Paid: 0.00105808000105808 ETH (1058080 gas * avg 1.000000001 gwei)
                                                                                                                                                                                                                                                         

==========================

ONCHAIN EXECUTION COMPLETE & SUCCESSFUL.

Transactions saved to: /root/web3/Web3Sounds/w3contract/contract/broadcast/DeployFileStorage.s.sol/31337/run-latest.json

Sensitive values saved to: /root/web3/Web3Sounds/w3contract/contract/cache/DeployFileStorage.s.sol/31337/run-latest.json