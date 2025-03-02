package tavarua

import (
	"fm-go-bin/internal/v4l2"
	"unsafe"
)

func (tuner TavaruaRadio) getRdsGroupOptions() (uint8, error) {
	control := v4l2.Control{
		Id: CID_RDS_GROUP_PROC,
	}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_G_CTRL,
		unsafe.Pointer(&control),
	)

	if err != nil {
		return 0, err
	}

	return uint8(control.Value), nil
}
