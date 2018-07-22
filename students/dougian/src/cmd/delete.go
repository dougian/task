package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/gophercises/task/students/dougian/src/db"
	"strconv"
	"os"
)

func init() {
	rootCmd.AddCommand(deleteCommand)
}

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task given an id",
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			id, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			db.DeleteTask(uint64(id))
		}
	},
}
