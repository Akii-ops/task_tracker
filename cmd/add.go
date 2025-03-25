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
		fmt.Println(args[0])

		tasktable.TaskList = append(tasktable.TaskList, Task{
			ID:       tasktable.CurrentID,
			Desc:     args[0],
			Status:   StatusMAP["TODO"],
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdateAT: time.Now().Format("2006-01-02 15:04:05"),
		})

		tasktable.CurrentID += 1
		fmt.Println("1")
		for _, v := range tasktable.TaskList {
			fmt.Println(v)
		}
	},
}

func init() {

	rootcmd.AddCommand(task_cli_add)

	// task_cli_add.Flags().StringP("description", "d", "", "the description of this task, default value is empty")
	// task_cli_add.Flags().StringP("status","s","TODO","to mark the status of string")
}
