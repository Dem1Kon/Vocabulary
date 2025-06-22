package remove

import (
	"github.com/spf13/cobra"
	"log"
	"vocabulary/internal/json"
)

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "This command removes a word from your vocabulary",
	Example: `remove word
remove слово`,
	Aliases: []string{"rm"},
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		json, err := json.Init()
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Remove(args[0])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
