package cmd

import (
	"os"
	"vocabulary/internal/knowledgeCheck/test"
	"vocabulary/internal/words/add"
	"vocabulary/internal/words/remove"
	"vocabulary/internal/words/show"
	"vocabulary/internal/words/update"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vocabulary",
	Short: "Here you can train your English Vocabulary",
	Long: `There is a list of words that can help to train your English Vocabulary.
Enter 'English' to show English words to translate in Russian
If you enter 'Russian' from Russian to English.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(show.ShowCmd)
	rootCmd.AddCommand(update.UpdateCmd)
	rootCmd.AddCommand(test.TestCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
