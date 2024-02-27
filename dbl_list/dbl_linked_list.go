package dbl_list

import (
	"fmt"
	"reflect"
)

type List[T comparable] struct {
	head, tail *node[T]
	length int
	max int
}

type node[T comparable] struct {
	next *node[T]
	prev *node[T]
	blob *T
}

// Creates a new double linked list with a max length
// maxLen of <= 0 means no maximum
func New[T comparable](maxLen int) *List[T] {
	if maxLen <= 0 {
		maxLen = 0
	}

  	return &List[T]{
		max: maxLen,
  	}
}

// Add the value to the tail of the list
// return false if list is full
func (l *List[T])PushTail(v T) bool {
	if l.max > 0 && l.length >= l.max {
		return false
	}

	newNode := &node[T]{blob: &v, prev: l.tail}

	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}

	l.tail = newNode
	l.length++

	return true
}

// Add the value to the head of the list
// return false if list is full
func (l *List[T])PushHead(v T) bool {
	if l.max > 0 && l.length >= l.max {
		return false
	}

	newNode := &node[T]{blob: &v, next: l.head }

	if l.head == nil {
		l.tail = newNode
	} else {
		l.head.prev = newNode
	}

	l.head = newNode
	l.length++

	return true
}

// The lenght of the list
func (l *List[T])Length() int {
	return l.length
}

// Remove the value from the head of the list
func (l* List[T])PopHead() *T {
	var p *node[T]

	if l.length <= 0 {
		return nil
	}

	p = l.head

	// single node
	if l.length == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}

	l.length--

	return p.blob
}

// Remove the value from the tail of the list
func (l* List[T])PopTail() *T {
	var p *node[T]

	if l.length <= 0 {
		return nil
	}

	p = l.tail

	// single node
	if l.length == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	l.length--

	return p.blob
}

// return the value of the idx'th node but do not remove it
// return of nil will indicate not found or idx out of range
// search from head or tail depending if which is closest to idx
func (l *List[T])Seek(idx int) *T {
	if idx < 0 || idx >= l.length || l.length == 0 {
		return nil
	}

	var p *node[T]

	// seek from head or tail, whichever is shorter
	if idx < (l.length / 2) {
		for p = l.head ; p != nil ; p = p.next {
			idx -= 1
			if idx < 0 {
				break
			}
		}
	} else {
		idx = l.length - idx - 1
		for p = l.tail ; p != nil ; p = p.prev {
			idx -= 1
			if idx < 0 {
				break
			}
		}
	}

	return p.blob
}

// traverse and print the list head to tail to stdout for inspection
// debug true will print out detailed information object information
// Note if blob is large this may not be desirable
func (l *List[T])Print(debug bool) {
	fmt.Print("--Double Linked List--\n")
	if debug {
		fmt.Printf("  TypeOf %v Object %v\n",reflect.TypeOf(l),l)
		fmt.Printf("  Length: %v of %v Head: %v Tail: %v\n", l.length, l.max, l.head, l.tail)
	} else {	
		fmt.Printf("  Length: %v of %v Head: %p Tail: %p\n", l.length, l.max, l.head, l.tail)
	}

	idx := 0
	for p := l.head ; p != nil ; p = p.next {
		if debug {
			fmt.Printf("    Idx: %v node: %v next: %v prev: %v blob: %p \n", idx, p, p.next, p.prev, p.blob)
		} else {
			fmt.Printf("    Idx: %v node: %p blob: %p\n", idx,p, p.blob)
		}

		if (idx >= l.max) {
			// structure is corrupted, not good
			// void the potential overflow or infinite loop
			fmt.Println("ERROR: list corrupted, goes past maximum length")
			return
		}
		idx++
	}
	
	fmt.Println("--End--")
}