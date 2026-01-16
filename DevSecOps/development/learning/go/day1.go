package main

import "fmt"

func main(){

	var name string
	
	fmt.Println("Enter your Name: ")
	fmt.Scanln(&name)

	fmt.Println("Hello",name)
}