package threadtimesync

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
)

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

	start := time.Now()

	for th.queue.Size > 0 {
		node, err := th.queue.Pop()

		if err != 0 {
			return
		}

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

	th.waiter.AddThread(1)

	go func() {
		time.Sleep(time.Duration(th.sleepTime) * time.Second)
		handler(ag)
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
