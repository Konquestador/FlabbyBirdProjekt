package main

import (
    "fmt"
    "gfx2" // Import the gfx2 library (replace "your-username" with the actual package path)
    "time"
)

func main() {
	var timeInterval float64
	timeInterval = float64(0.1)
	var acceleration float64 = -9.81
	var speed float64 = 0
		
		
		
    // Open the graphics window with specified dimensions
    gfx2.Fenster(1000, 800)
	fmt.Println(gfx2.Grafikspalten())
    // Load an image file for the sprite
    gfx2.LadeBild(uint16(100), uint16(100), "Frame-1.bmp") // Convert int to uint16
    
    

    // Set the sprite's initial position
    var spriteX, spriteY uint16 = 100, 100 // Use uint16 for sprite position

    // Set the sprite's transparency (0-255)
    gfx2.Transparenz(0)

    // Set the sprite's initial velocity
    var _, velocityY float64 = 0, -90// Use int for velocity
	var positionY float64
	positionY = float64(spriteY)
    // Main game loop
    for {
	
		speed += timeInterval * acceleration
		positionY -= speed * timeInterval
        // Clear the screen
        gfx2.UpdateAus()
        gfx2.Stiftfarbe(0,0,0)
        gfx2.Cls()

        // Update the sprite's position
       // spriteX += uint16(velocityX) // Convert int to uint16
       
			taste, status,_,_:=gfx2.MausLesen1()
		
			if status == 1 && taste == 1 {
				positionY += velocityY // Convert int to uint16
			

        // Check for collision with screen edges
      // if spriteX < 0 || spriteX > 800 {
       //     velocityX = -velocityX // Reverse X velocity on collision
        }
        if positionY < 0 || positionY > 630 {
            velocityY = -velocityY // Reverse Y velocity on collision
        }
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
        time.Sleep(1000 / 60 * time.Millisecond)
        
        
    }
}

