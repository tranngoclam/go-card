package luhn

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckLuhnSuccess(t *testing.T) {
	t.Parallel()

	cardNumbers := []string{
		"4242424242424242",
		"4000056655665556",
		"5555555555554444",
		"2223003122003222",
		"5200828282828210",
		"5105105105105100",
		"378282246310005",
		"371449635398431",
		"6011111111111117",
		"6011000990139424",
		"30569309025904",
		"38520000023237",
		"3566002020360505",
		"6200000000000005",
	}

	for _, cardNumber := range cardNumbers {
		t.Run(cardNumber, func(t *testing.T) {
			t.Parallel()

			valid, err := CheckLuhn(cardNumber)
			require.NoError(t, err)
			require.True(t, valid)
		})
	}
}
