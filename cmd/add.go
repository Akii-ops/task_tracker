package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var task_cli_add = &cobra.Command{

	Use:   "add",
	Short: "add a task",
	Args:  cobra.ExactArgs(1),
	Long:  `add a task with a line of description`,
	Run: func(cmd *cobra.Command, args []string) {

		tasktable.TaskList[tasktable.CurrentID] = Task{
			ID:       tasktable.CurrentID,
			Desc:     args[0],
			Status:   TODO,
			CreateAt: time.Now().Format(TIME_LAYOUT),
			UpdateAT: time.Now().Format(TIME_LAYOUT),
		}
		fmt.Printf("Task (ID : %d ) has added successfully\n", tasktable.CurrentID)
		tasktable.CurrentID++

	},
}

func init() {

	rootcmd.AddCommand(task_cli_add)

	// task_cli_add.Flags().StringP("description", "d", "", "the description of this task, default value is empty")
	// task_cli_add.Flags().StringP("status","s","TODO","to mark the status of string")
}
