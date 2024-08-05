// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/FileStorage.sol";
import "../src/proxy/FileStorageProxy.sol";

contract DeployFileStorage is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address admin = vm.addr(deployerPrivateKey);

        vm.startBroadcast(deployerPrivateKey);

        // 部署实现合约
        FileStorage fileStorageImpl = new FileStorage();

        // 准备初始化数据
        bytes memory data = abi.encodeWithSelector(
            FileStorage.initialize.selector,
            admin
        );

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