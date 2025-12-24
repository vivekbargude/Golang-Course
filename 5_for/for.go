package main 

import "fmt"

//for -> only loop construct in Go

func main() {
	
	//while loop
	count := 1
	for count <= 5 {
		fmt.Println("Count:", count)
		count++
	}

	//Infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	//for loop with init and post statements
	for i := 1; i <= 5; i++ {
		//break //to exit the loop
		if i == 4 {
			break //exit loop when i is 4
		}
		//continue //to skip the current iteration
		if i == 3 {
			continue //skip iteration when i is 3
		}
		fmt.Println("Iteration:", i)
	}

	//Range-based for loop

	for i := range 3 {
		fmt.Println("Index:", i)
	}

	numbers := []int{10, 20, 30, 40, 50}	
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}