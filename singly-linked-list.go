/**
Implementation of singly linked list
cmd: go run .\singly-linked-list.go
*/
package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	ll := createLinkedList()

	for i := 1; i <= 5; i++ {
		ll.insertNextNode(&Node{Data: i, Next: nil})
	}

	ll.insertStartNode(&Node{Data: "New node at start", Next: nil})
	ll.insertNodeAtPos(&Node{Data: "New node at Pos 3", Next: nil}, 3)

	ll.removeLastNode()
	ll.removeStartNode()
	ll.removeNodeAtPos(2)

	ll.print()
}

func createLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func (ll *LinkedList) print() {
	fmt.Println("##### Printing linked list #####")
	curr := ll.Head
	if curr != nil {
		if curr.Next == nil {
			fmt.Println(curr.Data) // Print when only 1 node
		} else {
			fmt.Println(curr.Data)
			for curr.Next != nil {
				curr = curr.Next
				fmt.Println(curr.Data)
			}
		}
	} else {
		fmt.Println("List is empty")
	}
}

func (ll *LinkedList) insertNextNode(node *Node) {
	fmt.Println("Inserting node with data", node.Data)
	curr := ll.Head
	if curr != nil {
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node // add node to existing list
	} else {
		ll.Head = node // add node to empty list
	}
}

func (ll *LinkedList) insertStartNode(node *Node) {
	fmt.Println("Inserting node at start with data", node.Data)
	node.Next = ll.Head
	ll.Head = node
}

func (ll *LinkedList) insertNodeAtPos(node *Node, pos int) {
	if ll.Head != nil {
		if pos < 1 {
			fmt.Println("Insert fail, Pos must be >= 1")
		}
		if pos == 1 {
			ll.insertStartNode(node)
			return
		}
		curr := ll.Head
		count := 1
		for count < pos-1 && curr != nil {
			count++
			curr = curr.Next
		}
		if curr != nil {
			fmt.Println("Inserting node with data", node.Data, "at pos", pos)
			node.Next = curr.Next
			curr.Next = node
		} else {
			fmt.Println("Insert fail, pos out of bound")
		}
	} else {
		fmt.Println("Insert fail, list is empty")
	}
}

func (ll *LinkedList) removeLastNode() {
	if ll.Head != nil {
		curr := ll.Head
		prev := ll.Head
		if curr.Next == nil { // last node left in the list
			ll.Head = nil
		}
		for curr.Next != nil {
			prev = curr      // n-1 node
			curr = curr.Next // nth node
		}
		fmt.Println("Removing last node with data", curr.Data)
		prev.Next = nil

	} else {
		fmt.Println("Remove fail, list is empty")
	}
}

func (ll *LinkedList) removeStartNode() {
	if ll.Head != nil {
		fmt.Println("Removing start node with data", ll.Head.Data)
		ll.Head = ll.Head.Next
	} else {
		fmt.Println("Remove fail, list is empty")
	}
}

func (ll *LinkedList) removeNodeAtPos(pos int) {
	if ll.Head != nil {
		curr := ll.Head
		if pos < 1 {
			fmt.Println("Remove fail, Pos must be >= 1")
		}
		if pos == 1 {
			ll.removeStartNode()
			return
		}
		if curr.Next == nil { // last node in the list
			ll.Head = nil
		} else {
			count := 1
			for count < pos-1 && curr != nil {
				count++
				curr = curr.Next
			}
			if curr != nil {
				fmt.Println("Removing node with data", curr.Next.Data, "at pos", pos)
				curr.Next = curr.Next.Next
			} else {
				fmt.Println("Remove fail, pos out of bound")
			}
		}
	} else {
		fmt.Println("Remove fail, list is empty")
	}
}
