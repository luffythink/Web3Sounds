export const Espoir = {
  ABI: [
    'function writeChunk(bytes memory name, bytes memory fileType, uint256 chunkId, bytes calldata data) public payable',
    'function remove(bytes memory name) external returns (uint256)',
    'function removes(bytes[] memory names) public',
    'function countChunks(bytes memory name) external view returns (uint256)',
    'function getChunkHash(bytes memory name, uint256 chunkId) public view returns (bytes32)',
    'function getAuthorFiles(address author) public view returns (uint256[] memory times,bytes[] memory names,bytes[] memory types,string[] memory urls)'
  ] as const,
  ADDRESS: '0xbdEdC0Ae1F17Cabf18CCcC517c176F00213Efd9C' as `0x${string}`
};
