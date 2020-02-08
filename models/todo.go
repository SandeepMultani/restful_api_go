package models

import (
	"errors"
	"fmt"
)

type ToDo struct {
	ID          int
	Task        string
	IsCompleted bool
}

var (
	todos  []*ToDo
	nextID = 1
)

func GetToDos() []*ToDo {
	return todos
}

func AddTodo(t ToDo) (ToDo, error) {
	if t.ID != 0 {
		return ToDo{}, errors.New("New ToDo must not include an id or it must be set to zero.")
	}
	t.ID = nextID
	nextID++
	todos = append(todos, &t)
	return t, nil
}

func GetToDoByID(id int) (ToDo, error) {
	for _, t := range todos {
		if t.ID == id {
			return *t, nil
		}
	}
	return ToDo{}, fmt.Errorf("ToDo with ID '%v' not found", id)
}

func UpdateToDo(t ToDo) (ToDo, error) {
	for i, todo := range todos {
		if todo.ID == t.ID {
			todos[i] = &t
			return t, nil
		}
	}
	return ToDo{}, fmt.Errorf("ToDo with ID '%v' not found", t.ID)
}

func DeleteToDoByID(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("ToDo with ID '%v' not found", id)
}
