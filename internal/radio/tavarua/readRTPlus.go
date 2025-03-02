package tavarua

import (
	"fmt"
)

func (tuner TavaruaRadio) readRTPlus() (string, error) {
	buf, len, err := tuner.readRawBuffer(RT_PLUS_IND)

	if err != nil {
		fmt.Printf("err(RT Plus): %v", err)
		return "", err
	}

	if len == 0 {
		return "", fmt.Errorf("invalid rt plus buffer length (%d)", len)
	}

	// [8 0 1 23 16 4 9 11 8 0 1 23 16 4 9 11]

	const MAX_RT_PLUS_TAGS = 2

	fmt.Printf("rtp(%d) %v\n", len, buf)

	if buf[1] == 0 && buf[0] > 0 && buf[0] <= MAX_RT_PLUS_TAGS {

	}

	return string(buf), nil
}
