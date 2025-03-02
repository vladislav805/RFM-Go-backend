package tavarua

func (tuner TavaruaRadio) SetRdsSystem(system uint32) error {
	return tuner.Control.Set(
		CID_RDS_STD,
		system,
	)
}
