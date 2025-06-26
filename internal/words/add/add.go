package add

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"vocabulary/internal/json"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds new words to your Vocabulary",
	Long:    `This command adds new words to your Vocabulary.`,
	Example: `vocabulary add house дом`,
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 {
			fmt.Println("If your unit from pair has more than one word you have to use this form 'your word'!")
			return
		}

		JSON, err := json.Init()
		if err != nil {
			log.Fatal(err)
		}

		err = JSON.Add(args[0], args[1])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {

}
