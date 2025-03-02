package radio_state

import "fm-go-bin/internal/rds"

type RadioGlobalState struct {
	Enabled bool `json:"enabled"`

	// Current frequency in kHz (90.6 MHz = 90600 kHz)
	Frequency uint32 `json:"frequency"`

	// Is stereo audio?
	Stereo bool `json:"stereo"`

	// Signal strength
	SignalStrength uint32 `json:"signalStrength"`

	// Is RDS available?
	RdsAvailable bool `json:"isRdsAvailable"`

	// RDS PS (Program Service)
	RdsPS rds.PS `json:"rdsPS"`

	// RDS RT (Radio Text)
	RdsRT rds.RT `json:"rdsRT"`

	// RDS PI (Program ID)
	RdsPI rds.PI `json:"rdsPI"`

	// RDS PTY (Program Type)
	RdsPTY rds.PTY `json:"rdsPTY"`

	// AF: Alternative frequencies, array of frequencies (kHz)
	RdsAF rds.AF `json:"rdsAF"`
}
