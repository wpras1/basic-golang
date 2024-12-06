package main

import (
	calculation "basic-golang/entities"
	"fmt"
)

func main() {
	fmt.Println("Berikut Hasil Dari Kalkulasi")

	result := calculation.Add(9, 9)

	fmt.Println("9 dikali 9 =", result)

	//Variable
	name := "Rohman Hidayat"
	age := 24
	ipk := 200
	var grade string

	fmt.Println("nama saya", name, "umur saya", age)

	//If Else
	if age > 10 {
		fmt.Println("Saya boleh minum khamr")
	} else {
		fmt.Println("Saya tidak boleh minum khamr")
	}

	if ipk <= 275 {
		grade = "Bodoh"
	} else if ipk <= 375 {
		grade = "Biasa Saja"
	} else if ipk < 400 {
		grade = "Pandai"
	} else {
		grade = "Error"
	}

	fmt.Println("Saya merupakan mahasiswa berpredikat", grade)

	//Switch Case
	number := 4

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	case 4:
		fmt.Println("Empat")
	default:
		fmt.Println("Default")
	}

	//For (Looping)
	for i := 5; i < 10; i++ {
		fmt.Println("Saya seorang programmer :", i)
	}
}
