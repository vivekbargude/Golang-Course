package main 

import "fmt"
//Package-level constant declaration
const name = "GolangApp"


func main() {
	
	//Constant declarations that cannot be changed
	const pi = 3.14
	const appName = "GolangApp"
	const isActive = true
	const maxUsers = 1000

	//Mutliple constant declarations
	const (
		version = "1.0.0"
		author  = "Vivek"
		license = "MIT"
	)

	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println("App Name:", appName)
	fmt.Println("Pi Value:", pi)
	fmt.Println("Is Active:", isActive)
	fmt.Println("Max Users:", maxUsers)
	fmt.Println("Package-level Constant Name:", name)

	fmt.Println("Version:", version)
	fmt.Println("Author:", author)
	fmt.Println("License:", license)

	fmt.Println("Host:", host)
	fmt.Println("Port:", port)

}