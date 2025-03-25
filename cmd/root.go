package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{

	Use:   "rootCMD",
	Short: "root CMD",
	Long:  "THIS IS ROOT CMD",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nothing but a root cmd output")
		fmt.Println(args)
	},
}

var verbose bool

func init() {
	rootcmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "详细输出模式")
}

func Execute() {

	if err := rootcmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
