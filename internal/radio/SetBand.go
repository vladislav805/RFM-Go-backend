package radio

import (
	"fm-go-bin/internal/v4l2"
	"fm-go-bin/internal/v4l2utils"
	"unsafe"
)

func (tuner RadioTuner) SetBand(lower uint32, upper uint32) error {
	ctl := &v4l2.Tuner{
		Index:     0,
		Signal:    0,
		RangeLow:  v4l2utils.KHzToTuneFrequency(lower),
		RangeHigh: v4l2utils.KHzToTuneFrequency(upper),
		AudMode:   1,
	}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_S_TUNER,
		unsafe.Pointer(ctl),
	)

	// FIXME
	// ret = set_v4l2_ctrl(V4L2_CID_PRIVATE_TAVARUA_REGION, band)

	return err
}
