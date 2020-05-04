package devmem_cgo

//#include"devmem.h"
import "C"

func Openfile(target int32) {
	C.Openfile(C.long(target))
}

func Closefile() {
	C.Closefile()
}
func Writebit(offset int, bitsize int, value byte) {
	C.Writebit(C.int(offset), C.int(bitsize), C.char(value))
}
