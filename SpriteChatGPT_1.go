package main

import (
    "fmt"
    "gfx2" // Import the gfx2 library (replace "your-username" with the actual package path)
    
    "time"
)

func main() {
	
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
    var velocityX, velocityY int = 10, 10 // Use int for velocity

    // Main game loop
    for {
        // Clear the screen
        gfx2.UpdateAus()
        gfx2.Cls()

        // Update the sprite's position
        spriteX += uint16(velocityX) // Convert int to uint16
        spriteY += uint16(velocityY) // Convert int to uint16

        // Check for collision with screen edges
        if spriteX < 0 || spriteX > 800 {
            velocityX = -velocityX // Reverse X velocity on collision
        }
        if spriteY < 0 || spriteY > 630 {
            velocityY = -velocityY // Reverse Y velocity on collision
        }

        // Draw the sprite at its new position
        gfx2.LadeBild(spriteX, spriteY, "Frame-1.bmp")

        // Update the graphics window
        gfx2.UpdateAn()

        // Delay for a short time (e.g., 60 frames per second)
        time.Sleep(1000 / 60 * time.Millisecond)
    }
}
