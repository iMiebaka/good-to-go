package main

import (
	i "example/basic/internal"
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

// This is for variables and stuff
func variableTest() {
	i.LoggerFunT()
	var alpha string = "My string"
	fmt.Print(alpha, "\n")

	alpha = "This is the alpha string"
	fmt.Print(alpha, "\n")

	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	/*
		Arrays
	*/
	var arr1 = [3]int{1, 2, 3}
	// arr1 := [3]int{1,2,3} //This is the alternative
	// arr1 := [...]int{1,2,3} //This is a inferred length

	fmt.Println(arr1)
}

func structFun() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	fmt.Println(pers1.age + pers2.age)
	fmt.Println(pers1)
}

func operatorsFun() {
	/*
		Arithmetic operator - Reference: https://www.w3schools.com/go/go_arithmetic_operators.php
		Operator	Name			Description								Example
		+			Addition		Adds together two values				x + y
		-			Subtraction		Subtracts one value from another		x - y
		*			Multiplication	Multiplies two values					x * y
		/			Division		Divides one value by another			x / y
		%			Modulus			Returns the division remainder			x % y
		++			Increment		Increases the value of a variable by 1	x++
		--			Decrement		Decreases the value of a variable by 1	x--
	*/

	/*
		Assignment operator - Reference: https://www.w3schools.com/go/go_assignment_operators.php
		Operator	Example	    Same As
		=			x = 5   	x = 5
		+=			x += 3  	x = x + 3
		-=			x -= 3  	x = x - 3
		*=			x *= 3  	x = x * 3
		/=			x /= 3  	x = x / 3
		%=			x %= 3  	x = x % 3
		&=			x &= 3  	x = x & 3
		|=			x |= 3  	x = x | 3
		^=			x ^= 3  	x = x ^ 3
		>>=			x >>= 3  	x = x >> 3
		<<=			x <<= 3  	x = x << 3
	*/

	/*
		Comparison operator - Reference: https://www.w3schools.com/go/go_comparison_operators.php
		Operator	Name						Example
		==			Equal to					x == y
		!=			Not equal					x != y
		>			Greater than				x > y
		<			Less than					x < y
		>=			Greater than or equal to	x >= y
		<=			Less than or equal to		x <= y
	*/

	/*
		Logical operator - Reference: https://www.w3schools.com/go/go_logical_operators.php
		Operator	Name			Description													Example
		&& 			Logical and		Returns true if both statements are true					x < 5 &&  x < 10
		|| 			Logical or		Returns true if one of the statements is true				x < 5 || x < 4
		!			Logical not		Reverse the result, returns false if the result is true		!(x < 5 && x < 10)
	*/

}

func main() {
	operatorsFun()
	fmt.Print("************\n\n")
	variableTest()
	fmt.Print("************\n\n")
	structFun()
	fmt.Print("************\n\n")
}
