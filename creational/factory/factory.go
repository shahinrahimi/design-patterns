package factory

import (
	"fmt"
)

// Shape is the interface that concrete shapes will implement
type Shape interface {
  Draw() string
}

// Circle is the concrete implementation of the Shape interface
type Circle struct {}

func (c *Circle) Draw() string {
  return "Drawing a Circle"
}
// Square is the concrete implementation of the Shape interface
type Square struct {}

func (s *Square) Draw() string {
  return "Drawing a Square"
}
// ShapeFactory is the factory method that return the appropriate shape.
func ShapeFactory(shapeType string) Shape {
  switch shapeType {
    case "circle":
      return &Circle{}
    case "square":
      return &Square{}
    default:
      return nil
  } 
}



func Run() {
  circle := ShapeFactory("circle")
  if circle != nil {
    fmt.Println(circle.Draw())
  }
  square := ShapeFactory("square")
  if square != nil {
    fmt.Println(square.Draw())
  }

  invalidShape := ShapeFactory("triangle")
  if invalidShape == nil {
    fmt.Println("Invalid shape type!")
  }
}
