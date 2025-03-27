import { useState, useEffect, useRef } from "react";

const useWebSocketConn = (url: string) => {
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    socketRef.current = new WebSocket(url);

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
      if (socketRef.current) {
        socketRef.current.close();
      }
    };
  }, [url]);

  return { isConnected };
};

export default useWebSocketConn;
