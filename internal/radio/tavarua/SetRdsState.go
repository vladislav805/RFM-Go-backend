package tavarua

func (tuner TavaruaRadio) SetRdsState(enable bool) error {
	var value uint32 = 0

	if enable {
		value = 1
	}

	return tuner.Control.Set(
		CID_RDS_ON,
		value,
	)
}
