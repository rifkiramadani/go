package main

import "fmt"

// STRUCT
// type Person struct {
// 	name string
// 	age  uint
// 	bio  string
// }

// func introduction(p Person) string {
// 	output := fmt.Sprintf(`Hi my name is %s, i am %d years old, my biodata is %s`, p.name, p.age, p.bio)
// 	return output
// }

// NESTED STRUCT
// type Jobs struct {
// 	name     string
// 	position string
// }

// type People struct {
// 	name     string
// 	age      uint
// 	listJobs []Jobs
// }

// INTERFACE
type Shape interface {
	getPerimeter() uint
}

type Square struct {
	width uint
}

type Triangle struct {
	a uint
	b uint
	c uint
}

func (t Triangle) getPerimeter() uint {
	return t.a * t.b * t.c
}

func (s Square) getPerimeter() uint {
	return 4 * s.width
}

func main() {
	// INTERFACE
	var shape1 Shape = Triangle{2, 2, 2}
	fmt.Println(shape1)
	fmt.Println(shape1.getPerimeter())

	var shape2 Shape = Square{10}
	fmt.Println(shape2.getPerimeter())

	// NESTED STRUCT
	// p3 := People{name: "rifky", age: 17, listJobs: []Jobs{
	// 	{
	// 		"Meta", "Frontend Dev",
	// 	},
	// 	{
	// 		"Google", "Lead Dev",
	// 	},
	// }}

	// fmt.Println(p3)

	// STRUCT
	// p2 := Person{name: "rifky", age: 17, bio: "Hello world"}
	// fmt.Println(introduction(p2))

	// p1Empty := Person{}
	// fmt.Println(p1Empty)

	// p1 := Person{"rifky", 17, "Hello world"}
	// p1.name = "putro"
	// fmt.Println(p1)
}
