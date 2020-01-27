package main

import (
	"container/heap"
	"fmt"
)

func main() {
	slice := []uint8 {'(',')','[',']','{','}'}
	//map1 := map[uint8]uint8 {'(':')',')':'(',
	//						'[':']',']':'[',
	//						'{':'}','}':'{'}
	for _, value := range slice {
		fmt.Printf("char : %c Value : %d\n", value, value)
	}
	isValid("()[]{}")
}

func isValid(s string) bool {
	chars := new(heap.Interface)
	for _, value := range s {

	}
	fmt.Printf("%v", chars)
	return false
}
