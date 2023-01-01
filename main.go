package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	fmt.Printf("How many lessons do you have: ")
	var numOfLessons int
	fmt.Scan(&numOfLessons)

	sum := 0
	for i := 0; i < numOfLessons; i++ {
		fmt.Printf("%v . enter your grade: ", i)
		var not int
		fmt.Scan(&not)
		sum = sum + not
	}

	average := int(sum) / int(numOfLessons)
	fmt.Printf("Average: %v\n", average)
}
