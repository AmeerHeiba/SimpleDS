package linkedlists

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) AddNode(v string) {
	newHead := &Node{
		Value: v,
		Next:  l.Head,
	}
	l.Head = newHead
	l.Length++
}

func (l *LinkedList) DeleteNode(v string) {
	if l.Length == 0 {
		return
	}
	if l.Length == 1 && l.Head.Value == v {
		l.Head = nil
		return
	}

	if l.Head.Value == v {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	currentNode := l.Head
	counter := l.Length

	for counter != 0 {
		if currentNode.Next.Value == v {
			currentNode.Next = currentNode.Next.Next
			l.Length--
			return
		}
		currentNode = currentNode.Next
		counter--
	}
	fmt.Print("Not Found")
}

func (l LinkedList) PrintList() {
	if l.Length == 0 {
		return
	}
	currentNode := l.Head
	count := l.Length
	for count > 0 && currentNode != nil {
		fmt.Printf("%q ", currentNode.Value)
		currentNode = currentNode.Next
		count--
	}
	fmt.Println("")
}

func (l *LinkedList) SearchList(v string) *Node {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == v {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}
