package saeulen

type Saeule interface {

//	Vor.: -
//	Erg.: Eine neue Säule entsteht
//	func New ()

//Hier kommen die Funktionen aus der sauelenimpl rein
//Datentyp aber hinter der Funktion abgeben
//Beispiel:
//	Vor.: -
//	Erg.: Zeichnet den oberen Teil der Säule
//	ObereErstellen() uint

SetzeWerte(x,hoehe,loch uint16)

Draw()
 
String () string
 
Move ()

SetzeZufallswerte()

GibXWert() uint16
}


func MoveList(l []Saeule){
	for i:=0;i<len(l);i++{
		l[i].Move()
	}
}
