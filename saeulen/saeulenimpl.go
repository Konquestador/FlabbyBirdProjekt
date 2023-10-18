package saeulen

import (
    "gfx2"
    "fmt"
    "zufallszahlen"
    //"time"
)

type data struct{
	xWert uint16
	hoehe uint16
	loch uint16
	r,g,b,hr,hg,hb uint8
	highlight bool
	geschwindigkeit uint16
}

func New () *data {
	var s *data
	s = new(data)
	return s
}

func (s *data) SetzeZufallswerte() {
	zufallszahlen.Randomisieren()
	s.xWert = gfx2.Grafikspalten()
	s.hoehe = uint16(zufallszahlen.Zufallszahl(100,500))
	s.loch = uint16(zufallszahlen.Zufallszahl(200,300))
}

func (s *data) GibXWert()  uint16{
	return s.xWert
}

func (s *data) SetzeWerte(xWert,hoehe,loch uint16) {
	s.xWert = xWert
	s.hoehe = hoehe
	s.loch = loch
}  

func (s *data) String () string {
	var erg string
	erg = fmt.Sprintln(s.xWert,s.hoehe,s.loch,s.r,s.g,s.b,s.highlight)
	return erg
}

func (s *data) Move (sleep_time int) {
	s.geschwindigkeit = 3
	if s.xWert != 0 {
		s.xWert = s.xWert - s.geschwindigkeit
		 }	
}

func (s *data) Draw() {
	var fensterhoehe uint16
	if gfx2.FensterOffen() {
		fensterhoehe = gfx2.Grafikzeilen()
	}
	if s.highlight {
		gfx2.Stiftfarbe(s.hr,s.hg,s.hb)
	} else{
		gfx2.Stiftfarbe(s.r,s.g,s.b)
	}
	gfx2.Vollrechteck(s.xWert,0,100,s.hoehe)
	gfx2.Vollrechteck(s.xWert,s.hoehe+s.loch,100,fensterhoehe-(s.hoehe+s.loch))
    gfx2.LadeBildMitColorKey(s.xWert-186, s.hoehe, "saeule.bmp", uint8(237), uint8(28), uint8(36))
}

//~ func (s *data) Kollision() {
	//~ if s.xWert == 
