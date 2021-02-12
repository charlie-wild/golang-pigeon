package main

import (
	"fmt"

	"github.com/charlie-wild/golang-pigeon/pigeon"
)

func main() {

	pigeonOne := pigeon.Pigeon{
		Name:         "BigWing",
		Colour:       "blue",
		HomeLocation: "Chorlton",
	}

	fmt.Println(pigeonOne.GetLocation())

}
