package main


//Imports
import (
	"fmt"
    "gfx2" // Import the gfx2 library (replace "your-username" with the actual package path)
    "time"
    //"image"
	"os"
)



//Main
func main(){
	// Variablen erstellen
	var speed float64 = 0
	var birdposX float64 = 100
	var birdposY float64 = 100
	var timeInterval float64 = 0.1
	var acceleration float64 = -20
	var click int
	ch := make(chan int)
	var windowX int = 1000
	var windowY int = 800

	//Scale Images
	scale_Image(windowY)
		
	// Erzeuge Fenster
	gfx2.Fenster(windowX, windowY)
	gfx2.Fenstertitel ("Flappy Bird")
	
	//UpdateAus() um Vogel-Fehler zu vermeiden
	gfx2.UpdateAus()
	
    //Hintergrund
	gfx2.Stiftfarbe(127,255,212)
	gfx2.Cls()
	
	// Vogel reinladen
	// gfx2.LadeBildMitColorKey(birdposx, birdposy, "Frame-1.bmp",0,0,0)
	
	gfx2.LadeBildInsClipboard("Frame-1.bmp")
	gfx2.Clipboard_einfuegenMitColorKey (uint16(birdposX), uint16(birdposY), 255,0,0) 
	
	gfx2.UpdateAn()
	
	
	//Thread starten
	go Mauslesen(ch)
	
	// Main Loop
	
    for {
		select {
			case click = <-ch:
				if click == 1 {
					
				speed = 40
				birdposY -= float64(speed)
			
				}
			default:
				speed += timeInterval * acceleration
				birdposY -= float64(speed * timeInterval)
				
		
        // Clear the screen
        gfx2.UpdateAus()
        gfx2.Cls()
			
				
        if birdposY < 0 {
			birdposY = 0
			
		}else if birdposY > 630 {
			birdposY = 630
		}
		
        // Draw the sprite at its new position
        gfx2.LadeBild(uint16(birdposX), uint16(birdposY), "Frame-1.bmp")

        // Update the graphics window
        gfx2.UpdateAn()
        

        // Delay for a short time (e.g., 60 frames per second)
        time.Sleep(1000 / 1000 * time.Millisecond)
		gfx2.LadeBildInsClipboard("Frame-2.bmp")
		gfx2.Clipboard_einfuegenMitColorKey (uint16(birdposX), uint16(birdposY), 255,0,0)	
		  time.Sleep(1000 / 5000 * time.Millisecond)
		  
		gfx2.LadeBildInsClipboard("Frame-3.bmp")
		gfx2.Clipboard_einfuegenMitColorKey (uint16(birdposX), uint16(birdposY), 255,0,0)
		  time.Sleep(1000 / 5000 * time.Millisecond)
		  
		gfx2.LadeBildInsClipboard("Frame-4.bmp")
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

func scale_Image(windowY int){
		image_list := []string{"./Frame-1.bmp", "./Frame-2.bmp", "./Frame-3.bmp", "./Frame-4.bmp"}
		
		for _, value := range image_list {
			
			//Open File
			file, err := os.Open(value)
			
			if err != nil{
				fmt.Println("Öffnen der Datei fehlgeschlagen!")
				defer file.Close()
				}
			
			//Decode BMP File
			img, _, err := image.Decode(file)
			
			if err != nil{
				fmt.Println("Entschlüsseln der Datei fehlgeschlagen!")
				defer file.Close()
				}
				
			//Resize Image
			resizedImage := imaging.Resize()
			
			}
}



