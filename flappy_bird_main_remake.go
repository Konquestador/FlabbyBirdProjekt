package main

//Imports
import (
	"fmt"
    "gfx2"
    "time"
    "image"
	"imaging"
	"os"
	"saeulen"
	"strconv"
)

//Main
func main(){
	// Variablen erstellen
	
	//Festgelegte Größen:
	var windowX int = 1000 //Wird nur fürs Fenster genutzt (Fenster nimmt x als uint16)
	var windowY int = 800 //Wird unter Anderem für Scale Funktion genutzt, deswegen int 
	var height int
	var birdposX int = 100
	var birdposY int = windowY / 10 //Damit der Vogel anfangs im oberen Zehntel spwant
	
	//Zeitberechnungen:
	start := time.Now() //Einmalige Startzeit, um die duration-Berechnung am Anfang der for-Schleife zu ermöglichen
	
	//Da jeder Prozess (Gravity, die Bewegung der Säulen und das Erstellen neuer Säulen) ab einer eigenen Anzahl an Microsekunden erst
	//ausgeführt wird und danach die Variable wieder gleich 0 gesetzt werden muss, erhält jeder Prozess seine eigen Variable für die
	//Summe der Durchlaufsdauern. Obwohl die Variablennamen relativ eindeutig sind, ist trotzdem beschrieben wofür sie sind.
	
	var durationmsall int = 0 //(Nur verwendet für Gravity)
	var pillarspwanms int = 0 //(Nur verwendet fürs Erstellen neuer Säulen)
	var pillarmovems int = 0 //(Nur verwendet fürs Bewegen der Säulen)
	
	//Gravity
	var gravity float32 = 2.0
	var factor float32 = 1.0
	var speed int = 0
	
	var counter int = 0
	
	//Boolesche Werte
	var update bool = false
	var jump bool = false
	var firstRound bool = true
	var end bool = false
	var touchAbove bool = false
	var touchBelow bool = false
	var touchoutside bool = false
	//~ var birdInHole bool = false
	
	
	
	var counterstr string
	var endbirdposY int
	
	var liste []saeulen.Saeule
	
	//Channel 
	click_channel := make(chan int, 1)
	
	//Scale Images
	height, width := scale_Image(windowY)
	
	// Erzeuge Fenster
	gfx2.Fenster(uint16(windowX), uint16(windowY))
	gfx2.Fenstertitel ("Flappy Bird")
	
	//Schriftart festlegen
	gfx2.SetzeFont("./schriftart.ttf", 100)
	
	//UpdateAus() um Vogel-Fehler zu vermeiden
	gfx2.UpdateAus()
	
    //Hintergrund
	gfx2.Stiftfarbe(127,255,212)
	gfx2.Cls()
	
	//Vogel in die Mitte des Fensters laden
	gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-1.bmp", uint8(255), uint8(255),uint8(0))
	
	gfx2.UpdateAn()
	
	//Threads starten
	go mauslesen(click_channel)

	for{
		duration := time.Since(start) //Berechnung der Dauer der Berechnung bzw. einem Loop-Durchlauf
		durationms := int(duration.Microseconds()) //Umwandlung der Zeit in Microsekunden (Da der Computer so schnell rechnet, kann man nur mit diesen kleinen Werten arbeiten)
		durationmsall += durationms	
		pillarspwanms += durationms
		pillarmovems += durationms
		start = time.Now()
		
		if firstRound{
			durationmsall = 0
			pillarspwanms = 0
			pillarmovems = 0
			firstRound = false
		}	
		
		//Berechnung ob der Vogel sich im Loch befindet. Wichtig um nachher bei der Kollision herauszuinden, ob der Vogel von außen oder im Loch gegen die Säule traf
		if len(liste) != 0{ //Wenn es überhaupt eine Säule gibt:
			for i:=0;i<len(liste);i++{ //Berechne es für jede Säule, die in der Liste ist und sich damit im Fenster befindet
					if birdposY >= int(liste[i].GibHoehe()) && birdposY + height <= int(liste[i].GibHoehe() + liste[i].GibLoch()){//Wenn der Vogel im Loch ist (Oberer Rand des Vogels muss größer sein, als die Höhe der Säule, aber kleiner als Die Höhe der Säule + die Lochgröße)
						liste[i].InHole()//Eigenschaft der Säule wird neu gesetzt: Vogel befindet sich im Loch
					}else{
						liste[i].NotInHole()//Eigenschaft der Säule wird neu gesetzt: Vogel befindet sich nicht im Loch
					}
			}
		}
		
		select{
			case  <- click_channel:	
						
				speed = 0
				durationmsall = -75000
				birdposY -= 70
				
				if birdposY < 0 {
					birdposY = 0
				}
				update = true
				jump = true

			default:				
				if durationmsall >= 40000 {
					jump = false 

					factor = float32(durationmsall) / 25000.0					
					durationmsall = 0

					speed += int(gravity * factor)
					birdposY += int(speed)
					
					if birdposY > (windowY-height) {
						birdposY = windowY - height
						speed = 0
					}

					update = true
				}
		}
		if pillarspwanms > 1250000{
			gfx2.Stiftfarbe(34, 139, 34)
			var s saeulen.Saeule = saeulen.New()
			//~ s = saeulen.New()
			s.SetzeZufallswerte()
			liste = append(liste,s)
			pillarspwanms = 0
		}
		
		if pillarmovems >= 10000 && len(liste) > 0{

			var nliste []saeulen.Saeule
			
			for i:=0;i<len(liste);i++{
				liste[i].Move()
				
				if liste[i].GibXWert() < 100 && liste[i].GibPassed() == false{
					counter ++
					counterstr = strconv.Itoa(counter)
					liste[i].Passed()
					}
					
				if liste[i].GibXWert() < 10000 {
					nliste = append(nliste,liste[i])
				}
				
			}
			liste = nliste
			pillarmovems = 0
			update = true
		}
		
		for i:=0;i<len(liste);i++{ //Alle Säulen werden auf kollision mit dem Vogel überprüft. Es kann nur eine Säule gleichzeitig geben
			touchAbove, touchBelow, touchoutside = collision(birdposX, birdposY, width, height,touchAbove, touchBelow, touchoutside, liste[i])
			if touchAbove || touchBelow || touchoutside{
				if touchAbove{
					endbirdposY = int(liste[i].GibHoehe())
					end = true
					break
				}else if touchBelow{
					endbirdposY = int(liste[i].GibHoehe() + liste[i].GibLoch()) - height
					end = true
					break
				}else if touchoutside{
					endbirdposY = birdposY
					end = true
					break
				
				}
				break //Schleifendurchlauf kann unterbrochen werden, da das Spiel bei einer Kollision sofort beendet wird und es keine zweite Säule geben kann und darf, die berührt wird
			}
			fmt.Println("Keine collision")
		}
		
		if update && end == false{ 
			
			gfx2.Stiftfarbe(135,206,250)
			gfx2.Cls()
			
			if jump{
				gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-2.bmp", uint8(135), uint8(206),uint8(250))
			}else{
				gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-1.bmp", uint8(135), uint8(206),uint8(250))			
			}
			gfx2.Stiftfarbe(0,0,0)

			for i:=0;i<len(liste);i++{
			   liste[i].Draw(birdposX, width)
			}

			gfx2.Stiftfarbe(255,255,85)
			gfx2.SchreibeFont(uint16(450), uint16(25), counterstr)
			
			gfx2.UpdateAn()
			gfx2.UpdateAus()
			update = false
			
			}
		if end{
			fmt.Println("Ende erreicht")
			fmt.Println("touchAbove:", touchAbove)
			fmt.Println("touchBelow:", touchBelow)
			fmt.Println("touchoutside:", touchoutside)
			
			gfx2.SetzeFont("./schriftart.ttf", 150)
			
			gfx2.Stiftfarbe(135,206,250)
			gfx2.Cls()
			
			gfx2.Stiftfarbe(0,0,0)
			for i:=0;i<len(liste);i++{
			   liste[i].Draw(birdposX, width)
			}
			
			
			gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(endbirdposY) , "./images/resized/Frame-1.bmp", uint8(135), uint8(206),uint8(250))			
			
			gfx2.Stiftfarbe(255,0,0)
			gfx2.SchreibeFont(uint16(70), uint16(76), "GAME OVER")
			
			gfx2.Stiftfarbe(255,255,0)
			gfx2.SchreibeFont(uint16(400), uint16(326), counterstr)
			
			gfx2.LadeBildMitColorKey(uint16(251), uint16(450), "./playbutton.bmp", uint8(255), uint8(0),uint8(0))
			gfx2.LadeBildMitColorKey(uint16(501), uint16(450), "./quitbutton.bmp", uint8(255), uint8(0),uint8(0))
			gfx2.UpdateAn()
					
			for{
				taste, status,_,_:=gfx2.MausLesen1()
				if status == 1 && taste == 1 {
					break
				}
			}
			
			
			
			}//If end{}
		
		
	}//For Loop Ende	
}//Func Main Ende
	
func collision(birdposX int, birdposY int, width int, height int, touchAbove bool, touchBelow bool, touchoutside bool, s saeulen.Saeule) (bool, bool, bool){	

	if birdposX + width >= int(s.GibXWert()) && birdposX <= int(s.GibXWert() + s.GibBreite()){ //x-Koordinaten werden vgerglichen --> Wahr, wenn der  Vogel und die Säule sich irgendwelche x-Koordinaten teilen
		if birdposY < int(s.GibHoehe()) || birdposY + height > int(s.GibHoehe() + s.GibLoch()){ //y-Koordinaten werden verglichen --> Wahr, wenn sich der Vogel nicht im Loch befindet --> Vogel muss die Säule berühren
			if s.GibbirdInHole(){//alter Wert wird überprüft --> Wahr, wenn der Vogel sich vor den neu berechneten Werten im Loch aufhielt --> Vogel muss aus dem Loch aus in die Säule gesprungen/gefallen sein)
				if birdposY < int(s.GibHoehe()){//Wahr, wenn der Vogel die Säule oben berührt (y-Koordinaten des Vogels müssen kleiner sein als die Säulenhöhe)
					fmt.Println("Kollision oben")
					touchAbove = true
				}else if birdposY + height > int(s.GibHoehe() + s.GibLoch()){//Wahr, wenn der Vogel die Säule unten berührt
					fmt.Println("Kollision oben")
					touchBelow = true
				}
			}else{//Wird ausgeführt, wenn der alte Wert falsch war --> Vogel ist nicht aus einem Loch in die Säule gespringen/gefallen
				touchoutside = true
				s.Touch()
			}
		}
	}
	return touchAbove, touchBelow, touchoutside //Gibt immer die Werte zurück
}

func mauslesen(click_channel chan int){
	for{	
		taste, status,_,_:=gfx2.MausLesen1()
			if status == 1 && taste == 1 {
				click_channel <- 1
				}
	}
}

func scale_Image(windowY int) (int,int){
	
		image_list := []string{"./images/original/Frame-1.bmp", "./images/original/Frame-2.bmp"}
		rescaled_image_list := []string{"./images/resized/Frame-1.bmp", "./images/resized/Frame-2.bmp"}
		
		//Verhältnis WindowY <--> BirdX = 0,1375 == 1375 / 10000 (um float zu vermeiden)
		//Verhältnis WindowY <--> BirdY = 0,1275 == 1275 / 1000 (um float zu vermeiden)
		
		//Festlegen der neuen Größe des Vogels:
		var width int = (windowY * 1375) / 10000
		var height int = (windowY * 1275) / 10000
		
		//Loop zum scalen aller Bilder
		for i:=0; i < 2; i++ {
			
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
