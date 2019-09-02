package luhn

import "strconv"

// CheckLuhn checks validity of credit card number using Luhn algorithm
// Reference: https://en.wikipedia.org/wiki/Luhn_algorithm
func CheckLuhn(cardNumber string) (bool, error) {
	var sum int
	var alternate bool

	numberLen := len(cardNumber)

	for i := numberLen - 1; i >= 0; i-- {
		mod, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false, err
		}
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate
		sum += mod
	}

	return sum%10 == 0, nil
}
