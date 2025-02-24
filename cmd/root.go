package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

type Todo struct {
	Id           uint8
	Task         string
	CreationDate int64
	Completion   bool
}

func (t Todo) ToSlice() []string {
	return []string{strconv.FormatInt(int64(t.Id), 10), t.Task,
		strconv.FormatInt(t.CreationDate, 10), strconv.FormatBool(t.Completion)}
}

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is used to create simple tasks in your CLI",
	Long:  "jfldkslkjds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root executed")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
