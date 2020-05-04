package main

import (
	"fmt"
	"os"
	"syscall"
)

const (
	devName    = "/dev/mem"
	gpiobase   = 0x01C20800
	gedatabase = gpiobase + 4*0x24 + 0x10
	gecfgbase  = gpiobase + 0x90
)

// 实现了对gpio131---PE3的data控制
func main() {
	fd, err := os.OpenFile(devName, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()

	gpioData, err := syscall.Mmap(int(fd.Fd()), 0x1c208a0&(^4095), 4096,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("len(gpioData):", len(gpioData))

	if len(gpioData) < 4096 {
		fmt.Println("len(gpioData) < 4096")
		return
	}

	fmt.Println(gpioData)
	fmt.Println(gpioData[2208])
	gpioData[2208] = 0xff

	err = syscall.Munmap(gpioData)
	fmt.Println(err)
}
