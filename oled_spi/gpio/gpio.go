package gpio

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const gpioExportFile = "/sys/class/gpio/export"
const gpioDirFile = "/sys/class/gpio/gpioN/direction"
const gpioValueFile = "/sys/class/gpio/gpioN/value"

var GpioFileHandle map[string]*os.File

func init() {
	GpioFileHandle = make(map[string]*os.File)
}

//gpioNum like 131,132....
func ExportGpio(gpioNum ...string) error {
	if len(gpioNum) <= 0 {
		return errors.New("len(gpioNum)<=0")
	}

	file, err := os.OpenFile(gpioExportFile, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range gpioNum {
		file.WriteString(v)
	}
	return nil
}

type GPIO struct {
	GpioN string //"gpio131","gpio132"...
	Dir   string //"out" or "in"
}

//gpio linke gpio131,gpio132...
func InitGpio(gpios []GPIO) error {
	if len(gpios) <= 0 {
		return errors.New("len(gpios)<=0")
	}

	gdFile := ""
	for _, v := range gpios {
		gdFile = strings.Replace(gpioDirFile, "gpioN", v.GpioN, -1)

		file, err := os.OpenFile(gdFile, os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer file.Close()

		file.WriteString(v.Dir)
	}
	return nil
}

func GpioNSetValue(gpioN string, value string) error {
	// gvFile := ""
	// gvFile = strings.Replace(gpioValueFile, "gpioN", gpioN, -1)
	// file, err := os.OpenFile(gvFile, os.O_WRONLY, 0666)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()
	if file, ok := GpioFileHandle[gpioN]; ok {
		file.WriteString(value)
	}
	return nil
}

func OpenGpioFile(gpioN string) {
	gvFile := strings.Replace(gpioValueFile, "gpioN", gpioN, -1)
	file, err := os.OpenFile(gvFile, os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	GpioFileHandle[gpioN] = file
}

func CloseGpioFile() {
	for _, v := range GpioFileHandle {
		v.Close()
	}
}

