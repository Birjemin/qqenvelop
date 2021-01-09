package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRand(t *testing.T) {
	ast := assert.New(t)

	billNo := GenerateBillNo()
	ast.Equal(32, len(billNo))

	num := RandNum(100, 1000)
	ast.Equal(0, num/1000)
}
