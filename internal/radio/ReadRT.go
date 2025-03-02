package radio

import (
	"fm-go-bin/internal/radio/tavarua"
	"fmt"
)

func (tuner RadioTuner) ReadRT() (string, uint16, uint8, error) {
	buf, len, err := tuner.readRawBuffer(tavarua.BUF_RT_RDS)

	if err != nil {
		fmt.Printf("err(RT): %v", err)
		return "", 0, 0, err
	}

	if len == 0 {
		return "", 0, 0, fmt.Errorf("invalid rt buffer length (%d)", len)
	}

	/*
	 * Radio text
	 * | buf[0]  | buf[1]  | buf[2]  | buf[3]  | buf[4]  | buf[5]  |
	 * |.... ....|.... ....|.... ....|.... ....|.... ....|.... ....|
	 * |1111 1111|         |         |         |         |         | length (buf[0] + 5)
	 *?|         |0001 1111|         |         |         |         | program_type (buf[1] & 0x1f)
	 *?|         |         |1111 1111|1111 1111|         |         | program_id (((buf[2] & 0xff) << 8) | (buf[3] & 0xff))
	 * |         |         |         |         |         |1111 1111| radio_text (buf[5] with length buf[0] + 5)
	 */
	var shift byte = 5
	length := buf[0] + shift

	// Radio text
	rt := string(buf[shift : shift+length])

	// Program Type
	pty := (buf[1] & 0x1F)

	// Program ID
	pi := ((uint16(buf[2]) & 0xFF) << 8) | (uint16(buf[3]) & 0xFF)

	return string(rt), pi, pty, nil
}
