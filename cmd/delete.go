package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var task_cli_delete = &cobra.Command{

	Use:   "delete",
	Short: "delete a task",
	Long:  `delete a task `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println("delete")
	},
}

func init() {

	rootcmd.AddCommand(task_cli_delete)

}
