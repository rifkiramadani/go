package main

import "fmt"

// BASIC ERROR HANDLING
// func divine(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("Tidak bisa dibagi nol")
// 	}
// 	return a / b, nil //nil == kosong
// }

// BASIC ERROR HANDLING 2
// func getUser(id int) error {
// 	if id <= 0 {
// 		return fmt.Errorf("User dengan id %d tidak ditemukan", id)
// 	}
// 	return nil
// }

// DEFER, PANIC, RECOVER
// func handlePanic() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Program Berhasil dipulihkan dari panic")
// 	}
// }

// func bagi(a, b int) {
// 	fmt.Println("--Error Handling--")

// 	defer handlePanic()

// 	if b == 0 {
// 		panic("Tidak Bisa Di bagi 0")
// 	}

// 	fmt.Println("Hasil Bagi : ", a/b)
// }

// BASIC GENERIC
// func PrintData[T any](data T) {
// 	fmt.Println(data)
// }

// STRUCT GENERIC
type Response[T any] struct {
	Data    T
	message string
}

func main() {
	// STRUCT GENERIC
	user1 := Response[string]{
		Data:    "Rifky",
		message: "Success",
	}

	response2 := Response[int]{
		Data:    200,
		message: "100",
	}

	fmt.Println(user1)
	fmt.Println(response2)

	// BASIC GENERIC
	// PrintData("Rifky")
	// PrintData(7)
	// PrintData(true)

	// DEFER, PANIC, RECOVER
	// bagi(10, 1)
	// bagi(10, 0)
	// fmt.Println("Program Selesai")

	// fmt.Println("--Error Handling--")

	// BASIC ERROR HANDLING 2
	// fmt.Println(getUser(-7))

	// BASIC ERROR HANDLING
	// hasil, err := divine(5, 1)

	// if err != nil { // kalau tidak error atau error tidak kosong
	// 	fmt.Println("Error: ", err)
	// }

	// fmt.Println("Hasil : ", hasil)
}
