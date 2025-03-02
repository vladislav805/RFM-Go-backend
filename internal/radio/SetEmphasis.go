package radio

import "fm-go-bin/internal/radio/tavarua"

func (tuner RadioTuner) SetEmphasis(emphasis uint32) error {
	return tuner.Control.Set(
		tavarua.CID_EMPHASIS,
		emphasis,
	)
}
