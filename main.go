package main

import (
	calculation "basic-golang/entities"
	"fmt"
)

func main() {
	fmt.Println("Berikut Hasil Dari Kalkulasi")

	result := calculation.Add(9, 9)

	fmt.Println("9 dikali 9 =", result)

	fmt.Println("=====================")
	//Variable
	name := "Rohman Hidayat"
	age := 24
	ipk := 200
	var grade string

	fmt.Println("nama saya", name, "umur saya", age)

	fmt.Println("=====================")
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

	fmt.Println("=====================")
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

	fmt.Println("=====================")
	//For (Looping)
	for i := 5; i < 10; i++ {
		fmt.Println("Saya seorang programmer :", i)
	}

	title := "Golang is the best language"

	for i, letter := range title {
		if i%2 == 0 { // "%" merupakan modulus (i dibagi 2 = 0) kalau ga bisa dibagi 2 contoh 5 % 2 maka hasilnya 4/2 sisa 1 maka tulisannya (5%2 == 1)
			fmt.Println("letter :", string(letter), "index :", i)
		}
	}

	fmt.Println("=====================")
	//Array => Hanya String atau int saja

	// var typescript [5]string
	// typescript[0] = "GO"
	// typescript[1] = "Vue"
	// typescript[2] = "Express"
	// typescript[3] = "Laravel"
	// typescript[4] = "Angular"

	// length := len(typescript)
	// fmt.Println(typescript, length)

	mahasiswa := [...]string{
		"Lulu",
		"Rei",
		"Pei",
		"Numad",
		"Onplo"}

	// length := len(role)
	// fmt.Println(role, length)

	for i, nama := range mahasiswa {
		fmt.Println(" index : ", i, " nama : ", nama)
	}

	fmt.Println("=====================")
	// Mapping / Map

	var myMap map[string]int
	myMap = map[string]int{}

	myMap["Ruby"] = 0

	fmt.Println(myMap)

	// Mapping #2

	namaSiswa := map[string]string{
		"siswa 1": "Ahmad",
		"siswa 2": "Rohmad",
		"siswa 3": "Rendi",
		"siswa 4": "Omar",
	}
	fmt.Println(namaSiswa)

	for key, value := range namaSiswa {
		fmt.Println("key :", key, "value :", value)

		delete(namaSiswa, "siswa 3")
	}

	fmt.Println("=====================")

	myResultData()
}

func myResultData() {
	fmt.Println("skadosaidjsajdapda")
}
