package saeulen

import (
    "gfx2"
    "fmt"
    "zufallszahlen"
)

type data struct{
	xWert uint16
	breite uint16
	rXWert int
	hoehe uint16
	loch uint16
	geschwindigkeit uint16
	passed bool
	touch bool
	birdInHole bool
}

func New () *data {
	var s *data
	s = new(data)
	s.breite=100
	s.passed= false
	return s
}

func (s *data) SetzeZufallswerte() {
	zufallszahlen.Randomisieren()
	s.xWert = gfx2.Grafikspalten()
	s.hoehe = uint16(zufallszahlen.Zufallszahl(50,500))
	s.loch = uint16(zufallszahlen.Zufallszahl(200,300))
}

func (s *data) GibXWert()  uint16{
	return s.xWert
}

func (s * data) GibPassed() bool{
	return s.passed
	}

func (s *data) GibBreite() uint16 {
	return s.breite
}

func (s *data) GibHoehe() uint16 {
	return s.hoehe
}

func (s *data) GibLoch() uint16 {
	return s.loch
}

func (s * data) GibbirdInHole() bool{
	return s.birdInHole
	}

func (s * data) InHole(){
	s.birdInHole = true
}

func (s * data) NotInHole(){
	s.birdInHole = false
}

func (s * data) Passed(){
	s.passed = true
}

func (s * data) Touch(){
	s.touch = true
}

func (s *data) SetzeWerte(xWert,hoehe,loch uint16) {
	s.xWert = xWert
	s.hoehe = hoehe
	s.loch = loch
}  

func (s *data) String () string {
	var erg string
	erg = fmt.Sprintln(s.xWert,s.hoehe,s.loch)
	return erg
}

func (s *data) Move () {
	s.geschwindigkeit = 5
	nx:=int(s.xWert)-int(s.geschwindigkeit)+s.rXWert
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

func (s *data) Draw(birdposX int, width int) {
	var fensterhoehe uint16
	if gfx2.FensterOffen() {
		fensterhoehe = gfx2.Grafikzeilen()
	}
	if s.touch{
		gfx2.Vollrechteck(uint16(birdposX + width + 1),0,s.breite,s.hoehe)
		gfx2.Vollrechteck(uint16(birdposX + width + 1),s.hoehe+s.loch,s.breite,fensterhoehe-(s.hoehe+s.loch))
	}else{
		gfx2.Vollrechteck(s.xWert,0,s.breite,s.hoehe)
		gfx2.Vollrechteck(s.xWert,s.hoehe+s.loch,s.breite,fensterhoehe-(s.hoehe+s.loch))
	}
}
