package test

import (
	"github.com/spf13/cobra"
	"log"
	"vocabulary/internal/json"
)

var (
	amountFlag = 0
	modeFlag   string
)

// TestCmd represents the test command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Testing knowledge",
	Long:  `This command calls a knowledge test.`,
	Example: `vocabulary test
vocabulary test -m New -a 15`,
	Run: func(cmd *cobra.Command, args []string) {
		Json, err := json.Init()
		if err != nil {
			log.Fatalln(err)
		}
		testing(Json, modeFlag, amountFlag)
	},
}

func init() {
	TestCmd.PersistentFlags().IntVarP(&amountFlag, "amount", "a", 0, "Flag for set amount of words")
	TestCmd.PersistentFlags().StringVarP(&modeFlag, "mode", "m", "", "Flag for set mode[Level of word's knowledge]")
}

func testing(J *json.JSON, mode string, amount int) {

}
