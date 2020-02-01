package main

import (
	"fmt"
	"leetcodetop100/algirithm"
)

func main() {
	node := algirithm.ConstructorA()
	node.Insert("apple")
	fmt.Println(node.Search("apple"))
	fmt.Println(node.Search("app"))
	fmt.Println(node.StartsWith("app"))
	node.Insert("app")
	fmt.Println(node.Search("app"))
}
