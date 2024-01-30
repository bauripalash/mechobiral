package main

import (
	"fmt"
	"mechobiral/lib/alphabet"
	"os"
)

var AboutMechobiral string = `Mechobiral v0.1
Bengali Phonetic Library
(C) 2024 Palash Bauri`

func main() {
	fmt.Printf("Total Vowels -> %d\n" , len(alphabet.VOWELS_ALL) )
	fmt.Printf("Total Consonants -> %d\n" , len(alphabet.CONSONANTS_ALL))

	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "-h" {
			fmt.Println(AboutMechobiral)
		} else if arg == "-v" {
			fmt.Println("v0.1")
		}

	}
}
