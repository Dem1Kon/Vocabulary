package training

import (
	"github.com/spf13/cobra"
	"log"
	"vocabulary/internal/json"
	"vocabulary/internal/knowledge/test"
)

// TrainingCmd represents the training command
var TrainingCmd = &cobra.Command{
	Use:   "training",
	Short: "train of your knowledge",
	Long: `This command calls a knowledge training
without influence on word's rating'.`,
	Example: `vocabulary training
vocabulary training -m New -a 15`,
	Run: func(cmd *cobra.Command, args []string) {
		test.IsTrainMode = true

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

		test.Testing(Json, mode, amount)
	},
}

func init() {
	TrainingCmd.Flags().BoolVarP(&test.IsNative, "native", "n", false, "Training with native words")
	TrainingCmd.Flags().BoolVarP(&test.IsForeign, "foreign", "f", false, "Training with foreign words")
	TrainingCmd.PersistentFlags().IntP("amount", "a", 10, "Flag for set amount of words")
	TrainingCmd.Flags().StringP("mode", "m", "Any", "Flag for set mode[Level of word's knowledge]")
}
