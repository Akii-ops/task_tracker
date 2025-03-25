package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var task_cli_update = &cobra.Command{

	Use:   "update",
	Short: "update a task",
	Long:  `update a task with a line of description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("update")
	},
}

func init() {

	rootcmd.AddCommand(task_cli_update)

}
