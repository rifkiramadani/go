package main

import "fmt"

func main() {
	// fmt.Println("---Array---")

	// nilai := [5]uint{50, 66, 85, 67, 70}
	// panjangNilai := len(nilai)
	// nilai4 := nilai[3]

	// fmt.Println(nilai)
	// fmt.Println("Panjang Nilai : ", panjangNilai)
	// fmt.Println("Nilai Ke-4 : ", nilai4)

	// kumpulanNilai := [2][5]int{{50, 66, 85, 67, 70}, {1, 2, 3, 4, 5}}
	// kumpulanNilai1 := kumpulanNilai[0][4]

	// kumpulanNilai[1] = [5]int{2, 5, 6, 8, 3}

	// fmt.Println(kumpulanNilai)
	// fmt.Println(kumpulanNilai1)

	// fmt.Println("---Slice---")

	// buah := []string{"Banana", "Apple", "Orange", "Grape", "Melon"}
	// buah1 := buah[:3]
	// buah2 := buah[2:]
	// buah[0] = "Pisang"
	// addBuah := append(buah, "Pineapple", "Wartermelon", "Kiwi")

	// fmt.Println(buah)
	// fmt.Println("Ganti Buah Indeks ke 0 : ", buah[0])
	// fmt.Println("3 Buah Pertama : ", buah1)
	// fmt.Println("Buah Dari Indeks Ke 2 : ", buah2)
	// fmt.Println("Buah 3 : ", buah[2])
	// fmt.Println("Jumlah Buah : ", len(buah))
	// fmt.Println("Tambah Buah : ", addBuah)

	fmt.Println("---Map---")

	cars := map[string]string{
		"car1": "Avanza",
		"car2": "Xenia",
		"car3": "Innova",
		"car4": "Pajero",
		"car5": "Marcedez",
	}

	countCars := len(cars)
	cars["car4"] = "Veloz"
	delete(cars, "car5")

	fmt.Println("List Mobil : ", cars)
	fmt.Println("Jumlah Mobil : ", countCars)
	fmt.Println("Car 1 : ", cars["car1"])

}
