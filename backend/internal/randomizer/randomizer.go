package randomizer

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/BurntSushi/toml"
)

type ConfigSymbols struct {
	Symbols []string `toml:"symbols`
}

func LoadSymbols(filename string) ([]string, error) {
	var config ConfigSymbols
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return nil, err
	}
	return config.Symbols, nil
}

func GenerateRandomRunes(symbols []string, size int) []rune {
	randomRunes := make([]rune, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		randomIndex := rand.Intn(len(symbols))
		randomRunes[i] = []rune(symbols[randomIndex])[0]
	}
	return randomRunes
}

func Randomize(rand *rand.Rand, size int, latin []rune) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString(string(latin[rand.Intn(len(latin))]))
	}
	s := buffer.String()
	return s
}

/*
Frage 1: Zusätzliche Parameter für die Spielschwierigkeit
Code-Wiederholungen: Erlaube versteckte Wiederholungen von Symbolen im Code. Zum Beispiel könnte ein Code „3-1-1-2“ sein, was zusätzliche Komplexität hinzufügt.

Verschiedene Symbole: Füge neue Symboltypen wie Sonderzeichen oder Buchstaben hinzu, um die Vielfalt der Codes zu erhöhen.

Zeitdruck: Reduziere die verfügbare Zeit pro Runde bei steigendem Level, um den Druck zu erhöhen.

Fehleranfälligkeit: Erlaube dem Spieler, nur eine begrenzte Anzahl von Fehlversuchen pro Level, oder mache die Fehlerkosten höher (weniger Punkte pro Fehler).

Restriktionen bei Guessing: Füge Einschränkungen hinzu, wie zum Beispiel, dass die Spieler nicht mehr als eine bestimmte Anzahl von gleichen Symbolen in einer Rate-Sequenz verwenden dürfen.

Dynamische Code-Komplexität: Variiere die Code-Länge und die Anzahl der möglichen Symbole dynamisch, abhängig von der bisherigen Leistung des Spielers.

Frage 2: Skalierung des Systems bis ins Unendliche
Modulares Design: Entwickle das System mit klaren Modulen, die leicht erweitert werden können. Jedes Modul könnte zum Beispiel für das Hinzufügen neuer Symbole, das Anpassen der Code-Länge oder das Implementieren neuer Vergleichsmechanismen verantwortlich sein.

Dynamische Schwierigkeitskurve: Implementiere ein dynamisches Schwierigkeitsgrad-System, das sich an die Spielerleistung anpasst. Dies könnte so gestaltet sein, dass es die Schwierigkeit in kleinen Schritten erhöht, basierend auf dem Spielergebnis und der benötigten Zeit.

Konfigurable Parameter: Halte alle schwierigen Parameter in einer zentralen, konfigurierbaren Übersicht. Auf diese Weise kannst du die Level-Parameter einfach anpassen, ohne dass die Codebasis geändert werden muss.

Event-Driven Architecture: Verwende eine event-gesteuerte Architektur, wo Spielereignisse und Schwierigkeitsanpassungen über Events kommuniziert werden, um die Logik voneinander zu entkoppeln.

Zufällige Elemente: Integriere Zufälligkeit nicht nur in der Code-Generierung, sondern auch in der Bestimmung der nächsten Levelparameter. Zum Beispiel kann ein Spieler in einem Level zusätzlichen Bonus oder eine Herausforderung erhalten, die nicht vorhersehbar ist.

Hold & Scale Prinzip: Implementiere die Möglichkeit, dass Spieler bestimmte Levels "halten" können, bevor sie zu schwierigeren übergehen. Dies sorgt für kontinuierliche Herausforderungen und gibt Spielern das Gefühl von Fortschritt.


*/
