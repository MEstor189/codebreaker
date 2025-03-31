
import './SubmitHistory.css';
import React, { useState, useRef, useLayoutEffect } from "react";

interface SubmitHistoryProps {
  submittedHistory: { symbols: string[], correctPositions: number[] }[];
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
    ...Array(Math.max(0, MAX_ROWS - submittedHistory.length)).fill({ symbols: Array(MAX_SYMBOLS).fill(null) }),
  ];

  return (
    <div className="submit-history" ref={historyRef}>
      {filledHistory.map((entry, index) => (
        <div key={index} className="history-row">

          <div className={`history-box left-box ${entry.symbols.some((s: any) => s) ? '' : 'placeholder-box'}`}> 
                {entry.symbols.some((s: any) => s) ? 'L' : ''}
          </div>

          <div className="symbols-array">
            {entry.symbols.map((symbol: any, index2: React.Key | null | undefined) => (
              <span key={index2} className={`symbol-itemHistory ${symbol ? '' : 'placeholder'}`}>
                {symbol || ''}
              </span>
            ))}
          </div>
          <div className={`history-box right-box ${entry.symbols.some((s: any) => s) ? '' : 'placeholder-box'}`}> 
                {entry.symbols.some((s: any) => s) ? 'R' : ''}
          </div>
        </div>
      ))}
    </div>
  );
};

export default SubmitHistory;