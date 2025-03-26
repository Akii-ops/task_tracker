package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var task_cli_update = &cobra.Command{

	Use:   "update",
	Short: "update a task",

	Long: `update a task with a line of description`,
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		var nDesc string

		id, _ = cmd.Flags().GetInt("ID")
		if id == 0 && len(args) > 0 {
			id, _ = strconv.Atoi(args[0])
		}

		nDesc, _ = cmd.Flags().GetString("desc")
		if nDesc == "" && len(args) > 1 {
			nDesc = args[1] // error
		}

		if _, ex := tasktable.TaskList[id]; !ex {
			fmt.Println("ERROR: Task id not exist !")
			return
		}

		entry := tasktable.TaskList[id]
		entry.Desc = nDesc
		entry.UpdateAT = time.Now().Format(TIME_LAYOUT)

		tasktable.TaskList[id] = entry
		fmt.Printf("Task (ID: %d) has been updated successfully\n", id)
	},
}

func init() {

	rootcmd.AddCommand(task_cli_update)

	var id int
	task_cli_update.Flags().IntVarP(&id, "ID", "I", 0, "the id of the task to be updated")

	var nDesc string
	task_cli_update.Flags().StringVarP(&nDesc, "desc", "d", "", "the new description of the task updated")

}
