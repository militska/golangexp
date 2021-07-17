package main

import "fmt"

type BaseModelV2 struct {
	errors map[string][]Error
}

type Error struct {
	Field       string
	Message     string
	TechMessage string
}

type Task struct {
	Name        string
	Description string
	BaseModelV2
}

func (receiver *BaseModelV2) AddError(field string, message string) {
	receiver.errors[field] = append(receiver.errors[field], Error{Field: field, Message: message})
}

func (receiver *BaseModelV2) GetAllErrors() map[string][]Error {
	return receiver.errors
}

func (receiver *BaseModelV2) GetErrors(field string) []Error {
	return receiver.errors[field]
}

func (receiver *BaseModelV2) existsErrors() bool {
	return len(receiver.errors) > 0
}

func main() {

	errors := make(map[string][]Error)

	task := Task{
		Name:        "New task",
		Description: "ddd",
		BaseModelV2: BaseModelV2{},
	}

	task.BaseModelV2.errors = errors

	task.AddError("Name", "Название должно быть длиннее")
	task.AddError("Name", "Название должно быть написано кирилецей")
	task.AddError("Description", "Неверно заполнено описание")

	fmt.Println(len(task.errors))
	if task.existsErrors() {
		fmt.Println(task.GetAllErrors())
		fmt.Println(task.GetErrors("Name"))
	}
}
