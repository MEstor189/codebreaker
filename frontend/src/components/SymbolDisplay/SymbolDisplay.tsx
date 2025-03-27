import React from 'react';
import './SymbolDisplay.css';

interface SymbolDisplayProps {
  symbols: Array<string>;
}

const SymbolDisplay: React.FC<SymbolDisplayProps> = ({ symbols }) => {
  return (
    <div className="symbol-display">
      {symbols.map((symbol, index) => (
        <div key={index} className="symbol-item">
          {symbol}
        </div>
      ))}
    </div>
  );
};

export default SymbolDisplay;