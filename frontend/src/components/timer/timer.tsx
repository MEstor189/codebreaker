import React, { useState, useEffect } from 'react';

interface TimerProps {
    initialTime: number; // Zeit in Sekunden
    onExpire?: () => void;
    onTimeLeftChange?: (timeLeft: number) => void;
}

const Timer: React.FC<TimerProps> = ({ initialTime, onExpire, onTimeLeftChange }) => {
    const [timeLeft, setTimeLeft] = useState(initialTime);
    const [isRunning] = useState(true);

    useEffect(() => {
        setTimeLeft(initialTime);
    }, [initialTime]);

    useEffect(() => {
        let timer: ReturnType<typeof setInterval> | undefined;

        if (isRunning && timeLeft > 0) {
            timer = setInterval(() => {
                setTimeLeft(prevTime => {
                    const newTime = prevTime - 1;
                    if (onTimeLeftChange) onTimeLeftChange(newTime);
                    return newTime;
                });
            }, 1000);
        } else if (timeLeft === 0) {
            clearInterval(timer);
            if (onExpire) onExpire();
        }
        return () => clearInterval(timer); 
    }, [isRunning, timeLeft, onExpire]);

    const formatTime = (totalSeconds: number) => {
        const minutes = Math.floor(totalSeconds / 60);
        const seconds = totalSeconds % 60;
        return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
    };

    return (
        <span className="info-value">{formatTime(timeLeft)}</span>
    );
};

export default Timer;