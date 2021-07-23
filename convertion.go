package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("Nilai 32 = ", nilai32)
	fmt.Println("Nilai 64 = ", nilai64)
	fmt.Println("Nilai 8 = ", nilai8) // -96, cuz kembali ke nilai bawah

	var firstValue int32 = 127 //int8 limit 127
	var conValue int8 = int8(firstValue)

	fmt.Println("First Value = ", firstValue)
	fmt.Println("Convertion Value = ", conValue)

	var name = "Serius"
	var cName = name[0]
	var rName = string(cName)
	var random = string(83)

	fmt.Println("Name : ", name)
	fmt.Println("Character Name : ", cName)
	fmt.Println("Return Name : ", rName)
	fmt.Println("Return Random : ", random)

	for i := 0; i < 10000; i++ {
		fmt.Println(i, ". => ", string(i))
	}

}
