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

	var tn *ThreadsNode = &ThreadsNode{
		DataExecute: data,
		Argument:    ag,
		Next:        nil,
	}

	if this.Head == nil {
		this.Head = tn
		this.Size++
		return
	}

	var pointer *ThreadsNode = this.Head

	for pointer.Next != nil {
		pointer = pointer.Next
	}

	pointer.Next = tn
	this.Size++
}

func (this *ThreadsQueue) Pop() (ThreadsNode, int) {
	//var pointer *ThreadsNode = this.Head
	err := 0

	temp := *this.Head

	if this.Size == 0 {
		err = 1
		return temp, err
	}

	this.Head = this.Head.Next

	this.Size--

	return temp, err
}
