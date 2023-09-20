package main

import (
    "fmt"
    "gfx2" // Import the gfx2 library (replace "your-username" with the actual package path)
 //   "time"
//	"math/rand"
//  "zufallszahlen"
)

func main() {
	var timeInterval float64
	timeInterval = float64(0.1)
	var acceleration float64 = -8
	var speed float64 = 0
	var TasteGedrueckt int = 0
	ch := make(chan int)	
//	var x uint16 = 1000
//	var screenWidth uint16  = 1000
//	var   screenHeight uint16 = 800
//	var   pillarWidth uint16  = 30
//	var  holeWidth uint16 = 400
//	var  pillarSpeed uint16 = 5
//	var holeY uint16 = 40	
    // Open the graphics window with specified dimensions
  //  gfx2.Stiftfarbe(0,0,0)
    gfx2.Fenster(1000, 800)
    gfx2.Stiftfarbe(127,255,212)
    gfx2.Cls()
//	fmt.Println(gfx2.Grafikspalten())
    // Load an image file for the sprite
    gfx2.LadeBild(uint16(100), uint16(100), "Frame-1.bmp") // Convert int to uint16
    
    

    // Set the sprite's initial position
    var spriteX, spriteY uint16 = 100, 100 // Use uint16 for sprite position

    // Set the sprite's transparency (0-255)
    gfx2.Transparenz(0)

    // Set the sprite's initial velocity
    var _, velocityY float64 = 0, -30// Use int for velocity
	var positionY float64
	positionY = float64(spriteY)
    // Main game loop
    go Mauslesen(ch)
    
    for {
		select {
			case TasteGedrueckt := <-ch:
				if TasteGedrueckt == 1 {
			
				speed = 25
				positionY -= speed
			
				}
			default:
//				gfx2.Cls()
		//		rand.Seed(time.Now().UnixNano())
//		holeY =	zufallszahlen.Randomisieren(100-700)	
				        // Clear the screen
		//	gfx2.Cls()
//			  gfx2.Stiftfarbe(0,0,0)
			// Generate a random position for the hole in the pillar
		//	holeY = rand.Intn(800 - 100)

			// Draw the left part of the pillar
	//		gfx2.Vollrechteck(x, 0, 150, holeY)

			// Draw the right part of the pillar
//			gfx2.Vollrechteck(x, 440, 150, 360)

			// Move the pillar to the left
//			x = x-8

			// Check if the pillar is out of the screen
	//		if x <= 0 {
 //           x = 1000 // Reset the pillar's position when it goes off-screen
//			}
//			gfx2.Stiftfarbe(127,255,212)
			// Update the graphics window
			

			// Delay for a short time (e.g., 60 frames per second)
			
				
				
				
				
				speed += timeInterval * acceleration
		positionY -= speed * timeInterval
        // Clear the screen
		gfx2.UpdateAus()
		gfx2.Cls()
			
				

        // Check for collision with screen edges
      // if spriteX < 0 || spriteX > 800 {
       //     velocityX = -velocityX // Reverse X velocity on collision
        
       // if positionY < 0 || positionY > 630 {
       //     velocityY = -velocityY // Reverse Y velocity on collision
       // }
        if positionY < 0 {
			spriteY = 0
			positionY = 0
		}else if positionY > 630 {
			spriteY = 630
			positionY = 630
		} else {
			spriteY = uint16(positionY)
		}
		fmt.Println(velocityY,spriteY)
        // Draw the sprite at its new position
        gfx2.LadeBild(spriteX, spriteY, "Frame-1.bmp")

        // Update the graphics window
        gfx2.UpdateAn()
        

        // Delay for a short time (e.g., 60 frames per second)
 //       time.Sleep(1000 / 1000 * time.Millisecond)
		gfx2.LadeBild(uint16(100), uint16(positionY), "Frame-2.bmp")	
//		  time.Sleep(1000 / 5000 * time.Millisecond)
		gfx2.LadeBild(uint16(100), uint16(positionY), "Frame-3.bmp")
//		  time.Sleep(1000 / 5000 * time.Millisecond)
		gfx2.LadeBild(uint16(100), uint16(positionY), "Frame-4.bmp")
//		  time.Sleep(1000 / 5000 * time.Millisecond)
	//	//TasteGedrueckt = <-ch
		fmt.Println(TasteGedrueckt)
			
	 }	}}
        
        
    

func Mauslesen(ch chan int){
	for{	
		taste, status,_,_:=gfx2.MausLesen1()
		
			if status == 1 && taste == 1 {
				ch <- 1
				
}	
	
}
}
