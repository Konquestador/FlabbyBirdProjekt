package thread


import (
    "fmt"
    "gfx2" 
    "time"
)


type data struct {
	var speed float64 = 0
	var positionY float64
	var ja bool
}


func New() *data {
	var p *data
	p = new(data)
	p.ja = true
	return p
}


func (p *data) aenderRichtung() {
	for {
			taste, status,_,_:=gfx2.MausLesen1()
		
			if status == 1 && taste == 1 {
				speed = 20
				positionY -= speed
				
			}
		}
	}

