package gocard

import (
	"regexp"
)

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
// Reference: https://gist.github.com/michaelkeevildown/9096cd3aac9029c4e6e05588448a8841
var SupportedCardBrands = map[string]CardBrand{
	AmericanExpress: {
		Code:  AmericanExpress,
		Name:  "American Express",
		Regex: regexp.MustCompile(`^3[47][0-9]{13}$`),
	},
	DinersClub: {
		Code:  DinersClub,
		Name:  "Diners Club",
		Regex: regexp.MustCompile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`),
	},
	Discover: {
		Code:  Discover,
		Name:  "Discover",
		Regex: regexp.MustCompile(`^65[4-9][0-9]{13}|64[4-9][0-9]{13}|6011[0-9]{12}|(622(?:12[6-9]|1[3-9][0-9]|[2-8][0-9][0-9]|9[01][0-9]|92[0-5])[0-9]{10})$`),
	},
	JCB: {
		Code:  JCB,
		Name:  "JCB",
		Regex: regexp.MustCompile(`^(?:2131|1800|35\d{3})\d{11}$`),
	},
	MasterCard: {
		Code:  MasterCard,
		Name:  "Master Card",
		Regex: regexp.MustCompile(`^5[1-5][0-9]{14}$`),
	},
	UnionPay: {
		Code:  UnionPay,
		Name:  "Union Pay",
		Regex: regexp.MustCompile(`^(62[0-9]{14,17})$`),
	},
	Visa: {
		Code:  Visa,
		Name:  "Visa",
		Regex: regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`),
	},
}

// UnknownCardBrand defines unknown card brand
var UnknownCardBrand = CardBrand{
	Code:  Unknown,
	Name:  "Unknown Card Brand",
	Regex: nil,
}

// Validate checks validity of a card number whether it matches regex or not
func (brand CardBrand) Validate(cardNumber string) bool {
	return brand.Regex.MatchString(cardNumber)
}

// IsCardBrandSupported returns true if card brand is supported
func IsCardBrandSupported(code string) (supported bool) {
	_, supported = SupportedCardBrands[code]
	return supported
}
