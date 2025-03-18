import React, { useEffect, useState } from 'react';
import axios from 'axios';

const App: React.FC = () => {
    const [status, setStatus] = useState<string | null>(null);

    useEffect(() => {
        const fetchHealth = async () => {
            try {
                const response = await axios.get('/api/health');
                setStatus(response.data.status);
            } catch (error) {
                console.error('Error fetching health:', error);
            }
        };

        fetchHealth();
    }, []);

    return (
        <div>
            <h1>Codebreaker Arena</h1>
            {status ? <p>Status: {status}</p> : <p>Loading...</p>}
        </div>
    );
};

export default App;
