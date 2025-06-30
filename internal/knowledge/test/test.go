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
	isForeign   bool
	isNative    bool
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
	TestCmd.PersistentFlags().IntP("amount", "a", 10, "Flag for set amount of words")
	TestCmd.Flags().StringP("mode", "m", "Any", "Flag for set mode[Level of word's knowledge]")
	TestCmd.Flags().BoolVarP(&isForeign, "foreign", "f", false, "Show only foreign words")
	TestCmd.Flags().BoolVarP(&isNative, "native", "n", false, "Show only native words")
}

func testing(J *json.JSON, mode string, amount int) {
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
		case !isForeign && !isNative && (b.Status.Rate == mode || mode == "Any"):

			r := rand.Intn(2)
			switch r {
			case 0:
				fmt.Print("Translate '", b.Foreign, "': ")
				return b, b.Foreign, nil
			case 1:
				fmt.Print("Translate '", b.Translate, "': ")
				return b, b.Translate, nil
			}

		case isForeign && !isNative && (b.Status.Rate == mode || mode == "Any"):

			fmt.Print("Translate '", b.Foreign, "': ")
			return b, b.Foreign, nil
		case isNative && !isForeign && (b.Status.Rate == mode || mode == "Any"):

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
		fmt.Println("Well done")
		if !IsTrainMode {
			pair.Status.Attempts++
			pair.Status.Good++
		}
		time.Sleep(time.Second)
		utils.ClearTerminal()
		return true
	}

	fmt.Println("Wrong answer")
	time.Sleep(time.Second)
	utils.ClearTerminal()
	pair.Status.Attempts++

	return false
}
