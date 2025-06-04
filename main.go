package main

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type Task struct {
	taskName        string
	taskDescription string
	taskDaone       bool
}

type TaskManager struct {
	task Task
}

func main() {
	pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithStyle("TERMINAL", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Task Manager", pterm.FgMagenta.ToStyle())).Render()

	viewTask := TaskManager{
		task: addTask(),
	}

	fmt.Println(viewTask)
}

func addTask() Task {
	Name, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Name Task").Show()

	taskDes, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Task Description").Show()

	taskAdd := Task{
		taskName:        Name,
		taskDescription: taskDes,
		taskDaone:       false,
	}

	return taskAdd
}
