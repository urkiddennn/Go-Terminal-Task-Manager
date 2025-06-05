package main

import (
	"fmt"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type Task struct {
	taskName        string
	taskDescription string
	taskDaone       bool
}

type TaskManager struct {
	Tasks []Task
}

func main() {
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("TERMINAL", pterm.FgCyan.ToStyle()), putils.LettersFromStringWithStyle("Task Manager", pterm.FgMagenta.ToStyle())).Render()
	tm := loadTask()
	listTask(tm)
	for {
		choice, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"Add Task", "List Task", "Mark Done", "Delete Task", "Exit"}).Show()

		switch choice {
		case "Add Task":
			addTask(&tm)
		case "List Task":
			listTask(tm)
		case "Mark Done":
			markDone(&tm)
		case "Exit":
			break
		}
	}
}

// Add new Task
func addTask(tm *TaskManager) {
	Name, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Name Task").Show()

	taskDes, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Task Description").Show()

	taskAdd := Task{
		taskName:        Name,
		taskDescription: taskDes,
		taskDaone:       false,
	}

	tm.Tasks = append(tm.Tasks, taskAdd)
	pterm.Success.Println("Added new Task!")
}

// View list task
func listTask(tm TaskManager) {
	table := pterm.TableData{{"Task Name", "Task Description", "Task Done"}}
	for _, task := range tm.Tasks {
		taskBool := strconv.FormatBool(task.taskDaone)
		table = append(table, []string{task.taskName, task.taskDescription, taskBool})
	}
	pterm.DefaultTable.WithHasHeader().WithData(table).Render()
}

func markDone(tm *TaskManager) {
	var options []string

	for _, tasks := range tm.Tasks {
		options = append(options, tasks.taskName)
	}

	fmt.Println("All options: ", options)

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	fmt.Println("The selectedOption is ", selectedOption)
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}

// Load Task
func loadTask() TaskManager {
	var tm TaskManager
	fmt.Println(tm)

	return tm
}
