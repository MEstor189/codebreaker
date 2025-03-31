import SymbolButton from "../SymbolButton/SymbolButton";
import "./SymbolButtonField.css"

type SymbolButtonFieldProps = {
    symbols: number[];
    count: number;
    onClick: (symbol: string) => void;
  };
  
  export default function SymbolButtonField({ symbols, count,  onClick }: SymbolButtonFieldProps) {
    const symbolsToStr: string[] = [];
    for (let i = 0; i < symbols.length; i++) {
      const element = symbols[i];
      symbolsToStr.push(String.fromCharCode(element))
    }


    return (
      <div className="symbol-button-field">
        {[...Array(count)].map((_, index) => (
          <SymbolButton
            key={index}
            symbol={(symbolsToStr[index]).toString()}
             onClick={() => onClick((symbolsToStr[index]).toString())} 
          />
        ))}
      </div>
    );
  }