package radio

func (tuner RadioTuner) GetSignalStrength() (uint32, error) {
	info, err := tuner.getTuner()

	if err == nil {
		return info.Signal, nil
	} else {
		return 0, err
	}
}
