package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
)

var (
	FileOpenError   = errors.New("file open error")
	FileDecodeError = errors.New("file decode error")
	FileWriteError  = errors.New("file write error")
	FileEncodeError = errors.New("file encode error")
	NotFoundError   = errors.New("not found error")
)

type JSON struct {
	Amount int
	Pairs  map[int]*Pair
}

func Init() (*JSON, error) {
	file, err := os.OpenFile("vocabulary.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, FileOpenError
	}
	defer file.Close()
	Json := &JSON{Pairs: make(map[int]*Pair), Amount: 0}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Json)
	if err != nil {
		if !(errors.Is(err, io.EOF)) {
			return nil, FileDecodeError
		} else {
			err = nil
		}
	}
	return Json, err
}

func (J *JSON) Add(english, russian string) error {
	J.Pairs[J.Amount] = &Pair{Foreign: english, Translate: russian, Status: Status{Rate: "New"}}
	J.Amount++

	err := J.WriteToAFile()
	if err != nil {
		return FileWriteError
	}

	return nil
}

func (J *JSON) Remove(rm string) error {
	IsFound := false

	for a, b := range J.Pairs {
		if b.Foreign == rm || b.Translate == rm {
			IsFound = true
			delete(J.Pairs, a)
		}
	}

	err := J.WriteToAFile()
	if err != nil {
		return FileWriteError
	}

	if !IsFound {
		fmt.Println(NotFoundError)
		return NotFoundError
	}
	return nil
}

func (J *JSON) Update(rm string, append []string) error {
	err := J.Remove(rm)
	if err != nil {
		return err
	}

	err = J.Add(append[0], append[1])
	if err != nil {
		return err
	}

	return nil
}

func (J *JSON) Show() {
	if len(J.Pairs) == 0 {
		fmt.Println("No words found")
		return
	}

	for _, pair := range J.Pairs {
		pair.Rate()
		fmt.Printf("\t%s\t- \t%s \t\t", pair.Foreign, pair.Translate)
		switch pair.Status.Rate {
		case "New":
			color.Red("%s\n", pair.Status.Rate)
		case "Familiar":
			color.Yellow("%s\n", pair.Status.Rate)
		case "Known":
			color.Blue("%s\n", pair.Status.Rate)
		case "Well-known":
			color.Green("%s\n", pair.Status.Rate)
		}
	}
	err := J.WriteToAFile()
	if err != nil {
		log.Fatalln(err)
	}
}

func (J *JSON) WriteToAFile() error {
	file, err := os.OpenFile("vocabulary.json", os.O_WRONLY, 0666)
	if err != nil {
		return FileOpenError
	}
	defer file.Close()

	err = CleanFile(file)

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(J)
	if err != nil {
		return FileEncodeError
	}
	return nil
}

func CleanFile(file *os.File) error {
	file.Seek(0, 0)

	file.Truncate(0)

	err := file.Truncate(0)
	if err != nil {
		return FileWriteError
	}
	return nil
}
