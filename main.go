package main

import (
	"strconv"
	"strings"
)

func main() {
	numberOfCards := 6
	cards := []string{
		"4253-6258-796-15787",
		"4424444424442444",
		"5122-2368-7954 - 3214",
		"44244x4424442444",
		"0525362587961578",
		"5123 - 3567 - 8912 - 3456",
	}

	validateCreditCard(numberOfCards, cards)
}

func validateCreditCard(numberOfCards int, cards []string) {
	for _, card := range cards {
		if checkCard(card) {
			println("valid")
		} else {
			println("invalid")
		}

	}
}

func checkCard(card string) bool {
	if len(card) == 19 {
		groups := strings.Split(card, "-")

		for _, group := range groups {
			if len(group) != 4 {
				return false
			}
		}
	}

	cardDigits := strings.ReplaceAll(card, "-", "")

	if len(cardDigits) != 16 {
		return false
	}

	for _, digit := range cardDigits {
		if digit < '0' || digit > '9' {
			return false
		}
	}

	digit, _ := strconv.Atoi(string(cardDigits[0]))
	if digit != 4 && digit != 5 && digit != 6 {
		return false
	}

	counter := 1
	for i := 0; i < len(cardDigits)-1; i++ {
		if counter == 4 {
			return false
		}
		if cardDigits[i] == cardDigits[i+1] {
			counter++
		} else {
			counter = 1
		}
	}

	return true
}

