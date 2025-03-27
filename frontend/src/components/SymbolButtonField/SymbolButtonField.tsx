import SymbolButton from "../SymbolButton/SymbolButton";

type ButtonLeisteProps = {
    symbols: number[];
    count: number;
    onClick: (symbol: string) => void;
  };
  
  export default function ButtonLeiste({ symbols, count, onClick }: ButtonLeisteProps) {
    const symbolsToStr: string[] = [];
    for (let i = 0; i < symbols.length; i++) {
      const element = symbols[i];
      symbolsToStr.push(String.fromCharCode(element))
    }


    return (
      <div style={{ display: "flex", gap: "10px", justifyContent: "center" }}>
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