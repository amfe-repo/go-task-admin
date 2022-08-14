package helpers

import (
	"github.com/charmbracelet/bubbles/list"

	sysproc "github.com/go-system-tasks/sys-proc"
	"github.com/go-system-tasks/threadtimesync"
)

func GetListServices(svcs *[]list.Item) {

	list, err := sysproc.GetServices()

	if err != nil {
		panic(err)
	}

	*svcs = list

}

func GetListProccesses(procs *[]list.Item) {

	list, err := sysproc.GetProccesses()

	if err != nil {
		panic(err)
	}

	*procs = list

}

func GetSystemProcSvcsInfo(svcs, procs *[]list.Item) {
	th := threadtimesync.Threads{}

	th.InsertThread(GetListServices, svcs)
	th.InsertThread(GetListProccesses, procs)

	th.StartThreads()
}

func Refresh() ([]list.Item, []list.Item) {

	var svcs, procs []list.Item

	GetSystemProcSvcsInfo(&svcs, &procs)

	return svcs, procs
}

func KillProcess(pid uint32) {
	sysproc.CreateAndKillProcess(pid)
}
