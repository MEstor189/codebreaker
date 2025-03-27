import React from 'react';
import './SubmitButton.css';

interface SubmitButtonProps {
  onClick: () => void; // Callback-Funktion für den Button-Klick
}

const SubmitButton: React.FC<SubmitButtonProps> = ({ onClick }) => {
  return (
    <button className="submit-button" onClick={onClick}>
      Submit
    </button>
  );
};

export default SubmitButton;