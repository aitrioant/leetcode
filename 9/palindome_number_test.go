package palindromeNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumberSum(t *testing.T) {
	isPalindrome := isPalindrome(101)

	assert.Equal(t, true, isPalindrome)
}

func TestPalindromeNumberSum2(t *testing.T) {
	isPalindrome := isPalindrome(-101)

	assert.Equal(t, false, isPalindrome)
}

func TestPalindromeNumberSum3(t *testing.T) {
	isPalindrome := isPalindrome(10)

	assert.Equal(t, false, isPalindrome)
}
