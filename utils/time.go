package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func CurrentDate() string {
	CurrentTime := time.Now()
	dateOnly := CurrentTime.Format("2006-01-02")
	return dateOnly
}

func DateOnly(date string) (string, error) {
	cutIndex := strings.Index(date, "T")

	if cutIndex == -1 {
		fmt.Printf("Character 'T' not found in the string.\n")
		err := errors.New("character 'T' not found in the string")
		return "", err
	}
	trimmedStr := date[:cutIndex]
	return trimmedStr, nil
}
