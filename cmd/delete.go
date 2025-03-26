package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var task_cli_delete = &cobra.Command{

	Use: "delete",

	Short: "delete a task",
	Long:  `delete a task `,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("ID")
		if id == 0 && len(args) > 0 {
			id, _ = strconv.Atoi(args[0])
		}
		if _, ex := tasktable.TaskList[id]; !ex {
			fmt.Println("ERROR : Task ID not exist!")
			return
		}
		delete(tasktable.TaskList, id)
		fmt.Printf("TASK (ID : %d) has been deleted \n", id)
	},
}

func init() {

	rootcmd.AddCommand(task_cli_delete)
	var id int
	task_cli_delete.Flags().IntVarP(&id, "ID", "I", 0, "the id of tasked to be deleted")

}
