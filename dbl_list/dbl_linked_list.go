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

// maxLen of <= 0 means bounded only by memory
func New[T comparable](maxLen int) *List[T] {
	if maxLen <= 0 {
		maxLen = 0
	}

  	return &List[T]{
		head: nil,
		tail: nil,
		length: 0,
		max: maxLen,
  	}
}

func (l *List[T])Append(v T) bool {
	ele := node[T]{blob: &v, next: nil, prev: nil}

	if l.max > 0 && l.length >= l.max {
		return false
	}

	if l.length == 0 {
		l.head = &ele
		l.tail = &ele
	} else {
		ele.prev = l.tail
		l.tail.next = &ele
		l.tail = &ele
	}

	l.length++

	return true
}

func (l *List[T])Prepend(v T) bool {
	ele := node[T]{blob: &v, next: nil, prev: nil }

	if l.max > 0 && l.length >= l.max {
		return false
	}

	if l.length == 0 {
		l.head = &ele
		l.tail = &ele
	} else {
		ele.next = l.head
		l.head.prev = &ele
		l.head = &ele
	}

	l.length++

	return true
}

func (l *List[T])Length() int {
	return l.length
}

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

// return the value of the idx'th node, false will indicate not found and return an empty value
// serach from head or tail depending if idx , 1/2 of length
func (l *List[T])GetIndex(idx int) *T {

	if idx < 0 || idx >= l.length || l.length == 0 {
		return nil
	}

	var p *node[T]

	if idx < (l.length / 2) {
		// search from head, its shorter
		for p = l.head ; p != nil ; p = p.next {
			idx -= 1
			if idx < 0 {
				break
			}
		}
	} else {
		idx = l.length - idx - 1
		// search from tail, its shorter
		for p = l.tail ; p != nil ; p = p.prev {
			idx -= 1
			if idx < 0 {
				break
			}
		}
	}

	return p.blob
}

// traverse and print the list to the tail from the node at idx, indexed by 0
func (l *List[T])Print(debug bool) {
	fmt.Print("--Double Linked List--\n")
	if debug {
		fmt.Printf("  TypeOf %v Object %v\n",reflect.TypeOf(l),l)
	}
	fmt.Printf("  Head %v Tail %v Length %v\n", l.head, l.tail, l.Length())

	idx := 0
	for p := l.head ; p != nil ; p = p.next {
		if debug {
			fmt.Printf("    Idx %v this %v prev %v next %v blob %v\n", idx, p, p.prev, p.next, p.blob)
		} else {
			fmt.Printf("    Idx %v blob %v\n", idx, p.blob)
		}
		idx++
	}

}