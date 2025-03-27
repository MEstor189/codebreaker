import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import StartScreen from './pages/StartScreen';
import Game from './pages/Game';
import { WebSocketProvider } from './components/WebSocket/WebSocketContext';

function App() {
    return (
        <WebSocketProvider>        
            <Router>
                <Routes>
                    <Route path="/StartScreen"
                        element={<StartScreen />} />
                    <Route path="/Game"
                        element={<Game />} />
                </Routes>
            </Router>
        </WebSocketProvider>

    );
  };
  
  export default App;