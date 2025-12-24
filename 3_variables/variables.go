package main

import "fmt"

func main() {
	
	//Strict type declaration
		//var name string = "Vivek"
		//var age int = 21
		var isStudent bool = true
		var height float64 = 5.9


	//Automatic type inference
		var name = "Vivek"
		var isEmployed = false
		var age = 21

	//Short variable declaration
		//name := "Vivek"
		//age := 21
		//isStudent := true
		//height := 5.9
	
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Height:", height)
	fmt.Println("Is Employed:", isEmployed)

}