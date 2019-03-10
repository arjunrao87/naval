package main

import (
	"fmt"

	"github.com/arjunrao87/naval/commands"
)

func main() {
	err := commands.NavalCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
