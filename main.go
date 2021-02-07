//Use recover to regain control of a panicking program (like try - catch)
package main

import "fmt"

var food = [4]string{"pizza", "ice cream", "hamburger", "tacos"}
var ban bool

func checkRecover() {
	var r = recover() //Takes message of panic
	if r != nil {
		fmt.Println("Recovered:", r)
	}
}

func checkData(num int) {
	if num >= len(food) {
		panic("num is out of bound")
	}
	fmt.Printf("%s %d: %s\n", "Element", num, food[num])
}

func enterData() {
	defer checkRecover()
	var num int
	fmt.Println("Enter element of array")
	fmt.Scan(&num)
	checkData(num)
	ban = true //Ends func enterData
}

func main() {
	enterData()
	if ban == true {
		fmt.Println("Ends func enterData and main")
	} else {
		fmt.Println("Ends func main")
	}
}
