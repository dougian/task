package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/gophercises/task/students/dougian/src/db"
	"strconv"
	"os"
)

func init() {
	rootCmd.AddCommand(completeCommand)
}

var completeCommand = &cobra.Command{
	Use:   "do",
	Short: "Marks the given id as completed",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		for _, s := range args {
			id, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			db.CompleteTask(uint64(id))
		}
	},
}
