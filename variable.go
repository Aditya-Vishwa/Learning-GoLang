package main

import "fmt"

func main(){
	var v1 = "this is the first variable"
	const c1 = "this is the first constant"
	fmt.Println("I am creating a variable and",v1,"creation in go.")
	fmt.Println(v1)
	fmt.Println("I am creating a constant and",c1,"creation in go.")
	fmt.Println(c1)
}