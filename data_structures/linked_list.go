package algorithmsunimi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

// It creates a new node
func newNode(val int) *listNode {
	return &listNode{val, nil}
}

// It adds a new node to the linked list
func (l *linkedList) push(val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

// It deletes a node of the linked list
func (l *linkedList) delete(val int) {
	p := l.head
	var previous *listNode
	for p != nil {
		if p.item == val {
			if previous == nil { //If the node to delete is the first one
				l.head = p.next
			} else {
				previous.next = p.next
			}
			return
		}
		previous = p
		p = p.next
	}
}

// It prints the nodes of the linked list
func (l *linkedList) print() {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ->")
		p = p.next
	}
	if p == nil {
		fmt.Print(p)
	}
	fmt.Println()
}

// It searches a node in a linkedlist
func (l *linkedList) search(val int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == val {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

// Counts the nodes of a linked list
func (l *linkedList) count() int {
	p := l.head
	counter := 0
	for p != nil {
		p = p.next
		counter++
	}
	return counter + 1
}

// It prints the nodes in the opposite order of their insertion
func (l *linkedList) reversePrint() {
	var elements []*listNode
	p := l.head
	for p != nil {
		p = p.next
		elements = append([]*listNode{p}, elements...) //prepending
	}
	if p == nil {
		fmt.Print(p)
	}
	for i := 0; i < len(elements); i++ {
		fmt.Print("<- ", elements[i].item)
	}
	fmt.Println()
}

func readUserInput() {
	var myLinkedList *linkedList
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		switch command[0] {
		case "+":
			numb, _ := strconv.Atoi(command[1])
			if exist, _ := myLinkedList.search(numb); !exist {
				myLinkedList.push(numb)
			}
		case "-":
			numb, _ := strconv.Atoi(command[1])
			if exist, _ := myLinkedList.search(numb); exist {
				myLinkedList.delete(numb)
			}
		case "?":
			numb, _ := strconv.Atoi(command[1])
			if exist, _ := myLinkedList.search(numb); !exist {
				fmt.Printf("%d does not belong to the set.\n", numb)
			}
		case "c":
			myLinkedList.count()
		case "p":
			myLinkedList.print()
		case "o":
			myLinkedList.reversePrint()
		case "d":
			myLinkedList.head = nil
		case "f":
			break
		}
	}
}
