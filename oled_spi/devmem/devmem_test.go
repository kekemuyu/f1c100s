package devmem

import (
	"testing"
)

func TestWriteBit(t *testing.T) {

	t.Log(WriteBit(0x1c20890, 12, 1))
	// t.Log(WriteBit(0x1c20890, 13, 0))
	 //t.Log(WriteBit(0x1c20890, 14, 0))
}

