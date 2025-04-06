import React from 'react';
import './SolvedPopUp.css'; 
import { IoHomeSharp } from "react-icons/io5";
import { FaRunning } from "react-icons/fa";
import { IoMdArrowRoundForward } from "react-icons/io";

interface SolvedPopUpProps {
    isOpen: boolean;          
    onClose: () => void;  
    onNext:() => void;    
    level: string;            
    trys: string;
    time: number;
    score:string;
}


const SolvedPopUp: React.FC<SolvedPopUpProps> = ({ isOpen, onClose, onNext, level,trys,time, score }) => {
    if (!isOpen) return null;

    const formatTime = (totalSeconds: number) => {
        const minutes = Math.floor(totalSeconds / 60);
        const seconds = totalSeconds % 60;
        return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
    };
    return (
        <div className="popup-overlay" onClick={onClose}>
            <div className="popup-content" onClick={e => e.stopPropagation()}>
                <div id="title">
                <h2 id='title-info'>Level<span className="title-highlight">{level}</span>cleared</h2>
                </div>
                <div className="popup-body">
                    <div className='pop-info-item'>
                        <span className="pop-info-label">Trys:</span>
                        <span className="pop-info-value">{trys}</span>
                    </div>
                    <div className='pop-info-item'>
                        <span className="pop-info-label">Time left:</span>
                        <span className="pop-info-value">{formatTime(time)}</span>
                    </div>
                    <div className='pop-info-item'>
                        <span id="pop-score-label">Level Score:</span>
                        <span id="pop-score-value">{score}</span>
                    </div>
                </div>
                <div className='buttonDiv'>
                    <button id='pop-HomeButton' onClick={onClose}><IoHomeSharp /></button>
                    <button id="pop-NextButton" onClick={onNext}>
                        <span className="icon-wrapper">
                            <FaRunning />
                            <IoMdArrowRoundForward />
                        </span>
                    </button>
                </div>
            </div>
        </div>
    );
};

export default SolvedPopUp;