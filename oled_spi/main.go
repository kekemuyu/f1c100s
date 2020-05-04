package main

import (
	"fmt"
	"oled_spi/oled_cgo"
	"time"
)

func main() {
	oled_cgo.Init()
	fmt.Println("begintime:", time.Now())
	var flag = false
	for i := 0; i < 100; i++ {
		if flag {
			flag = false
			oled_cgo.Unclear()
		} else {
			flag = true
			oled_cgo.Clear()
		}

	}
	fmt.Println("endtime:", time.Now())
	//	oled_cgo.Close()
}
