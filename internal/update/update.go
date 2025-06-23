/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"vocabulary/internal/json"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update word",
	Long:  `This command needs to update words.`,
	Example: `vocabulary update
vocabulary update [word to update (in any language)]
vocabulary update [word to update (in any language)] [new word] [new word's translate]`,
	Run: func(cmd *cobra.Command, args []string) {
		Json, err := json.Init()
		if err != nil {
			log.Fatalln(err)
		}
		switch len(args) {
		case 0:
			err := Json.Update(getAllArgs())
			if err != nil {
				log.Fatalln(err)
			}
		case 1:
			err = Json.Update(args[0], getTranslate())
			if err != nil {
				log.Fatalln(err)
			}
		case 3:
			err = Json.Update(args[0], []string{args[1], args[2]})
			if err != nil {
				log.Fatalln(err)
			}
		default:
			fmt.Println("Gde-to ti naebal!")
		}
	},
}

func init() {}

func getAllArgs() (string, []string) {
	toUpdate := getUpdateWord()
	newPair := getTranslate()
	return toUpdate, newPair
}

func getUpdateWord() string {
	var toUpdate string
	fmt.Println("Enter the word you need to update: ")
	fmt.Scan(&toUpdate)

	return toUpdate
}

func getTranslate() []string {
	TranslatePair := make([]string, 2)

	fmt.Println("Enter the word: ")
	fmt.Scan(&TranslatePair[0])

	fmt.Println("Enter the word's translate: ")
	fmt.Scan(&TranslatePair[1])

	return TranslatePair
}
