package gocard

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCardBrand_Validate(t *testing.T) {
	t.Parallel()

	masterCard := SupportedCardBrands[MasterCard]
	require.True(t, masterCard.Validate("5555555555554444"))
	require.False(t, masterCard.Validate("4242424242424242"))
}

func TestIsCardBrandSupported(t *testing.T) {
	t.Parallel()

	require.True(t, IsCardBrandSupported(JCB))
	require.False(t, IsCardBrandSupported(Unknown))
	require.False(t, IsCardBrandSupported("..."))
}
