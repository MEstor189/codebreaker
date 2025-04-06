
import './SubmitHistory.css';
import React, { useRef, useLayoutEffect } from "react";

export interface SubmittedHistoryEntry {
  symbols: string[]; // Ein einzelner Symbolstring
  correctPositions: number; // Die korrekte Position
  correctCount: number; // Die Anzahl der richtigen Antworten
}

interface SubmitHistoryProps {
  submittedHistory: SubmittedHistoryEntry[]; 
}

const MAX_ROWS = 7;
const MAX_SYMBOLS = 4;

const SubmitHistory: React.FC<SubmitHistoryProps> = ({ submittedHistory }) => {
  const historyRef = useRef<HTMLDivElement>(null);

  useLayoutEffect(() => {
    if (historyRef.current) {
      historyRef.current.scrollTop = historyRef.current.scrollHeight;
    }
  }, [submittedHistory]);

  const filledHistory = [
    ...submittedHistory,
    ...Array(Math.max(0, MAX_ROWS - submittedHistory.length)).fill({ symbols: Array(MAX_SYMBOLS).fill(null), correctPositions: 0, correctCount: 0 }),
  ];

  return (
    <div className="submit-history" ref={historyRef}>
      {filledHistory.map((entry, index) => (
        <div key={index} className="history-row">

          <div className={`history-box left-box ${entry.symbols.some((s: any) => s) ? '' : 'placeholder-box'}`}> 
                {entry.symbols.some((s: any) => s) ? entry.correctPositions : ''}
          </div>

          <div className="symbols-array">
            {entry.symbols.map((symbol: any, index2: React.Key | null | undefined) => (
              <span key={index2} className={`symbol-itemHistory ${symbol ? '' : 'placeholder'}`}>
                {symbol || ''}
              </span>
            ))}
          </div>
          <div className={`history-box right-box ${entry.symbols.some((s: any) => s) ? '' : 'placeholder-box'}`}> 
                {entry.symbols.some((s: any) => s) ? entry.correctCount : ''}
          </div>
        </div>
      ))}
    </div>
  );
};

export default SubmitHistory;