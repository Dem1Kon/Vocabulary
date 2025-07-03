package test

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"os"
	"time"
	"vocabulary/internal/json"
	"vocabulary/utils"
)

var (
	IsForeign   bool
	IsNative    bool
	IsTrainMode bool
)

// TestCmd represents the test command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "testing knowledge",
	Long:  `This command calls a knowledge test.`,
	Example: `vocabulary test
vocabulary test -m New -a 15`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ClearTerminal()

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

		Testing(Json, mode, amount)
	},
}

func init() {
	TestCmd.PersistentFlags().IntP("amount", "a", 10, "Flag for set amount of words")
	TestCmd.Flags().StringP("mode", "m", "Any", "Flag for set mode[Level of word's knowledge]")
	TestCmd.Flags().BoolVarP(&IsForeign, "foreign", "f", false, "Show only foreign words")
	TestCmd.Flags().BoolVarP(&IsNative, "native", "n", false, "Show only native words")
}

func Testing(J *json.JSON, mode string, amount int) {
	if len(J.Pairs) < amount {
		amount = len(J.Pairs)
	}
	for i := 0; i < amount; i++ {
		pair, matched, err := choiceAWord(J, mode)
		if err != nil {
			log.Fatalln(err)
		}

		checkMatch(pair, matched)
	}
	err := J.WriteToAFile()
	if err != nil {
		log.Fatalln(err)
	}
}

func choiceAWord(Json *json.JSON, mode string) (*json.Pair, string, error) {
	var b *json.Pair
	for _, b = range Json.Pairs {
		b.Rate()
		switch {
		case !IsForeign && !IsNative && (b.Status.Rate == mode || mode == "Any"):

			r := rand.Intn(2)
			switch r {
			case 0:
				fmt.Print("Translate '", b.Foreign, "': ")
				return b, b.Foreign, nil
			case 1:
				fmt.Print("Translate '", b.Translate, "': ")
				return b, b.Translate, nil
			}

		case IsForeign && !IsNative && (b.Status.Rate == mode || mode == "Any"):

			fmt.Print("Translate '", b.Foreign, "': ")
			return b, b.Foreign, nil
		case IsNative && !IsForeign && (b.Status.Rate == mode || mode == "Any"):

			fmt.Print("Translate '", b.Translate, "': ")
			return b, b.Translate, nil
		}

	}

	return nil, "", errors.New("could not find a matching pair")
}

func checkMatch(pair *json.Pair, matched string) bool {
	if pair == nil {
		return false
	}

	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	answer = scanner.Text()

	if (pair.Foreign == matched && pair.Translate == answer) || (pair.Translate == matched && pair.Foreign == answer) {

		if !IsTrainMode {
			pair.Status.Attempts++
			pair.Status.Good++
		}
		fmt.Println("Well done")
		waiting(pair)

		return true
	}

	fmt.Println("Wrong answer")
	waiting(pair)

	if !IsTrainMode {
		pair.Status.Attempts++
	}

	return false
}

func waiting(pair *json.Pair) {
	if IsTrainMode {
		fmt.Println(pair.Foreign, "-", pair.Translate)
	}

	time.Sleep(time.Second * 2)
	utils.ClearTerminal()
}
