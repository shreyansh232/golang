package main

var url = "https://www.google.com" // package level variable or global variable

func main(){
	//function scope variables
	var message string = "Hello, world!" // when we declare a variable, we need to specify the type, but when we initialize it, we can omit the type
	var price = 24 // type inference will automatically assign the type of the value
	title := "Golang Programming" // without using var keyword we can declare and initialize a variable using :=
	const maxSpeed int = 100 // constant variable
	print(message, price, title, maxSpeed, url)
	PrintValue() 
}