package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello world")
	var myName string = "john Doe"
	const mySecondName string = "john Doe"

	myThirdName := "Bob Doe"
	fmt.Println(myName)
	fmt.Println(mySecondName)
	fmt.Println(myThirdName)
}
