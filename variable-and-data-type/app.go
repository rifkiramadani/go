package main

import "fmt"

var name string = "Rifky"
var age uint = 17
var isComplete bool = true
var emoji rune = '😹'

func main() {
	fullName := "Muhammad Rifky Ramadani"
	grade := 100
	isPass := true
	height := 179.9

	fmt.Println("My Name 		: ", name)
	fmt.Println("I Am 			: ", age, "Years Old")
	fmt.Println("IsComplete 	: ", isComplete)

	fmt.Println("My Full Name 	: ", fullName)
	fmt.Println("My Grade 		: ", grade)
	fmt.Println("isPass 		: ", isPass)
	fmt.Println("My Height 		: ", height)
	fmt.Println("Fav Emoji 		: ", string(emoji))

}
