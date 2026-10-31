package main

import "fmt"

func main() {
	// fmt.Println("Looping Biasa")
	// for i := 1; i <= 9; i += 2 {
	// 	fmt.Println("Perulangan Ke-,", i)
	// }

	// fmt.Println("Looping Mirip While")
	// i := 1

	// for i <= 5 {
	// 	fmt.Println("Perulangan Ke-, ", i)
	// 	i++
	// }

	// fmt.Println("Looping Break")
	// for i := 0; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Perulangan Ke -, ", i)
	// }

	// fmt.Println("Looping Continue")
	// for i := 0; i <= 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println("Perulangan Ke -, ", i)
	// }

	// fmt.Println("Array Looping")
	// buah := []string{"Apple", "Banana", "Mango", "Strawberry", "Watermelon"}

	// // WITH INDEX
	// // for index, value := range buah {
	// // 	fmt.Println(index, value)
	// // }

	// // WITHOUT INDEX
	// for _, value := range buah {
	// 	fmt.Println("Fruit Name : ", value)
	// }

	fmt.Println("---Map Looping---")

	bioUser := map[string]string{
		"Name":        "Muhammad Rifky Ramadani",
		"Fav Animal":  "Fish",
		"Fav Food":    "Fried Chicken",
		"Fav Tv Show": "Breaking Bad",
	}

	for key, value := range bioUser {
		fmt.Println(key, ":", value)
	}
}
