package radio

import (
	"fm-go-bin/internal/radio/tavarua"
	"fmt"
)

func (tuner RadioTuner) ReadAF() ([]uint32, error) {
	buf, len, err := tuner.readRawBuffer(tavarua.BUF_AF_LIST)

	if err != nil {
		fmt.Printf("err(AF): %v", err)
		return []uint32{}, err
	}

	if len == 0 {
		return []uint32{}, fmt.Errorf("invalid rt buffer length (%d)", len)
	}

	// buf[4] | (buf[5] << 8)
	// buf[6] = count of frequencies
	// buf[($index * 4) + 6 + (1...4)] with shift = one frequency (uint32)
	size := int(buf[6] & 0xff)

	if size <= 0 || size > 25 {
		return []uint32{}, fmt.Errorf("af list length invalid: 4=%d, 4&0ff=%d, 6=%d", buf[4], buf[4]&0xff, size)
	}

	af := make([]uint32, size)

	for i := range size {
		shift := 6 + i*4

		var freq uint32 = (uint32(buf[shift+1]) & 0xff) |
			(uint32((buf[shift+2])&0xff) << 8) |
			(uint32((buf[shift+3])&0xff) << 16) |
			(uint32((buf[shift+4])&0xff) << 24)

		af[i] = freq
	}

	return af, nil
}
