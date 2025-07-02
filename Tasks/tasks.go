package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type stage struct {
	TaskDesc string
	Status   bool
}

var tasks []stage
var dinput string
var loop bool

func appendingTasks() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	input = strings.TrimSpace(input)

	dinput = input

	if strings.HasPrefix(input, "d ") {
		markTaskDone()
		return
	} else if strings.HasPrefix(input, "x ") {
		delete()
		return
	} else if strings.HasPrefix(input, "wq") {
		saveTask()
		fmt.Println("File saved. Quit action initaited...")
		os.Exit(0)
		return
	} else if strings.HasPrefix(input, "w") {
		saveTask()
		fmt.Println("File saved.")
	} else {

		realTask := stage{TaskDesc: input, Status: false}

		tasks = append(tasks, realTask)

		fmt.Println("Task added.")

		for i, task := range tasks {
			statusui := "[]"
			if task.Status {
				statusui = "[x]"
			}
			fmt.Printf("%d. %s %s\n", i+1, statusui, task.TaskDesc)
		}
	}
}

func markTaskDone() {
	doneinput := dinput[2:]
	doneinputint, _ := strconv.Atoi(doneinput)
	fmt.Println(doneinput)

	for i := 0; i < len(tasks); i++ {
		if i+1 == doneinputint {
			tasks[i].Status = true
			fmt.Println(doneinput, " marked as done\n")
			break
		}
	}
}

func delete() {
	doneinput := dinput[2:]
	doneinputint, _ := strconv.Atoi(doneinput)
	//	doneinputint = doneinputint - 1

	tasks = append(tasks[:doneinputint-1], tasks[doneinputint:]...)
}

func saveTask() {
	taskData, _ := json.MarshalIndent(tasks, "", " ")
	err := ioutil.WriteFile("Savedtasks.json", taskData, 0644)
	if err != nil {
		fmt.Println("There es an error in saving the file")
	}
}

func LoadTask() {
	loadedTask, _ := ioutil.ReadFile("Savedtasks.json")

	json.Unmarshal(loadedTask, &tasks)
}

func main() {
	LoadTask()
	for loop == false {
		appendingTasks()

	}
}
