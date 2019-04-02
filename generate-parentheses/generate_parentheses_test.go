package  generate_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (int)[]string) {
	assert.Equal(t, []string{"((()))","(()())","(())()","()(())","()()()"}, f(3))
}

func Test(t *testing.T) {
	testCase(t, generateParenthesis)
}
