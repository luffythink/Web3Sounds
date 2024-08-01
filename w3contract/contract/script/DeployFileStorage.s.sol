// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/FileStorage.sol";

contract DeployFileStorage is Script {
    function run() external {
        // 获取部署者的私钥
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        
        // 获取部署者的地址
        address deployerAddress = vm.addr(deployerPrivateKey);

        // 开始广播交易
        vm.startBroadcast(deployerPrivateKey);

        // 部署 FileStorage 合约
        // FileStorage fileStorage = new FileStorage(address(this));
        FileStorage fileStorage = new FileStorage(deployerAddress);

        // 可以在这里添加一些初始化操作，比如：
        // fileStorage.transferOwnership(newOwnerAddress);

        // 停止广播交易
        vm.stopBroadcast();

        // 输出部署的合约地址
        console.log("FileStorage deployed at:", address(fileStorage));
        console.log("Deployed By:",deployerAddress);
    }
}