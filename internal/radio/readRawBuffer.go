package radio

import (
	"fm-go-bin/internal/v4l2"
	"fmt"
	"unsafe"
)

func (tuner RadioTuner) readRawBuffer(index uint32) ([]byte, uint32, error) {
	var size uint32 = 128

	buf := make([]byte, size)

	v4l2_buf := v4l2.Buffer{
		Index:   index,
		Type:    v4l2.BUF_TYPE_PRIVATE,
		Memory:  v4l2.MEMORY_USERPTR,
		Userptr: unsafe.Pointer(&buf),
		Length:  size,
	}

	err := tuner.Control.Ioctl(
		v4l2.VIDIOC_DQBUF,
		unsafe.Pointer(&v4l2_buf),
	)

	if err != nil {
		return []byte{}, 0, fmt.Errorf("readEvents: ioctl read 0x%x failed, err = %v", index, err)
	}

	return buf[0:v4l2_buf.Bytesused], v4l2_buf.Bytesused, nil
}
