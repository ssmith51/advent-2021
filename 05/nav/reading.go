package nav

type Coordinate struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

type VentGrid struct {
	MaxX        int
	MaxY        int
	Coordinates []Coordinate
}

func NewReading(x1 int, y1 int, x2 int, y2 int) Coordinate {
	reading := Coordinate{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
	return reading
}
