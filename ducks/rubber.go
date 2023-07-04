package ducks

import "fmt"

type RubberDuck struct {
	Duck
}

func (r *RubberDuck) Display() {
	fmt.Println("Резиновая утка", r.GetName(), " на экране!!!!")
}

func (r *RubberDuck) Quack() {
	fmt.Println("Резиновая утка", r.GetName(), " пищит!!!!")
}
