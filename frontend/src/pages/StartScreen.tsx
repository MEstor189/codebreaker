import React, { useState } from 'react';
import './StartScreen.css';
import { useNavigate } from 'react-router-dom';
import { useWebSocket } from '../components/WebSocket/WebSocketContext';


function StartScreen() {
    const navigate = useNavigate();
    const { isConnected, sendMessage } = useWebSocket();

    const goToGameScreen = () => {
        //handleStart();
        console.log("Rüber")
        navigate('/game');

    };

    return (
        <div className="start-screen">
            <h1 className='title'>DiCodeBreaker</h1>
            <h1>WebSocket Verbindung: {isConnected ? "Verbunden ✅" : "Getrennt ❌"}</h1>
            <button className="play-button" onClick={goToGameScreen}>
                Start
            </button>
        </div>
    );
};

export default StartScreen;