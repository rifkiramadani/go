package main

import (
	"fmt"
	"time"
)

// GO ROUTINE
// SYNCHRONOUS (BLOCKING)
// func sendMail(target string) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Success Sending Email To : ", target)
// }

// ASYNCHRONOUS (NON-BLOCKING)
// func sendMail(target string) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Success Sending Email To : ", target)
// }

// CHANNEL
func fetchDataAPI(ch chan string) {
	fmt.Println("Sedang Mengambil Data Dari API")
	time.Sleep(2 * time.Second)

	ch <- "Data User Berhasil Di Ambil"
	fmt.Println("Data terkirim ke channel")
}

func main() {
	// CHANNEL
	fmt.Println("--Channel--")

	promise := make(chan string)

	go fetchDataAPI(promise)

	fmt.Println("Loading Tunggu API")

	dataDariApi := <-promise

	fmt.Println("Hasil", dataDariApi)

	// GO ROUTINE
	// ASYNCHRONOUS (NON-BLOCKING)
	// go sendMail("Rifky@mail.com")
	// go sendMail("Putro@mail.com")
	// go sendMail("Purot@mail.com")
	// go sendMail("Dwi@mail.com")
	// go sendMail("Saputro@mail.com")

	// fmt.Println("All Email Has Been Sent")

	// time.Sleep(2 * time.Second)

	// SYNCHRONOUS (BLOCKING)
	// sendMail("Rifky@mail.com")
	// sendMail("Putro@mail.com")
	// sendMail("Purot@mail.com")
	// sendMail("Dwi@mail.com")
	// sendMail("Saputro@mail.com")

	// fmt.Println("All Email Has Been Sent")

}
