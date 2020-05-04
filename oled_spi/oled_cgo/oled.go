package oled_cgo

import (
	"oled_spi/devmem_cgo"
	"time"
)

const (
	OLED_CLK = 3
	OLED_DI  = 4
	OLED_RST = 5
	OLED_DC  = 10
	OLED_CS  = 11

	OLED_CMD  = 0 //写命令
	OLED_DATA = 1 //写数据

	SIZE       = 16
	Max_Column = 128
	Max_Row    = 64

	GPIO_BASE = 0x1c20800
)

func Init() {
	devmem_cgo.Openfile(GPIO_BASE)
	devmem_cgo.Writebit(0x90, 12, 1) //pe3 out
	devmem_cgo.Writebit(0x90, 13, 0)
	devmem_cgo.Writebit(0x90, 14, 0)

	devmem_cgo.Writebit(0x90, 16, 1) //pe4 out
	devmem_cgo.Writebit(0x90, 17, 0)
	devmem_cgo.Writebit(0x90, 18, 0)

	devmem_cgo.Writebit(0x90, 20, 1) //pe5 out
	devmem_cgo.Writebit(0x90, 21, 0)
	devmem_cgo.Writebit(0x90, 22, 0)

	devmem_cgo.Writebit(0x94, 8, 1) //pe10 out
	devmem_cgo.Writebit(0x94, 9, 0)
	devmem_cgo.Writebit(0x94, 10, 0)

	devmem_cgo.Writebit(0x94, 12, 1) //pe11 out
	devmem_cgo.Writebit(0x94, 13, 0)
	devmem_cgo.Writebit(0x94, 14, 0)

	SetRst()
	time.Sleep(time.Millisecond * 100)
	ClrRst()
	time.Sleep(time.Millisecond * 100)
	SetRst()

	WriteByte(0xAE, OLED_CMD) //--turn off oled panel
	WriteByte(0x00, OLED_CMD) //---set low column address
	WriteByte(0x10, OLED_CMD) //---set high column address
	WriteByte(0x40, OLED_CMD) //--set start line address  Set Mapping RAM Display Start Line (0x00~0x3F)
	WriteByte(0x81, OLED_CMD) //--set contrast control register
	WriteByte(0x66, OLED_CMD) // Set SEG Output Current Brightness
	WriteByte(0xA1, OLED_CMD) //--Set SEG/Column Mapping     0xa0左右反置 0xa1正常
	WriteByte(0xC8, OLED_CMD) //Set COM/Row Scan Direction   0xc0上下反置 0xc8正常
	WriteByte(0xA6, OLED_CMD) //--set normal display
	WriteByte(0xA8, OLED_CMD) //--set multiplex ratio(1 to 64)
	WriteByte(0x3f, OLED_CMD) //--1/64 duty
	WriteByte(0xD3, OLED_CMD) //-set display offset    Shift Mapping RAM Counter (0x00~0x3F)
	WriteByte(0x00, OLED_CMD) //-not offset
	WriteByte(0xd5, OLED_CMD) //--set display clock divide ratio/oscillator frequency
	WriteByte(0x80, OLED_CMD) //--set divide ratio, Set Clock as 100 Frames/Sec
	WriteByte(0xD9, OLED_CMD) //--set pre-charge period
	WriteByte(0xF1, OLED_CMD) //Set Pre-Charge as 15 Clocks & Discharge as 1 Clock
	WriteByte(0xDA, OLED_CMD) //--set com pins hardware configuration
	WriteByte(0x12, OLED_CMD)
	WriteByte(0xDB, OLED_CMD) //--set vcomh
	WriteByte(0x40, OLED_CMD) //Set VCOM Deselect Level
	WriteByte(0x20, OLED_CMD) //-Set Page Addressing Mode (0x00/0x01/0x02)
	WriteByte(0x02, OLED_CMD) //
	WriteByte(0x8D, OLED_CMD) //--set Charge Pump enable/disable
	WriteByte(0x14, OLED_CMD) //--set(0x10) disable
	WriteByte(0xA4, OLED_CMD) // Disable Entire Display On (0xa4/0xa5)
	WriteByte(0xA6, OLED_CMD) // Disable Inverse Display On (0xa6/a7)
	WriteByte(0xAF, OLED_CMD) //--turn on oled panel

	WriteByte(0xAF, OLED_CMD) /*display ON*/
	Clear()
	SetPos(0, 0)
}

func SetClk() {
	devmem_cgo.Writebit(0xa0, 3, 1)
}
func ClrClk() {
	devmem_cgo.Writebit(0xa0, 3, 0)
}

func SetDi() {
	devmem_cgo.Writebit(0xa0, 4, 1)
}

func ClrDi() {
	devmem_cgo.Writebit(0xa0, 4, 0)
}

func SetRst() {
	devmem_cgo.Writebit(0xa0, 5, 1)
}

func ClrRst() {
	devmem_cgo.Writebit(0xa0, 5, 0)
}
func SetDc() {
	devmem_cgo.Writebit(0xa0, 10, 1)
}

func ClrDc() {
	devmem_cgo.Writebit(0xa0, 10, 0)
}
func SetCs() {
	devmem_cgo.Writebit(0xa0, 11, 1)
}

func ClrCs() {
	devmem_cgo.Writebit(0xa0, 11, 0)
}

func WriteByte(dat, cmd byte) {
	if cmd != 0 {
		SetDc()
	} else {
		ClrDc()
	}
	ClrCs()

	for i := 0; i < 8; i++ {
		ClrClk()
		if dat&0x80 != 0 {
			SetDi()
		} else {
			ClrDi()
		}
		SetClk()
		dat <<= 1
	}
	SetCs()
	SetDc()
}

func SetPos(x, y byte) {
	WriteByte(0xb0+y, OLED_CMD)
	WriteByte(((x&0xf0)>>4)|0x10, OLED_CMD)
	WriteByte((x&0x0f)|0x01, OLED_CMD)
}

func Clear() {
	for i := byte(0); i < 8; i++ {
		WriteByte(0xb0+i, OLED_CMD) //设置页地址（0~7）
		WriteByte(0x00, OLED_CMD)   //设置显示位置—列低地址
		WriteByte(0x10, OLED_CMD)   //设置显示位置—列高地址
		for n := 0; n < 128; n++ {
			WriteByte(0, OLED_DATA) //更新显示
		}
	}
}

func Unclear() {
	for i := byte(0); i < 8; i++ {
		WriteByte(0xb0+i, OLED_CMD) //设置页地址（0~7）
		WriteByte(0x00, OLED_CMD)   //设置显示位置—列低地址
		WriteByte(0x10, OLED_CMD)   //设置显示位置—列高地址
		for n := 0; n < 128; n++ {
			WriteByte(1, OLED_DATA) //更新显示
		}
	}
}

func ShowChar(x, y, chr byte) {
	c := chr - ' ' //得到偏移后的值
	if x > Max_Column-1 {
		x = 0
		y = y + 2
	}
	if SIZE == 16 {

		SetPos(x, y)
		for i := 0; i < 8; i++ {
			WriteByte(F8X16[int(c)*16+i], OLED_DATA)
		}

		SetPos(x, y+1)
		for i := 0; i < 8; i++ {
			WriteByte(F8X16[int(c)*16+i+8], OLED_DATA)
		}

	} else {
		SetPos(x, y+1)

	}
}

func ShowString(x, y byte, chrs string) {
	for _, v := range []byte(chrs) {
		ShowChar(x, y, v)
		x += 8
		if x > 120 {
			x = 0
			y += 2
		}

	}

}
