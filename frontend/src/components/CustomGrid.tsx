import React, { useState } from 'react';

const getRandomDarkGreenColor = (): string => {
  const greenValue = Math.floor(Math.random() * 128) + 128;
  const redBlueValue = Math.floor(Math.random() * 128);
  return `rgba(${redBlueValue}, ${greenValue}, ${redBlueValue}, 0.5)`;
};

const CustomGrid: React.FC = () => {
  const rows = 9;
  const cols = 49;

  const initialGridData: string[][] = Array.from({ length: rows }, () =>
    Array.from({ length: cols }, () => getRandomDarkGreenColor())
  );

  const [gridData, setGridData] = useState(initialGridData);

  const mainDeepColor = '#1b632e';
  // W
  gridData[2][4] = mainDeepColor;
  gridData[3][4] = mainDeepColor;
  gridData[4][4] = mainDeepColor;
  gridData[5][4] = mainDeepColor;
  gridData[6][4] = mainDeepColor;
  gridData[6][5] = mainDeepColor;
  gridData[3][6] = mainDeepColor;
  gridData[4][6] = mainDeepColor;
  gridData[5][6] = mainDeepColor;
  gridData[6][6] = mainDeepColor;
  gridData[6][7] = mainDeepColor;
  gridData[2][8] = mainDeepColor;
  gridData[3][8] = mainDeepColor;
  gridData[4][8] = mainDeepColor;
  gridData[5][8] = mainDeepColor;
  gridData[6][8] = mainDeepColor;
  // E
  gridData[2][10] = mainDeepColor;
  gridData[3][10] = mainDeepColor;
  gridData[4][10] = mainDeepColor;
  gridData[5][10] = mainDeepColor;
  gridData[6][10] = mainDeepColor;
  gridData[2][11] = mainDeepColor;
  gridData[2][12] = mainDeepColor;
  gridData[2][13] = mainDeepColor;
  gridData[4][11] = mainDeepColor;
  gridData[4][12] = mainDeepColor;
  gridData[4][13] = mainDeepColor;
  gridData[6][11] = mainDeepColor;
  gridData[6][12] = mainDeepColor;
  gridData[6][13] = mainDeepColor;
  // B
  gridData[2][15] = mainDeepColor;
  gridData[3][15] = mainDeepColor;
  gridData[4][15] = mainDeepColor;
  gridData[5][15] = mainDeepColor;
  gridData[6][15] = mainDeepColor;
  gridData[2][16] = mainDeepColor;
  gridData[2][17] = mainDeepColor;
  gridData[2][18] = mainDeepColor;
  gridData[4][16] = mainDeepColor;
  gridData[4][17] = mainDeepColor;
  gridData[4][18] = mainDeepColor;
  gridData[6][16] = mainDeepColor;
  gridData[6][17] = mainDeepColor;
  gridData[6][18] = mainDeepColor;
  gridData[3][18] = mainDeepColor;
  gridData[4][18] = mainDeepColor;
  gridData[5][18] = mainDeepColor;
  gridData[6][18] = mainDeepColor;
  // 3
  gridData[2][20] = mainDeepColor;
  gridData[2][21] = mainDeepColor;
  gridData[2][22] = mainDeepColor;
  gridData[2][23] = mainDeepColor;
  gridData[4][20] = mainDeepColor;
  gridData[4][21] = mainDeepColor;
  gridData[4][22] = mainDeepColor;
  gridData[4][23] = mainDeepColor;
  gridData[6][20] = mainDeepColor;
  gridData[6][21] = mainDeepColor;
  gridData[6][22] = mainDeepColor;
  gridData[3][23] = mainDeepColor;
  gridData[4][23] = mainDeepColor;
  gridData[5][23] = mainDeepColor;
  gridData[6][23] = mainDeepColor;
  gridData[3][25] = mainDeepColor;
  gridData[5][25] = mainDeepColor;
  gridData[2][27] = mainDeepColor;
  gridData[3][27] = mainDeepColor;
  gridData[4][27] = mainDeepColor;
  gridData[5][27] = mainDeepColor;
  gridData[6][27] = mainDeepColor;
  // :
  gridData[2][29] = mainDeepColor;
  gridData[3][29] = mainDeepColor;
  gridData[4][29] = mainDeepColor;
  gridData[5][29] = mainDeepColor;
  gridData[6][29] = mainDeepColor;
  // U
  gridData[2][31] = mainDeepColor;
  gridData[3][31] = mainDeepColor;
  gridData[4][31] = mainDeepColor;
  gridData[5][31] = mainDeepColor;
  gridData[6][31] = mainDeepColor;
  gridData[6][32] = mainDeepColor;
  gridData[6][33] = mainDeepColor;
  gridData[2][34] = mainDeepColor;
  gridData[3][34] = mainDeepColor;
  gridData[4][34] = mainDeepColor;
  gridData[5][34] = mainDeepColor;
  gridData[6][34] = mainDeepColor;
  // R
  gridData[2][36] = mainDeepColor;
  gridData[3][36] = mainDeepColor;
  gridData[4][36] = mainDeepColor;
  gridData[5][36] = mainDeepColor;
  gridData[6][36] = mainDeepColor;
  gridData[2][37] = mainDeepColor;
  gridData[2][38] = mainDeepColor;
  gridData[2][39] = mainDeepColor;
  gridData[3][39] = mainDeepColor;
  gridData[4][37] = mainDeepColor;
  gridData[4][38] = mainDeepColor;
  gridData[4][39] = mainDeepColor;
  gridData[5][38] = mainDeepColor;
  gridData[6][39] = mainDeepColor;
  // L
  gridData[2][41] = mainDeepColor;
  gridData[3][41] = mainDeepColor;
  gridData[4][41] = mainDeepColor;
  gridData[5][41] = mainDeepColor;
  gridData[6][41] = mainDeepColor;
  gridData[6][42] = mainDeepColor;
  gridData[6][43] = mainDeepColor;
  gridData[6][44] = mainDeepColor;

  const handleCellClick = (rowIndex: number, colIndex: number) => {
    // alert(`Clicked cell at row ${rowIndex}, column ${colIndex}`);
    alert('声音播放待开发……');
  };

  return (
    <div className="flex items-center justify-center">
      <div
        className="grid gap-1"
        style={{ gridTemplateColumns: `repeat(${cols}, minmax(0, 1fr))` }}
      >
        {gridData.map((row, rowIndex) =>
          row.map((color, colIndex) => (
            <div
              key={`${rowIndex}-${colIndex}`}
              className="h-3 w-3 cursor-pointer rounded-sm"
              style={{ backgroundColor: color }}
              onClick={() => handleCellClick(rowIndex, colIndex)}
            ></div>
          ))
        )}
      </div>
    </div>
  );
};

export default CustomGrid;
