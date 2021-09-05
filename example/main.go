package main

import (
	"fmt"

	namer "github.com/voidshard/namer"
)

func main() {
	n, err := namer.New()
	if err != nil {
		panic(err)
	}

	// choose names at random
	fmt.Println("town:", n.Town())

	a, b := n.Male()
	fmt.Println("character [male]:", a, b)

	a, b = n.Female()
	fmt.Println("character [female]:", a, b)

	a, b = n.River()
	fmt.Println("river:", a, b)

	num := 5
	for _, tag := range n.Tags() {
		fmt.Printf("Tag [%s]\n", tag)

		fmt.Printf("\tTowns\n")
		for i := 0; i < num; i++ {
			fmt.Printf("\t\t%s\n", n.Tag(tag).Town())
		}

		fmt.Printf("\tCharacter [male]\n")
		for i := 0; i < num; i++ {
			a, b := n.Tag(tag).Male()
			fmt.Printf("\t\t%s %s\n", a, b)
		}

		fmt.Printf("\tCharacter [female]\n")
		for i := 0; i < num; i++ {
			a, b := n.Tag(tag).Female()
			fmt.Printf("\t\t%s %s\n", a, b)
		}

		fmt.Printf("\tRiver\n")
		for i := 0; i < num; i++ {
			a, b := n.Tag(tag).River()
			fmt.Printf("\t\t%s %s\n", a, b)
		}
	}

}
