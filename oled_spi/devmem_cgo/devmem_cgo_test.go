package devmem_cgo

import (
	"testing"
)

func TestWritebit(t *testing.T) {
	Writebit(0x1c20890, 12, 1)
	Writebit(0x1c20890, 13, 0)
	Writebit(0x1c20890, 14, 0)
	
}

