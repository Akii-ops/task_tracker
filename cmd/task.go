package cmd

type TaskStatus uint8

const (
	TODO TaskStatus = iota
	IN_PROGRESS
	DONE
)

type Task struct {
	ID int

	Desc string

	Status TaskStatus

	CreateAt string

	UpdateAT string
}

var StatusMAP = map[string]TaskStatus{
	"DONE": DONE,
	"done": DONE,

	"in-progress": IN_PROGRESS,
	"INPROGRESS":  IN_PROGRESS,

	"todo": TODO,
	"TODO": TODO,
}

type TaskTable struct {
	TaskList []Task

	CurrentID int
}




var tasktable = &TaskTable{
	CurrentID: 1,
	
}

var _ = tasktable
