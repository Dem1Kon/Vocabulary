package test

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"strings"
	"vocabulary/internal/json"
)

var (
	isForeign bool
	isNative  bool
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
				fmt.Print("Translate '", b.English, "': ")
				return b, b.English, nil
			case 1:
				fmt.Print("Translate '", b.Russian, "': ")
				return b, b.Russian, nil
			}

		case isForeign && !isNative && (b.Status.Rate == mode || mode == "Any"):

			fmt.Print("Translate '", b.English, "': ")
			return b, b.English, nil
		case isNative && !isForeign && (b.Status.Rate == mode || mode == "Any"):

			fmt.Print("Translate '", b.Russian, "': ")
			return b, b.Russian, nil
		}
	}

	return nil, "", errors.New("could not find a matching pair")
}

func checkMatch(pair *json.Pair, matched string) bool {
	if pair == nil {
		return false
	}

	var answer string

	fmt.Scan(&answer)
	answer = strings.ToLower(answer)

	if (pair.English == matched && pair.Russian == answer) || (pair.Russian == matched && pair.English == answer) {
		fmt.Println("Well done")
		pair.Status.Attempts++
		pair.Status.Good++
		return true
	}

	fmt.Println("Wrong answer")
	pair.Status.Attempts++

	return false
}
