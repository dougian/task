package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gophercises/task/students/dougian/src/db"
	"fmt"
)


func init() {
	rootCmd.AddCommand(listCommand)
}
var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists the task that exist in the db",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		for _, t := range db.ListTasks() {
			if !t.Status {
				fmt.Println(t.Id, " - ", t.Value)
			}
		}
	},
}
