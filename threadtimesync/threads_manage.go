package threadtimesync

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
)

// Threads Structure
type Threads struct {
	waiter      Wait
	queue       ThreadsQueue
	sleepTime   int
	timeElapsed time.Duration
}

func (th *Threads) init() {
	th.waiter = Wait{}
	th.queue.Size = 0
	th.sleepTime = 0
	th.timeElapsed = 0
}

//Insert thread
func (th *Threads) InsertThread(handler Func, ag *[]list.Item) {
	th.queue.Push(handler, ag)
}

//Init threads
func (th *Threads) StartThreads() {
	th.timeElapsed = 0

	// Count time when threads init to execute
	start := time.Now()

	for th.queue.Size > 0 {
		// Get first element of queue
		node, err := th.queue.Pop()

		if err != 0 {
			return
		}

		// Create concurrent thread
		th.createThreads(node.DataExecute, node.Argument)
	}

	//Wait threads
	th.waitThreads()

	th.timeElapsed = time.Since(start)
}

//Set time to sleep
func (th *Threads) SetTimeSleep(time int) {
	th.sleepTime = time
}

//Create function and becomes a thread
func (th *Threads) createThreads(handler Func, ag *[]list.Item) {

	// Add count of thread in execution
	th.waiter.AddThread(1)

	// Execute concurrently the thread
	go func() {
		time.Sleep(time.Duration(th.sleepTime) * time.Second)
		handler(ag)
		// Delete count of thread
		th.waiter.FinishThread()
	}()

}

//Waiting to all threads finish
func (th *Threads) waitThreads() {
	th.waiter.Waiting()
}

//Obtain time elapsed of last threads
func (th *Threads) GetTimeElapsed() time.Duration {
	return th.timeElapsed
}
