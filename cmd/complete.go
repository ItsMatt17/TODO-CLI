package cmd

import (
	// "fmt"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Update the status of a task based on ID to turn to complete",
	Long:  "jfldkslkjds",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		completeTask(cmd, args)
	},

	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string

		val, err := strconv.ParseInt(args[0], 10, 8)

		if err != nil {
			comps = cobra.AppendActiveHelp(comps, "You must provide a valid ID for a task")
		} else if _, _, err := FetchTaskById(uint8(val)); err != nil {
			comps = cobra.AppendActiveHelp(comps, "You must a valid ID for a task")
		}

		return comps, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeTask(_ *cobra.Command, args []string) {
	if !VerifyDatabase() {
		Must(CreateDatabase())
	}

	id, _ := strconv.ParseInt(args[0], 10, 8)
	num := uint8(id)

	task, idx, err := FetchTaskById(num)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	task.Completion = true
	UpdateTask(task, idx)

}
