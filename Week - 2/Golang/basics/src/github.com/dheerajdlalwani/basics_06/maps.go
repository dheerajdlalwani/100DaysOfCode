package main

import (
	"fmt"
)

func main() {

	// var score map[string]int //this used 'new' that means only alloc of mem & not init
	// score["Dheeraj"] = 99
	// fmt.Println(score)  //this will give error
	score := make(map[string]int) //this used 'make' that means alloc + init
	score["Dheeraj"] = 77
	fmt.Println(score) //this will not give error
	score["Jay"] = 100
	score["Devika"] = 85
	score["Tanya"] = 90
	fmt.Println(score)
	getTanyaScore := score["Tanya"]
	fmt.Println(getTanyaScore)
	delete(score, "Devika")
	fmt.Println(score["Devika"])
	fmt.Println(score)

}
