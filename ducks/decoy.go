package ducks

import "fmt"

var decoy decoyDuck

type decoyDuck struct {
	Duck
}

func NewDecoyDuck() decoyDuck {
	decoy.name = "Деревянная утка"
	decoy.flyBehavior = &NoFly{}
	decoy.quackBehavior = &NoQuack{}
	return decoy
}

func (d *decoyDuck) Display() {
	fmt.Println("Утка маллард дак на экране!!!!", d.GetName())
}
