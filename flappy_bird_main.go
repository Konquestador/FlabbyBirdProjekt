package main


//Imports
import (
	//"fmt"
    "gfx2" // Import the gfx2 library (replace "your-username" with the actual package path)
    //"time"
    //"math"
)

// Graviation


//Main
func main(){
	// variablen erstellen
	//var speed uint16 = 5
	//var gravity uint16 = 5
	var birdposx uint16 = 100
	var birdposy uint16 = 100
	//var previousTime = time.Now()
    //var currentTime = time.Now()
    //var deltaTime float64
    //var zaehler float64
	
//	var direction uint16 = 0
	
	
	// Erzeuge Fenster
	gfx2.Fenster(1000, 800)
	
	//UpdateAus() um Vogel-Fehler zu vermeiden
	gfx2.UpdateAus()
	
    //Hintergrund
	gfx2.Stiftfarbe(127,255,212)
	gfx2.Cls()
	
	// Vogel reinladen
//    gfx2.LadeBildMitColorKey(birdposx, birdposy, "Frame-1.bmp",0,0,0)
	gfx2.Transparenz(0)
	
	gfx2.LadeBildInsClipboard("Frame-1.bmp")
	gfx2.Clipboard_einfuegenMitColorKey (birdposx, birdposy, 255,0,0) 

	gfx2.MausLesen1()
	
	gfx2.UpdateAn() 
	//for {
		//gfx2.UpdateAus()
		//gfx2.Cls()
		//currentTime = time.Now()
		//deltaTime = currentTime.Sub(previousTime).Seconds()
		
		//zaehler += deltaTime
		
		//if zaehler >= 1 {
			//zaehler = math.Round(zaehler)
			//speed += uint16(zaehler) * gravity
			//birdposy += speed
			//}		
		
		////Ausgabe der Koordinaten zur Kotrolle
		//fmt.Println(birdposy)
		
		
		////Vogel mit neuen koordinaten reinladen
		//gfx2.Clipboard_einfuegenMitColorKey (birdposx, birdposy, 255,0,0)
		
		//previousTime = currentTime
		
		//gfx2.UpdateAn()
		//}
}
