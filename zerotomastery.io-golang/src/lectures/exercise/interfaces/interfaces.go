//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.

//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name

//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	smallLift = iota
	standardLift
	largeLift
)

type lift int

type LiftPicker interface {
	PickLift() lift
}

type Truck string
type Car string
type Motorcicle string

func (t Truck)  String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (c Car)  String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (m Motorcicle) String() string {
	return fmt.Sprintf("Motorcicle: %v", string(m))
}

func (t Truck)  PickLift() lift {
	return largeLift
}

func (c Car)  PickLift() lift {
	return standardLift
}

func (m Motorcicle) PickLift() lift {
	return smallLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift(){
	case smallLift:
		fmt.Printf("send %v to small lifts\n", p)
	case standardLift:
		fmt.Printf("send %v to standard lifts\n", p)
	case largeLift:
		fmt.Printf("send %v to large lifts\n", p)
	}
}

func main() {
	car := Car("Ferrari")
	truck := Truck("Volvo")
	motorcicle := Motorcicle("Honda")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcicle)
}
