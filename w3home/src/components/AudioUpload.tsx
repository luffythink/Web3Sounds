import type { ChangeEvent } from 'react';

import React, { useRef } from 'react';

interface AudioUploadProps {
  accept?: string;
  onFileChange?: (file: File) => void;
}

const AudioUpload: React.FC<AudioUploadProps> = ({ accept = 'audio/*', onFileChange }) => {
  const inputRef = useRef<HTMLInputElement>(null);

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (onFileChange && event.target.files) {
      const file = event.target.files[0];

      if (file) {
        onFileChange(file);
      }
    }
  };

  return <input ref={inputRef} accept={accept} type="file" onChange={handleInputChange} />;
};

export default AudioUpload;
