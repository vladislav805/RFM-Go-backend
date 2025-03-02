package radio

import (
	"fm-go-bin/internal/radio/tavarua"
	"fmt"
)

func (tuner RadioTuner) SetRdsGroupOptions() error {
	err := tuner.Control.Set(
		tavarua.CID_RDSGROUP_PROC,
		0xff,
	)

	if err != nil {
		return fmt.Errorf("failed to set RDSGROUP_PROC: %v", err)
	}

	err = tuner.Control.Set(
		tavarua.CID_PSALL,
		0xffff,
	)

	if err != nil {
		return fmt.Errorf("failed to set PSALL: %v", err)
	}

	return nil
}
