package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add          string
	Del          int
	Edit         string
	Complete     int
	CompleteAll  bool
	List         bool
	ToggleStatus int
}

func NewCmdFlags() *CmdFlags {
	c := CmdFlags{}

	flag.StringVar(&c.Add, "add", "", "add todo item")
	flag.IntVar(&c.Del, "del", -1, "delete todo item")
	flag.StringVar(&c.Edit, "edit", "", "edit todo item")
	flag.IntVar(&c.Complete, "complete", -1, "complete todo item")
	flag.BoolVar(&c.CompleteAll, "complete-all", false, "complete all todo items")
	flag.BoolVar(&c.List, "list", false, "list todo items")
	flag.IntVar(&c.ToggleStatus, "toggle-status", -1, "toggle todo item status")

	flag.Parse()

	return &c
}

func (c *CmdFlags) Execute(todos *Todos) error {
	switch {
	case c.Add != "":
		todos.add(c.Add)
	case c.Del != -1:
		return todos.delete(c.Del)
	case c.Edit != "":
		parts := strings.SplitN(c.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid edit format")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index")
			os.Exit(1)
		}
		return todos.edit(index, parts[1])
	case c.Complete != -1:
		return todos.toggleStatus(c.Complete)
	case c.CompleteAll:
		todos.completeAll()
	case c.List:
		todos.print()
	case c.ToggleStatus != -1:
		return todos.toggleStatus(c.ToggleStatus)
	default:
		fmt.Println("Error: Invalid command")
		os.Exit(1)
	}
	return nil
}
