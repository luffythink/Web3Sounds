'use client';

import React from 'react';

import { ethers } from 'ethers';
import { upload } from 'ethfs-sdk';

import AudioUpload from '@/components/AudioUpload';

export default function UploaderPage() {
  const directoryAddress = '0x27F34A93A30B86b3Af43E958259CD78d1688FC7f';

  async function onInputChange(file: File) {
    await uploadFile(directoryAddress, file);
  }

  const readFile = (file: File) => {
    return new Promise((resolve) => {
      const reader = new FileReader();

      reader.onload = (res) => {
        console.log('readFile', res);
        resolve(Buffer.from(res.target.result));
      };
      reader.readAsArrayBuffer(file);
    });
  };

  // callback, can be null
  const onProgress = (chunkIndex, totalChunk, fileName) => {
    console.log('progress');
  };
  const onSuccess = (fileName) => {
    console.log('success!', fileName);
  };
  const onError = (message) => {
    console.log('Error', message);
  };

  // const uploadFile = async (directoryAddress: string, rawFile: File) => {
  //   try {
  //     console.log('uploadFile', rawFile);
  //     const directoryPath = rawFile.name;
  //     const fileSize = rawFile.size;
  //     const content = await readFile(rawFile);
  //     const provider = new ethers.providers.Web3Provider(window.ethereum);
  //     // const provider = new ethers.BrowserProvider(window.ethereum);
  //     const signer = provider.getSigner();

  //     console.log('1');
  //     // const directoryAddress = await createDirectory(signer);

  //     await ethfs.upload(
  //       signer,
  //       directoryAddress,
  //       directoryPath,
  //       fileSize,
  //       content,
  //       (args) => {
  //         console.log('onProgress', args);
  //       },
  //       (args) => {
  //         console.log('onSuccess', args);
  //       },
  //       (args) => {
  //         console.log('onError', args);
  //       }
  //     );
  //   } catch (error) {
  //     console.log('try', error);
  //   }
  // };
  const uploadFile = async (directoryAddress: string, rawFile: File) => {
    const directoryPath = rawFile.name;
    const fileSize = rawFile.size;
    const content = await readFile(rawFile);

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
  };

  return (
    <div className="flex min-h-screen w-full flex-col items-center justify-center bg-[url('/images/upload-background.png')] bg-cover bg-center">
      <AudioUpload onFileChange={onInputChange} />
    </div>
  );
}
