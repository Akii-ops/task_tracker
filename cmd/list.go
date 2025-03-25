package cmd

import (
	"github.com/spf13/cobra"
)

var task_cli_list = &cobra.Command{

	Use:   "list",
	Short: "list a task",
	Long:  `list a task with a line of description`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	rootcmd.AddCommand(task_cli_list)

}
