// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/FileStorage.sol";

contract FileStorageTest is Test {
    FileStorage public fileStorage;
    address public owner;
    address public user1;
    address public user2;

    function setUp() public {
        owner = address(this);
        user1 = address(0x1);
        user2 = address(0x2);
        fileStorage = new FileStorage(address(this));
    }

    function testUploadFile() public {
        vm.prank(user1);
        fileStorage.uploadFile("test.txt", 100, keccak256("test"));
        
        FileStorage.File[] memory files = fileStorage.getUserFiles();
        assertEq(files.length, 1);
        assertEq(files[0].name, "test.txt");
        assertEq(files[0].size, 100);
        assertEq(files[0].hash, keccak256("test"));
    }

    function testUploadDuplicateFile() public {
        vm.startPrank(user1);
        fileStorage.uploadFile("test.txt", 100, keccak256("test"));
        
        vm.expectRevert("File already exists");
        fileStorage.uploadFile("test.txt", 100, keccak256("test"));
        vm.stopPrank();
    }

    function testModifyFile() public {
        vm.startPrank(user1);
        bytes32 fileHash = keccak256("test");
        fileStorage.uploadFile("test.txt", 100, fileHash);
        
        fileStorage.modifyFile(fileHash, "updated.txt", 200);
        
        FileStorage.File[] memory files = fileStorage.getUserFiles();
        assertEq(files[0].name, "updated.txt");
        assertEq(files[0].size, 200);
        vm.stopPrank();
    }

    function testModifyNonExistentFile() public {
        vm.prank(user1);
        vm.expectRevert("File does not exist");
        fileStorage.modifyFile(keccak256("nonexistent"), "test.txt", 100);
    }

    function testDeleteFile() public {
        vm.startPrank(user1);
        bytes32 fileHash = keccak256("test");
        fileStorage.uploadFile("test.txt", 100, fileHash);
        
        fileStorage.deleteFile(fileHash);
        
        FileStorage.File[] memory files = fileStorage.getUserFiles();
        assertEq(files.length, 0);
        vm.stopPrank();
    }

    function testDeleteNonExistentFile() public {
        vm.prank(user1);
        vm.expectRevert("File does not exist");
        fileStorage.deleteFile(keccak256("nonexistent"));
    }

    function testGetUserFiles() public {
        vm.startPrank(user1);
        fileStorage.uploadFile("test1.txt", 100, keccak256("test1"));
        fileStorage.uploadFile("test2.txt", 200, keccak256("test2"));
        
        FileStorage.File[] memory files = fileStorage.getUserFiles();
        assertEq(files.length, 2);
        assertEq(files[0].name, "test1.txt");
        assertEq(files[1].name, "test2.txt");
        vm.stopPrank();
    }

    function testGetAllUserFiles() public {
        vm.prank(user1);
        fileStorage.uploadFile("user1.txt", 100, keccak256("user1"));
        
        vm.prank(user2);
        fileStorage.uploadFile("user2.txt", 200, keccak256("user2"));
        
        (address[] memory users, FileStorage.File[][] memory allFiles) = fileStorage.getAllUserFiles();
        
        assertEq(users.length, 2);
        assertEq(allFiles.length, 2);
        assertEq(allFiles[0].length, 1);
        assertEq(allFiles[1].length, 1);
    }

    function testOnlyOwnerCanGetAllUserFiles() public {
        vm.prank(user1);
        vm.expectRevert("Ownable: caller is not the owner");
        fileStorage.getAllUserFiles();
    }

    function testInvalidFileData() public {
        vm.startPrank(user1);
        
        vm.expectRevert("File name cannot be empty");
        fileStorage.uploadFile("", 100, keccak256("test"));
        
        vm.expectRevert("File size must be greater than 0");
        fileStorage.uploadFile("test.txt", 0, keccak256("test"));
        
        vm.expectRevert("File hash cannot be empty");
        fileStorage.uploadFile("test.txt", 100, bytes32(0));
        
        vm.stopPrank();
    }
}