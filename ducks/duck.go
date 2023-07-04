package ducks

import "fmt"

type Duck struct {
	name          string
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) Swim() {
	fmt.Println("Утка ", d.GetName(), "  плавает")
}

func (d *Duck) Display() {
	fmt.Println("Утка ", d.GetName(), "  перед нами")
}

func (d *Duck) GetName() string {
	return d.name
}

func (d *Duck) PerformQuack() {
	d.quackBehavior.Quack()
}

func (d *Duck) PerformFly() {
	d.flyBehavior.Fly(d.name)
}

func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb QuackBehavior) {
	d.quackBehavior = qb
}
