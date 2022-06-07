package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (element Square) End() Point {
	return Point{x: element.start.x + int(element.a), y: element.start.y - int(element.a)}
}

func (element Square) Area() uint {
	return element.a * element.a
}

func (element Square) Perimeter() uint {
	return 4 * element.a
}
