package datafile

import (
	"encoding/json"
	"os"
)

func Read(name string, v any) (err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(v)
	return
}

func Write(name string, v any) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(v)
	return
}
