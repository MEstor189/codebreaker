export interface WebSocketMessage {
    type: MessageType;      // Type der Nachricht
    content?: string;       // Optionaler Inhalt für die Nachricht
    score?: number;         // Optionaler Score
    highScores?: number[];  // Optionales Array für Highscores
}

export type MessageType = 
    | 'start'                  // Nachricht zum Starten eines Spiels
    | 'guess'                  // Nachricht für einen Rateversuch
    | 'highscore'              // Nachricht zum Abrufen der Highscores
    | 'message'                // Allgemeine Nachricht (z.B. bei einem Fehler)
    | 'update'                 // Nachricht zur Aktualisierung des Spiels
    | 'end';                   // Nachricht zum Beenden des Spiels
