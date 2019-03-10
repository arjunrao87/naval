package commands

import (
	"fmt"

	"github.com/arjunrao87/naval/helpers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of Naval's quotes",
	Long:  ``,
	Run:   list,
}

func init() {
}

func list(ccmd *cobra.Command, args []string) {
	var response []string
	response = helpers.Read()
	for index, element := range response {
		output := fmt.Sprintf("%d : %s", index+1, element)
		fmt.Println(output)
	}
}
