// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";


// import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
// import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
// import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract FileStorage is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    struct File {
        string name;
        uint256 size;
        bytes32 hash;
        uint256 uploadTime;
        address uploader;
        bool exists;
    }

    mapping(bytes32 => File) private files;
    mapping(address => bytes32[]) private userFileList;
    bytes32[] private allFileHashes;
    uint256 private fileCount;


    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) public initializer {
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();
    }
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    event FileUploaded(
        address indexed user,
        bytes32 indexed fileHash,
        string name
    );
    event FileModified(
        address indexed user,
        bytes32 indexed fileHash,
        string name
    );
    event FileDeleted(address indexed user, bytes32 indexed fileHash);

    // constructor(address initialOwner) __Ownable_init(initialOwner) {}

    modifier fileExists(bytes32 _hash) {
        require(files[_hash].exists, "File does not exist");
        _;
    }

    modifier onlyUploader(bytes32 _hash) {
        require(
            files[_hash].uploader == msg.sender,
            "Only uploader can modify or delete the file"
        );
        _;
    }

    modifier validFileData(
        string memory _name,
        uint256 _size,
        bytes32 _hash
    ) {
        require(bytes(_name).length > 0, "File name cannot be empty");
        require(_size > 0, "File size must be greater than 0");
        require(_hash != bytes32(0), "File hash cannot be empty");
        _;
    }

    function uploadFile(
        string memory _name,
        uint256 _size,
        bytes32 _hash
    ) external validFileData(_name, _size, _hash) {
        require(!files[_hash].exists, "File already exists");

        files[_hash] = File(
            _name,
            _size,
            _hash,
            block.timestamp,
            msg.sender,
            true
        );
        userFileList[msg.sender].push(_hash);
        allFileHashes.push(_hash);
        fileCount++;

        emit FileUploaded(msg.sender, _hash, _name);
    }

    function modifyFile(
        bytes32 _hash,
        string memory _newName,
        uint256 _newSize
    )
        external
        fileExists(_hash)
        onlyUploader(_hash)
        validFileData(_newName, _newSize, _hash)
    {
        File storage file = files[_hash];
        file.name = _newName;
        file.size = _newSize;
        file.uploadTime = block.timestamp;

        emit FileModified(msg.sender, _hash, _newName);
    }

    function deleteFile(
        bytes32 _hash
    ) external fileExists(_hash) onlyUploader(_hash) {
        delete files[_hash];
        _removeFileFromList(msg.sender, _hash);
        _removeFileFromAllHashes(_hash);
        fileCount--;

        emit FileDeleted(msg.sender, _hash);
    }

    function getUserFiles() external view returns (File[] memory) {
        return _getFiles(msg.sender);
    }

    function getAllUserFiles() external view onlyOwner returns (File[] memory) {
        File[] memory allFiles = new File[](fileCount);
        for (uint256 i = 0; i < allFileHashes.length; i++) {
            bytes32 fileHash = allFileHashes[i];
            if (files[fileHash].exists) {
                allFiles[i] = files[fileHash];
            }
        }
        return allFiles;
    }

    function _removeFileFromList(address user, bytes32 _hash) private {
        bytes32[] storage userFiles = userFileList[user];
        for (uint i = 0; i < userFiles.length; i++) {
            if (userFiles[i] == _hash) {
                userFiles[i] = userFiles[userFiles.length - 1];
                userFiles.pop();
                break;
            }
        }
    }

    function _removeFileFromAllHashes(bytes32 _hash) private {
        for (uint i = 0; i < allFileHashes.length; i++) {
            if (allFileHashes[i] == _hash) {
                allFileHashes[i] = allFileHashes[allFileHashes.length - 1];
                allFileHashes.pop();
                break;
            }
        }
    }

    function _getFiles(address user) private view returns (File[] memory) {
        bytes32[] memory hashes = userFileList[user];
        File[] memory userFiles = new File[](hashes.length);
        uint256 validFileCount = 0;

        for (uint i = 0; i < hashes.length; i++) {
            if (files[hashes[i]].exists) {
                userFiles[validFileCount] = files[hashes[i]];
                validFileCount++;
            }
        }

        // Resize the array to remove any empty slots
        assembly {
            mstore(userFiles, validFileCount)
        }

        return userFiles;
    }
}
