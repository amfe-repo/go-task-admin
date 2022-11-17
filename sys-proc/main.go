package sysproc

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"

	"github.com/shirou/gopsutil/process"

	"golang.org/x/sys/windows/svc/mgr"

	models "github.com/go-system-tasks/models"
)

// Get Services of OS
func GetServices() ([]list.Item, interface{}) {
	items := []list.Item{}

	// Connect to services
	m, err := mgr.Connect()

	if err != nil {
		return nil, err
	}

	// List services
	svcs, _ := m.ListServices()

	for _, i := range svcs {
		// Open particular service
		s, err := m.OpenService(i)

		if err == nil {

			// Get info of particular service
			query, err := s.Query()

			if err != nil {
				continue
			}

			// Calling all info of services with SERVICE ID [PID]
			proc_info, err := process.NewProcess(int32(query.ProcessId))

			if err == nil {
				cpu, _ := proc_info.CPUPercent()

				mem, _ := proc_info.MemoryPercent()

				// Creating item [service structure] and saving
				item := models.Item(models.Item{
					Name: s.Name,
					Pid:  query.ProcessId,
					Cpu:  fmt.Sprintf("%.4f", cpu),
					Mem:  fmt.Sprintf("%.4f", mem),
				})

				items = append(items, item)

			}
		}
	}

	return items, nil
}

// Get Processes of OS
func GetProccesses() ([]list.Item, interface{}) {
	items := []list.Item{}

	//Get process of OS
	p, err := process.Processes()

	if err != nil {
		return nil, err
	}

	// Iterating all proccesses of OS
	for _, proccess := range p {

		name, _ := proccess.Name()
		pid := proccess.Pid
		cpu, _ := proccess.CPUPercent()
		mem, _ := proccess.MemoryPercent()

		// Creating item [process structure] and saving
		item := models.Item(models.Item{
			Name: name,
			Pid:  uint32(pid),
			Cpu:  fmt.Sprintf("%.6f", cpu),
			Mem:  fmt.Sprintf("%.6f", mem),
		})

		items = append(items, item)

	}

	return items, nil
}

// Kill Process of OS
func CreateAndKillProcess(pid uint32) {
	p, err := process.NewProcess(int32(pid))

	if err != nil {
		panic(err)
	}

	p.Kill()
}
