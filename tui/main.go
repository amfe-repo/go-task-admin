package tui

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/go-system-tasks/helpers"
	models "github.com/go-system-tasks/models"
)

const listHeight = 30

var (
	titleStyle         = lipgloss.NewStyle().MarginLeft(2)
	itemStyle          = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle  = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	selectedItemStyle2 = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("180"))
	paginationStyle    = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle          = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle      = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

var selection bool = false

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(models.Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i.Name)

	var element models.Item = i

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			pid := element.Pid
			cpu := element.Cpu
			mem := element.Mem
			str := selectedItemStyle2.Render(fmt.Sprintf("[%s]\n>", s))
			info := selectedItemStyle.Render(fmt.Sprintf("PID:[%d] CPU:[%s] MEM:[%s]", pid, cpu, mem))
			return str + info
		}
	}

	fmt.Fprintf(w, fn(str))
}

type model struct {
	list     list.Model
	list2    list.Model
	items    []models.Item
	selected bool
	choice   models.Item
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

var m1 = model{}

func createList(items []list.Item, defaultWidth int, title string, changeOption [2]string) list.Model {

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = title //"Proccess in your machine\n\n[Pid] [Name] [Cpu] [Mem]"

	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.KeyMap.GoToStart.SetHelp("enter", "kill proccess")
	l.KeyMap.GoToStart.SetEnabled(true)
	l.KeyMap.ShowFullHelp.SetHelp(changeOption[0], changeOption[1])
	l.KeyMap.ShowFullHelp.SetEnabled(true)

	return l
}

func fullModelPrincipal(list list.Model, list2 list.Model) {
	m1 = model{
		list:  list,
		list2: list2,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(models.Item)
			if ok {
				m.choice = models.Item(i)
				selection = true
			}
			return m, nil

		case "-":
			m1.selected = !m1.selected
			m.selected = m1.selected

			m.list = m1.list2

			return m, nil

		case "+":
			m1.selected = !m1.selected
			m.selected = m1.selected

			m.list = m1.list

			return m, nil
		}

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {

	if selection {

		helpers.KillProcess(m.choice.Pid)
		m.list.RemoveItem(m.list.Index())

		selection = false
	}
	if m.quitting {
		return quitTextStyle.Render("Bye bye!")
	}

	return "\n" + m.list.View()
}

func Start(items []list.Item, items2 []list.Item) {

	const defaultWidth = 700

	l := createList(items, defaultWidth, "Services in your machine", [2]string{"-", "process view"})

	k := createList(items2, defaultWidth, "Proccess in your machine", [2]string{"+", "service view"})

	fullModelPrincipal(l, k)

	m := model{
		list:  l,
		list2: k,
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
