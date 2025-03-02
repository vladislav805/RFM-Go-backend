package tavarua

import (
	"fm-go-bin/internal/v4l2"
	"unsafe"
)

func (tuner TavaruaRadio) getTuner() (v4l2.Tuner, error) {
	info := v4l2.Tuner{}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_G_TUNER,
		unsafe.Pointer(&info),
	)

	return info, err
}
