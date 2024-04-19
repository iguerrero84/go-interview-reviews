package designpatterns

// Define an interface for creating an object, but let subclasses decide which class to instanciate

type Shape interface {
	Draw() string
}

type Circle struct{}

func (c *Circle) Draw() string {
	return "Drawing Circle"
}

type Square struct{}

func (s *Square) Draw() string {
	return "Drawing Square"
}

func GetShape(shapeType string) Shape {
	if shapeType == "circle" {
		return &Circle{}
	} else if shapeType == "square" {
		return &Square{}
	}
	return nil
}
