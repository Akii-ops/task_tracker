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
	ID int `json:"id"`

	Desc string `json:"description"`

	Status TaskStatus `json:"status"`

	CreateAt string `json:"createAt"`

	UpdateAT string `json:"updateAt"`
}

func (t Task) String() string {
	return fmt.Sprintf("|%-6d|%-10s|%-15s|%-20s|%-20s|\n", t.ID, t.Status.String(), t.Desc, t.CreateAt, t.UpdateAT)
}

var StatusMAP = map[TaskStatus]string{
	DONE:        "DONE",
	IN_PROGRESS: "IN-PROGRESS",
	TODO:        "TODO",
}

func (T TaskStatus) String() string {
	return StatusMAP[T]
}

var TIME_LAYOUT = "2006-01-02 15:04:05"

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
	if len(data) <= 0 {
		return nil
	}
	return json.Unmarshal(data, &tasktable)

}
