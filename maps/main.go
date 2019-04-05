package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#4bf745",
		"white": "#fffff",
	}
	fmt.Println(colors)

	// One empty map
	var colors2 map[string]string
	fmt.Println(colors2)

	// Another empty map
	colors3 := make(map[string]string)
	fmt.Println(colors3)

	// Access a map property
	colors4 := make(map[string]string)
	colors4["white"] = "#ffffff"

	fmt.Println(colors4)

	// int to string map
	colors5 := make(map[int]string)
	colors5[10] = "#ffffff"

	fmt.Println(colors5)

	// Delete map entry
	delete(colors5, 10)

	fmt.Println(colors5)

	// Print initial map
	printMap(colors)

	stuff := stuff{}
	stuff.thing = "a" == "a"
	fmt.Println(stuff)

	num := 15
	switch num {

	case 10:
	case 15:
		fmt.Println("it's stacked!")
	case 5:
		fmt.Println("you got it!")
	}
}

type stuff struct {
	thing bool
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
