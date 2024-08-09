import React, { useCallback, useState } from 'react';

import { Upload } from 'lucide-react';

interface UploadDraggerProps {
  enable?: boolean;
  onFilesHandled?: (files: FileList) => void;
  onClick?: () => void;
}

const UploadDragger: React.FC<UploadDraggerProps> = ({
  enable = true,
  onClick,
  onFilesHandled
}) => {
  const [dragging, setDragging] = useState(false);

  const handleDragEnter = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    setDragging(true);
    e.stopPropagation();
    e.preventDefault();
  }, []);

  const handleDragLeave = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    setDragging(false);
    e.stopPropagation();
    e.preventDefault();
  }, []);

  const handleDragOver = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.stopPropagation();
    e.preventDefault();
  }, []);

  const handleDrop = useCallback(
    (e: React.DragEvent<HTMLDivElement>) => {
      setDragging(false);
      e.stopPropagation();
      e.preventDefault();
      const files = e.dataTransfer.files;

      onFilesHandled && onFilesHandled(files);
    },
    [onFilesHandled]
  );

  return enable ? (
    <div
      className={`flex h-44 w-1/2 flex-col items-center justify-center border-2 border-dashed ${dragging ? 'border-green-500 bg-gray-200' : 'border-gray-400'} cursor-pointer hover:border-green-500`}
      onClick={onClick}
      onDragEnter={handleDragEnter}
      onDragLeave={handleDragLeave}
      onDragOver={handleDragOver}
      onDrop={handleDrop}
    >
      <Upload size={18} />
      <div className="">
        <span>Drop file here or click to upload</span>
      </div>
    </div>
  ) : (
    <div className="flex h-44 flex-col items-center justify-center border-2 border-dashed border-gray-300">
      <Upload size={18} />
      <div className="text-gray-300">
        <span>Drop file here or click to upload</span>
      </div>
    </div>
  );
};

export default UploadDragger;
