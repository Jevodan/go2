package ducks

import "fmt"

type RedDuck struct {
	Duck
}

func (r *RedDuck) Display() {
	fmt.Println("Красная утка", r.GetName(), " на экране!!!!")
}

func (r *RedDuck) Quack() {
	fmt.Println("Утка ", r.GetName(), " крякает")
}

func (r *RedDuck) Fly() {
	fmt.Println("Утка ", r.GetName(), " летает")
}
