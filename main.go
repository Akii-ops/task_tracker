package main

import (
	"task-cli/cmd"
)

func main() {

	err := cmd.LoadFromJson("tasklist.json")
	if err != nil {
		panic(err)
	}

	cmd.Execute()

	err = cmd.SaveToJson("tasklist.json")
	if err != nil {
		panic(err)
	}
}
