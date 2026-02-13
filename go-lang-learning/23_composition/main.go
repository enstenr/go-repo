package main

import (
	"fmt"
	"strconv"
)

type Drivable interface {
	Start()
}
type Engine struct {
	model    string
	makeyear string
}

type Vehicle struct {
	Engine
	speed   float64
	vehType string
}
type Aeroplane struct {
	Vehicle
}

var makeYearStr = strconv.Itoa(2026)

func NewVehicle(model string, speed float64, vehType string) Vehicle {
	return Vehicle{Engine: Engine{model, makeYearStr}, speed: speed, vehType: vehType}
}
func NewAeroplane(model string, speed float64) Aeroplane {
	return Aeroplane{Vehicle: NewVehicle(model, speed, "aeroplane")}
}
func (v *Vehicle) Start() {
	//	fmt.Printf(" Engine Start: %s\n Max speed %f \n Model %s\n Make Year %s \n", v.vehType, v.speed, v.model, v.makeyear)
}
func (v *Vehicle) Stop() {
	//fmt.Printf(" Engine Stop: %s\n", v.vehType)
}
func (a *Aeroplane) Takeoff() {
	//fmt.Printf(" Engine Takeoff: %s\n", a.vehType)
}

// This function "Accepts an Interface"
// It doesn't care if it's a Car, Plane, or Boat.
// As long as it can Start(), this function can "Launch" it.
func Launch(d Drivable) {
	//fmt.Println("System Check...")
	d.Start()
}
func main() {
	//engine := &Vehicle{Engine: Engine{model: "Ritz"}, speed: 80, vehType: "Car"}
	//engine := NewVehicle("Ritz", 80, "Car")
	//engine.Start()
	//aeroplane :=  &Aeroplane{Vehicle: Vehicle{Engine: Engine{model: "Ritz"}, speed: 800, vehType: "Aeroplane"}}
	//aeroplane := NewAeroplane("Airbus", 800)
	//aeroplane.Start()
	//aeroplane.Takeoff()
	v1 := NewVehicle("Ritz", 80, "Car")
	p1 := NewAeroplane("Airbus", 800)
	v2 := NewVehicle("Tesla", 120, "Electric Car")
	fleet := []Drivable{
		&v1, // Get the address so it satisfies (*Vehicle).Start()
		&p1, // Get the address so it satisfies (*Aeroplane).Start()
		&v2,
	}

	fmt.Println("--- Fleet Status ---")
	for _, d := range fleet {
		//println()
		Launch(d)

		if plane, ok := d.(*Aeroplane); ok {
			plane.Takeoff()

		}

	}
}
