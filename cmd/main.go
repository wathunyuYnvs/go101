package main

import (
	"os"
	"task"
)

func main() {
	args := os.Args[1:]
	switch args[0] {
	case "add":
		task.Add(args[1], args[2])
	case "list":
		task.List()
	}
}
