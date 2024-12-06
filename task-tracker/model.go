package main

import (
	"fmt"
	"time"
)

type Status int

const (
	ToDo Status = iota
	InProcess
	Done
)

func (task Task) Show(){
	fmt.Println("Task description: ", task.Description)
	fmt.Println("Task ID: ", task.ID)
	fmt.Println("Task status: ", task.Status.convert())
	fmt.Println("Task created: ", task.CreatedAt.Format("2006-01-02"))
	fmt.Println("Task updated: ", task.UpdatedAt.Format("2006-01-02"))
}


func (s Status) convert() string{
	switch s{
	case InProcess:return "InProcess"
	case Done: return "Done"
	default: return "unknown status"
	}
}

type Task struct {
	ID          int        	`json:"id"`
	Description string		`json:"description"`
	Status      Status		`json:"status"`
	CreatedAt   time.Time	`json:"createdAt"`
	UpdatedAt 	time.Time	`json:"updatedAt"`
}