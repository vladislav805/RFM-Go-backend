package v4l2utils

/*
 * Multiplying factor to convert to Radio frequency
 * The frequency is set in units of 62.5 Hz when using V4L2_TUNER_CAP_LOW,
 * 62.5 kHz otherwise.
 * The tuner is able to have a channel spacing of 50, 100 or 200 kHz.
 * tuner->capability is therefore set to V4L2_TUNER_CAP_LOW
 * The TUNE_MULT is then: 1 MHz / 62.5 Hz = 16000
 */
const TUNE_MULT = 16000

/**
 * Multiplying factor to convert to Radio frequency
 * The frequency is set in units of 62.5 Hz when using V4L2_TUNER_CAP_LOW,
 * 62.5 kHz otherwise.
 * The tuner is able to have a channel spacing of 50, 100 or 200 kHz.
 * tuner->capability is therefore set to V4L2_TUNER_CAP_LOW
 * The FREQ_MUL is then: 1 MHz / 62.5 Hz = 16000
 */
func TuneFrequencyToKHz(freq uint32) uint32 {
	return (freq * 1000) / TUNE_MULT
}

func KHzToTuneFrequency(freq uint32) uint32 {
	return (freq * TUNE_MULT / 1000)
}
