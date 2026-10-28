package main

import "fmt"

func multiple(num1, num2 int) string {
	kali := num1 * num2
	output := fmt.Sprintf("angka %d dikali %d hasilnya adalah %d", num1, num2, kali)
	return output
}

func main() {
	fmt.Println(multiple(7, 7))
	fmt.Println(multiple(5, 5))
	fmt.Println(multiple(6, 7))
}
