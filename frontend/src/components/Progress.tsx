import React, { useEffect, useRef } from 'react';

interface ProgressComponentProps {
  percent?: number;
  chunks?: number;
}

const ProgressComponent: React.FC<ProgressComponentProps> = ({ chunks = 1, percent = 0 }) => {
  const innerRefs = useRef<HTMLDivElement[]>([]);

  useEffect(() => {
    const updateWidths = () => {
      innerRefs.current.forEach((ref) => {
        if (ref) {
          ref.style.width = `${(1 / chunks) * 100 - 0.1}%`;
        }
      });
    };

    updateWidths();
  }, [chunks]);

  return (
    <div className="flex w-full space-x-2">
      {Array.from({ length: chunks }).map((_, index) => (
        <div
          key={index}
          ref={(el) => {
            if (el && !innerRefs.current.includes(el)) {
              innerRefs.current.push(el);
            }
          }}
          className="transition-width h-1 rounded-full bg-gray-300 duration-200"
        >
          {index < percent && <div className="bg-primary h-full w-full rounded-full" />}
        </div>
      ))}
    </div>
  );
};

export default ProgressComponent;
