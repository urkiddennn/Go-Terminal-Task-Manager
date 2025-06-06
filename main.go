package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type Task struct {
	TaskName        string `json:"taskName"`
	TaskDescription string `json:"taskDescription"`
	TaskDone        bool   `json:"taskDone"` // Fixed field name to match JSON tag
}

type TaskManager struct {
	Tasks []Task `json:"tasks"`
}

func main() {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("TERMINAL", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Task Manager", pterm.FgMagenta.ToStyle()),
	).Render()

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
		case "Delete Task":
			deleteTask(&tm) // Added placeholder for delete functionality
		case "Exit":
			if err := saveTask(tm); err != nil {
				pterm.Error.Printfln("Failed to save tasks: %v", err)
			} else {
				pterm.Success.Println("Tasks saved successfully!")
			}
			return
		}
	}
}

// Add new Task
func addTask(tm *TaskManager) {
	name, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Task Name").Show()
	taskDes, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Task Description").Show()

	taskAdd := Task{
		TaskName:        name,
		TaskDescription: taskDes,
		TaskDone:        false,
	}

	tm.Tasks = append(tm.Tasks, taskAdd)
	if err := saveTask(*tm); err != nil {
		pterm.Error.Printfln("Failed to save task: %v", err)
		return
	}
	pterm.Success.Println("Added new task!")
}

// View list of tasks
func listTask(tm TaskManager) {
	if len(tm.Tasks) == 0 {
		pterm.Info.Println("No tasks available.")
		return
	}
	table := pterm.TableData{{"Task Name", "Task Description", "Task Done"}}
	for _, task := range tm.Tasks {
		taskBool := strconv.FormatBool(task.TaskDone)
		table = append(table, []string{task.TaskName, task.TaskDescription, taskBool})
	}
	pterm.DefaultTable.WithHasHeader().WithData(table).Render()
}

// Mark task as done
func markDone(tm *TaskManager) {
	if len(tm.Tasks) == 0 {
		pterm.Warning.Println("No tasks available")
		return
	}

	var options []string
	for _, task := range tm.Tasks {
		options = append(options, task.TaskName)
	}

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	for i, task := range tm.Tasks {
		if task.TaskName == selectedOption {
			tm.Tasks[i].TaskDone = true
			if err := saveTask(*tm); err != nil {
				pterm.Error.Printfln("Failed to save tasks: %v", err)
				return
			}
			pterm.Success.Printfln("Task '%s' marked as done!", selectedOption)
			return
		}
	}
	pterm.Error.Printfln("Task '%s' not found!", selectedOption)
}

// Delete task (placeholder implementation)
func deleteTask(tm *TaskManager) {
	if len(tm.Tasks) == 0 {
		pterm.Warning.Println("No tasks available")
		return
	}

	var options []string
	for _, task := range tm.Tasks {
		options = append(options, task.TaskName)
	}

	selectedOption, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Enter Task Name to Delete").Show()

	for i, task := range tm.Tasks {
		if task.TaskName == selectedOption {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			// tm.Tasks = slices.Delete(tm.Tasks[:i], tm.Tasks[i+1:])
			if err := saveTask(*tm); err != nil {
				pterm.Error.Printfln("Failed to save tasks: %v", err)
				return
			}
			pterm.Success.Printfln("Task '%s' deleted!", selectedOption)
			return
		}
	}
	pterm.Error.Printfln("Task '%s' not found!", selectedOption)
}

// Load tasks from file
func loadTask() TaskManager {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			pterm.Info.Println("No tasks file found, starting with empty task list.")
			return TaskManager{}
		}
		pterm.Error.Printfln("Error reading tasks file: %v", err)
		return TaskManager{}
	}

	var tm TaskManager
	if err := json.Unmarshal(data, &tm); err != nil {
		pterm.Error.Printfln("Error parsing tasks file: %v", err)
		return TaskManager{}
	}
	return tm
}

// Save tasks to file
func saveTask(tm TaskManager) error {
	data, err := json.MarshalIndent(tm, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling tasks: %v", err)
	}
	if err := os.WriteFile("tasks.json", data, 0644); err != nil {
		return fmt.Errorf("error writing tasks file: %v", err)
	}
	return nil
}
