package main

import (
	"fmt"

	"github.com/pyromage/dll_example/dbl_list"
)

// The basic operations
func main(){

	list := dbl_list.New[string](10)

	ok := list.PushTail("Tail")

	if !ok {
		fmt.Println("Failed to PushTail")
	}

	ok = list.PushHead("Head")

	if !ok {
		fmt.Println("Failed to PushHead")
	}
 
	// Print the full debug output, should be the full list
	list.Print(true)


	ele := list.Seek(1)
	
	if ele == nil {
		fmt.Println("Failed to get by index")
	} else if *ele != "Tail" {
		fmt.Println("Get returned wrong value")
	}

	l := list.Length()

	if l != 2 {
		fmt.Printf("List lenght is incorrect, should be 2 got %d", l)
	}

	ele = list.PopHead()

	if ele == nil {
		fmt.Println("Failed to pop the head node")
	} else if *ele != "Head" {
		fmt.Printf("PopHead returned incorrect blob, expected: %s got %s ", "Head", *ele)
	}

	ele = list.PopTail()

	if ele == nil {
		fmt.Println("Failed to pop the tail node")
	} else if *ele != "Tail" {
		fmt.Printf("PopTail returned incorrect blob, expected: %s got %s ", "Tail", *ele)
	}

	// Print the full debug output, should be the empty list
	list.Print(true)

}