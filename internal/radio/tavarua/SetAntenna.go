package tavarua

func (tuner TavaruaRadio) SetAntenna(antenna uint32) error {
	return tuner.Control.Set(
		CID_ANTENNA,
		antenna,
	)
}
