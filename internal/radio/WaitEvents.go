package radio

import "fm-go-bin/internal/radio/tavarua"

func (tuner RadioTuner) WaitEvents() []byte {
	buf, len, err := tuner.readRawBuffer(tavarua.BUF_EVENTS)

	if err != nil || len == 0 {
		return []byte{}
	}

	return buf
}
