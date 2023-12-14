package saeulen

import (
	"gfx2"
	"fmt"
	"zufallszahlen"
)

// Saeule repräsentiert eine Struktur zur Verwaltung von Säulen im Spiel.
type Saeule struct {
	xWert          uint16 // X-Koordinate der Säule
	breite         uint16 // Breite der Säule
	rXWert         int    // Relative X-Koordinate für die Bewegung
	hoehe          uint16 // Höhe der Säule
	loch           uint16 // Höhe der Lücke in der Säule
	geschwindigkeit uint16 // Geschwindigkeit der Säule
	passed         bool   // Flag, ob der Spieler die Säule passiert hat
	touch          bool   // Flag, ob die Säule berührt wurde
	birdInHole     bool   // Flag, ob der Vogel in der Lücke der Säule ist
}

// New erstellt eine neue Säule mit Standardwerten.
func New() *Saeule {
	var s *Saeule
	s = new(Saeule)
	s.breite = 100
	s.passed = false
	return s
}

// SetzeZufallswerte setzt zufällige Werte für die x-Koordinate, Höhe und Lückenhöhe der Säule.
func (s *Saeule) SetzeZufallswerte() {
	zufallszahlen.Randomisieren()
	s.xWert = gfx2.Grafikspalten()
	s.hoehe = uint16(zufallszahlen.Zufallszahl(50, 500))
	s.loch = uint16(zufallszahlen.Zufallszahl(200, 300))
}

// GibXWert gibt die x-Koordinate der Säule zurück.
func (s *Saeule) GibXWert() uint16 {
	return s.xWert
}

// GibPassed gibt zurück, ob der Spieler die Säule passiert hat.
func (s *Saeule) GibPassed() bool {
	return s.passed
}

// GibBreite gibt die Breite der Säule zurück.
func (s *Saeule) GibBreite() uint16 {
	return s.breite
}

// GibHoehe gibt die Höhe der Säule zurück.
func (s *Saeule) GibHoehe() uint16 {
	return s.hoehe
}

// GibLoch gibt die Höhe der Lücke in der Säule zurück.
func (s *Saeule) GibLoch() uint16 {
	return s.loch
}

// GibbirdInHole gibt zurück, ob der Vogel in der Lücke der Säule ist.
func (s *Saeule) GibbirdInHole() bool {
	return s.birdInHole
}

// InHole setzt das Flag birdInHole auf true.
func (s *Saeule) InHole() {
	s.birdInHole = true
}

// NotInHole setzt das Flag birdInHole auf false.
func (s *Saeule) NotInHole() {
	s.birdInHole = false
}

// Passed setzt das Flag passed auf true.
func (s *Saeule) Passed() {
	s.passed = true
}

// Touch setzt das Flag touch auf true.
func (s *Saeule) Touch() {
	s.touch = true
}

// SetzeWerte setzt die Werte der x-Koordinate, Höhe und Lückenhöhe für die Säule.
func (s *Saeule) SetzeWerte(xWert, hoehe, loch uint16) {
	s.xWert = xWert
	s.hoehe = hoehe
	s.loch = loch
}

// String gibt eine Zeichenkettenrepräsentation der Werte der Säule zurück.
func (s *Saeule) String() string {
	return fmt.Sprintln(s.xWert, s.hoehe, s.loch)
}

// Move aktualisiert die Position der Säule basierend auf ihrer Geschwindigkeit.
func (s *Saeule) Move() {
	s.geschwindigkeit = 5
	nx := int(s.xWert) - int(s.geschwindigkeit) + s.rXWert
	if nx > 0 {
		s.xWert = s.xWert - s.geschwindigkeit
	} else if nx < -100 {
		s.xWert = 20000
	} else {
		s.xWert = 0
		s.breite = 100 - uint16(-1*nx)
		s.rXWert = nx
	}
}

// Draw erstellt die Säule im Grafikfenster.
func (s *Saeule) Draw(birdposX int, width int) {
	var fensterhoehe uint16
	if gfx2.FensterOffen() {
		fensterhoehe = gfx2.Grafikzeilen()
	}
	if s.touch {
		gfx2.Vollrechteck(uint16(birdposX+width+1), 0, s.breite, s.hoehe)
		gfx2.Vollrechteck(uint16(birdposX+width+1), s.hoehe+s.loch, s.breite, fensterhoehe-(s.hoehe+s.loch))
	} else {
		gfx2.Vollrechteck(s.xWert, 0, s.breite, s.hoehe)
		gfx2.Vollrechteck(s.xWert, s.hoehe+s.loch, s.breite, fensterhoehe-(s.hoehe+s.loch))
	}
}
