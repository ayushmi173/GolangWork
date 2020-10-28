package main

import (
	"fmt"
)

//Node is a struct
type Node struct {
	value interface{}
	next  *Node
}

//It is for storing many nodes...
type Linkedlist struct {
	head *Node
	len  int
}

//Insert element in Golang
func (l *Linkedlist) Insert(val int) {

	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head

	//insert in end
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

//Searching th element in ll
func (l *Linkedlist) Search(val int) {
	length := 0
	ptr := l.head
	if l.len == 0 {
		fmt.Println("Linkedlist has no nodes")

	}
	for i := 0; i < l.len; i++ {
		if i+1 == val {
			fmt.Println(val," position has ",ptr.value)
			length++
			return
		}
		ptr = ptr.next
	}
	if length == 0 {
		fmt.Println("Searching element is not present in the linkedlist")
	}
}

//For printing the  linkedlist
func (l *Linkedlist) Print() {

	if l.len == 0 {
		fmt.Println("Linkedlist has No nodes")
		return
	}

	ptr := l.head
	for ptr != nil {
		fmt.Print(ptr.value, "-->")
		ptr = ptr.next
	}
	fmt.Println("Null")
}

//Edit the element by position
func (l *Linkedlist) Edit(position int) {
	ptr := l.head

	if l.len < position {
		fmt.Println("\nEnter Right Position")
		return
	}

	for i := 1; i <= position; i++ {
		if i == position {
			var value int
			fmt.Println("\nEnter the new value that you want to put : ")
			fmt.Scanln(&value)
			ptr.value = value
			return
		}
		ptr = ptr.next
	}
}

// remove last element
func (l *Linkedlist) Remove() {

	var ptr *Node
	ptr = l.head
	if ptr == nil || ptr.next == nil {
		l.head =nil
		return
	}
	if l.len >= 2 {
		for ptr.next.next != nil && ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = nil
	}
	ptr.next = nil
	l.len--

}

func main() {

	var l Linkedlist

	for {
		fmt.Println("\n\nHappy Coding With Linked List")
		fmt.Println(" 1: Push element in the end \n 2: Search element by position\n 3: Remove the last element \n 4: Edit element by position\n 5: get the length of list \n 6: Print List \n 7: Exit")

		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("how many elements do you want to add? Enter Please...")
			var iteration int
			fmt.Scanln(&iteration)
			for i := 0; i < iteration; i++ {
				fmt.Println("\nEnter Value : ", i+1)
				var value int
				fmt.Scanln(&value)
				l.Insert(value)
			}
		case 2:
			fmt.Println("\nEnter the any position & get element of that position ")
			var value int
			fmt.Scanln(&value)
			l.Search(value)

		case 3:
			l.Remove()

		case 4:

			fmt.Println("\nEnter the any position for editing the value")
			var value int
			fmt.Scanln(&value)
			l.Edit(value)
		case 5:
			fmt.Println("\nLength is: ",l.len)
		case 6:
			l.Print()
		case 7:
			break
		
	default:
			fmt.Println("\n Please Enter valid number")
	}
		if input == 7 {
			break
		}
	}
}
