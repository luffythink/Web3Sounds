'use client';

import React, { useRef, useState } from 'react';

import copy from 'clipboard-copy';
import { keccak_256 as sha3 } from 'js-sha3';

import { request } from '@/utils/request';

import UploadDragger from './UploadDragger';
import UploadList from './UploadList';

type Props = {
  account?: string;
  fileContract?: string;
  flatDirectory?: string;
  dirPath?: string;
  beforeUpload?: () => boolean;
  onChange?: (file: any, files: any[]) => void;
  onSuccess?: (response: any, file: any, files: any[]) => void;
  onError?: (error: any, file: any, files: any[]) => void;
  onProgress?: (event: any, file: any, files: any[]) => void;
  onExceed?: (rawFiles: File[], files: any[]) => void;
  accept?: string;
  multiple?: boolean;
  customRequestClint?: (options: any) => Promise<void>;
  limit?: number;
  drag?: boolean;
  showList?: boolean;
  children?: React.ReactNode;
};

const W3qDeployer: React.FC<Props> = ({
  accept = 'audio/*',
  account = '',
  beforeUpload,
  children,
  customRequestClint = request,
  dirPath = '',
  drag = true,
  fileContract = '',
  flatDirectory = '',
  limit,
  multiple = false,
  onChange = () => {},
  onError = () => {},
  onExceed = () => {},
  onProgress = () => {},
  onSuccess = () => {},
  showList = true
}) => {
  const [files, setFiles] = useState<any[]>([]);
  const [currentReq, setCurrentReq] = useState<any>(null);
  const reqs = useRef<Record<string, any>>({});
  const inputRef = useRef<HTMLInputElement>(null);
  const chunkLength = 24 * 1024;

  // const enable = fileContract !== null;
  const enable = true;

  const onClickTrigger = () => {
    if (enable && inputRef.current) {
      inputRef.current.click();
    }
  };

  const onInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const rawFiles = Array.from(e.target.files || []);

    uploadFiles(rawFiles);
  };

  const uploadFiles = (rawFiles: File[]) => {
    rawFiles = clearFile(rawFiles);
    const filesLen = rawFiles.length + files.length;

    if (limit && limit < filesLen) {
      return onExceed(rawFiles, files);
    }
    startUpload(rawFiles);
  };

  const clearFile = (rawFiles: File[]) => {
    const newFiles: File[] = [];

    rawFiles.forEach((rawFile) => {
      const uid = sha3(rawFile.name + rawFile.size + rawFile.type);

      if (!files.some((file) => file.uid === uid)) {
        newFiles.push(rawFile);
      }
    });

    return newFiles;
  };

  const startUpload = (rawFiles: File[]) => {
    rawFiles.forEach((rawFile) => {
      const file = normalizeFiles(rawFile);

      normalizeReq(file);
    });
    autoUpload();
  };

  const normalizeFiles = (rawFile: File) => {
    let chunkSize = 1;

    if (rawFile.size > chunkLength) {
      chunkSize = Math.ceil(rawFile.size / chunkLength);
    }
    const file = {
      name: rawFile.name,
      percent: 0,
      raw: rawFile,
      size: rawFile.size,
      status: 'init',
      totalChunks: chunkSize,
      type: rawFile.type,
      uid: sha3(rawFile.name + rawFile.size + rawFile.type)
    };

    setFiles((prevFiles) => [...prevFiles, file]);

    return file;
  };

  const normalizeReq = (file: any) => {
    const { uid } = file;

    reqs.current[uid] = {
      account,
      chunkLength,
      contractAddress: fileContract,
      dirPath,
      file,
      flatDirectoryAddress: flatDirectory,
      onError: (error: any) => handleError(file, error),
      onProgress: (event: any) => handleProgress(file, event),
      onSuccess: (response: any) => handleSuccess(file, response)
    };
  };

  const getFirstReq = () => {
    const keys = Object.keys(reqs.current);

    return keys.length > 0 ? reqs.current[keys[0]] : null;
  };

  const autoUpload = async () => {
    if (currentReq) {
      return;
    }
    if (!beforeUpload || beforeUpload()) {
      let req = getFirstReq();

      setCurrentReq(req);
      while (req) {
        const file = req.file;

        file.status = 'pending';
        onChange(file, files);
        await customRequestClint(req);

        req = getFirstReq();
        setCurrentReq(req);
      }
    }
  };

  const handleError = (file: any, error: any) => {
    const { uid } = file;

    delete reqs.current[uid];
    file.status = 'failure';
    onError(error, file, files);
  };

  const handleSuccess = (file: any, response: any) => {
    const { uid } = file;

    delete reqs.current[uid];
    file.status = 'success';
    file.url = response.path;
    onChange(file, files);
    onSuccess(response, file, files);
  };

  const handleProgress = (file: any, event: any) => {
    file.percent = event.percent;
    onChange(file, files);
    onProgress(event, file, files);
  };

  const onCopy = (url: string) => {
    copy(url);
  };

  const onDelete = (file: any) => {
    const index = files.indexOf(file);

    setFiles((prevFiles) => {
      const updatedFiles = [...prevFiles];

      updatedFiles.splice(index, 1);

      return updatedFiles;
    });
    abort(file);
  };

  const onReUpload = (file: any) => {
    file.status = 'init';
    file.percent = 0;
    onChange(file, files);
    normalizeReq(file);
    autoUpload();
  };

  const abort = (file: any) => {
    const { uid } = file;

    if (reqs.current[uid]) {
      delete reqs.current[uid];
    }
  };

  return (
    <div className="flex w-full flex-col items-center">
      <input
        ref={inputRef}
        accept={accept}
        className="hidden w-full"
        multiple={multiple}
        type="file"
        onChange={onInputChange}
      />
      {drag ? (
        <UploadDragger
          enable={enable}
          onClick={onClickTrigger}
          onFilesHandled={(files) => uploadFiles(Array.from(files))}
        />
      ) : (
        <div
          className="cursor-pointer rounded-lg border border-dashed border-gray-300 p-4 text-center"
          onClick={onClickTrigger}
        >
          {children}
        </div>
      )}
      {!!enable && !!showList && (
        <UploadList files={files} onCopy={onCopy} onDelete={onDelete} onUpload={onReUpload} />
      )}
    </div>
  );
};

export default W3qDeployer;
