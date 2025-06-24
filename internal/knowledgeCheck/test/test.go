package test

import (
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"vocabulary/internal/json"
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

		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			log.Fatalln(err)
		}

		amount, err := cmd.Flags().GetInt("amount")
		if err != nil {
			log.Fatalln(err)
		}

		testing(Json, mode, amount)
	},
}

func init() {
	TestCmd.PersistentFlags().IntP("amount", "a", 0, "Flag for set amount of words")
	TestCmd.Flags().StringP("mode", "m", "", "Flag for set mode[Level of word's knowledge]")
	TestCmd.Flags().BoolP("foreign", "f", false, "Show only foreign words")
	TestCmd.Flags().BoolP("native", "n", false, "Show only native words")

}

func testing(J *json.JSON, mode string, amount int) {

	isForeign, _ := TestCmd.PersistentFlags().GetBool("foreign")
	isNative, _ := TestCmd.PersistentFlags().GetBool("native")

	for i := 0; i < amount; i++ {
		switch {
		default:
			ChoiceTheWord(append(J.Foreign, J.Native...), mode)
		case isForeign && !isNative:
			ChoiceTheWord(J.Foreign, mode)
		case isNative && !isForeign:
			ChoiceTheWord(J.Native, mode)
		}
	}
}

func ChoiceTheWord(words []string, mode string) {
	for {
		r := rand.Intn(len(words))
		switch mode {
		case "New":
		case "Familiar":
		case "Known":
		case "Well-known":
		}
	}

}
