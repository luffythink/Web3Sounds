import { Buffer } from 'buffer';

import { ethers } from 'ethers';
import { keccak_256 as sha3 } from 'js-sha3';

import { Espoir } from '@/contracts/Espoir';

// 检查是否在浏览器环境中
if (typeof window !== 'undefined') {
  (window as any).Buffer = Buffer;
}

const stringToHex = (s: string): string => ethers.hexlify(ethers.toUtf8Bytes(s));

export const FileContract = async (address: string) => {
  // console.log('address', address);
  const provider = new ethers.BrowserProvider(window.ethereum);
  const contract = new ethers.Contract(Espoir.ADDRESS, Espoir.ABI, provider);

  return contract.connect(await provider.getSigner());
};

const bufferChunk = (buffer: Buffer, chunkSize: number): Buffer[] => {
  let i = 0;
  const result: Buffer[] = [];
  const len = buffer.length;
  const chunkLength = Math.ceil(len / chunkSize);

  while (i < len) {
    result.push(buffer.slice(i, (i += chunkLength)));
  }

  return result;
};

const clearOldFile = async (
  fileContract: ethers.Contract,
  chunkSize: number,
  hexName: string
): Promise<boolean> => {
  try {
    const oldChunkSize = await fileContract.countChunks(hexName);

    if (oldChunkSize > chunkSize) {
      // 删除旧文件
      const tx = await fileContract.remove(hexName);

      // console.log(`Remove file: ${hexName}`);
      // console.log(`Transaction Id: ${tx.hash}`);
      const receipt = await tx.wait();

      return receipt.status === 1;
    }
  } catch (e) {
    return false;
  }

  return true;
};

interface RequestParams {
  account: string;
  chunkLength: number;
  contractAddress: string;
  dirPath: string;
  file: {
    raw: File;
  };
  flatDirectoryAddress: string;
  onError: (error: Error) => void;
  onProgress: (progress: { percent: number }) => void;
  onSuccess: (success: { path: string }) => void;
}

export const request = async ({
  account,
  chunkLength,
  contractAddress,
  dirPath,
  file,
  flatDirectoryAddress,
  onError,
  onProgress,
  onSuccess
}: RequestParams) => {
  const rawFile = file.raw;
  // let content = await rawFile.arrayBuffer();

  // content = Buffer.from(new Uint8Array(content));
  let content: Buffer = Buffer.from(new Uint8Array(await rawFile.arrayBuffer()));

  // 文件名
  const name = dirPath + rawFile.name;
  const hexName = stringToHex(name);
  const hexType = stringToHex(rawFile.type);
  const fileSize = rawFile.size;
  let chunks: Buffer[] = [];

  if (fileSize > chunkLength) {
    const chunkSize = Math.ceil(fileSize / chunkLength);

    // ts-ignore
    chunks = bufferChunk(content, chunkSize);
  } else {
    // ts-ignore
    chunks.push(content);
  }

  const fileContract = (await FileContract(contractAddress)) as ethers.Contract;
  const clear = await clearOldFile(fileContract, chunks.length, hexName);

  if (!clear) {
    onError(new Error('Check Old File Fail!'));

    return;
  }

  let uploadState = true;

  for (const [index, chunk] of Array.from(chunks.entries())) {
    const hexData = '0x' + chunk.toString('hex');
    const localHash = '0x' + sha3(chunk);
    const hash = await fileContract.getChunkHash(hexName, index);

    if (localHash === hash) {
      // console.log(`File ${name} chunkId: ${index}: The data is not changed.`);
      onProgress({ percent: index + 1 });
      continue;
    }

    try {
      // 文件被移除或改变
      const tx = await fileContract.writeChunk(hexName, hexType, index, hexData);

      // console.log(`Transaction Id: ${tx.hash}`);
      const receipt = await tx.wait();

      if (!receipt.status) {
        uploadState = false;
        break;
      }
      onProgress({ percent: index + 1 });
    } catch (e) {
      // console.log(e);
      uploadState = false;
      break;
    }
  }
  if (uploadState) {
    const url = 'https://' + flatDirectoryAddress + '.3336.w3link.io/' + account + '/' + name;

    onSuccess({ path: url });
  } else {
    onError(new Error('Upload request failed!'));
  }
};
