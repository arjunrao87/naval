package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arjunrao87/naval/helpers"
	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Call a Naval quote randomly",
	Long:  ``,
	Run:   random,
}

func init() {}

func random(ccmd *cobra.Command, args []string) {
	var response []string
	response = helpers.Read()
	var index = getRandomIndex() % len(response)
	fmt.Println(response[index])
}

func getRandomIndex() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}
