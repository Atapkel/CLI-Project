package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tasks := []Task{}
	file, err := os.OpenFile("data/data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("error on finding data file: ", err)
		return
	}
	defer file.Close()

	
	decode := json.NewDecoder(file)
	err = decode.Decode(&tasks)


	if err != nil {
		fmt.Println("error on getting data from json", err)
		return
	}
	check := true
	for check{
		fmt.Printf("Write HELP to get: %s! ", "HELP")
		req := getInput("Enter your request", reader)

		

		switch getFirstWord(req){
		case "HELP":
			fmt.Printf("To add task write: add\n", )
			fmt.Printf("To update task write: update \n" )
			fmt.Printf("To mark task as in-progress write: %s %s\n", "mark-in-progress", "your_task_id")
			fmt.Printf("To mark task as done write: %s %s\n", "mark-done", "your_task_id")
			fmt.Printf("To see all task write: %s\n", "list")
			fmt.Printf("To get tasks that done type: %s\n", "list-done")
			fmt.Printf("To get tasks that in-progress type: %s\n", "list-in-process")
			fmt.Printf("To delete task type: %s\n", "del")

		case "add":
			taskDescription := getInput("Enter your task that you will added", reader)
			task := Task{
				ID: tasks[len(tasks) - 1].ID + 1,
				Description: taskDescription,
				Status: InProcess,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			tasks = append(tasks, task)
			fmt.Println("Task added with id:", task.ID)
		case "update":
			taskIDStr := getInput("Enter your task_id that you want update", reader)
			taskID, err := strconv.Atoi(taskIDStr)
			if err != nil {
				fmt.Println("Invalid task Id")
				continue
			}
			var taskToUpdate Task
			for _, task := range tasks{
				if task.ID == taskID {
					taskToUpdate = task
				}
			}
		
			if taskToUpdate.Description == "" {
				fmt.Println("Your task is not found")
				continue
			}

			fmt.Println("your task details")
			taskToUpdate.Show()

			taskStatusStr := getInput("Enter task status: done or in-process", reader)
			if taskStatusStr == "done" {
				taskToUpdate.Status = Done
				fmt.Println("Task done, GREAT BRO!")
			}else if taskStatusStr == "inprocess" {
				taskToUpdate.Status = InProcess
				fmt.Println("Task In Proccess!")
			}else{
				fmt.Println("invalid task status")
				continue
			}

		case "list":
			for _, task := range tasks{
				task.Show()
			}
		case "list-done":
			for _, task := range tasks{
				if task.Status == Done {
					task.Show()
					fmt.Println("statussss", task.Status)
				}
			}
		case "list-in-process":
			for _, task := range tasks{
				if task.Status == InProcess {
					task.Show()
					fmt.Println("task status is ",task.Status)
				}
			}
		case "delete":
			idStr := getInput("Enter id of task that you want to delete", reader)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Your ID invalid")
				continue
			}

			index := -1
			for i, el := range tasks{
				if el.ID == id {
					index = i
					break
				}
			} 
			if index == -1 {
				fmt.Printf("Your task with id %s is not find\n", string(id))
				continue
			}

			tasks = delete(tasks, index)

		case "quit":
			fmt.Println("thanks for using")
			check = false
			
		}
	
	}
	file.Seek(0, 0)
	file.Truncate(0)
	encode := json.NewEncoder(file)
	err = encode.Encode(tasks)
	if err != nil {
		fmt.Println("tasks not saved ", err)
	}
}

func delete(tasks []Task, id int) (ans []Task){
	ans = append(tasks[:id],tasks[id+1:]...)
	return 
}


func getFirstWord(text string) string{
	return strings.Split(text, " ")[0]
	
}


func getInput(prompt string, reader *bufio.Reader) string{
	fmt.Print(prompt,": ")
	request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error on getting user input")
	}
	return strings.TrimSpace(request)
}