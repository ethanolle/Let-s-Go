package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":"lol",
		"green":"amazing",
		"fun":"LOl",
	}
	loopAndPrint(colors)
}

func loopAndPrint(c map[string]string){
	for color,joke:= range c{
		fmt.Println(color,joke)
	}
}