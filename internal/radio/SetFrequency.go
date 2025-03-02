package radio

import (
	"fm-go-bin/internal/v4l2"
	"fm-go-bin/internal/v4l2utils"
	"fmt"
	"unsafe"
)

func (tuner RadioTuner) SetFrequency(kHz uint32) error {
	control := v4l2.Frequency{
		Type:      v4l2.V4L2_TUNER_RADIO,
		Frequency: v4l2utils.KHzToTuneFrequency(kHz),
	}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_S_FREQUENCY,
		unsafe.Pointer(&control),
	)

	if err != nil {
		return fmt.Errorf("failed to set frequency to %d: errno = %d", kHz, err)
	}

	return nil
}
