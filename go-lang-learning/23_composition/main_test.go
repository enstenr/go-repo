package main

import (
	"testing"
)

var globalAeroplane Aeroplane

// BenchmarkNewAeroplane tests the speed of object allocation and initialization
func BenchmarkNewAeroplane(b *testing.B) {

	for i := 0; i < b.N; i++ {
		globalAeroplane = NewAeroplane("Airbus A380", 900.0)
	}

}

// BenchmarkLaunchInterface tests the overhead of calling a method through an interface
// This is often slightly slower than a direct method call due to "dynamic dispatch"
func BenchmarkLaunchInterface(b *testing.B) {
	plane := NewAeroplane("Boeing 747", 950.0)

	b.ResetTimer() // Don't count the setup time above
	for i := 0; i < b.N; i++ {
		Launch(&plane)
	}
}

// BenchmarkTypeAssertion tests the performance cost of checking
// if a Drivable is an *Aeroplane (the "ok" check)
func BenchmarkTypeAssertion(b *testing.B) {
	var p = NewAeroplane("Cessna", 200.0)
	var d Drivable = &p
	for i := 0; i < b.N; i++ {
		if plane, ok := d.(*Aeroplane); ok {
			_ = plane // Use the variable to prevent compiler optimization
		}
	}
}

func BenchmarkFleetLaunch(b *testing.B) {
	// Creating a slice of interfaces usually forces heap escape
	fleet := []Drivable{
		&Aeroplane{Vehicle: NewVehicle("Airbus", 800, "Plane")},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Launch(fleet[0])
	}
}
