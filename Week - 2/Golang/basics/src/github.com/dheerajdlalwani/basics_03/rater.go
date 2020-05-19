package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var name string
	var userRating string

	//Frontend
	//Take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter full name: ")
	name, _ = reader.ReadString('\n')

	//Take rating from user
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate my code between 1 & 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Backend
	fmt.Printf("Hello %v, \n Thanks for rating my code with %v star rating. \n\n Your rating was recorded in our system at %v", name, mynumrating, time.Now().Format(time.Stamp))

	if mynumrating == 5 {
		fmt.Println("O My God! A 5 star rating!! Thank you. I will continue to be the best")
	} else if mynumrating == 4 || mynumrating == 3 {
		fmt.Println("I am learning and improving!")
	} else if mynumrating < 3 {
		fmt.Println("I need to improve more. Please drop in your suggestions.")
	}

}
