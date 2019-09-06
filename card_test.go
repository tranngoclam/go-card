package gocard

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCard_IsExpired(t *testing.T) {
	t.Parallel()

	expiredCard := Card{
		ExpireAt: time.Now().AddDate(0, 0, -1),
	}
	require.True(t, expiredCard.IsExpired())

	validCard := Card{
		ExpireAt: time.Now().AddDate(1, 0, 0),
	}
	require.False(t, validCard.IsExpired())
}

func TestStandardizeCardNumber(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Actual   string
		Expected string
	}{
		{"4242 4242 4242 4242", "4242424242424242"},
		{"4242  4242   4242 4242", "4242424242424242"},
		{"4242-4242-4242-4242", "4242424242424242"},
		{" 4242 4242 4242 4242 ", "4242424242424242"},
		{" 4242 -4242-4242- 4242 ", "4242424242424242"},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.Actual, func(t *testing.T) {
			t.Parallel()

			standardizedCardNumber := StandardizeCardNumber(tc.Actual)
			require.Equal(t, tc.Expected, standardizedCardNumber)
		})
	}
}

func TestGetCardBrand(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Number string
		Brand  string
	}{
		{"4242424242424242", Visa},
		{"4000056655665556", Visa},
		{"5555555555554444", MasterCard},
		{"5200828282828210", MasterCard},
		{"5105105105105100", MasterCard},
		{"378282246310005", AmericanExpress},
		{"371449635398431", AmericanExpress},
		{"6011111111111117", Discover},
		{"6011000990139424", Discover},
		{"30569309025904", DinersClub},
		{"38520000023237", DinersClub},
		{"3566002020360505", JCB},
		{"6200000000000005", UnionPay},
		{"0001000200030004", Unknown},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.Number, func(t *testing.T) {
			t.Parallel()

			actualBrand := GetCardBrand(tc.Number)
			expectedBrand, ok := SupportedCardBrands[tc.Brand]
			if ok {
				require.Equal(t, expectedBrand, actualBrand)
			} else {
				require.Equal(t, UnknownCardBrand, actualBrand)
			}
		})
	}
}
