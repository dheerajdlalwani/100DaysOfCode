package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Test")
	var thingsIWantToLearn = []string{"Django REST API Framework", "Vue.js", "Golang"}
	fmt.Println(thingsIWantToLearn)
	thingsIWantToLearn = append(thingsIWantToLearn, "Hardware")
	fmt.Println(thingsIWantToLearn)

	thingsIWantToLearn = append(thingsIWantToLearn[2:])
	fmt.Println(thingsIWantToLearn)

	//thingsIWantToLearn = append(thingsIWantToLearn[1:-1]) //negative index not allowed

	languages := make([]string, 3, 3)
	languages[0] = "python"
	languages[1] = "golang"
	languages[2] = "java"
	fmt.Println(languages)
	languages = append(languages, "assembly")
	languages = append(languages, "brainfuck")
	fmt.Println(cap(languages))

	marks := []int{50, 47, 35, 33, 50, 49}
	isSorted := sort.IntsAreSorted(marks)
	fmt.Println("Are marks sorted: ", isSorted)

	sort.Ints(marks)
	isSortedNow := sort.IntsAreSorted(marks)
	fmt.Println("Are marks sorted: ", isSortedNow)

}
