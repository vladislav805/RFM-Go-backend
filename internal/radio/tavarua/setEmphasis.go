package tavarua

func (tuner TavaruaRadio) setEmphasis(emphasis uint32) error {
	return tuner.Control.Set(
		CID_EMPHASIS,
		emphasis,
	)
}
