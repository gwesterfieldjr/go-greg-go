package main

import (
	"fmt"
)

func main() {

	//var colors map[string]string

	colors := make(map[string]string)

	//colors := map[string]string{
	//"red": "#ff0000",
	//}

	colors["red"] = "#ff0000"
	colors["yellow"] = "#ff0000"
	colors["blue"] = "#ff0000"

	fmt.Println(colors)

	delete(colors, "red")

	fmt.Println(colors)

	printMap((colors))
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Color is " + color)
		fmt.Println("Hex is " + hex)
	}
}
