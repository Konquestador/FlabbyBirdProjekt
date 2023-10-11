package main

//Imports
import (
	"fmt"
    "gfx2"
    "time"
    //"math"
    "image"
	"imaging"
	"os"
)

//Main
func main(){
	// Variablen erstellen
	var speed int = 0
	var birdposX int = 100
	var birdposY int = 100
	var timeInterval int = 10
	var acceleration int = -20
	var click int
	ch := make(chan int)
	var windowX uint16 = 1000 //Wird nur fürs Fenster genutzt (Fenster nimmt x,y als uint16)
	var windowY int = 800 //Wird unter Anderem für Scale Funktion genutzt, deswegen int 
	var height int
	
	//Scale Images
	height = scale_Image(windowY)
		
	// Erzeuge Fenster
	gfx2.Fenster(windowX, uint16(windowY))
	gfx2.Fenstertitel ("Flappy Bird")
	
	//UpdateAus() um Vogel-Fehler zu vermeiden
	gfx2.UpdateAus()
	
    //Hintergrund
	gfx2.Stiftfarbe(127,255,212)
	gfx2.Cls()
	
	// Vogel reinladen
	// gfx2.LadeBildMitColorKey(birdposx, birdposy, "Frame-1.bmp",0,0,0)
	
	//Clipboard_einfuegenMitColorKey nimmt y-Wert als uint16
	gfx2.LadeBildInsClipboard("./images/Frame-1.bmp")
	gfx2.Clipboard_einfuegenMitColorKey(uint16(birdposX), uint16(birdposY), 255,0,0) 
	
	gfx2.UpdateAn()
	
	
	//Thread starten
	go Mauslesen(ch)
	
	// Main Loop
	
    for {
		select {
			case click = <-ch:
				if click == 1 {
					
				speed = 40
				birdposY -= speed
			
				}
			default:
				speed += acceleration / timeInterval 
				birdposY -= speed / timeInterval
				
		
        // Clear the screen
        gfx2.UpdateAus()
        gfx2.Cls()
			
		if birdposY < 0 {
			birdposY = 0
		
		} else if birdposY > (windowY-height) {
			birdposY = windowY - height
		}

        // Draw the sprite at its new position
        gfx2.LadeBild(uint16(birdposX), uint16(birdposY), "./images/Frame-1.bmp")

        // Update the graphics window
        gfx2.UpdateAn()
        

        // Delay for a short time (e.g., 60 frames per second)
        time.Sleep(1000 / 1000 * time.Millisecond)
		gfx2.LadeBildInsClipboard("./images/Frame-2.bmp")
		gfx2.Clipboard_einfuegenMitColorKey (uint16(birdposX), uint16(birdposY), 255,0,0)	
		  time.Sleep(1000 / 5000 * time.Millisecond)
		  
		gfx2.LadeBildInsClipboard("./images/Frame-3.bmp")
		gfx2.Clipboard_einfuegenMitColorKey (uint16(birdposX), uint16(birdposY), 255,0,0)
		  time.Sleep(1000 / 5000 * time.Millisecond)
		  
		gfx2.LadeBildInsClipboard("./images/Frame-4.bmp")
		gfx2.Clipboard_einfuegenMitColorKey (uint16(birdposX), uint16(birdposY), 255,0,0)
		  time.Sleep(1000 / 5000 * time.Millisecond)
		  
		}
		}//?? Gehört zu select-case statement
        
    }

func Mauslesen(ch chan int){
	for{	
		taste, status,_,_:=gfx2.MausLesen1()
		
			if status == 1 && taste == 1 {
				ch <- 1
				}	
	}
}

func scale_Image(windowY int) int{
	
		image_list := []string{"./Frame-1.bmp", "./Frame-2.bmp", "./Frame-3.bmp", "./Frame-4.bmp"}
		rescaled_image_list := []string{"./images/Frame-1.bmp", "./images/Frame-2.bmp", "./images/Frame-3.bmp", "./images/Frame-4.bmp"}
		
		//Verhältnis WindowY <--> BirdX = 0,25 == 25 / 100 (um float zu vermeiden)
		//Verhältnis WindowY <--> BirdY = 0,225 == 225 / 1000 (um float zu vermeiden)
		
		//Festlegen der neuen Größe des Vogels:
			
		var width int = (windowY * 25) / 100 
		var height int = (windowY * 225) / 1000
		
		//Ausgeben zum Debuggen
		fmt.Println(width)
		fmt.Println(height)
		
		//Loop zum scalen aller Bilder
		for i:=0; i < 4; i++ {
			
			//Datei Öffnen
			file, err := os.Open(image_list[i])
			
			if err != nil{
				fmt.Println("Öffnen der Datei fehlgeschlagen:", err)
				defer file.Close()
				}
			
			//BMP Datei entschlüsseln
			img, _, err := image.Decode(file)
			
			if err != nil{
				fmt.Println("Entschlüsseln der Datei fehlgeschlagen:", err)
				defer file.Close()
				}
				
			//Bildgröße verändern
			resizedImage := imaging.Resize(img, width, height, imaging.NearestNeighbor)
			
			//Neues Bild speichern
			err = imaging.Save(resizedImage, rescaled_image_list[i])
			if err != nil{
				fmt.Println("Speichern der Datei fehlgeschlagen:", err)
				}		
			}
		return height
}



