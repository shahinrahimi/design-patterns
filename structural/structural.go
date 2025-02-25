package structural

import (
	"design-patters/structural/adaptor"
	"design-patters/structural/bridge"
	"design-patters/structural/composite"
	"design-patters/structural/decorator"
	"design-patters/structural/facade"
)

func Run() {
  adaptor.Run()
  bridge.Run()
  composite.Run()
  decorator.Run()
  facade.Run()
}

