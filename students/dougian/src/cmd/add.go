package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gophercises/task/students/dougian/src/db"
	"strings"
)
func init() {
	rootCmd.AddCommand(addCommand)
}
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		s := strings.Join(args, " ")
		db.AddTask(db.Task{0, s, false})
	},
}

