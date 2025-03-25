package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskStatus uint8

const (
	TODO TaskStatus = iota
	IN_PROGRESS
	DONE
)

type Task struct {
	ID int `json:"id":`

	Desc string `json:"description"`

	Status TaskStatus `json:"status"`

	CreateAt string `json:"createAt"`

	UpdateAT string `json:"updateAt"`
}

var TIME_LAYOUT = "2006-01-02 15:04:05"

var StatusMAP = map[string]TaskStatus{
	"DONE": DONE,
	"done": DONE,

	"in-progress": IN_PROGRESS,
	"INPROGRESS":  IN_PROGRESS,

	"todo": TODO,
	"TODO": TODO,
}

type TaskTable struct {
	TaskList  map[int]Task `json:"tasklist"`
	CurrentID int          `json:"currentID"`
}

var tasktable = TaskTable{
	TaskList:  map[int]Task{},
	CurrentID: 1,
}

var _ = tasktable

func SaveToJson(filename string) error {
	data, err := json.MarshalIndent(tasktable, "", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.WriteFile(filename, data, 0644)

}

func LoadFromJson(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return json.Unmarshal(data, &tasktable)

}
