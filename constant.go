package main

import "fmt"

func main() {
	//Const adalah var yang tidak bisa diubah lagi nilainya
	const firstName string = "Seriusman"

	//tanpa menyisipkan tipe data
	const lastName = "Waruwu"

	fmt.Println("First Name = ", firstName)
	fmt.Println("Last Name = ", lastName)

	//contoh bedanya const ==> tidak akan bisa di ubah
	// firstName = "New Name"
	// lastName = "New Last Name"

	// Note :
	// Deklarasi Pertama kali const harus ada nilai/value
	// walaupun tidak digunakan, maka const tidak masalah
}
