package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Taking user input")

	// var myString string
	// fmt.Scanln(&myString) // giving reference to var using '&'
	// fmt.Println(myString) // problen is nothing is stored after space

	var name string = "dheeraj"
	var a, _ = fmt.Println(name)
	fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter full name: ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumrating + 2)

}
