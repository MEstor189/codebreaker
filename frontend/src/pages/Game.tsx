import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './Game.css';
import SymbolDisplay from '../components/SymbolDisplay/SymbolDisplay';
import SubmitHistory, { SubmittedHistoryEntry } from '../components/SubmitHistory/SubmitHistory';
import SymbolButtonField from '../components/SymbolButtonField/SymbolButtonField';
import { useWebSocket } from '../components/WebSocket/WebSocketContext';

import { BsBoxArrowLeft } from "react-icons/bs";
import SolvedPopUp from '../components/PopUp/SolvedPopUp';
import Timer from '../components/timer/timer';
import EndRoundPopUp from '../components/PopUp/EndRoundPopUp';


const Game: React.FC = () => {
  const navigate = useNavigate();
  const [pressedSymbols, setPressedSymbols] = useState<string[]>([]);
  const [submittedHistory, setSubmittedHistory] = useState<SubmittedHistoryEntry[]>([]);
  const { sendMessage } = useWebSocket();
  const [loadingRound, setLoadingRound] = useState(true);
  const [roundObj, setRoundObj] = useState<any>(null); 
  const [ ,setIsNextButtonVisible] = useState(false);
  const [ ,setPopupOpen] = useState(false);
  const [timerValue, setTimerValue] = useState(0);
  const [timerStarted, setTimerStarted] = useState(false);
  const [remainingTime, setRemainingTime] = useState<number>(0);
  const [score, setScore] = useState<string>("0")
  const [roundScore, setRoundScore] = useState<string>("0")
  const [lvlCleared, setLevelCleared] = useState(false);
  const [showEndRoundPopUp, setShowEndRoundPopUp] = useState(false)


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

        setRoundObj(data.state)
        setTimerStarted(true)

        correctGuess(data.state.Solved )

        setSubmittedHistory(prevHistory => [...prevHistory, { 
          symbols: pressedSymbols,
          correctPositions: data?.state?.ComparisonResultNormal?.Positions ?? 0,
          correctCount:data?.state?.ComparisonResultNormal?.Contains ?? 0
        }
      ]);

    } catch (error) {
        console.error("Fehler beim Senden der Startnachricht:", error);
    }
    setPressedSymbols([]);
};

  const goToStartScreen = () =>{
    navigate('/StartScreen');
  }

  const handleNextLevel = async () => {
    console.log("neue runde")
    setLoadingRound(true)
    setLevelCleared(false)
    try{
      const response = await sendMessage("nextLevel", { nextLevel:true  });
      const data = typeof response === "string" ? JSON.parse(response) : response;
      console.log("Response: ",response)
      setRoundObj(data.state)

      console.log("Obj: ",roundObj)
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
      setLevelCleared(true) 

    }else{
      setIsNextButtonVisible(false);
      setPopupOpen(false)

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

  const handleTimeLeftChange = (timeLeft: number) => {
    setRemainingTime(timeLeft);
  };

  useEffect(() => {
    const startGame = async () => {
      await startRound();
       
    };
    startGame();
  }, []); 


  useEffect(() => {
    if (roundObj !== null && !lvlCleared) {
      setTimerValue(Number(roundObj.Difficulty.Timer));
      setRemainingTime(timerValue)
    }
  }, [roundObj]);

  useEffect(() => {
    if (roundObj && roundObj.Level && pressedSymbols.length === roundObj.Difficulty.CodeLength) {
      handleSubmit();
    }
  }, [pressedSymbols, roundObj]);

  useEffect(() => {
    if (lvlCleared) {
      setScore(roundObj.LvLScore)
      setRoundScore(roundObj.RoundScore)
      setIsNextButtonVisible(true);
      setPopupOpen(true)
      setTimerStarted(false)
      console.log(score)
    }
}, [lvlCleared]);


const handleNameSubmit = async (name: string, score: string) => {
  console.log("neue runde")
  try{
    const response = await sendMessage("highscoreEntry", {name,score});
    console.log("Submit: ",response)
  }catch (error){

    console.error(error)
  }
  console.log("Eingegebener Name:", name);
};

const goToEndPopUp = () => {
  setShowEndRoundPopUp(true);
};




  const isLoading = loadingRound ? "Lade" : null;


  
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
              <span className="info-value">{loadingRound ? "Lade..." : roundObj ? roundObj.Level : 0}</span>
            </div>
            <div className="info-item">
              <span className="info-label">Timer</span>
              {timerStarted && (
              <Timer 
              initialTime={timerValue}
              onExpire={undefined}
              onTimeLeftChange={handleTimeLeftChange}
              />)}
              
            </div>
            <div className="info-item">
              <span className="info-label">Score</span>
              <span className="info-value">{roundScore}</span>

            </div>
          </div>
          <div className='history'>
              <SubmitHistory submittedHistory={submittedHistory}></SubmitHistory>
          </div>
          <div className='display'>
            <SymbolDisplay symbols={pressedSymbols} onRemove={handleRemoveSymbol}></SymbolDisplay>

          </div>
          <div className="symbol-button-field-wrapper">
            <SymbolButtonField symbols={loadingRound? "Lade..." : roundObj.Runes} count={loadingRound? "": roundObj.Difficulty.PSC} onClick={handleSymbolClick}  />
          </div>
          {lvlCleared && (
          <SolvedPopUp
              isOpen={isLoading || roundObj.Solved}
              onClose={goToEndPopUp}
              onNext={handleNextLevel}
              level={isLoading || roundObj.Level}
              trys={isLoading || roundObj.Trys}
              time={remainingTime}
              score={score}>
          </SolvedPopUp>
          )}
          <EndRoundPopUp isOpen={showEndRoundPopUp} onClose={goToStartScreen} onSubmit={handleNameSubmit} level={isLoading || roundObj.Level} score={score}></EndRoundPopUp>
        </div>
      </div>
      <div className="right"></div>
    </div>
  );
};

export default Game;