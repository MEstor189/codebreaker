import React from "react";
import "./SymbolButton.css";

interface SymbolButtonProps {
  symbol: string;
  onClick?: () => void;
}

const SymbolButton: React.FC<SymbolButtonProps> = ({ symbol, onClick }) => {
  return (
    <button className="symbol-button" onClick={onClick}>
      {symbol}
    </button>
  );
};

export default SymbolButton;