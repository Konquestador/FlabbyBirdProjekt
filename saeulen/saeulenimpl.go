package saeulen

import (
    "gfx2"
    "fmt"
    "zufallszahlen"
    //"time"
)

type data struct{
	xWert uint16
	breite uint16
	rXWert int
	hoehe uint16
	loch uint16
	r,g,b,hr,hg,hb uint8
	highlight bool
	geschwindigkeit uint16
}

func New () *data {
	var s *data
	s = new(data)
	s.breite=100
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

func (s *data) Move () {
	s.geschwindigkeit = 1
	nx:=int(s.xWert)-int(s.geschwindigkeit)+s.rXWert
	fmt.Println(nx)
	if nx>0 {
		s.xWert = s.xWert - s.geschwindigkeit
	}else if nx < (-100) {
		s.xWert = 20000
	}else{
		s.xWert=0
		s.breite=100-uint16(-1*nx)
		s.rXWert = nx
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
	gfx2.Vollrechteck(s.xWert,0,s.breite,s.hoehe)
	gfx2.Vollrechteck(s.xWert,s.hoehe+s.loch,s.breite,fensterhoehe-(s.hoehe+s.loch))
    gfx2.LadeBildMitColorKey(s.xWert-186, s.hoehe, "saeule.bmp", uint8(237), uint8(28), uint8(36))
}

//~ func (s *data) Kollision() {
	//~ if s.xWert == 
