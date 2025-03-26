package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var mark_done = &cobra.Command{

	Use: "mark-done",

	Short: "mark a task done",
	Long:  `mark a task done `,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("ID")
		if id == 0 && len(args) > 0 {
			id, _ = strconv.Atoi(args[0])
		} else {
			fmt.Println("ERROR: Task ID invalid !")
			return
		}

		entry := tasktable.TaskList[id]
		entry.Status = DONE
		entry.UpdateAT = time.Now().Format(TIME_LAYOUT)
		tasktable.TaskList[id] = entry

	},
}

func init() {

	rootcmd.AddCommand(mark_in_progress)
	var id int
	mark_in_progress.Flags().IntVarP(&id, "ID", "I", 0, "the id of tasked to be deleted")

}
