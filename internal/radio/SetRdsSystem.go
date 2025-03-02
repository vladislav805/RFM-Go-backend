package radio

import "fm-go-bin/internal/radio/tavarua"

func (tuner RadioTuner) SetRdsSystem(system uint32) error {
	return tuner.Control.Set(
		tavarua.CID_RDS_STD,
		system,
	)
}
