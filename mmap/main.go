package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	f, err := os.OpenFile("mmap.bin", os.O_RDWR|os.O_CREATE, 0644)

	if nil != err {

		log.Fatalln(err)

	}

	// extend file

	if _, err := f.WriteAt([]byte{byte(0)}, 1<<8); nil != err {

		log.Fatalln(err)

	}

	data, err := syscall.Mmap(int(f.Fd()), 0, 1<<8, syscall.PROT_WRITE, syscall.MAP_SHARED)

	if nil != err {

		log.Fatalln(err)

	}

	if err := f.Close(); nil != err {

		log.Fatalln(err)

	}

	for i, v := range []byte("hello syscall") {

		data[i] = v

	}

	if err := syscall.Munmap(data); nil != err {

		log.Fatalln(err)

	}
}
