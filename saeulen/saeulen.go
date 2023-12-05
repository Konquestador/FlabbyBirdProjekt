package saeulen

type Saeule interface {
	SetzeWerte(x, hoehe, loch uint16)
	Draw(birdposX int, width int)
	String() string
	Move()
	SetzeZufallswerte()
	GibXWert() uint16
	GibPassed() bool
	Passed()
	GibBreite() uint16
	GibHoehe() uint16
	GibLoch() uint16
	GibbirdInHole() bool
	NotInHole()
	InHole()
	Touch()
}
