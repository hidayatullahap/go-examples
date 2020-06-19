package main

import (
	"fmt"

	"github.com/hidayatullahap/go-examples/pattern/hook/db"
)

func main() {
	var todo Todo

	con := &db.Db{}
	con.Save(&todo)
}

type Todo struct {
	Name string
}

func (t *Todo) BeforeSave() {
	fmt.Println("*Todo.BeforeSave")
}

func (t *Todo) AfterSave() {
	fmt.Println("*Todo.AfterSave")
}
