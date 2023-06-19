package main

import "fmt"

func main (){

	/*
	Write a Go program that takes a string as input and counts the number of vowels (a, e, i, o, u) present in the string. Ignore case sensitivity.

	For example, if the input string is "Hello World", the program should output 3 since there are three vowels ('e', 'o', 'o') in the string.

	You can take user input from the command line or hardcode a string inside the program.

	Take your time to solve the problem, and feel free to ask for hints or clarification if needed!
	*/

	fmt.Println("Enter some verbiage for vowels to be found")

	var input string 
	fmt.Scanln(&input)
	
	fmt.Println(GetVowels(input))




}
