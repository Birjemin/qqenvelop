package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRand(t *testing.T) {
	ast := assert.New(t)
	num := RandNum(100, 1000)
	ast.Equal(0, num/1000)
}
