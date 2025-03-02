package radio

import (
	"fm-go-bin/internal/radio/tavarua"
	"fm-go-bin/internal/v4l2"
	"unsafe"
)

func (tuner RadioTuner) GetRdsGroupOptions() (uint8, error) {
	control := v4l2.Control{
		Id: tavarua.CID_RDSGROUP_PROC,
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
