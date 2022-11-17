package threadtimesync

import "github.com/charmbracelet/bubbles/list"

//import "fmt"

type Func func(*[]list.Item)

type ThreadsQueue struct {
	Head *ThreadsNode
	Size int
}

type ThreadsNode struct {
	DataExecute Func
	Argument    *[]list.Item
	Next        *ThreadsNode
}

// Push node of data
func (this *ThreadsQueue) Push(data Func, ag *[]list.Item) {

	// Creating node
	var tn *ThreadsNode = &ThreadsNode{
		DataExecute: data,
		Argument:    ag,
		Next:        nil,
	}

	// Verify if this node is the first
	if this.Head == nil {
		this.Head = tn
		this.Size++
		return
	}

	var pointer *ThreadsNode = this.Head

	// Go through the whole queue
	for pointer.Next != nil {
		pointer = pointer.Next
	}

	// Insert the node in final of queue
	pointer.Next = tn

	// Agument size of queue
	this.Size++
}

func (this *ThreadsQueue) Pop() (ThreadsNode, int) {
	//var pointer *ThreadsNode = this.Head
	err := 0

	// Saving head of queue
	temp := *this.Head

	if this.Size == 0 {
		err = 1
		return temp, err
	}

	// Referencing the head to next pointer
	this.Head = this.Head.Next

	// Decrement size of queue
	this.Size--

	// Return last head of the queue
	return temp, err
}
