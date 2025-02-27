package cmd

import (
	// "fmt"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Task is used to create simple tasks in your CLI",
	Long:  "jfldkslkjds",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addTask(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(_ *cobra.Command, args []string) {
	if !VerifyDatabase() {
		Must(CreateDatabase())
	}

	task := Todo{
		Id:           uint8(fetchNextId()),
		Task:         args[0],
		CreationDate: time.Now().UnixMicro(),
		Completion:   false,
	}

	CreateNewTask(&task)

}
