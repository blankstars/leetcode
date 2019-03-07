package len_of_longest_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (string)int) {
	assert.Equal(t, 3, f("abcabcbb"))
	assert.Equal(t, 1, f(" "))
	assert.Equal(t, 2, f("au"))
	assert.Equal(t, 2, f("aab"))
	assert.Equal(t, 3, f("dvdf"))
	assert.Equal(t, 0, f(""))
}

func Test(t *testing.T) {
	testCase(t, lengthOfLongestSubstring)
}
