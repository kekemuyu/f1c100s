package main

import (
	"oled_spi/oled_cgo"
)

func main() {
	oled_cgo.Init()
	oled_cgo.ShowString(2, 3, "oled go")
	//	oled_cgo.Close()
}
