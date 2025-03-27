import React from 'react';
import './SymbolButton.css';

interface SymbolButtonProps {
  symbol: string; // Der Pfad zum Symbol
  onClick: () => void; // Callback-Funktion für Klick-Ereignis
}

const SymbolButton: React.FC<SymbolButtonProps> = ({ symbol, onClick }) => {
  return (
    <button className="symbol-button" onClick={onClick}>
      {symbol}
    </button>
  );
};

export default SymbolButton;