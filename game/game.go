package main

import "fmt"

type Location struct {
	X int
	Y int
}

type Player struct {
	Name     string
	Location // Player Embeds Location
}

type Mover interface {
	Move(int, int)
}

const (
	maxX = 1000
	maxY = 600
)

// "l" is called the reciever
func (l Location) Move(x, y int) {
	l.X = x
	l.Y = y
}

// func NewLocation(x, y int) Location
// func NewLocation(x, y int) (Location, error)
// func NewLocation(x, y int) *Location
func NewLocation(x, y int) (*Location, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	loc := Location{
		X: x,
		Y: y,
	}
	return &loc, nil // Go does escape analysis, allocate loc on heap
}

func main() {
	var loc Location
	fmt.Printf("loc  (v): %v\n", loc)
	fmt.Printf("loc (+v): %+v\n", loc)
	fmt.Printf("loc (#v): %#v\n", loc)

	loc = Location{1, 2} // must specify all fields
	fmt.Println(loc)

	loc = Location{
		Y: 1,
		//X: 2,
	}
	fmt.Println(loc)
	fmt.Println(NewLocation(100, 200))
	fmt.Println(NewLocation(100, -200))

	loc.Move(200, 600)

	p1 := Player{
		Name:     "Parazival",
		Location: Location{100, 200},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.Location.X: %#v\n", p1.Location)
	p1.Move(600, 500)
	fmt.Printf("p1.X: %#v\n", p1)
	ms := []Mover{
		&loc,
		&p1,
	}
	moveAll(ms, 1, 2)
	fmt.Printf("loc: %#v, p1: %#v\n", loc, p1)

}

/* Thought experiment: Sortable interface

func Sort(s Sortable) { }

type Sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

*/

func moveAll(ms []Mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}
