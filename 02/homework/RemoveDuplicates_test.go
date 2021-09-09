package homework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	a := []int{1,2,5,2,6,2,5}
	result := RemoveDuplicates(a)
	assert.NotEmpty(t, result)
}
