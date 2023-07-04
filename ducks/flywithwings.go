package ducks

import (
	"fmt"
)

type FlyWithWings struct {
	FlyBehavior
}

func (f *FlyWithWings) Fly(name string) {
	fmt.Println("Утка", name, "летает")
}
