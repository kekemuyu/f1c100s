package devmem_cgo

//#include"devmem.h"
import "C"

func Openfile(){
   C.Openfile()
}

func Closefile(){
    C.Closefile()
}
func Writebit(target int32, bitsize int, value byte) {
	C.Writebit(C.long(target), C.int(bitsize), C.char(value))
}

