package main

import (. "gfx2"
		"time"
		)
func main () {
	Fenster (700, 800)
	Stiftfarbe (0,255,0)
	Vollrechteck (250,300,125,300)
	time.Sleep(5 * time.Second)
	}
