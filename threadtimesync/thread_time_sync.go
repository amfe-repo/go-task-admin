package threadtimesync

import "time"

type Wait struct {
	actualThreads    int
	completedThreads int
}

// Add thread reference number
func (w *Wait) AddThread(i int) {
	w.actualThreads += i
}

// Add thread reference finished
func (w *Wait) FinishThread() {
	w.completedThreads += 1
}

// Waiting for reference finished
func (w *Wait) Waiting() {
	for !w.VerifyFinished() {
		//fmt.Println("Actual Threads: ", w.actualThreads, "Completed Threads: ", w.completedThreads)
		time.Sleep(500)
	}
}

// Verify if all threads have terminated
func (w *Wait) VerifyFinished() bool {
	return w.completedThreads == w.actualThreads
}
