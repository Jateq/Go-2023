package main

import "fmt"

func main() {
	// var var_type *var-type
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("Adress of a variable: %x\n", &a)
	fmt.Printf("Adress stored in ip variable: %x\n", ip)
	fmt.Printf("Value of *ip variable: %d\n", *ip)
	
}