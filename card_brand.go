package gocard

import "regexp"

// CardBrand defines card brand properties
type CardBrand struct {
	Code  string
	Name  string
	Regex *regexp.Regexp
}

// Constants of Card Brands
const (
	AmericanExpress = "american_express"
	DinersClub      = "diners_club"
	Discover        = "discover"
	JCB             = "jcb"
	MasterCard      = "master_card"
	UnionPay        = "union_pay"
	Visa            = "visa"
	Unknown         = "unknown"
)

// SupportedCardBrands defines a list of supported card brands
var SupportedCardBrands = map[string]CardBrand{
	AmericanExpress: {
		Code:  AmericanExpress,
		Name:  "American Express",
		Regex: regexp.MustCompile("^3[47][0-9]{13}$"),
	},
}

// UnknownCardBrand defines unknown card brand
var UnknownCardBrand = CardBrand{
	Code:  Unknown,
	Name:  "Unknown Card Brand",
	Regex: nil,
}

// IsCardBrandSupported returns true if card brand is supported
func IsCardBrandSupported(code string) (supported bool) {
	_, supported = SupportedCardBrands[code]
	return supported
}

// Validate checks validity of a card number whether it matches regex or not
func (brand CardBrand) Validate(cardNumber string) bool {
	return brand.Regex.MatchString(cardNumber)
}
