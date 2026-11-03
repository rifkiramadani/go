package main

import "fmt"

// OBJECT POINTER
type Hero struct {
	health uint
}

func AttackByEnemy(p *Hero) {
	p.health -= 10
}

func main() {
	// OBJECT POINTER
	allies := Hero{health: 100}

	AttackByEnemy(&allies)
	fmt.Println(allies.health)

	// BASIC POINTER
	// fmt.Println("--Pointer--")

	// x := 10

	// var p *int = &x

	// fmt.Println("Nilai x : ", x)
	// fmt.Println("Alamat x : ", &x)
	// fmt.Println("Nilai Pointer : ", p)
	// fmt.Println("Nilai yang ditunjuk : ", *p)
}
