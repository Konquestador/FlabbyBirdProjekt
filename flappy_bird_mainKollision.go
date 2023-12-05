package main

//Imports
import (
	"fmt"
    "gfx2"
    //~ "time"
    "math"
    "image"
	"imaging"
	"os"
	"saeulen"
)

//Main
func main(){
	// Variablen erstellen
	var speed int = 0
	var birdposX int = 100
	var birdposY int = 100
	var birdCenterX int
	var birdCenterY int
	var birdRadius int = 50
	var timeInterval int = 10
	var acceleration int = -18
	var click int
	ch := make(chan int)
	var windowX uint16 = 1000 //Wird nur fürs Fenster genutzt (Fenster nimmt x als uint16)
	var windowY int = 800 //Wird unter Anderem für Scale Funktion genutzt, deswegen int 
	var height int
	var width int
	//Säulen
	var zähler int
	var liste []saeulen.Saeule
	//Scale Images
	height, width = scale_Image(windowY)
	fmt.Println(width)
	// Erzeuge Fenster
	gfx2.Fenster(windowX, uint16(windowY))
	gfx2.Fenstertitel ("Flappy Bird")
	
	//UpdateAus() um Vogel-Fehler zu vermeiden
	gfx2.UpdateAus()
	
    //Hintergrund
	gfx2.Stiftfarbe(127,255,212)
	gfx2.Cls()
	
	//Vogel in die Mitte des Fensters laden
	gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(windowY / 2) , "./images/Frame-1.bmp", uint8(255), uint8(0),uint8(0))
	
	
	gfx2.UpdateAn()
	
	
	//Thread starten
	go Mauslesen(ch)
	
	// Main Loop
	
    for {
		if zähler%600==0{
			var s saeulen.Saeule
			s = saeulen.New()
			s.SetzeZufallswerte()
			liste = append(liste,s)
			}			
				
		gfx2.UpdateAus()
		gfx2.Stiftfarbe(255,255,255)
		gfx2.Cls()
				
		//Zeichnet die Säule auf der Höhe des Fensters
		for i:=0;i<len(liste);i++{
			liste[i].Draw()
			}
			
		gfx2.UpdateAn()
			   
		var nliste []saeulen.Saeule
			   
		for i:=0;i<len(liste);i++{
			liste[i].Move(0)
			if liste[i].GibXWert() < 10000 {
				nliste = append(nliste,liste[i])
			}	    
		}
				
		liste = nliste
		zähler++
				
		select{
			case click = <-ch:
				if click == 1 {
					
					speed = 40
					birdposY -= speed
					
					fmt.Println("Click")

					gfx2.UpdateAn()

					  
					gfx2.LadeBildMitColorKey(uint16(birdposX), uint16(birdposY), "./images/resized/Frame-3.bmp", uint8(255), uint8(0),uint8(0))
				}
				
			default:
				speed += acceleration / timeInterval //-18 : 10 
				birdposY -= speed / timeInterval^^	 //speed : 10
				
				if Kollision(birdCenterX, birdCenterY, birdRadius, birdposY, birdposX) {
					fmt.Println("Kollision")
					gfx2.LadeBild(0,0, "./failscreen.bmp")
					
					}
				// Clear the screen
				//~ gfx2.UpdateAus()
				//~ gfx2.Cls()		
				
				if birdposY < 0 {
					birdposY = 0
				
				} else if birdposY > (windowY-height) {
					birdposY = windowY - height
				}
				
				gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY), "./images/resized/Frame-1.bmp", uint8(255), uint8(0),uint8(0))
		}
	}
				
}

func Mauslesen(ch chan int){
	for{	
		taste, status,_,_:=gfx2.MausLesen1()
		
			if status == 1 && taste == 1 {
				ch <- 1
				}	
	}
}

func scale_Image(windowY int) (int, int) {
	
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
		return height, width
}

func Kollision(birdCenterX int, birdCenterY int, birdRadius int, birdposY int, birdposX int) bool {
	birdCenterX = birdposX + birdRadius
	birdCenterY = birdposY + birdRadius
	//~ gfx2. Stiftfarbe(255,0,0)
	//~ gfx2.Vollkreis(uint16(birdCenterX), uint16(birdCenterY), uint16(birdRadius))
	// Check the points exactly on the radius
	for angle := 0; angle <= 360; angle += 15 { // Adjust the step based on your needs
		x := int(float64(birdCenterX) + float64(birdRadius)*math.Cos(float64(angle)*math.Pi/180))
		y := int(float64(birdCenterY) + float64(birdRadius)*math.Sin(float64(angle)*math.Pi/180))

		if x >= 0 && x < int(gfx2.Grafikspalten()) && y >= 0 && y < int(gfx2.Grafikzeilen()) {
			// Überprüfe die Farbe des Punktes
			colorAtPointR, colorAtPointG, colorAtPointB := gfx2.GibPunktfarbe(uint16(x), uint16(y))
			if colorAtPointR == 0 && colorAtPointG == 0 && colorAtPointB == 0 {
				return true // Kollision
			}
		}
	}

	return false
}


