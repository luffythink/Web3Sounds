// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/FileStorage.sol";
import "../src/proxy/FileStorageProxy.sol";

contract DeployFileStorage is Script {
    function run() external {
        require(vm.envUint("OWNER_PRIVATE_KEY") != 0, "PRIVATE_KEY not set");

        uint256 deployerPrivateKey = vm.envUint("OWNER_PRIVATE_KEY");
        address admin = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);
        // vm.txGasPrice(15 gwei);
        // vm.setNonce(msg.sender, 11);

        console.log("Deploying implementation contract...");
        // 部署实现合约
        FileStorage fileStorageImpl = new FileStorage();

        console.log("Implementation deployed at:", address(fileStorageImpl));

        // 准备初始化数据
        bytes memory data = abi.encodeWithSelector(
            FileStorage.initialize.selector,
            admin
        );

        console.log("Deploying proxy contract...");

        // 部署代理合约
        FileStorageProxy proxy = new FileStorageProxy(
            address(fileStorageImpl),
            admin,
            data
        );

        vm.stopBroadcast();

        console.log("FileStorage Proxy deployed at:", address(proxy));
    }
}
