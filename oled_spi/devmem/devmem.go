package devmem

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

const (
	MAP_SIZE = 4096
	MAP_MASK = MAP_SIZE - 1
	devName  = "/dev/mem"
)

//access_type:b-byte;h-short(2 byte);w-word(4 byte)
func Write(oft int64, value []byte) error {
	file, err := os.OpenFile(devName, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	re, err := syscall.Mmap(int(file.Fd()), oft&(^MAP_MASK), MAP_SIZE,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(value) <= 0 {
		return errors.New("len(value)<=0")
	}

	for k, v := range value {
		re[int((oft&MAP_MASK))+k] = v
	}

	err = syscall.Munmap(re)
	if err != nil {
		return err
	}
	return nil
}

func Read(oft int64, access_type string) ([]byte, error) {
	file, err := os.OpenFile(devName, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	re, err := syscall.Mmap(int(file.Fd()), oft&(^MAP_MASK), MAP_SIZE,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	out := make([]byte, 0)
	switch access_type {
	case "b":
		out = append(out, re[oft&MAP_MASK])
	case "h":
		out = append(out, re[oft&MAP_MASK], re[oft&MAP_MASK+1])
	case "w":
		out = append(out, re[oft&MAP_MASK], re[oft&MAP_MASK+1], re[oft&MAP_MASK+2], re[oft&MAP_MASK+3])
	}

	err = syscall.Munmap(re)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//offBitSize:offBitSize is the position of Bit from start
//value must be 1 or 0
func WriteBit(oft int64, offBitSize int, value byte) error {
	file, err := os.OpenFile(devName, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	re, err := syscall.Mmap(int(file.Fd()), 
        	oft&(int64(^uint32(MAP_MASK))), MAP_SIZE,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return err
	}

	
	nbyte := offBitSize / 8
	sbyte := byte(offBitSize % 8)
	fmt.Println(nbyte,sbyte)
		
	oft=oft&MAP_MASK+int64(nbyte)	
	fmt.Println(oft)
	fmt.Println(oft,re[2193:])
	

	if value == 0 {
		re[oft] &= ^(byte(1 << sbyte))
		fmt.Println("value==0:",^byte((1 << sbyte)))
	} else {
		fmt.Println("value==1:",re[oft]|1 << sbyte, 1 << sbyte)
		re[oft] =16 // re[oft] |1 << sbyte

	}
		re[2193] =re[2193] |1 << sbyte
	fmt.Println(re[2193])	
	fmt.Println(2192,re[2192:])
	err = syscall.Munmap(re)
	if err != nil {
		return err
	}
	return nil
}

func ReadBit(oft int64, offBitSize int) (byte, error) {
	file, err := os.OpenFile(devName, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer file.Close()

	re, err := syscall.Mmap(int(file.Fd()), oft&(^MAP_MASK), MAP_SIZE,
		syscall.PROT_WRITE|syscall.PROT_READ,
		syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	out := byte(0)

	nbyte := offBitSize / 8
	sbyte := offBitSize % 8

	re[oft&MAP_MASK+int64(nbyte)] &= (1 << sbyte)
	re[oft&MAP_MASK+int64(nbyte)] = re[oft&MAP_MASK+int64(nbyte)] >> sbyte

	err = syscall.Munmap(re)
	if err != nil {
		return 0, err
	}
	return out, nil
}

