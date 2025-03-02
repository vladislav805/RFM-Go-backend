package radio

import "fm-go-bin/internal/radio_state"

type RadioInterface interface {
	StartTuner() error
	Close()

	GetFrequency() (uint32, error)
	SetFrequency(kHz uint32) error

	GetSignalStrength() (uint32, error)

	Seek(direction int8) error

	SetSpacing(spacing uint32) error
	SetAntenna(antenna uint32) error
	SetBand(lower uint32, upper uint32) error
	SetRdsState(state bool) error

	// Background loop for polling events
	PollEvents(state *radio_state.RadioGlobalState, onEvent func())
}
