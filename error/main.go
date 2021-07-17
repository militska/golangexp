package main

import "fmt"

type ErrorModel struct {
	errors map[string][]Error
}

type BaseModel struct {
	ErrorModel
}

type Error struct {
	Field       string
	Message     string
	TechMessage string
}

type Task struct {
	Name        string
	Description string
	BaseModel
}

func (receiver *ErrorModel) AddError(field string, message string) {
	receiver.errors[field] = append(receiver.errors[field], Error{Field: field, Message: message})
}

func (receiver *ErrorModel) GetAllErrors() map[string][]Error {
	return receiver.errors
}

func (receiver *ErrorModel) GetErrors(field string) []Error {
	return receiver.errors[field]
}

func (receiver *ErrorModel) existsErrors() bool {
	return len(receiver.errors) > 0
}

func New() BaseModel {
	base := BaseModel{
		ErrorModel: ErrorModel{},
	}
	base.errors = make(map[string][]Error)
	return base
}
func main() {

	task := Task{
		Name:        "tt",
		Description: "descrr",
		BaseModel:   New(),
	}

	task.AddError("Name", "Название должно быть длиннее")
	task.AddError("Name", "Название должно быть написано кирилец")
	task.AddError("Description", "Неверно заполнено описание")

	fmt.Println(len(task.errors))
	if task.existsErrors() {
		fmt.Println(task.GetAllErrors())
		fmt.Println(task.GetErrors("Name"))
	}
}
