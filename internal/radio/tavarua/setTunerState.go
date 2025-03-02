package tavarua

func (tuner TavaruaRadio) setTunerState(state uint32) error {
	return tuner.Control.Set(
		CID_STATE,
		state,
	)
}
