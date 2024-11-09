package main

import "fmt"

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	err := cmdFlags.Execute(&todos)
	if err != nil {
		fmt.Println(err)
	}
	storage.Save(todos)
}
