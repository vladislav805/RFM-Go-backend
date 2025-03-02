package radio

import "fm-go-bin/internal/radio/tavarua"

func (tuner RadioTuner) SetSpacing(spacing uint32) error {
	return tuner.Control.Set(
		tavarua.CID_SPACING,
		spacing,
	)
}
