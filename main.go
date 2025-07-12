package main

import (
	"fmt"

	"simplecode/internal/hashtables"
)

func main() {

	testHash := hashtables.HashMap{}
	hashtables.Init(&testHash, 3)
	testHash.Insert("Ameer")
	testHash.Insert("Ameera")
	testHash.Insert("Ameerb")
	testHash.Insert("Ameerc")

	testHash.PrintMap()

	ameer := testHash.Get("Ameerb")
	if ameer != nil {
		fmt.Println("Found:", ameer.Value)
	} else {
		fmt.Println("Not found")
	}

	testHash.Delete("Ameer")
	testHash.PrintMap()

}
