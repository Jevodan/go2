package ducks

import "fmt"

type Squeak struct {
	QuackBehavior
}

func (q *Squeak) Quack() {
	fmt.Println("Утка  пищщщит")
}
