package utils

import (
	"log"
	"os"
)

func Logger() *os.File {
	file, err := os.OpenFile("logger.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(file)

	return file
}
