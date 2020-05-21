package main

import (
	"fmt"
)

//User struct
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	//THIS IS RECOMMENDED
	Sam := User{"Sam", "sam@real.com", 18}
	fmt.Printf("%+v\n", Sam) // %+v for extra info
	fmt.Printf("%v\n", Sam.Email)

	//THIS IS ALSO RECOMMENDED
	var Jacob = new(User) //refrence
	Jacob.Name = "Jacob"
	Jacob.Email = "jacob@real.com"
	Jacob.Age = 19
	fmt.Printf("%+v\n", Jacob) // will have an '&' as prefix

	var Tom = &User{"Tom", "tom@real.com", 20}
	fmt.Printf("%+v\n", Tom)

}
