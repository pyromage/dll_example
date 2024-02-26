package dbl_list

import (
	"fmt"
	"math/rand"
	"testing"
)

type testNode struct {
		num int
		str string
		b bool
	}

const maxArrSize = 10

type testLists struct {
	intList List[int]
	strList List[string]
	arrList List[[maxArrSize]int]
	stcList List[testNode]
	
}

// Good practice is unit tests should not use any other methods from the package, so
// create a test set of lists we will use for unit testing
func createUnitTestStaticList(sizeStaticList, max int) *testLists {
	var lists testLists

	if sizeStaticList > max && max != 0 {
		return nil
	} 
	for i := 0 ; i < sizeStaticList; i++ {

		blob_int := i
		blob_str := string(fmt.Sprintf("%d",i))
		blob_node := testNode{num: i, str: blob_str, b: true}
		blob_arr := [maxArrSize]int{}
		for j := 0 ; j < maxArrSize; j++ {
			blob_arr[j] = i*1000 + j
		}

		if i == 0 {
			lists.intList.head = &node[int]{next: nil, prev: nil, blob: &blob_int}
			lists.intList.tail = lists.intList.head
			lists.intList.length = 1
			lists.intList.max = max

			lists.strList.head = &node[string]{next: nil, prev: nil, blob: &blob_str}
			lists.strList.tail = lists.strList.head
			lists.strList.length = 1
			lists.strList.max = max

			lists.stcList.head = &node[testNode]{next: nil, prev: nil, blob: &blob_node}
			lists.stcList.tail = lists.stcList.head
			lists.stcList.length = 1
			lists.stcList.max = max

			lists.arrList.head = &node[[maxArrSize]int]{next: nil, prev: nil, blob: &blob_arr}
			lists.arrList.tail = lists.arrList.head
			lists.arrList.length = 1
			lists.arrList.max = max
		} else  {
			lists.intList.tail.next = &node[int]{next: nil, prev: lists.intList.tail, blob: &blob_int}
			lists.intList.tail = lists.intList.tail.next
			lists.intList.length += 1

			lists.strList.tail.next = &node[string]{next: nil, prev: lists.strList.tail, blob: &blob_str}
			lists.strList.tail = lists.strList.tail.next
			lists.strList.length += 1

			lists.arrList.tail.next = &node[[maxArrSize]int]{next: nil, prev: lists.arrList.tail, blob: &blob_arr}
			lists.arrList.tail = lists.arrList.tail.next
			lists.arrList.length += 1

			lists.stcList.tail.next = &node[testNode]{next: nil, prev: lists.stcList.tail, blob: &blob_node}
			lists.stcList.tail = lists.stcList.tail.next
			lists.stcList.length += 1
		}
	}

	if !lists.intList.isValid() || !lists.strList.isValid() || !lists.arrList.isValid() || !lists.stcList.isValid() {
		fmt.Printf("Did not create valid static list for testing size: %d max: %d\n",sizeStaticList,max)
		fmt.Printf("intList %v\n",lists.intList)
		fmt.Printf("strList %v\n",lists.strList)
		fmt.Printf("stcList %v\n",lists.stcList)
		fmt.Printf("arrList %v\n",lists.arrList)
		return nil
	}

	return &lists
}

// Verifies the integrity of any list by traversing it both ways, checking pointers etc.
func (l *List[T])isValid()bool{
	switch { 
	case l.head == nil && l.tail == nil :
		return true
	case l.head != nil && l.tail == nil :
		return false
	case l.head == nil && l.tail != nil :
		return false
	}

	// list is not empty, need to check pointers, we know l.head != nil && l.tail != nil 
	if l.head.prev != nil || l.tail.next != nil {
		// head is first node and tail is the last node
		return false
	}

	if l.head == l.tail && ( l.head.next != nil || l.head.prev != nil || l.tail.next !=nil || l.tail.prev != nil ){
		// single node, but head or tail point to other nodes
		return false
	}

	// now we will see if the head and tail are the real head and tail by traversing the list
	var tempHead *node[T]
	var tempTail *node[T]

	// traverse the list in both directions to confirm head and tail are corret
	// go to the tail first from the head
	for p := l.head ; p != nil ; p = p.next {
		tempTail = p
	}

	if tempTail != l.tail { 
		return false
	}

	// go to the head next from the tail
	for p := l.tail ; p != nil ; p = p.prev {
		tempHead = p
	}

	// small optimization
	return tempHead == l.head
}


func TestNew (t *testing.T) {

	max := 0
	list := New[int](max)
	
	if list.head != nil || list.tail != nil || list.length != 0 || list.max != max {
		t.Errorf("New of max %d failed", max)
	}

	max = -1
	list = New[int](max)
	
	if list.head != nil || list.tail != nil || list.length != 0 || list.max != 0 {
		t.Errorf("New of max %d failed", max)
	}

	max = 1
	list = New[int](1)
	
	if list.head != nil || list.tail != nil || list.length != 0 || list.max != max {
		t.Errorf("New of max %d failed", max)
	}

}

func TestAppend(t *testing.T) {
	var list List[int] 	// note max size will be 0 (infinite)

	if !list.isValid() {
		t.Errorf("list initialization failed %v", list)
		return 
	}

	for i := 0 ; i < 5 ; i++ {
		if !list.Append(i) {
			t.Errorf("Append(%d) failed %v",i,list)
			return
		}

		if !list.isValid() ||  *list.head.blob != 0 || *list.tail.blob != i {
			t.Errorf("Append(%d) list not correct: %v",i,list)
			return
		}
	}

	list2 := New[int](4)

	if !list2.isValid() {
		t.Errorf("list2 initialization failed %v", list2)
		return 
	}

	for i := 0 ; i < 4 ; i++ {
		if !list2.Append(i) || !list2.isValid() ||  *list2.head.blob != 0 || *list2.tail.blob != i {
			t.Errorf("Append(%d) failed %v",i,list2)
			return
		}
	}

	if list2.Append(5){
		t.Errorf("Append beyond max of size 4 should fail %v",list2)
	}
}

func TestPrepend(t *testing.T){
	var list List[string] 	// note max size will be 0 (infinite)

	if !list.isValid() {
		t.Errorf("list initialization failed %v", list)
		return 
	}

	for _,i := range []string{"1","2","3","4","5"} {
		if !list.Prepend(i) {
			t.Errorf("Prepend(%v) failed %v",i,list)
			return
		}
		if !list.isValid() ||  *list.head.blob != i || *list.tail.blob != "1" {
			t.Errorf("Prepend(%v) list not correct: %v",i,list)
			return
		}
	}

	list2 := New[string](4)

	if !list2.isValid() {
		t.Errorf("list2 initialization failed %v", list)
		return 
	}


	for _,i := range []string{"1","2","3","4"} {
		if !list2.Prepend(i) || !list2.isValid() ||  *list2.head.blob != i || *list2.tail.blob != "1" {
			t.Errorf("Prepend(%v) failed %v",i,list2)
			return
		}
	}

	if list2.Prepend("5"){
		t.Errorf("Prepend beyond max of size 4 should fail %v",list2)
	}
}

func TestLength(t *testing.T) {
//	type testNode [5]int
	var list List[testNode]

	// Length of the empty list
	n := list.Length()

	if n != 0 {
		t.Errorf("Length of empty created by var failed, expected 0 got %d", n)
		return
	}

	// See if new puts in the right length
	list2 := New[int](5)
	if list2 == nil {
		t.Error("creating list2 using new failed")
		return
	}
	
	n = list2.Length()
	if n != 0 {
		t.Errorf("Length of empty created by New failed, expected 0 got %d", n)
		return
	}
}

func TestPopHead(t *testing.T){
	var tmp List[int]

	if tmp.PopHead() != nil {
		t.Error("pop head in empty list failed")
	}

	// Test a single element list
	lists := createUnitTestStaticList(1,1)
	
	if lists == nil {
		t.Error("failed to create a static test set of lists")
		return
	}

	ele := lists.strList.PopHead()

	if ele == nil {
		t.Errorf("pop of single element list failed, got ele: %v",*ele)	
		return
	} 
	
	if *ele != "0" {
		t.Errorf("popped the wrong element, expected %s, got %v","1", *ele)
	}

	l := lists.strList.Length()

	if l != 0 {
		t.Errorf("Pop of single element list should have length 0, got %d",l)
	}

	// Test multiple element list, rebuild it
	const sizeList = 5
	lists = createUnitTestStaticList(sizeList,10)

	if lists == nil {
		t.Errorf("failed to create a static test set of lists of size %d",sizeList)
		return
	}

	// Pop all the elements
	for i := 0 ; i < sizeList ; i++ {	
		ele := lists.intList.PopHead()

		if ele == nil {
			t.Errorf("error popping element %d",i)
			return
		}

		if *ele != i {
			t.Errorf("popped the wrong element, expected %v, got %v",i, *ele)
			return
		}
	}

	l = lists.intList.Length()

	if l != 0 {
		t.Errorf("Pop of all elements should be empty but got length %d",l)
	}
}

func TestPopTail(t *testing.T){
	var tmp List[int]

	if tmp.PopTail() != nil {
		t.Error("pop tail in empty list failed")
	}

	// Test a single element list
	lists := createUnitTestStaticList(1,1)
	
	if lists == nil {
		t.Error("failed to create a static test set of lists")
		return
	}

	ele := lists.strList.PopTail()

	if ele == nil {
		t.Errorf("pop of single element list failed, got ele: %v",*ele)	
		return
	} 
	
	if *ele != "0" {
		t.Errorf("popped the wrong element, expected %s, got %v","1", *ele)
	}

	l := lists.strList.Length()

	if l != 0 {
		t.Errorf("Pop of single element list should have length 0, got %d",l)
	}

	// Test multiple element list, rebuild it
	const sizeList = 5
	lists = createUnitTestStaticList(sizeList,10)

	// Pop all the elements
	for i := sizeList-1 ; i >= 0 ; i-- {	
		ele := lists.intList.PopTail()

		if ele == nil {
			t.Errorf("error popping element %d",i)
			return
		}

		if *ele != i {
			t.Errorf("popped the wrong element, expected %v, got %v",i, *ele)
			return
		}
	}

	l = lists.intList.Length()

	if l != 0 {
		t.Errorf("Pop of all elements should be empty but got length %d",l)
	}
}

func TestGetIndex(t *testing.T) {
	var tmp List[[5]int]

	// test the empty list
	if tmp.GetIndex(0) != nil {
		t.Error("getting node idx for empty list failed")
	}

	lists := createUnitTestStaticList(1,1)

	// try to get the value of the element
	ele := lists.intList.GetIndex(0)

	if ele == nil {
		t.Error("could not get the element in a single element list")
	} else if *ele != 0 {
		t.Errorf("Get of single element list incorrect for idx %d, expected %v, got %v",0,1, ele)
	}

	// try to go past the end of the one element
	ele = lists.intList.GetIndex(1)

	if ele != nil {
		t.Error("getting idx 1 from list of only one element should fail")
	}

	// Create a list of multiple elements
	lists = createUnitTestStaticList(5,10)

	// check the indexes
	for i := 0; i < lists.intList.Length(); i++ {
		ele = lists.intList.GetIndex(i)

		if ele == nil {
			t.Errorf("GetValueByIndex[%d] failed",i)
		} else if *ele != i {
			t.Errorf("Get by index incorrect for idx %d, expected %v, got %v",i,i, *ele)
		}
	}

	// Negative test : idx is -1 and idx is 1 past the end
	ele = lists.intList.GetIndex(-1)

	if ele != nil {
		t.Error("GetByIdx -1 should always fail")
	}

	// indexed by 0, so length will be 1 past the end
	ele = lists.intList.GetIndex(5)

	if ele != nil {
		t.Error("GetByIdx past the end should always fail")
	}
}

func TestPrint(t *testing.T) {
	
	lists := createUnitTestStaticList(5,10)

	lists.intList.Print(false)
	lists.strList.Print(false)
	lists.stcList.Print(false)
	lists.arrList.Print(false)

	lists.intList.Print(true)
	lists.strList.Print(true)
	lists.stcList.Print(true)
	lists.arrList.Print(true)

}

func BenchmarkPushPop(b *testing.B) {
	var arrValues [maxArrSize]int

	lists := createUnitTestStaticList(1,1000)

	// bunch of appends/prepends
	for i := 0 ; i < b.N ; i ++{
		lists.arrList.Append(arrValues)
		lists.intList.Append(i)
		lists.strList.Append(fmt.Sprint(1000*i))
		lists.stcList.Append(testNode{i,fmt.Sprint(i),i%2==0})
		lists.arrList.Prepend(arrValues)
		lists.intList.Prepend(i)
		lists.strList.Prepend(fmt.Sprint(1000*i))
		lists.stcList.Prepend(testNode{i,fmt.Sprint(i),i%2==0})
	}

	// bunch of pops
	for i := 0 ; i < b.N ; i ++{
		lists.arrList.PopHead()
		lists.intList.PopHead()
		lists.strList.PopHead()
		lists.stcList.PopHead()
		lists.arrList.PopTail()
		lists.intList.PopTail()
		lists.strList.PopTail()
		lists.stcList.PopTail()
	}
}

func BenchmarkAppend(b *testing.B){
	var list List[int]

	for i := 0 ; i < b.N ; i++ {
		list.Append(i)
	}
}

func BenchmarkGetIndex( b *testing.B) {

	list := createUnitTestStaticList(100,200)

	for i := 0 ; i < b.N ; i++ {
		r := rand.Intn(100)
		list.intList.GetIndex(r)
	} 

}