package utils

import "os"

func ReadFile(filename string) (string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(f), nil
}
