package main

import "fmt"

func main() {
	var color string

	fmt.Print("Input Color : ")
	fmt.Scanln(&color)

	switch color {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Be Careful")
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Color Is Not Found")
	}
}

// func main() {
// 	var name string
// 	var age int

// 	fmt.Print("Input Your Name : ")
// 	fmt.Scanln(&name)
// 	fmt.Print("Input Your Age  : ")
// 	fmt.Scanln(&age)

// 	fmt.Println(name)
// 	fmt.Println(age)

// 	if name == "" {
// 		fmt.Println("Please Fill Your Name")
// 	} else if name != "" {
// 		if age > 0 && age <= 5 {
// 			fmt.Println("Hello", name, "You are toddler")
// 		} else if age > 5 && age <= 11 {
// 			fmt.Println("Hello", name, "You are children")
// 		} else if age > 11 && age <= 19 {
// 			fmt.Println("Hello", name, "You are teenager")
// 		} else if age > 20 && age <= 60 {
// 			fmt.Println("Hello", name, "You are adult")
// 		} else if age > 60 {
// 			fmt.Println("Hello", name, "You are elder")
// 		} else {
// 			fmt.Println("Hello", name, "Age must number")
// 		}
// 	}
// }
