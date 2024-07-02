package data

import "fmt"

var Countries [10]string // array is a value type. It is a fixed-size collection of items of the same type
var Slice []string // slice is a reference type, it is a wrapper around the array. It is a dynamic collection of items of the same type
var Codes map[int]string // map[keyType]valueType
func init() { // init function will be called before the main function
	Countries = [10]string{"India", "USA", "Japan", "Germany", "France", "Australia", "Canada", "Russia", "China", "Brazil"}

	length := len(Countries) // no objects in Go, so we can't use the dot operator to access the properties or methods of an object
	Slice := append(Slice, "Spain", "Italy")
	wellKnownPorts := map[int]string{80: "http", 443: "https", 21: "ftp", 22: "ssh"} // map literal

	fmt.Print(Countries, length, Slice, wellKnownPorts)

}

