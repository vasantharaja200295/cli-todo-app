package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}
	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) setCompleted(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}
	if !(*todos)[index].Completed {
		(*todos)[index].Completed = true
		completedTime := time.Now()
		(*todos)[index].CompletedAt = &completedTime
	}
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Id", "Title", "Completed", "Created At", "Completed At")
	for i, todo := range *todos {
		completed := "No"
		completedAt := "N/A"
		if todo.Completed {
			completed = "Yes"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC3339)
			}
		}
		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC3339), completedAt)
	}
	table.Render()
}

func (todos *Todos) revertCompleted(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}
	if (*todos)[index].Completed {
		(*todos)[index].Completed = false
		(*todos)[index].CompletedAt = nil
	}
	return nil
}

func (todos *Todos) toggleStatus(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}
	if !(*todos)[index].Completed {
		return todos.setCompleted(index)
	}
	return todos.revertCompleted(index)
}

func (todos *Todos) completeAll() {
	for i := 0; i < len(*todos); i++ {
		todos.setCompleted(i)
	}
}

// utils and helper functions
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}
	return nil
}
