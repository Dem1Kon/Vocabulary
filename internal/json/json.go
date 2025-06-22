package json

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

var (
	FileOpenError   = errors.New("file open error")
	FileDecodeError = errors.New("file decode error")
	FileWriteError  = errors.New("file write error")
	FileReadError   = errors.New("file read error")
	FileEncodeError = errors.New("file encode error")
)

type IJson interface {
	Add() error
	Remove() error
	Update() error
	writeToAFile() error
}

type JSON struct {
	English []string
	Russian []string
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
			err = nil
		}
	}
	return Json, err
}

func (J *JSON) Add(english, russian string) error {
	J.English = append(J.English, english)
	J.Russian = append(J.Russian, russian)

	err := J.writeToAFile()
	if err != nil {
		return FileWriteError
	}

	return nil
}

func (J *JSON) Remove() error {
	//TODO implement me
	panic("implement me")
}

func (J *JSON) Update() error {
	//TODO implement me
	panic("implement me")
}

func (J *JSON) writeToAFile() error {
	file, err := os.OpenFile("vocabulary.json", os.O_WRONLY, 0666)
	if err != nil {
		return FileOpenError
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(J)
	if err != nil {
		return FileEncodeError
	}
	return nil
}
