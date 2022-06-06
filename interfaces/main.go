package main

import "fmt"

const (
	SmallLift = iota
	StandarLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motocycle string
type Car string
type Truck string

func (m Motocycle) String() string {
	return fmt.Sprintf("Motocycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motocycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandarLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandarLift:
		fmt.Printf("send %v to standart lift\n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", p)
	}

}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motocycle := Motocycle("Croozer")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motocycle)
}
