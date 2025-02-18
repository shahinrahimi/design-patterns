package factory

import "fmt"

type Car interface {
  getPower() int
  getWeight() float32
}

type SportCar struct {
  power int
  weight float32
}

func NewSportCar(power int, weight float32) *SportCar {
  return &SportCar{
    power, weight,
  }
}

func (sc *SportCar) getWeight() float32 {
  return sc.weight
}

func (sc *SportCar) getPower() int {
  return sc.power
}

type CarFactory struct {}

func (cf *CarFactory) makeSportCar(power int, weight float32) *SportCar {
  return NewSportCar(power, weight) 
}

func RunSimpleFactory() {
  carFactory := &CarFactory{}
  sportCar := carFactory.makeSportCar(200, 2.15)
  fmt.Printf("Power of car: %d\n", sportCar.getPower())
  fmt.Printf("Weight of car: %f\n", sportCar.getWeight())
}

