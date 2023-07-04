package main

import "ducks/ducks"

func main() {

	mallard := ducks.NewMallardDuck()
	decoy := ducks.NewDecoyDuck()
	mallard.SetFlyBehavior(&ducks.NoFly{})

	var ducksArr []ducks.Duck
	ducksArr = append(ducksArr, mallard.Duck, decoy.Duck)

	for _, value := range ducksArr {
		value.PerformFly()
		value.Swim()
		value.PerformQuack()
		value.Display()
	}

}
