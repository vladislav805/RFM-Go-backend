package radio

import (
	"fm-go-bin/internal/v4l2"
	"unsafe"
)

func (tuner RadioTuner) QueryCapabilities() (*v4l2.Capability, error) {
	cap := &v4l2.Capability{}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_QUERYCAP,
		unsafe.Pointer(cap),
	)

	if err != nil {
		return nil, err
	}

	return cap, nil
}
