package homework

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax2Numbers(t *testing.T) {
	a := []int{2,1,3,4}
	maxNumber := Max2Numbers(a)
	assert.NotEmpty(t, maxNumber)
	fmt.Println(maxNumber)
}
