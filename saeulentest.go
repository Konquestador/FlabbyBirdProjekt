package main

import (
	"saeulen"
	"gfx2"
	//"fmt"
	//"zufallszahlen"
	//"time"
)

func main () {

	//Öffnet das Grafikfenster
    gfx2.Fenster(800, 600)

	var zähler int
	
	var liste []saeulen.Saeule

    for i:=0;i<=10000;i++{
		if zähler%500==0{
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
		zähler ++
	}
	   }

