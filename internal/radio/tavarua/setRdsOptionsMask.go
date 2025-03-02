package tavarua

import "fmt"

func (tuner TavaruaRadio) setRdsOptionsMask(mask uint32) error {
	err := tuner.Control.Set(
		CID_RDS_GROUP_MASK,
		mask,
	)

	if err != nil {
		return fmt.Errorf("failed to set RDSGROUP_MASK: %v", err)
	}

	return nil
}
