package pos

type Position struct {
	X int
	Y int
}
type Dimension = Position

type Direction struct {
	Dx     int
	Dy     int
	Symbol string
}

type Distance = Direction

var RIGHT = Direction{Dx: 1, Dy: 0, Symbol: ">"}
var DOWN = Direction{Dx: 0, Dy: 1, Symbol: "v"}
var LEFT = Direction{Dx: -1, Dy: 0, Symbol: "<"}
var UP = Direction{Dx: 0, Dy: -1, Symbol: "^"}

var Directions = map[string]Direction{
	UP.Symbol: UP, RIGHT.Symbol: RIGHT, DOWN.Symbol: DOWN, LEFT.Symbol: LEFT}

func (p Position) Equals(p2 Position) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p Position) Move(d Direction) Position {
	return Position{
		X: p.X + d.Dx,
		Y: p.Y + d.Dy,
	}
}

func (p Position) Add(d Distance) Position {
	return Position{
		X: p.X + d.Dx,
		Y: p.Y + d.Dy,
	}
}

func (d Direction) TurnRight() Direction {
	var result Direction
	switch d {
	case RIGHT:
		result = DOWN
	case DOWN:
		result = LEFT
	case LEFT:
		result = UP
	case UP:
		result = RIGHT
	}
	return result
}

func (p Position) DistanceTo(p2 Position) Distance {
	return Distance{Dx: p2.X - p.X, Dy: p2.Y - p.Y}
}

func AtPosition(lines []string, pos Position) uint8 {
	return lines[pos.Y][pos.X]
}
