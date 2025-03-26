package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var mark_in_progress = &cobra.Command{

	Use: "mark-in-progress",

	Short: "mark a task in_progress",
	Long:  `mark a task in_progress `,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("ID")
		if id == 0 && len(args) > 0 {
			id, _ = strconv.Atoi(args[0])
		} else {
			fmt.Println("ERROR: Task ID invalid !")
		}

		entry := tasktable.TaskList[id]
		entry.Status = IN_PROGRESS
		entry.UpdateAT = time.Now().Format(TIME_LAYOUT)
		tasktable.TaskList[id] = entry

	},
}

func init() {

	rootcmd.AddCommand(mark_done)
	var id int
	mark_done.Flags().IntVarP(&id, "ID", "I", 0, "the id of tasked to be deleted")

}
