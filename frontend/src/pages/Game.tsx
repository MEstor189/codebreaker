import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Game.css';
import SymbolDisplay from '../components/SymbolDisplay/SymbolDisplay';
import SubmitButton from '../components/SubmitButton/SubmitButton';
import SubmitHistory from '../components/SubmitHistory/SubmitHistory';
import ButtonLeiste from '../components/SymbolButtonField/SymbolButtonField';
import { useWebSocket } from '../components/WebSocket/WebSocketContext';


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

  const handleRemoveLastSymbol = () => {
    setPressedSymbols(prevSymbols => prevSymbols.slice(0, -1));
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

  return (
    <div id='gameScreen'>
      <div className='gridContainer'>
        <div className='gridItem' ></div>
        <div className='gridItem' id='lvlDiv'>
          <h1 id='lvl' >Level: {loadingRound ? "Lade..." : roundObj ? roundObj.Level.Lvl : 0}</h1>
        </div>
        <div className='gridItem' ></div>
        <div className='gridItem' ></div>
        <div className='gridItem' id='history'>
          <h1>Submitted Symbols History</h1>
          <SubmitHistory submittedHistory={submittedHistory} />
        </div>
        <div className='gridItem' ></div>
        <div className='gridItem'id='delDiv' >
          <button id='delButton' onClick={handleRemoveLastSymbol}>⬅</button>
        </div>
        <div className='gridItem' id='displayDiv' >
          <SymbolDisplay symbols={pressedSymbols} />
        </div>
        <div className='gridItem' >
          <SubmitButton onClick={handleSubmit} />
        </div>
        <div className='gridItem'id='backDiv' >
          <button id='backButton' onClick={goToStartScreen}>Back to Start </button>
        </div>
        <div className='gridItem' >
          <ButtonLeiste symbols={loadingRound? "Lade..." : roundObj.Level.Code.Runes} count={loadingRound? "": roundObj.Level.Difficulty.PSC} onClick={handleSymbolClick}></ButtonLeiste>
        </div>
        <div className='gridItem' >
          {isNextButtonVisible && <button onClick={handleNextLevel} >Next Level</button>}
        </div>
      </div>
    </div>
  );
};

export default Game;