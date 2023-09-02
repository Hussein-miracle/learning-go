package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#red",
		"blue":  "#blue",
		"green": "#green",
	}

	colors["me"] = "Abdullahi"
	colors["white"] = "#FFF"
	fmt.Println(colors)
	delete(colors, "me")
	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
