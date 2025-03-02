package v4l2

import (
	"fmt"
	"syscall"
)

type V4L2Ctl struct {
	fd int
}

func New(path string) (*V4L2Ctl, error) {
	fd, err := syscall.Open(path, syscall.O_RDWR|syscall.O_NONBLOCK, 0666)

	if err != nil {
		return nil, fmt.Errorf("failed to open device: %v", err)
	}

	return &V4L2Ctl{
		fd: fd,
	}, nil
}

func (ctl V4L2Ctl) Close() {
	syscall.Close(ctl.fd)
}
