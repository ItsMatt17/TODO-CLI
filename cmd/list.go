package cmd

import (
	// "fmt"
	// "time"
	"fmt"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Task is used to create simple tasks in your CLI",
	Long:  "jfldkslkjds",
	Run: func(cmd *cobra.Command, args []string) {
		listTask(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTask(_ *cobra.Command, _ []string) {
	if !VerifyDatabase() {
		return
	}

	tasks := MustFetchTasks()

	writer := tabwriter.NewWriter(os.Stdout, 8, 0, 2, ' ', 0)

	fmt.Fprintln(writer, "ID\tTask\tCreation Time\tCompletion")
	for _, task := range tasks {

		time := timediff.TimeDiff(time.UnixMicro(task.CreationDate))

		fmt.Fprintf(writer, "%d\t%s\t%t\t%s\t\n", task.Id, task.Task, task.Completion, time) //Other way to go to next line besides the last \n?

	}
	defer writer.Flush()

}
