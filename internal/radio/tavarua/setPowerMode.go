package tavarua

func (tuner TavaruaRadio) setPowerMode(mode PowerMode) error {
	return tuner.Control.Set(
		CID_LP_MODE,
		uint32(mode),
	)
}
