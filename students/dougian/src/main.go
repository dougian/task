package main

import (
	"github.com/gophercises/task/students/dougian/src/cmd"
	"github.com/gophercises/task/students/dougian/src/db"
	"fmt"
	"os"
)

func main() {
	err := db.InitBolt()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	cmd.Execute()

}

