import React from 'react';
import './SymbolDisplay.css';

interface SymbolDisplayProps {
  symbols: Array<string>;
  onRemove: (index: number) => void;
}

const MAX_SYMBOLS = 4;

const SymbolDisplay: React.FC<SymbolDisplayProps> = ({ symbols, onRemove }) => {
  const filledSymbols = [
    ...symbols,
    ...Array(MAX_SYMBOLS - symbols.length).fill(null),
  ];

  return (
    <div className="symbol-display">
      {filledSymbols.map((symbol, index) => (
        <div 
          key={index} 
          className={`symbol-item ${symbol ? '' : 'placeholder-display'}`} 
          onClick={symbol ? () => onRemove(index) : undefined}
        >
          {symbol || ''}
        </div>
      ))}
    </div>
  );
};

export default SymbolDisplay;