package gocard

import "strconv"

// CheckLuhn checks validity of credit card number using Luhn algorithm
// Reference: https://en.wikipedia.org/wiki/Luhn_algorithm
func CheckLuhn(cardNumber string) (bool, error) {
	var sum int
	var alt bool

	length := len(cardNumber)

	for i := length - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false, err
		}
		if alt {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alt = !alt
	}

	return sum%10 == 0, nil
}
