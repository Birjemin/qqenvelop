package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMd5(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("900150983cd24fb0d6963f7d28e17f72", Md5ByByte([]byte("abc")))
	ast.Equal("900150983cd24fb0d6963f7d28e17f72", GetMD5String("abc"))

	billNo := Hex(10)
	log.Println("billNo: ", billNo)
	ast.Equal(32, len(billNo))

}
