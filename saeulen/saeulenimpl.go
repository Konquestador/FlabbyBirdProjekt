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
	s.loch = uint16(zufallszahlen.Zufallszahl(25,150))
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
	if s.xWert != 0 {
		s.xWert--
		//time.Sleep(time.Duration(sleep_time)* time.Microsecond)
		}	else {
			s.xWert--
			fmt.Println("hfdjhkfdfdhkg", s.xWert)
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
	gfx2.Vollrechteck(s.xWert,0,50,s.hoehe)
	gfx2.Vollrechteck(s.xWert,s.hoehe+s.loch,50,fensterhoehe-(s.hoehe+s.loch))
}
