package v4l2

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

func (tuner V4L2Ctl) Ioctl(req uintptr, arg unsafe.Pointer) error {
	if _, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(tuner.fd),
		req,
		uintptr(arg),
	); errno != 0 {
		return errno
	}

	return nil
}
