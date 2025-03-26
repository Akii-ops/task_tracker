package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var task_cli_list = &cobra.Command{

	Use:   "list",
	Short: "list a task",
	Long:  `list a task with a line of description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(args) < 3 {
			PrintSelectedTasks(args...)

		} else {
			PrintAllTasks()
		}
	},
}

func PrintSelectedTasks(cond ...string) {

	fmt.Printf("|%-6s|%-10s|%-15s|%-20s|%-20s|\n", "ID", "Status", "Descrition", "Created Time", "Last Modified")

	fmt.Println("-----------------------------------------------------------------------------")

	for i, v := range cond {
		cond[i] = strings.ToUpper(v)
	}

	for _, entry := range tasktable.TaskList {

		if slices.Index(cond, entry.Status.String()) != -1 {
			fmt.Print(entry)
		}

	}
}

func PrintAllTasks() {

	fmt.Printf("|%-6s|%-10s|%-15s|%-20s|%-20s|\n", "ID", "Status", "Descrition", "Created Time", "Last Modified")

	fmt.Println("-----------------------------------------------------------------------------")
	for _, entry := range tasktable.TaskList {
		fmt.Print(entry)

	}
}

func init() {

	rootcmd.AddCommand(task_cli_list)

}
