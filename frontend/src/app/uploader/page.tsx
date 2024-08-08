'use client';

import React from 'react';

import { ethers } from 'ethers';
import { upload } from 'ethfs-sdk';

import AudioUpload from '@/components/AudioUpload';

export default function UploaderPage() {
  // const directoryAddress = '0xbdEdC0Ae1F17Cabf18CCcC517c176F00213Efd9C';
  const directoryAddress = '0x27F34A93A30B86b3Af43E958259CD78d1688FC7f';

  async function onInputChange(file: File) {
    await uploadFile(directoryAddress, file);
  }

  const readFile = (file: File) => {
    return new Promise<ArrayBuffer>((resolve) => {
      const reader = new FileReader();

      reader.onload = (event: ProgressEvent<FileReader>) => {
        const res = event.target;

        if (res && res.result) {
          resolve(res.result as ArrayBuffer);
        }
      };

      reader.readAsArrayBuffer(file);
    });
  };

  // callback, can be null
  const onProgress = (chunkIndex: number, totalChunk: number, fileName: string): void => {
    console.log('progress');
  };

  const onSuccess = (fileName: string): void => {
    console.log('success!', fileName);
  };

  const onError = (message: string): void => {
    console.log('Error', message);
  };

  const uploadFile = async (directoryAddress: string, rawFile: File) => {
    try {
      const directoryPath = rawFile.name;
      const fileSize = rawFile.size;
      const arrayBufferContent = await readFile(rawFile);
      const content = Buffer.from(arrayBufferContent);

      const provider = new ethers.providers.Web3Provider(window.ethereum);
      const signer = provider.getSigner();

      await upload(
        signer,
        directoryAddress,
        directoryPath,
        fileSize,
        content,
        onProgress,
        onSuccess,
        onError
      );
    } catch (error: any) {
      if (error.code === 'UNSUPPORTED_OPERATION' || error.message.includes('provider.getSigner')) {
        onError('Wallet is not connected');
      } else {
        onError(error.message);
      }
    }
  };

  return (
    <div className="flex min-h-screen w-full flex-col items-center justify-center bg-[url('/images/upload-background.png')] bg-cover bg-center">
      <AudioUpload onFileChange={onInputChange} />
    </div>
  );
}
