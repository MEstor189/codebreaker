import React, { useState } from 'react';
import './EndRoundPopUp.css'; 
import { IoHomeSharp, IoSave  } from "react-icons/io5";

interface EndRoundPopUpProps {
    isOpen: boolean;          
    onClose: () => void;  
    onSubmit:(name:string, score:string) => void;    
    level: string;            
    score:string;
}

const EndRoundPopUp: React.FC<EndRoundPopUpProps> = ({ isOpen, onClose, onSubmit,score }) => {
    const [username, setUsername] = useState(''); 
    const handleChange = (e: { target: { value: React.SetStateAction<string>; }; }) => {
        setUsername(e.target.value); 
    };
    if (!isOpen) return null;

    return (
        <div className="popup-overlay" onClick={onClose}>
            <div className="popup-content" onClick={e => e.stopPropagation()}>
                <div id="title">
                <h2 id='title-info'>Defeated</h2>
                </div>
                <div className="popup-body">
                    <div className='pop-info-item'>
                        <span id="pop-score-label">Score:</span>
                        <span id="pop-score-value">{score}</span>
                    </div>
                    <div className='pop-info-item'>
                    <span id="pop-score-label">Name:</span>
                        <input 
                        id='usernameInput' 
                        type='text'
                        value={username}
                        placeholder='name'
                        onChange={handleChange} >

                        </input>
                    </div>
                </div>
                <div className='buttonDiv'>
                    <button id='pop-HomeButton' onClick={onClose}><IoHomeSharp /></button>
                    <button id="pop-SubmitNameButton" onClick={() => onSubmit(username,score)}>
                        <IoSave/>
                    </button>
                </div>
            </div>
        </div>
    );
};

export default EndRoundPopUp;