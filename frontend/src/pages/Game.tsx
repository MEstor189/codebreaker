import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Game.css';
import SymbolDisplay from '../components/SymbolDisplay/SymbolDisplay';
import SubmitButton from '../components/SubmitButton/SubmitButton';
import SubmitHistory from '../components/SubmitHistory/SubmitHistory';
import SymbolButtonField from '../components/SymbolButtonField/SymbolButtonField';
import { useWebSocket } from '../components/WebSocket/WebSocketContext';
import SymbolButton from '../components/SymbolButton/SymbolButton';

import { BsBoxArrowLeft } from "react-icons/bs";


const Game: React.FC = () => {
  const navigate = useNavigate();
  const [pressedSymbols, setPressedSymbols] = useState<string[]>([]);
  const [submittedHistory, setSubmittedHistory] = useState<{ symbols: string[]; correctPositions: number[] }[]>([]);
  const { isConnected, sendMessage } = useWebSocket();
  const [loadingRound, setLoadingRound] = useState(true);
  const [roundObj, setRoundObj] = useState<any>(null); 
  const [isNextButtonVisible, setIsNextButtonVisible] = useState(false); 

   const handleSymbolClick = (symbol: string) => {
    setPressedSymbols(prevSymbols => [...prevSymbols, symbol]);
  }; 

  const handleRemoveSymbol = (index: number) => {
    setPressedSymbols((prevSymbols) => prevSymbols.filter((_, i) => i !== index));
  };


  const handleSubmit = async () => {
    if (pressedSymbols.length === 0) {
      console.warn("Leere Eingabe – Guess wird nicht abgeschickt.");
      return; 
  }
    console.log('Submitted Symbols:', pressedSymbols);
    
    try {
        const response = await sendMessage("guess", { pressedSymbols });

        console.log("Start-Antwort:", response);

        const data = typeof response === "string" ? JSON.parse(response) : response;

        const correctPositions = data?.evaluatedGuess?.CorrectPositions ?? [];
        correctGuess(data.solved )

        setSubmittedHistory(prevHistory => [...prevHistory, { symbols: pressedSymbols, correctPositions }]);

    } catch (error) {
        console.error("Fehler beim Senden der Startnachricht:", error);
    }
    setPressedSymbols([]);
};

  const goToStartScreen = () =>{
    navigate('/StartScreen');
  }

  const handleNextLevel = async () => {
    setLoadingRound(true)
    try{
      const response = await sendMessage("nextLevel", { nextLevel:true  });
      const data = typeof response === "string" ? JSON.parse(response) : response;
      setRoundObj(data.roundstate)
      setLoadingRound(false)
      setSubmittedHistory([]);
      setIsNextButtonVisible(false)

    }catch (error){
      setLoadingRound(false)
      console.error(error)
    }
  }



  const correctGuess = async (solved: Boolean) => {
    if(solved){
      setIsNextButtonVisible(true);
    }else{
      setIsNextButtonVisible(false);
    }
  }


  const startRound = async () => {
    try {
      const response = await sendMessage("start", { });
      console.log("Start-Antwort:", response);
      
      setRoundObj(response);
      setLoadingRound(false);

    } catch (error) {
      setLoadingRound(false);
      console.error("Fehler beim Senden der Startnachricht:", error);
    }
  };

  useEffect(() => {
    const startGame = async () => {
      await startRound(); 
    };
    startGame();
  }, []); 


  useEffect(() => {
    if (roundObj !== null) {
      console.log("Aktualisiertes GameObj:", roundObj); 
    }
  }, [roundObj]);

  useEffect(() => {
    if (roundObj && roundObj.Level && pressedSymbols.length === roundObj.Level.Difficulty.CodeLength) {
      handleSubmit();
    }
  }, [pressedSymbols, roundObj]);

  
  return (
    <div className="game-container">
      <div className="left">
        <button id='homeButton'>
        <BsBoxArrowLeft onClick={goToStartScreen} />
        </button>

      </div>
      <div className="center">
        <div className="content">
          <div className="game-info-bar">
            <div className="info-item">
              <span className="info-label">Level</span>
              <span className="info-value">{loadingRound ? "Lade..." : roundObj ? roundObj.Level.Lvl : 0}</span>
            </div>
            <div className="info-item">
              <span className="info-label">Timer</span>
              <span className="info-value">XX:XX</span>
            </div>
            <div className="info-item">
              <span className="info-label">Score</span>
              <span className="info-value">XXXX</span>

            </div>
          </div>
          <div className='history'>
              <SubmitHistory submittedHistory={submittedHistory}></SubmitHistory>
          </div>
          <div className='display'>
            <SymbolDisplay symbols={pressedSymbols} onRemove={handleRemoveSymbol}></SymbolDisplay>

          </div>
          <div className="symbol-button-field-wrapper">
            <SymbolButtonField symbols={loadingRound? "Lade..." : roundObj.Level.Code.Runes} count={loadingRound? "": roundObj.Level.Difficulty.PSC} onClick={handleSymbolClick}  />
          </div>
        </div>
      </div>
      <div className="right"></div>
    </div>
  );
};

export default Game;