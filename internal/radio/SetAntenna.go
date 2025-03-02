package radio

import "fm-go-bin/internal/radio/tavarua"

func (tuner RadioTuner) SetAntenna(antenna uint32) error {
	return tuner.Control.Set(
		tavarua.CID_ANTENNA,
		antenna,
	)
}
