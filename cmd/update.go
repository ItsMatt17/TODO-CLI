package cmd

import (
	"fmt"
	// "time"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the name/description of a task based on ID",
	Long:  "jfldkslkjds",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		updateTask(cmd, args)
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
	rootCmd.AddCommand(updateCmd)
}

func updateTask(_ *cobra.Command, args []string) {
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

	task.Task = args[1]

	UpdateTask(task, idx)

}
