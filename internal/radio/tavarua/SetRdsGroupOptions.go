package tavarua

import "fmt"

func (tuner TavaruaRadio) SetupRdsOptions() error {
	err := tuner.Control.Set(
		CID_RDS_GROUP_PROC,
		0xff,
	)

	if err != nil {
		return fmt.Errorf("failed to set RDSGROUP_PROC: %v", err)
	}

	err = tuner.Control.Set(
		CID_PS_ALL,
		0xffff,
	)

	if err != nil {
		return fmt.Errorf("failed to set PSALL: %v", err)
	}

	// https://github.com/Evervolv/android_vendor_qcom_opensource_libfmjni/blob/v-15.1/FmRadioController.cpp#L1133

	return nil
}
