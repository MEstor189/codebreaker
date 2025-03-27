export const connectWebSocket = (url: string): WebSocket => {
    const socket = new WebSocket(url);

    socket.onopen = () => {
        console.log("WebSocket-Verbindung hergestellt!");
    };

    socket.onclose = () => {
        console.log("WebSocket-Verbindung geschlossen!");
    };

    socket.onerror = (error) => {
        console.error("WebSocket-Fehler:", error);
    };

    return socket;
};

export const sendMessage = (socket: WebSocket, message: object) => {
    if (socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify(message));
    } else {
        console.error("WebSocket ist nicht geöffnet. Nachricht konnte nicht gesendet werden.");
    }
};