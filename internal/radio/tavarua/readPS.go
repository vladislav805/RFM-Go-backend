package tavarua

import (
	"fm-go-bin/internal/rds"
	"fmt"
)

func (tuner TavaruaRadio) readPS() (rds.PS, rds.PI, rds.PTY, error) {
	buf, len, err := tuner.readRawBuffer(BUF_PS_RDS)

	if err != nil {
		return "", 0, 0, err
	}

	if len == 0 {
		return "", 0, 0, fmt.Errorf("invalid ps buffer length (%d)", len)
	}

	/*
	 * TAVARUA_BUF_PS_RDS
	 * | buf[0]  | buf[1]  | buf[2]  | buf[3]  | buf[4]  | buf[5]  |
	 * |.... ....|.... ....|.... ....|.... ....|.... ....|.... ....|
	 * |0000 1111|         |         |         |         |         | numOfPs (buf[0] & 0x0f)
	 * |         |0001 1111|         |         |         |         | program_type (buf[1] & 0x1f)
	 * |         |         |1111 1111|1111 1111|         |         | program_id (((buf[2] & 0xff) << 8) | (buf[3] & 0xff))
	 * |         |         |         |         |         |1111 1111| program_name (buf[5] with length psLength)
	 */
	numOfPs := buf[0] & 0x0f // 0-15
	psLength := numOfPs * 8  // 0-120

	// Program Service
	ps := rds.PS(buf[5 : 5+psLength])

	// Program Type
	pty := rds.PTY(buf[1] & 0x1F)

	// Program ID
	pi := rds.PI(((uint16(buf[2]) & 0xFF) << 8) | (uint16(buf[3]) & 0xFF))

	return ps, pi, pty, nil
}
