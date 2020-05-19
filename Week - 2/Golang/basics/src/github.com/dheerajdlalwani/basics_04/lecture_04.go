package main

import (
	"fmt"
)

func main() {

	var balance float64 = 9567.9453 //declaring normal float variable
	balanceRef := &balance          //creating variable a reference to our previous variable

	fmt.Println(balance)     //printing the contents of the variable normally
	fmt.Println(balanceRef)  //printing the address of the variable
	fmt.Println(*balanceRef) //printing the contents of the location pointed by the pointer

	var zahlen [5]string //initializing the array
	zahlen[0] = "Null"
	zahlen[1] = "Eins"
	zahlen[2] = "Zwei" //declaring the elements of the array
	zahlen[3] = "Drei"
	zahlen[4] = "Vier"
	fmt.Println(zahlen)

	var farben = [4]string{"rot", "balu", "grün", "gelb"}
	fmt.Println(farben)
	fmt.Println(len(farben))
	fmt.Println(farben[2])

}
