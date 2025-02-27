package cmd

import (
	// "fmt"
	// "time"
	"github.com/spf13/cobra"
	"strconv"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Task is used to create simple tasks in your CLI",
	Long:  "jfldkslkjds",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		updateTask(cmd, args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string

		val, err := strconv.ParseInt(args[0], 10, 8)

		if err != nil {
			comps = cobra.AppendActiveHelp(comps, "You must provide a valid ID for a task")
		} else if _, err := FetchTaskById(uint8(val)); err != nil {
			comps = cobra.AppendActiveHelp(comps, "You must a valid ID for a task")
		}

		return comps, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateTask(_ *cobra.Command, args []string) {
	if !VerifyDatabase() {
		Must(CreateDatabase())
	}

	id, _ := strconv.ParseInt(args[0], 10, 8)
	num := uint8(id)

	task, _ := FetchTaskById(num)
	task.Completion = true
	CreateNewTask(task)

}
