
import './SubmitHistory.css';
import React, { useState, useRef, useLayoutEffect } from "react";

interface SubmitHistoryProps {
  submittedHistory: { symbols: string[], correctPositions: number[] }[];
}

const SubmitHistory: React.FC<SubmitHistoryProps> = ({ submittedHistory }) => {
  const historyRef = useRef<HTMLDivElement>(null);

  useLayoutEffect(() => {
    if (historyRef.current) {
      historyRef.current.scrollTop = historyRef.current.scrollHeight;
    }
  }, [submittedHistory]);

  return (
    <div className="submit-history">
      <div ref={historyRef} className="symbols-container">
        {submittedHistory.map((entry, index) => (
          <div key={index} className="symbols-array">
            {entry.symbols.map((symbol, index2) => {
              const isCorrect = entry.correctPositions.includes(index2 + 1);
              return (
                <span
                  key={index2}
                  className="symbol-itemHistory"
                  style={{
                    marginRight: "5px",
                    backgroundColor: isCorrect ? "green" : "transparent",
                    color: isCorrect ? "white" : "black",
                  }}
                >
                  {symbol}
                </span>
              );
            })}
          </div>
        ))}
      </div>
    </div>
  );
};

export default SubmitHistory;