package prototype

import "fmt"

// prototype
type Cloneable interface {
  Clone() Cloneable
}

// ConcretePrototype
type Car struct {
  Model string
  Year int 
}


func (c *Car) Clone() Cloneable {
  return &Car{
    Model: c.Model,
    Year: c.Year,
  }
}

func RunPrototype() {
  // Original car
  originalCar:= &Car{
    Model: "Tesla Model S",
    Year: 2022,
  }

  // Clone the original car
  clonedCar := originalCar.Clone().(*Car)

  // Modify the cloned car
  clonedCar.Year = 2023

  // Output
  fmt.Printf("Original car: %s, Year %d\n", originalCar.Model, originalCar.Year)
  fmt.Printf("Cloned car: %s, Year %d\n", clonedCar.Model, clonedCar.Year)


}
