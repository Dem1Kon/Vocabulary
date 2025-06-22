/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"vocabulary/cmd"
	"vocabulary/utils"
)

func main() {
	defer utils.Logger().Close()

	cmd.Execute()
}
