package entity

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type CPF struct {
	value string
	valid bool
}

func NewCPF(value string) Document {
	doc := CPF{}
	value, valid := doc.validate(value)
	doc.valid = valid
	doc.value = value
	return doc
}

func (c CPF) Valid() bool {
	return c.valid
}

func (c CPF) Value() string {
	return c.value
}

func (c *CPF) validate(value string) (string, bool) {
	value = c.extractDigits(value)
	if c.invalidLength(value) {
		return "", false
	}
	if c.isBlocked(value) {
		return "", false
	}
	digit1, err := c.calculateDigit(value, factorDigit1, maxDigits1)
	if err != nil {
		log.Println(fmt.Sprintf("could not calculate digit factor1: %s", err.Error()))
	}
	digit2, err := c.calculateDigit(value, factorDigit2, maxDigits2)
	if err != nil {
		log.Println(fmt.Sprintf("could not calculate digit factor2: %s", err.Error()))
	}
	calculateDigits := fmt.Sprintf("%d%d", digit1, digit2)
	return value, c.getCheckDigit(value) == calculateDigits
}

func (c CPF) extractDigits(value string) string {
	re := regexp.MustCompile(`[\d]*`)
	return strings.Join(re.FindAllString(value, -1), "")
}

func (c CPF) invalidLength(value string) bool {
	return len(value) != 11
}

func (c *CPF) isBlocked(value string) bool {
	digit := string(value[0])
	for _, it := range strings.Split(value, "") {
		if it != digit {
			return false
		}
	}
	return true
}

func (c *CPF) calculateDigit(value string, factor int, maxDigits int) (int, error) {
	var total int
	digitsSlice := strings.Split(value, "")[:maxDigits]
	for _, d := range digitsSlice {
		digit, err := strconv.Atoi(d)
		if err != nil {
			return 0, fmt.Errorf("could not parse string to int: %w", err)
		}
		total += digit * factor
		factor--
	}
	if total%11 < 2 {
		return 0, nil
	} else {
		return 11 - (total % 11), nil
	}
}

func (c *CPF) getCheckDigit(value string) string {
	return value[9:]
}
