package radio

import (
	"fm-go-bin/internal/v4l2"
	"fm-go-bin/internal/v4l2utils"
	"fmt"
	"unsafe"
)

func (tuner RadioTuner) GetFrequency() (uint32, error) {
	freq_struct := v4l2.Frequency{
		Type: v4l2.V4L2_TUNER_RADIO,
	}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_G_FREQUENCY,
		unsafe.Pointer(&freq_struct),
	)

	if err != nil {
		return 0, fmt.Errorf("failed to get frequency: errno = %d", err)
	}

	return v4l2utils.TuneFrequencyToKHz(freq_struct.Frequency), nil
}
