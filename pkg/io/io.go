package io

import (
	"fmt"
	"os"
)

func ReadExample() (string, error) {
	return Read("example")
}

func ReadProblem() (string, error) {
	return Read("problem")
}

func Read(name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v\n", err)
	}

	contentBytes, err := os.ReadFile(fmt.Sprintf("%s/%s.txt", wd, name))
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return "", err
	}

	return string(contentBytes), nil
}
