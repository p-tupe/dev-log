// This program shows how structs and intefaces can be embedded to form higher types.
//
//	Usage: go run main.go
package main

import "fmt"

// Consider a few of interfaces (abstract types), that have some behaviour
type Starter interface{ Start(model string) }
type Stopper interface{ Stop(model string) }
type Namer interface{ Name() string }

// These behaviours can be embedded into higher level interfaces
type Car interface {
	Starter
	Stopper
	Namer
	// This behaviour is specific to Car
	Drift()
}

type Plane interface {
	Starter
	Stopper
	Namer
	// This behaviour is specific to Plane
	Tilt()
}

// Some functions that expect these higher interfaces
func driveCar(c Car) {
	c.Start(c.Name())
	c.Drift()
	c.Stop(c.Name())
}

func flyPlane(p Plane) {
	p.Start(p.Name())
	p.Tilt()
	p.Stop(p.Name())
}

// Some structs (concrete types) implement the basic behaviours
type Engine struct{}

func (e Engine) Start(model string) { fmt.Println("Starting " + model) }
func (e Engine) Stop(model string)  { fmt.Println("Stopping " + model) }

type HandBrake struct{}

func (t HandBrake) Drift() { fmt.Println("Drifting") }

type Yoke struct{}

func (y Yoke) Tilt() { fmt.Println("Tilting...") }

// More concrete types that define an entity implementing above basic behaviours
type Corrolla struct {
	Engine
	HandBrake
	model string
}

func (c Corrolla) Name() string { return c.model }

type Boeing struct {
	Engine
	Yoke
	model string
}

func (b Boeing) Name() string { return b.model }

func main() {
	corolla := Corrolla{Engine{}, HandBrake{}, "corolla"}
	boeing := Boeing{Engine{}, Yoke{}, "boeing"}

	driveCar(corolla)
	flyPlane(boeing)
}
