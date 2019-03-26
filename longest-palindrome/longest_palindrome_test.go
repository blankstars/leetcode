package longest_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (string)string) {
	assert.Equal(t, "bb", f("cbbd"))
	assert.Equal(t, "bab", f("babad"))
	assert.Equal(t, "", f(""))
}

func Test(t *testing.T) {
	testCase(t, longestPalindrome)
}

