package radio

import "fm-go-bin/internal/radio/tavarua"

func (tuner RadioTuner) SetRdsState(enable bool) error {
	var value uint32 = 0

	if enable {
		value = 1
	}

	return tuner.Control.Set(
		tavarua.CID_RDSON,
		value,
	)
}
