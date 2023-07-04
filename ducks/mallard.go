package ducks

import (
	"fmt"
)

var mallard = mallardDuck{}

type mallardDuck struct {
	Duck
}

func NewMallardDuck() mallardDuck {
	mallard.name = "Утка мандаринка"
	mallard.flyBehavior = &FlyWithWings{}
	mallard.quackBehavior = &Squeak{}
	return mallard
}

func (m *mallardDuck) Display() {
	fmt.Println("Утка маллард дак на экране!!!!", m.GetName())
}
