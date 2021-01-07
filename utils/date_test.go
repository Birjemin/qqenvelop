package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(int(time.Now().Unix()), GetCurrTime())
}

func TestGetDateNum(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(8, len(GetDateNum()))
}
