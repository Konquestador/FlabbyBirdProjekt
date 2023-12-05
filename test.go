package main

import (
	"gfx2"
	//~ "fmt"
	//~ "time"
)

func main(){
	var hoehe = 400
	gfx2.Fenster(1000, 800)
	gfx2.LadeBildMitColorKey (uint16(100), uint16(hoehe) , "./images/original/Frame-1.bmp", uint8(255), uint8(255),uint8(0))
	
	//~ gfx2.Kreis(uint16(180), uint16(464) ,uint16(64))
	//~ r,g,b := gfx2.GibPunktfarbe(uint16(100),uint16(405))
	//~ fmt.Println(r,g,b)
	gfx2.TastaturLesen1 ()
}


