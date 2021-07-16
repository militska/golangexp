package main

import "fmt"

type BaseModel struct {
	errors []Error
}

type Error struct {
	Field   string
	Message string
}

type Task struct {
	Name        string
	Description string
	BaseModel
}

func (receiver *BaseModel) AddError(field string, message string) {
	receiver.errors = append(receiver.errors, Error{Field: field, Message: message})
}

func (receiver *BaseModel) GetErrors() []Error {
	return receiver.errors
}

func (receiver *BaseModel) existsErrors() bool {
	return len(receiver.errors) > 0
}

func main() {
	task := Task{
		Name:        "New task",
		Description: "ddd",
		BaseModel:   BaseModel{},
	}

	task.AddError("Name", "Название должно быть длиннее")
	task.AddError("Description", "Неверно заполнено описание")

	fmt.Println(len(task.errors))
	if task.existsErrors() {
		fmt.Println(task.GetErrors())
	}
}
