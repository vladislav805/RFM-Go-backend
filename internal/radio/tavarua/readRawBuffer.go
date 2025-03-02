package tavarua

import (
	"fm-go-bin/internal/v4l2"
	"fmt"
	"unsafe"
)

// https://github.com/bcyj/android_tools_leeco_msm8996/blob/f242858a8cbb2c9ad9b5b2e53d75380ae775727a/fm/hal_ss_test/FmRadioController_qcom.cpp#L1226-L1279
func (tuner TavaruaRadio) readRawBuffer(index uint32) ([]byte, uint32, error) {
	var size uint32 = 256

	buf := make([]byte, size)

	v4l2_buf := v4l2.Buffer{
		Index:   index,
		Type:    v4l2.BUF_TYPE_PRIVATE,
		Memory:  v4l2.MEMORY_USERPTR,
		UserPtr: unsafe.Pointer(&buf),
		Length:  size,
	}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_DQBUF,
		unsafe.Pointer(&v4l2_buf),
	)

	if err != nil {
		return []byte{}, 0, fmt.Errorf("readRawBuffer: ioctl read 0x%x failed, err = %v", index, err)
	}

	return buf[0:v4l2_buf.BytesUsed], v4l2_buf.BytesUsed, nil
}
