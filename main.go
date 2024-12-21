package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var jajanan string

	const NMAX = 20
	jajananTradisional := [NMAX]string{
		"Klepon",
		"Kue Cubir",
		"Pisang Goreng",
		"Martabak Manis",
		"Susi",
		"Pempek",
		"Dadar Gulung",
		"Lupis",
		"Gempol Puyuh",
		"Roti Bakar",
		"Lontong Balap",
		"Kue Lumpur",
		"Cendol",
		"Kue Pancong",
		"Tahu Tempe Bacem",
		"Bala-bala",
		"Es Doger",
		"Kue Apem",
		"Kue Serabi",
		"Kue Soes",
	}

	fmt.Print("Jajanan yang ingin anda cari: ")
	scanner.Scan()
	jajanan = scanner.Text()
	var lenData = len(jajananTradisional)

	// fmt.Println()
	fmt.Print("Hasil Sequential Search Iteratif: ")
	if iterativeSequentialSearch(jajanan, jajananTradisional[:], lenData) != -1 {
		fmt.Printf("ditemukan pada index ke-%d.\n", iterativeSequentialSearch(jajanan, jajananTradisional[:], lenData))
	} else {
		fmt.Printf("'%s' tidak ditemukan.\n", jajanan)
	}

	fmt.Print("Hasil Sequential Search Rekrusif: ")
	if recursiveSequentialSearch(jajanan, jajananTradisional[:], lenData, 0) != -1 {
		fmt.Printf("ditemukan pada index ke-%d.\n", recursiveSequentialSearch(jajanan, jajananTradisional[:], lenData, 0))
	} else {
		fmt.Printf("'%s' tidak ditemukan.\n", jajanan)
	}
}

// Fungsi Sequential Search Iteraif
func iterativeSequentialSearch(X string, Data []string, lenData int) int {
	for i := 0; i < lenData; i++ {
		if X == Data[i] {
			return i
		}
	}
	return -1
}

// Funsi Sequential Search Rekrusif
func recursiveSequentialSearch(X string, Data []string, lenData int, IDX int) int {
	if IDX >= lenData {
		return -1
	}
	if X == Data[IDX] {
		return IDX
	} else {
		return recursiveSequentialSearch(X, Data, lenData, IDX+1)
	}
}
