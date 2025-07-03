package cmd

import (
	"os"
	"vocabulary/internal/knowledge/test"
	"vocabulary/internal/knowledge/training"
	"vocabulary/internal/words/add"
	"vocabulary/internal/words/remove"
	"vocabulary/internal/words/show"
	"vocabulary/internal/words/update"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vocabulary",
	Short: "Here you can training your Foreign Vocabulary",
	Long: `There is a list of words that can help to training your Foreign Vocabulary.
Enter 'Foreign' to show Foreign words to translate in Native
If you enter 'Native' from Native to Foreign.`,
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
	rootCmd.AddCommand(training.TrainingCmd)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
