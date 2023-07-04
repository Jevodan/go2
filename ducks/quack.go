package ducks

import "fmt"

type Quack struct {
}

func (q *Quack) Quack() {
	fmt.Println("Утка  крякает")
}
