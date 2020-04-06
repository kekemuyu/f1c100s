package gpio_mmap

import (
	"oled_spi/devmem"
)

const (
	Gpiobase   = 0x01C20800
	PEcfgbase  = Gpiobase + 0x90 //pe config register
	PEdatabase = Gpiobase + 0xa0 //pe data register
)

func InitPE() {
	devmem.WriteBit(PEcfgbase, 12, 1)
	devmem.WriteBit(PEcfgbase, 13, 0)
	devmem.WriteBit(PEcfgbase, 14, 0)
}

func GpioNSetValue() {

}

func GpioNSetBit(gpioAddr int64, offbitsize int, value byte) {
	devmem.WriteBit(PEdatabase, offbitsize, value)
}

