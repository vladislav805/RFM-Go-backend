package v4l2

import (
	"fmt"
	"syscall"
	"unsafe"
)

func (ctl V4L2Ctl) Set(id uint32, value uint32) error {
	control := Control{
		Id:    id,
		Value: value,
	}

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(ctl.fd),
		VIDIOC_S_CTRL,
		uintptr(unsafe.Pointer(&control)),
	)
	if errno != 0 {
		return fmt.Errorf("failed to set control: control = 0x%x, value = %d, err = %d", id, value, errno)
	}

	return nil
}
