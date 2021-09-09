package homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxLengthElement(t *testing.T) {
	a := []string{"aba", "aa", "ad", "c", "vcd"}
	result := FindMaxLengthElement(a)
	assert.NotEmpty(t, result)
}