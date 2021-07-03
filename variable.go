package main

import "fmt"

func main() {

	//Deklarasi String
	var nama string = "Seriusman Waruwu"
	fmt.Println(nama)

	//Tanpa tipe data
	var lastName = "Waruwu"
	fmt.Println(lastName)

	//Deklarasi ':=' menggantikan `var`
	firstName := "Seriusman"
	fmt.Println(firstName)

	//buat variabel sekaligus

	var (
		age   int    = 20
		study string = "golang"
	)

	fmt.Println("Age = ", age)
	fmt.Println("Study = ", study)

}
