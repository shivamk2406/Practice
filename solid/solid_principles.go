package solid

import (
	"fmt"
	"math"
)

/*
S - Single-responsibility Principle --> A class should have one and only one reason to change, meaning that a class should have only one job.
O - Open-closed Principle
L - Liskov Substitution Principle
I - Interface Segregation Principle
D - Dependency Inversion Principle
*/

// ShapeArea interface - has only one responsibility: defining Area calculation behavior
type ShapeArea interface {
	AreaCalculator() float64
}

type Square struct {
	length int
}

func(s Square) AreaCalculator() float64{
	return math.Pow(float64(s.length),2.0)
}

type Circle struct {
	radius int
}

func(c Circle) AreaCalculator() float64{
	return math.Pi* math.Pow(float64(c.radius),2)
}


type AreaCalculator struct {
	shapes []ShapeArea
}

func Driver(){
	area:= AreaCalculator{
		[]ShapeArea{Circle{radius: 2}, Square{length: 4}},
	}

	for _, shape:= range area.shapes {
		fmt.Println(shape.AreaCalculator())

	}
}

//THE ABOVE CODE ADHERES SRP


// Now lets say one more requirement came where we need to add logic of adding perimeter as well. So, in this case 
// there are two ways of doing this
// If you add operations like perimeter calculation, 
// you might be tempted to modify existing structs like Square and Circle to implement a new method for calculating perimeters. However, this would:

// Violate SRP: Now, the structs are responsible for both area and perimeter calculations.
// Violate OCP: Existing code would need to be modified to accommodate the new functionality, which increases the risk of introducing bugs.
// So the solution is to create a Perimeter interface and then implement their logic so it would adhere to OCP and SRP.

type PerimeterCalc interface{
	PerimeterCalculate() float64
}

type PerimeterCalculator struct{
	shapes []PerimeterCalc
}

func(s Square) PerimeterCalculate() float64{
	return float64(4 * s.length)
}

func(c Circle) PerimeterCalculate() float64{
	return 2 * math.Pi * float64(c.radius)
}

func Driver2(){
	area:= AreaCalculator{
		[]ShapeArea{Circle{radius: 2}, Square{length: 4}},
	}

	for _, shape:= range area.shapes {
		fmt.Println(shape.AreaCalculator())

	}

	perims:=PerimeterCalculator{
		[]PerimeterCalc{
			Circle{
				radius: 2,
			},
			Square{
				length: 3,
			},
		},
	}

	for _, shape:= range perims.shapes {
		fmt.Println(shape.PerimeterCalculate())

	}
}