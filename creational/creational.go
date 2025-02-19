package creational

import (
	"design-patters/creational/singleton"
	"design-patters/creational/factory"
  "design-patters/creational/prototype"
)

func Run() {
	singleton.Run()
  factory.Run()
  factory.RunSimpleFactory()
  factory.RunAbstractFactory()
  factory.RunBuilder()
  prototype.RunPrototype()
}

