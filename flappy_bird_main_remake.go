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
	
	//Fenster & Vogel:
	var windowY int = 800 //Wird unter Anderem für Scale Funktion genutzt, deswegen int 
	var height int
	var birdposX int = 100
	var birdposY int = windowY / 10 //Damit der Vogel anfangs im oberen Zehntel spwant
	var endbirdposY int
	
	//Gravity
	var gravity float32 = 2.0
	var factor float32 = 1.0
	var speed int = 0	
	
	
	//Counter
	var counterstr string
	var counter int = 0
	
	//Zeitberechnungen:
	start := time.Now() //Einmalige Startzeit, um die duration-Berechnung am Anfang der for-Schleife zu ermöglichen
	
	//Da jeder Prozess (Gravity, die Bewegung der Säulen und das Erstellen neuer Säulen) ab einer eigenen Anzahl an Microsekunden erst
	//ausgeführt wird und danach die Variable wieder gleich 0 gesetzt werden muss, erhält jeder Prozess seine eigen Variable für die
	//Summe der Durchlaufsdauern. Obwohl die Variablennamen relativ eindeutig sind, ist trotzdem beschrieben wofür sie sind.
	
	var durationmsall int = 0 //(Nur verwendet für Gravity)
	var pillarspwanms int = 0 //(Nur verwendet fürs Erstellen neuer Säulen)
	var pillarmovems int = 0 //(Nur verwendet fürs Bewegen der Säulen)
	
	//Boolesche Werte
	var update bool = false
	var jump bool = false
	var firstRound bool = true
	var end bool = false
	var touchAbove bool = false
	var touchBelow bool = false
	var touchoutside bool = false
	
	
	//Säulen
	var liste []saeulen.Saeule
	var leereliste []saeulen.Saeule
	
	//Channel 
	click_channel := make(chan int, 1)
	
	//Scale Images
	height, width := scale_Image(windowY)
	
	// Erzeuge Fenster
	gfx2.Fenster(uint16(1000), uint16(800))
	gfx2.Fenstertitel ("Flappy Bird")
	
	//UpdateAus() um Anzeigefehler zu vermeiden
	gfx2.UpdateAus()
	
	//Startbildschirm:
	
	//Hintergrund
	clear()
	
	//Vogel und Buttons reinladen
	gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-1.bmp", uint8(135), uint8(206),uint8(250))
	gfx2.LadeBildMitColorKey(uint16(251), uint16(600), "./playbutton.bmp", uint8(255), uint8(0),uint8(0))
	gfx2.LadeBildMitColorKey(uint16(501), uint16(600), "./quitbutton.bmp", uint8(255), uint8(0),uint8(0))
	
	//Buttonbeschriftung
	gfx2.SetzeFont("./schriftart.ttf", 70)//Schritart und Größe festlegen
	gfx2.Stiftfarbe(255,230,5)
	gfx2.SchreibeFont(uint16(540), uint16(620), "QUIT")
	gfx2.UpdateAn()
	
	//Abfrage ob einer der Buttons gedrückt wurde
	for{
		mtaste, mstatus, mx, my:=gfx2.MausLesen1()
		if mstatus == 1 || mstatus == -1{
			if mtaste == 1 && mx >= 501 && mx <= 751 && my >= 600 && my <= 775{//Quit-Button
				gfx2.FensterAus() 
				break  
			}else if mtaste == 1 && mx >= 251 && mx <= 376 && my >= 600 && my <= 775{//Play-Button --> Es wird davon ausgegangen, dass man nicht ausversehen neben den Button klickt, ohne weiterspielen zu wollen, weswegen auf eine dreieckige Hitbox verzichtet wurde
				break
			}
		}
	}
	
	//Schriftart für den Counter festlegen
	gfx2.SetzeFont("./schriftart.ttf", 150)
	
    //Hintergrund
	clear()
	
	//Vogel in die Mitte des Fensters laden
	gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-12.bmp", uint8(135), uint8(206),uint8(250))
	
	gfx2.UpdateAn()
	
	//Threads starten
	go mauslesen(click_channel)
	go tastaturlesen(click_channel)
	
	//Starten des Spielloops
	for{
		duration := time.Since(start) //Berechnung der Dauer der Berechnung bzw. einem Loop-Durchlauf
		durationms := int(duration.Microseconds()) //Umwandlung der Zeit in Microsekunden (Da der Computer so schnell rechnet, kann man nur mit diesen kleinen Werten arbeiten)
		
		//Summe der einzelnen Durchlaufdauern --> siehe Zeile 39
		durationmsall += durationms	
		pillarspwanms += durationms
		pillarmovems += durationms
		start = time.Now() //Festlegen der Startzeit des Loops um ernuet um die Geschwindigkeit zu messen
		
		//Falsche Zeitwerte für die erste Runde eliminieren
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
		
		//Gravity
		select{
			case  <- click_channel:	//Wenn etwas im Channel liegt (Wert egal, da ein Wert im Channel ein Sprung bedeutet)	
				speed = 0 
				durationmsall = -75000 //Damit bleibt der Vogel 3fps stehen, bevor er wieder fällt --> natürlichere Bewegung & Flügelschlag sichtbar
				birdposY -= 70 //Y-Koordinate des Vogels niedriger setzen (Fenster hat oben 0)
				
				//Vogel soll die obere Grenze des Fensters nicht überschreiten
				//~ if birdposY < 0 {
					//~ birdposY = 0
				//~ }
				
				//Da sich etwas verändert hat, soll es geupdatet werden
				update = true
				jump = true

			default:				
				if durationmsall >= 40000 {//Nur alle 40 millisekunden ausführen 
					factor = float32(durationmsall) / 25000.0 //Falls der Loopdurchlauf länger gedauert hat, wird die Stärke des Fallens entsprechend angepasst				
					durationmsall = 0

					speed += int(gravity * factor) //Berechneter Faktor wird angewendet
					birdposY += int(speed) //Y-Koordinate des Vogels 
					
					//verhindern, dass der Vogel unten aus dem Fenster fällt
					if birdposY > (windowY-height) {
						birdposY = windowY - height
						speed = 0
					}
					//Da sich etwas verändert hat, soll es geupdatet werden
					update = true
					jump = false 
				}
		}
		
		if pillarspwanms > 1250000{ //Alle 1,25 Sekunden soll eine weitere Säule erschaffen werden
			//Neue Säule erschaffen und der Liste hinzufügen
			var s saeulen.Saeule = saeulen.New()
			s.SetzeZufallswerte()
			liste = append(liste,s)
			pillarspwanms = 0
		}
		
		if pillarmovems >= 10000 && len(liste) > 0{ //Alle 10 Millisekunden sollen die Säulen bewegt werden
			var nliste []saeulen.Saeule
			
			//Jedes Jede Säule aus der Liste bewegen
			for i:=0;i<len(liste);i++{
				liste[i].Move()
				
				if liste[i].GibXWert() < 100 && liste[i].GibPassed() == false{ //Abfrage ob die Säule den Vogel schon passiert hat
					//Akutalisiere den Counter
					counter ++
					counterstr = strconv.Itoa(counter)
					
					//Wert der Säule aktualisieren um das mehrmalige Zählen zu verhindern
					liste[i].Passed()
					}
				//Liste löschen, wenn diese aus dem Fenster verschwunden ist
				if liste[i].GibXWert() < 10000 {
					nliste = append(nliste,liste[i])
				}
				
			}
			liste = nliste
			pillarmovems = 0
			//Da sich etwas verändert hat, soll es geupdatet werden
			update = true
		}
		
		for i:=0;i<len(liste);i++{ //Alle Säulen werden auf kollision mit dem Vogel überprüft. Es kann nur eine Säule gleichzeitig geben
			touchAbove, touchBelow, touchoutside = collision(birdposX, birdposY, width, height,touchAbove, touchBelow, touchoutside, liste[i])
			if touchAbove || touchBelow || touchoutside{
				//Anpassen der Y-Werte des Vogels für den Endscreen, je nach Berührung
				if touchAbove{
					endbirdposY = int(liste[i].GibHoehe())
					end = true
					break //Schleifendurchlauf kann unterbrochen werden, da das Spiel bei einer Kollision sofort beendet wird und es keine zweite Säule geben kann und darf, die berührt wird
				}else if touchBelow{
					endbirdposY = int(liste[i].GibHoehe() + liste[i].GibLoch()) - height
					end = true
					break //Schleifendurchlauf kann unterbrochen werden, da das Spiel bei einer Kollision sofort beendet wird und es keine zweite Säule geben kann und darf, die berührt wird
				}else if touchoutside{
					endbirdposY = birdposY
					end = true 
					break //Schleifendurchlauf kann unterbrochen werden, da das Spiel bei einer Kollision sofort beendet wird und es keine zweite Säule geben kann und darf, die berührt wird
				
				}
				break //Schleifendurchlauf kann unterbrochen werden, da das Spiel bei einer Kollision sofort beendet wird und es keine zweite Säule geben kann und darf, die berührt wird
			}
		}
		
		if update && end == false{ //Wenn das Fenster nur geupdatet werden soll, das Spiel aber nicht vorbei ist
			//Hintergrund
			clear()
			
			//Vogel neu reinladen --> Es wird zwischen den beiden Bildern, je nach Sprung oder nicht unterschieden
			if jump{
				gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-2.bmp", uint8(135), uint8(206),uint8(250))
			}else{
				gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(birdposY) , "./images/resized/Frame-1.bmp", uint8(135), uint8(206),uint8(250))			
			}
			//Säulen zeichnen
			gfx2.Stiftfarbe(0,0,0)
			for i:=0;i<len(liste);i++{
			   liste[i].Draw(birdposX, width)
			}
			
			//Counter zeichnen
			gfx2.Stiftfarbe(255,255,85)
			gfx2.SchreibeFont(uint16(475), uint16(25), counterstr)
			
			//Updaten 
			gfx2.UpdateAn()
			gfx2.UpdateAus()
			
			update = false
		}
		
		if end{ //Endscreen
			//Hintergrund
			clear()
			
			//Säulen ein letztes mal zeichnen
			gfx2.Stiftfarbe(0,0,0)
			for i:=0;i<len(liste);i++{
			   liste[i].Draw(birdposX, width)
			}
			
			//Vogel mit den angepasster Y-Koordinate reinladen
			gfx2.LadeBildMitColorKey (uint16(birdposX), uint16(endbirdposY) , "./images/resized/Frame-1.bmp", uint8(135), uint8(206),uint8(250))			
			
			//Game Over Schrift reinschreiben
			gfx2.Stiftfarbe(255,0,0)
			gfx2.SchreibeFont(uint16(70), uint16(76), "GAME OVER")
			
			//Counter reinschreiben
			gfx2.Stiftfarbe(255,255,0)
			gfx2.SchreibeFont(uint16(400), uint16(326), counterstr)
			
			//Buttons reinladen
			gfx2.LadeBildMitColorKey(uint16(251), uint16(600), "./playbutton.bmp", uint8(255), uint8(0),uint8(0))
			gfx2.LadeBildMitColorKey(uint16(501), uint16(600), "./quitbutton.bmp", uint8(255), uint8(0),uint8(0))
			gfx2.SetzeFont("./schriftart.ttf", 70)
			gfx2.Stiftfarbe(255,230,5)
			gfx2.SchreibeFont(uint16(540), uint16(620), "QUIT")
			gfx2.SetzeFont("./schriftart.ttf", 150)
			gfx2.UpdateAn()
			
			//Abfrage und Verarbeitung der Mauseingabe
			for{
				mtaste, mstatus, mx, my:=gfx2.MausLesen1()
				if mstatus == 1 || mstatus == -1{
					if mtaste == 1 && mx >= 501 && mx <= 751 && my >= 600 && my <= 775{//Quit-Button --> Fenster schließen
						gfx2.FensterAus()
						break  
					}else if mtaste == 1 && mx >= 251 && mx <= 376 && my >= 600 && my <= 775{//Play-Button --> Variablen zurücksetzen und den Loop von Vorne beginnen 
						firstRound = true
						end = false
						touchAbove = false
						touchBelow = false
						touchoutside = false
						liste = leereliste
						duration = 0
						durationms = 0
						durationmsall = 0
						pillarspwanms = 0
						pillarmovems = 0
						birdposX = 100
						birdposY = windowY/10
						counterstr = ""
						counter = 0
						break
					}
				}
			}
		}
	}
}
	
func collision(birdposX int, birdposY int, width int, height int, touchAbove bool, touchBelow bool, touchoutside bool, s saeulen.Saeule) (bool, bool, bool){	

	if birdposX + width >= int(s.GibXWert()) && birdposX-5 <= int(s.GibXWert() + s.GibBreite()){ //x-Koordinaten werden vgerglichen --> Wahr, wenn der  Vogel und die Säule sich irgendwelche x-Koordinaten teilen
		if birdposY < int(s.GibHoehe()) || birdposY + height > int(s.GibHoehe() + s.GibLoch()){ //y-Koordinaten werden verglichen --> Wahr, wenn sich der Vogel nicht im Loch befindet --> Vogel muss die Säule berühren
			if s.GibbirdInHole(){//alter Wert wird überprüft --> Wahr, wenn der Vogel sich vor den neu berechneten Werten im Loch aufhielt --> Vogel muss aus dem Loch aus in die Säule gesprungen/gefallen sein)
				if birdposY < int(s.GibHoehe()){//Wahr, wenn der Vogel die Säule oben berührt (y-Koordinaten des Vogels müssen kleiner sein als die Säulenhöhe)
					touchAbove = true
				}else if birdposY + height > int(s.GibHoehe() + s.GibLoch()){//Wahr, wenn der Vogel die Säule unten berührt (y-Koordinaten des Vogels müssen größer sein als die Säulenhöhe + Lochgröße)
					touchBelow = true
				}
			}else{//Wird ausgeführt, wenn der alte Wert falsch war --> Vogel ist nicht aus einem Loch in die Säule gespringen/gefallen
				touchoutside = true
				s.Touch()//Wird verwendet um nachher die Säule zu identifizieren, die berührt wurde, um ein eine Überlappung des Vogels mit der Säule zu vermeiden
			}
		}
		
	}
	return touchAbove, touchBelow, touchoutside //Gibt immer die Werte zurück
}

//Mauseingaben abfragen und verarbeiten
func mauslesen(click_channel chan int){
	for{	
		mtaste, mstatus,_,_:=gfx2.MausLesen1()
		if mstatus == 1 && mtaste == 1{
			click_channel <- 1
		}
	}
}

//Tastatureingaben abfragen und verarbeiten
func tastaturlesen(click_channel chan int){
	for{
		ttaste, tstatus,_:= gfx2.TastaturLesen1()
		if tstatus == 1 && ttaste == 32 {
				click_channel <- 1
		}
	}
}

//Fenster aufräumen
func clear(){
	gfx2.Stiftfarbe(135,206,250)
	gfx2.Cls()
}

//Bilder an die Fenstergröße anpassen
func scale_Image(windowY int) (int,int){
		
		//Pfade der Bilder
		image_list := []string{"./images/original/Frame-12.bmp", "./images/original/Frame-22.bmp"}
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
