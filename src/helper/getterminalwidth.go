package helper

import (
	"syscall"
	"unsafe"
)

func GetTerminalWidth() int {

	layout := struct {
				  Row    uint16
				  Col    uint16
				  Xpixel uint16
				  Ypixel uint16
			  }{}

	syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdout), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&layout)))
	return int(layout.Col)
}
