package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")
var newLine rune = 10

func isFirstElementValid(element rune) error {
	if unicode.IsLetter(element) || element == newLine {
		return nil
	}
	return ErrInvalidString
}

func Unpack(inputString string) (string, error) {
	if inputString == "" {
		return inputString, nil
	}

	result := strings.Builder{}
	previousValue := ' '

	for index, value := range inputString {
		if index == 0 {
			if err := isFirstElementValid(value); err != nil {
				return "", err
			}
		}

		switch {
		case unicode.IsDigit(value):
			if unicode.IsDigit(previousValue) {
				return "", ErrInvalidString
			}

			count, _ := strconv.Atoi(string(value))
			if count-1 > 0 {
				result.WriteString(strings.Repeat(string(previousValue), count-1))
			} else {
				newResult := result.String()[:len(result.String())-1]
				result.Reset()
				result.WriteString(newResult)
			}
			previousValue = value

		case unicode.IsLetter(value) || value == newLine:
			result.WriteRune(value)
			previousValue = value
		default:
			return "", ErrInvalidString
		}
	}

	return result.String(), nil
}
