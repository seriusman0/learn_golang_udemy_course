package main

import "fmt"

func main() {
	//apa itu slice ==> sejenis array yang lebih fleksibel
	bulan := [12]string{
		"Januari", "Februari", "Maret",
		"April", "Mei", "Juni",
		"Juli", "Agustus", "September",
		"Oktober", "November", "Desember", "Tambahan",
	}

	fmt.Println(len(bulan))
	fmt.Println(cap(bulan))

}
