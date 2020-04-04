package main

import (
	"flag"

	// "time"

	"oled_spi/gpio"
	"oled_spi/oled"
)

//gpio map
//pe3  gpio131
//pe4  gpio132
//pe5  gpio133
//pe10 gpio138
//pe11 gpio139
const (
	PE3  = "gpio131"
	PE4  = "gpio132"
	PE5  = "gpio133"
	PE10 = "gpio138"
	PE11 = "gpio139"
)

var tag = flag.Bool("i", true, "-i=true or -i init gpio once")

func main() {

	defer gpio.CloseGpioFile() //at last close filehandle

	flag.Parse()
	if *tag {
		gpio.ExportGpio("131", "132", "133", "138", "139")
		gpiomode := make([]gpio.GPIO, 0)
		gpiomode = append(gpiomode,
			gpio.GPIO{"gpio131", "out"},
			gpio.GPIO{"gpio132", "out"},
			gpio.GPIO{"gpio133", "out"},
			gpio.GPIO{"gpio138", "out"},
			gpio.GPIO{"gpio139", "out"})
		gpio.InitGpio(gpiomode)
	}

	oled.Init()
	oled.ShowString(2, 3, "oled go")
	// for {
	// 	gpio.GpioNSetValue("gpio132", "1")
	// 	time.Sleep(time.Millisecond * 500)
	// 	gpio.GpioNSetValue("gpio132", "0")
	// 	time.Sleep(time.Millisecond * 500)
	// }
}
