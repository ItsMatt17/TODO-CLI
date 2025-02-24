package cmd

import (
	// "fmt"
	// "time"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Task is used to create simple tasks in your CLI",
	Long:  "jfldkslkjds",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listTask(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTask(_ *cobra.Command, args []string) {
	if !VerifyDatabase() {
		Must(CreateDatabase())
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(writer, "ID\tTask\tCreation Time\tCompletion")
	// for _, val := range {
	// 	fmt.Fprintln()
	// }

}
