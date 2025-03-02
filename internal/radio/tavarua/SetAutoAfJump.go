package tavarua

func (tuner TavaruaRadio) SetAutoAfJump(enable bool) error {
	var value uint32 = 0

	if enable {
		value = 1
	}

	return tuner.Control.Set(
		CID_AF_JUMP,
		value,
	)
}
