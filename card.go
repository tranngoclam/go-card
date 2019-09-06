package gocard

import (
	"strings"
	"time"
)

// Card defines card properties
type Card struct {
	Brand    CardBrand `json:"brand"`
	Number   string    `json:"number"`
	Name     string    `json:"name"`
	CVV      string    `json:"cvv"`
	ExpireAt time.Time `json:"expire_at"`
}

// Validate checks validity of a card
func (card Card) Validate() (valid bool, err error) {
	valid, err = CheckLuhn(card.Number)
	return valid, err
}

// IsExpired checks expiry of a card
func (card Card) IsExpired() bool {
	return time.Now().After(card.ExpireAt)
}

// StandardizeCardNumber standardizes card number by removing all spaces, hyphens
func StandardizeCardNumber(cardNumber string) string {
	cardNumber = strings.TrimSpace(cardNumber)
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	cardNumber = strings.ReplaceAll(cardNumber, "-", "")
	return cardNumber
}

// GetCardBrand returns card brand from a card number
func GetCardBrand(cardNumber string) CardBrand {
	cardNumber = StandardizeCardNumber(cardNumber)
	for _, brand := range SupportedCardBrands {
		if brand.Validate(cardNumber) {
			return brand
		}
	}

	return UnknownCardBrand
}
