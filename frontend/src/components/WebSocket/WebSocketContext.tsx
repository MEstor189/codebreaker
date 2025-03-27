import React, { createContext, useContext, useState, useEffect, useRef, ReactNode } from 'react';

interface WebSocketContextType {
  isConnected: boolean;
  sendMessage: (type: string, payload: any) => Promise<string>;
}

// WebSocketContext stellt den Kontext für die WebSocket-Verbindung bereit
const WebSocketContext = createContext<WebSocketContextType | undefined>(undefined);

// WebSocketProvider stellt die WebSocket-Verbindung und den Zustand zur Verfügung
interface WebSocketProviderProps {
  children: ReactNode;  // children muss den Typ ReactNode haben, der alle möglichen React-Komponenten abdeckt
}

export const WebSocketProvider: React.FC<WebSocketProviderProps> = ({ children }) => {
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    socketRef.current = new WebSocket('ws://localhost:8080/ws');

    socketRef.current.onopen = () => {
      setIsConnected(true);
      console.log("WebSocket connected");
    };

    socketRef.current.onclose = () => {
      setIsConnected(false);
      console.log("WebSocket disconnected");
    };

    socketRef.current.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    return () => {
      socketRef.current?.close();
    };
  }, []);

  const sendMessage = (type: string, payload: any) => {
    return new Promise<string>((resolve, reject) => {
      if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
        const message = JSON.stringify({ type, payload });
        socketRef.current.send(message);
  
        socketRef.current.onmessage = (event) => {
          // Hier wird event.data angenommen, dass es ein JSON-String ist
          try {
            const parsedResponse = JSON.parse(event.data); // Parsen des JSON-Strings
            resolve(parsedResponse); // Gibt das geparste Objekt zurück
          } catch (e) {
            reject("Fehler beim Parsen der Antwort: " + e);
          }
        };
      } else {
        reject("WebSocket is not connected");
      }
    });
  };

  return (
    <WebSocketContext.Provider value={{ isConnected, sendMessage }}>
      {children}  {/* Hier wird die children-Prop verwendet */}
    </WebSocketContext.Provider>
  );
};

export const useWebSocket = () => {
  const context = useContext(WebSocketContext);
  if (!context) {
    throw new Error('useWebSocket must be used within a WebSocketProvider');
  }
  return context;
};
