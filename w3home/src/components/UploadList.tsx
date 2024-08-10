import React from 'react';

import { BookCopy, FileAudio, FileVolume, Orbit, Upload, X } from 'lucide-react';

import MyProgress from './Progress';

interface File {
  uid: string;
  status: 'pending' | 'success' | 'failure';
  type: string;
  url: string;
  name: string;
  percent: number;
  totalChunks: number;
}

interface UploadListProps {
  files: File[];
  onDelete: (file: File) => void;
  onCopy: (url: string) => void;
  onUpload: (file: File) => void;
}

const UploadList: React.FC<UploadListProps> = ({ files, onCopy, onDelete, onUpload }) => {
  const isImage = (type: string) => {
    if (!type) return false;

    return type.includes('image');
  };

  return (
    <div className="mt-4 w-1/2 space-y-2">
      {files.map((file) => (
        <div
          key={file.uid}
          className={`flex items-center rounded border p-2 ${
            file.status === 'failure' ? 'border-red-500 text-red-500' : 'border-gray-300'
          }`}
        >
          <div className="flex h-12 w-12 flex-shrink-0 items-center justify-center">
            {file.status === 'pending' ? (
              <Orbit className="animate-spin text-xl" />
            ) : file.status === 'success' ? (
              isImage(file.type) ? (
                <Orbit />
              ) : (
                <FileVolume />
              )
            ) : file.status === 'failure' ? (
              <FileAudio className="text-2xl text-red-500" />
            ) : (
              <FileVolume />
            )}
          </div>

          <div className="ml-2 mr-2 flex-1 truncate font-bold">
            {file.status === 'success' ? (
              <a
                className="text-green-500 hover:text-green-400"
                href={file.url}
                rel="noopener noreferrer"
                target="_blank"
              >
                <span>{file.name}</span>
              </a>
            ) : (
              <span>{file.name}</span>
            )}
            {file.status === 'pending' && (
              <MyProgress chunks={file.totalChunks} percent={file.percent} />
            )}
          </div>

          {file.status === 'success' && (
            <span className="h-7 w-7 cursor-pointer text-lg" onClick={() => onCopy(file.url)}>
              <BookCopy />
            </span>
          )}
          {file.status === 'failure' && (
            <span className="h-7 w-7 cursor-pointer text-lg" onClick={() => onUpload(file)}>
              <Upload name="upload" />
            </span>
          )}
          {file.status !== 'pending' && file.status !== 'success' && (
            <span className="h-7 w-7 cursor-pointer text-lg" onClick={() => onDelete(file)}>
              <X />
            </span>
          )}
        </div>
      ))}
    </div>
  );
};

export default UploadList;
