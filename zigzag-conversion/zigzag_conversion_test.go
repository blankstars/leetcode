package zigzag_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (string, int)string) {
	assert.Equal(t, "PAHNAPLSIIGYIR", f("PAYPALISHIRING", 3))
}

func Test(t *testing.T) {
	testCase(t, convert)
}

