package radio

import (
	"fm-go-bin/internal/radio/tavarua"
	"fm-go-bin/internal/v4l2"
	"fmt"
	"unsafe"
)

func (tuner RadioTuner) Seek(direction int8) error {
	// вперёд - 1; назад - 0
	if direction != 1 {
		direction = 0
	}

	err := tuner.Control.Set(
		tavarua.CID_SRCHMODE,
		v4l2.SEARCH_SEEK,
	)

	if err != nil {
		fmt.Printf("err(seek): set SRCHMODE failed: %v\n", err)
		return err
	}

	err = tuner.Control.Set(
		tavarua.CID_SCANDWELL,
		7,
	)

	if err != nil {
		print("err(seek): set SCANDWELL failed: %v\n", err)
		return err
	}

	hw_seek := v4l2.HardwareFrequencySeek{
		Type:        v4l2.V4L2_TUNER_RADIO,
		Seek_upward: uint32(direction),
	}

	err = tuner.Control.Ioctl(
		v4l2.VIDIOC_S_HW_FREQ_SEEK,
		unsafe.Pointer(&hw_seek),
	)

	if err != nil {
		print("err(seek): search failed: %v\n", err)
		return err
	}

	return nil
}
