package main

import "fmt"

func main(){
	var firstname string
	fmt.Printf("Enter your first Name: ")
	fmt.Scan(&firstname) // Getting input
	var lastname string
	fmt.Printf("Enter your last Name: ")
	fmt.Scan(&lastname) // & is refered here a pointer
	var age int
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age) // In Go pointer is special type of varibale

	fmt.Printf("Hey! you're %v %v and you're %v years old",firstname,lastname,age)

}