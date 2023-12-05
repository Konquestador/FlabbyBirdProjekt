package main

//Imports
import (
	"fmt"
    "gfx2"
    "time"
    //~ //"math"
    "image"
	"imaging"
	"os"
	//~ "reflect"
	//~ "saeulen"
)

//Main
func main(){
	// Variablen erstellen
	//~ var speed int = 0
	var birdposX int = 100
	//~ var birdposY int = 100
	//~ var timeInterval int = 10
	//~ var acceleration int = -10
	//~ var click int

	var windowX uint16 = 1000 //Wird nur fürs Fenster genutzt (Fenster nimmt x als uint16)
	var windowY int = 800 //Wird unter Anderem für Scale Funktion genutzt, deswegen int 
	var height int
	//~ //Säulen
	//~ var zähler int
	//~ var liste []saeulen.Saeule
	
	//Channels
	click_channel := make(chan int, 1)
	//~ coordinates_channel := make(chan int)
	
	//Scale Images
	height = scale_Image(windowY)
	
	fmt.Println(height)
	// Erzeuge Fenster
	gfx2.Fenster(windowX, uint16(windowY))
	gfx2.Fenstertitel ("Flappy Bird")
	
	//UpdateAus() um Vogel-Fehler zu vermeiden
	gfx2.UpdateAus()
	
    //Hintergrund
	gfx2.Stiftfarbe(127,255,212)
	gfx2.Cls()
	
	//Vogel in die Mitte des Fensters laden
	gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(windowY / 2) , "./images/resized/Frame-1.bmp", uint8(255), uint8(0),uint8(0))
	gfx2.UpdateAn()
	gfx2.UpdateAus()
	
	//Threads starten
	go Mauslesen(click_channel)
	go bird_movement(click_channel)
	
	gfx2.TastaturLesen1 ()
	//~ // Main Loop
	
    //~ for {
		//~ if zähler%600==0{
			//~ var s saeulen.Saeule
			//~ s = saeulen.New()
			//~ s.SetzeZufallswerte()
			//~ liste = append(liste,s)
			//~ }			
				
		//~ gfx2.UpdateAus()
		//~ gfx2.Stiftfarbe(255,255,255)
		//~ gfx2.Cls()
				
		//~ //Zeichnet die Säule auf der Höhe des Fensters
		//~ for i:=0;i<len(liste);i++{
			//~ liste[i].Draw()
			//~ }
			
		//~ gfx2.UpdateAn()
			   
		//~ var nliste []saeulen.Saeule
			   
		//~ for i:=0;i<len(liste);i++{
			//~ liste[i].Move(0)
			//~ if liste[i].GibXWert() < 10000 {
				//~ nliste = append(nliste,liste[i])
			//~ }	    
		//~ }
				
		//~ liste = nliste
		//~ zähler++
				
		//~ select{
			//~ case click = <-ch:
				//~ if click == 1 {
					
					//~ speed = 30
					//~ birdposY -= speed
					
					//~ fmt.Println("Click")

					//~ gfx2.UpdateAn()

					  
					//~ gfx2.LadeBildMitColorKey(uint16(birdposX), uint16(birdposY), "./images/Frame-3.bmp", uint8(255), uint8(0),uint8(0))
				//~ }
				
			//~ default:
				//~ speed += acceleration / timeInterval 
				//~ birdposY -= speed / timeInterval
				

				//~ // Clear the screen
				//~ gfx2.UpdateAus()
				//~ gfx2.Cls()		
				
				//~ if birdposY < 0 {
					//~ birdposY = 0
				
				//~ } else if birdposY > (windowY-height) {
					//~ birdposY = windowY - height
				//~ }
				
				//~ gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY), "./images/Frame-1.bmp", uint8(255), uint8(0),uint8(0))
		//~ }
	//~ }
				
}

func bird_movement(click_channel chan int){
	var birdposY int = 50
	var gravity float32 = 15.0
	var factor float32 = 1.0
	var durationmsall int = 0
	var speed int = 0
	
	time.Sleep(2000 * time.Millisecond)
	for{
		start := time.Now() 
		select{
			case  <- click_channel:
				fmt.Println("Sprung")
				birdposY -= 50
				speed = 0
				gfx2.Cls()
				fmt.Println("Move")
				gfx2.LadeBildMitColorKey (uint16(100), uint16(birdposY) , "./images/resized/Frame-1.bmp", uint8(255), uint8(0),uint8(0))
				gfx2.UpdateAn()
						
			default:
				time.Sleep(10 * time.Millisecond)
				//~ fmt.Println(start)
				
				duration := time.Since(start)
				fmt.Println(duration)
				//~ fmt.Println("Typ von duration", reflect.TypeOf(duration))
				durationms := int(duration.Milliseconds())
				fmt.Println("Dauer des durchlaufs:", durationms)
				//~ fmt.Println("Typ von durationms:", reflect.TypeOf(durationms))
				durationmsall += durationms
				fmt.Println("duration all ", durationmsall)
				//~ fmt.Println("Typ von durationmsall:", reflect.TypeOf(durationmsall))
				if durationmsall >= 150 {
					fmt.Println("100ms erreicht, bedingugn erfüllt")
					
					factor = float32(durationmsall) / 150.0
					fmt.Println("Faktor:", factor)
					
					durationmsall = 0
					fmt.Println(factor)
					speed += int(gravity * factor)
					birdposY += int(speed)
					fmt.Println("Neue Koordinate des Vogels:", birdposY)
				}
				fmt.Println("speed:", speed)
						gfx2.Cls()
				fmt.Println("Move")
				gfx2.LadeBildMitColorKey (uint16(100), uint16(birdposY) , "./images/resized/Frame-1.bmp", uint8(255), uint8(0),uint8(0))
				gfx2.UpdateAn()
				fmt.Println("Gravity")
			
			}
				
	
		
		//Timer duration Counter
		
		
		

		
	}
}

func Mauslesen(click_channel chan int){
	for{	
		taste, status,_,_:=gfx2.MausLesen1()
			if status == 1 && taste == 1 {
				fmt.Println("Geklickt")
				fmt.Println(status)
	
				click_channel <- 1
				}
	}
}

//~ func saeulen_move(koordinaten chan int){
	//~ for{
	//~ bird.pos
	//~ }
//~ }

func scale_Image(windowY int) int{
	
		image_list := []string{"./images/original/Frame-1.bmp", "./images/original/Frame-2.bmp", "./images/original/Frame-3.bmp", "./images/original/Frame-4.bmp"}
		rescaled_image_list := []string{"./images/resized/Frame-1.bmp", "./images/resized/Frame-2.bmp", "./images/resized/Frame-3.bmp", "./images/resized/Frame-4.bmp"}
		
		//Verhältnis WindowY <--> BirdX = 0,1875 == 875 / 10000 (um float zu vermeiden)
		//Verhältnis WindowY <--> BirdY = 0,16 == 16 / 100 (um float zu vermeiden)
		
		//Festlegen der neuen Größe des Vogels:
			
		var width int = (windowY * 1875) / 10000
		var height int = (windowY * 16) / 100
		
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


