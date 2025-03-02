package tavarua

func (tuner TavaruaRadio) SetSpacing(spacing uint32) error {
	return tuner.Control.Set(
		CID_SPACING,
		spacing,
	)
}
