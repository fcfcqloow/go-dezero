package util

import (
	"encoding/gob"
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func SaveBinary(filename string, value interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("filed to create file: %w", err)
	}
	defer f.Close()

	if err := gob.NewEncoder(f).Encode(value); err != nil {
		return fmt.Errorf("cannnot encode: %w", err)
	}
	return nil
}

func LoadBinary(filename string, fileValue interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	if err := gob.NewDecoder(f).Decode(fileValue); err != nil {
		return fmt.Errorf("cannot decode: %w", err)
	}
	return nil
}
