package main

import "fmt"

func main() {
	var description string = "One of the best coders in the world."
	fmt.Println(description)

	length := 384 //AutoGuess the data-type of the variable
	println(length)
	fmt.Printf("%v, %T \n", length, length) //%v == value %T == type

	var username, password string = "NameIUse", "hackMeBeby!101"
	fmt.Println(username, password)

	var (
		email = "nameiuse@madafaka.com"
		phone = "9090909090"
		bonus = 9543.44
	)

	println(email)
	println(phone)
	println(bonus)

}
