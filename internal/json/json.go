package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	FileOpenError   = errors.New("file open error")
	FileDecodeError = errors.New("file decode error")
	FileWriteError  = errors.New("file write error")
	FileReadError   = errors.New("file read error")
	FileEncodeError = errors.New("file encode error")
	NotFoundError   = errors.New("not found error")
)

type JSON struct {
	English []string
	Russian []string
	Status  []string
}

func Init() (*JSON, error) {
	file, err := os.OpenFile("vocabulary.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, FileOpenError
	}
	defer file.Close()
	Json := &JSON{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Json)
	if err != nil {
		if !(errors.Is(err, io.EOF)) {
			return nil, FileDecodeError
		} else {
			Json.Status = make([]string, 0, 10)
			Json.English = make([]string, 0, 10)
			Json.Russian = make([]string, 0, 10)
			err = nil
		}
	}
	return Json, err
}

func (J *JSON) Add(english, russian string) error {
	J.English = append(J.English, english)
	J.Russian = append(J.Russian, russian)
	J.Status = append(J.Status, "New")

	err := J.writeToAFile()
	if err != nil {
		return FileWriteError
	}

	return nil
}

func (J *JSON) Remove(rm string) error {
	IsFound := false

	for a, b := range J.English {
		if J.Russian[a] == rm || b == rm {
			IsFound = true
			J.English = append(J.English[:a], J.English[a+1:]...)
			J.Russian = append(J.Russian[:a], J.Russian[a+1:]...)
			J.Status = append(J.Status[:a], J.Status[a+1:]...)
			break
		}
	}

	err := J.writeToAFile()
	if err != nil {
		return FileWriteError
	}

	if !IsFound {
		return NotFoundError
	}
	return nil
}

func (J *JSON) Update() error {
	//TODO implement me
	panic("implement me")
}

func (J *JSON) Show() {
	for i := 0; i < len(J.English); i++ {
		fmt.Printf("%s - %s Status: %s\n", J.English[i], J.Russian[i], J.Status[i])
	}
}

func (J *JSON) writeToAFile() error {
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
