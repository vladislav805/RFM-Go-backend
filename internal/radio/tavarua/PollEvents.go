package tavarua

import (
	"fm-go-bin/internal/env"
	"fm-go-bin/internal/radio_state"
	"fmt"
)

// Background loop for polling events
func (tuner TavaruaRadio) PollEvents(state *radio_state.RadioGlobalState, onEvent func()) {
	for {
		events, len, err := tuner.readRawBuffer(BUF_EVENTS)

		if err == nil && len > 0 {
			for _, eventCode := range events {
				tuner.handleRdsEvent(eventCode, state)
			}

			onEvent()
		}
	}
}

func _log(format string, a ...any) {
	if env.IsVerbose {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}

func (tuner TavaruaRadio) handleRdsEvent(eventCode byte, state *radio_state.RadioGlobalState) {
	switch eventCode {
	case EVT_RADIO_READY:
		{
			state.Enabled = true

			_log("🟢 Radio ready")
		}

	case EVT_RADIO_DISABLED:
		{
			state.Enabled = false

			_log("🟢 Radio disabled")
		}

	case EVT_TUNE_SUCCESS:
		{
			kHz, err := tuner.GetFrequency()

			// EVT_TUNE_SUCCESS => EVT_ABOVE_TH + EVT_MONO + EVT_RDS_NOT_AVAIL
			if err == nil {
				state.Frequency = kHz
				state.RdsAF = nil
			}

			_log("🟢 Tune to frequency %d kHz", kHz)
		}

	case EVT_SEEK_COMPLETE:
		{
			kHz, err := tuner.GetFrequency()

			if err == nil {
				state.Frequency = kHz
			}

			_log("🟢 Seek completed at frequency %d kHz", kHz)
		}

	case EVT_SCAN_NEXT:
		{
			kHz, _ := tuner.GetFrequency()

			_log("🟢 Event scan next, frequency %d", kHz)
		}

	case EVT_NEW_RAW_RDS:
		{
			_log("🟡 New raw RDS")
		}

	case EVT_NEW_RT_RDS:
		{
			rt, pi, pty, err := tuner.readRT()

			if err == nil {
				state.RdsRT = rt
				state.RdsPI = pi
				state.RdsPTY = pty
			}

			_log("🟢 New RT: '%v', PI: %x, PTY: %d, err: '%v'", rt, pi, pty, err)
		}

	case EVT_NEW_PS_RDS:
		{
			ps, pi, pty, err := tuner.readPS()

			if err == nil {
				state.RdsPS = ps
				state.RdsPI = pi
				state.RdsPTY = pty
			}

			_log("🟢 New PS: '%v', PI: %x, PTY: %d, err: '%v'", ps, pi, pty, err)
		}

	case EVT_ERROR:
		{
			_log("🟢 Error")
		}

	case EVT_BELOW_TH:
		{
			_log("🟢 Event below")
		}

	case EVT_ABOVE_TH:
		{
			_log("🟢 Event above")
		}

	case EVT_STEREO:
		{
			state.Stereo = true

			_log("🟢 Stereo = true")
		}

	case EVT_MONO:
		{
			state.Stereo = false

			_log("🟢 Stereo = false")
		}

	case EVT_RDS_AVAIL:
		{
			state.RdsAvailable = true

			_log("🟢 RDS available")
		}

	case EVT_RDS_NOT_AVAIL:
		{
			state.RdsAvailable = false

			_log("🟢 RDS not available")
		}

	case EVT_NEW_SEARCH_LIST:
		{
			_log("🟢 New search list")
		}

	case EVT_NEW_AF_LIST:
		{
			af, err := tuner.readAF()

			if err == nil {
				state.RdsAF = af
			}

			_log("🟢 New alternative frequencies list: %v, %v", af, err)
		}

	case EVT_NEW_RT_PLUS:
		{
			_log("🟡 EVT_NEW_RT_PLUS")
			// tuner.ReadRTPlus()
		}

	case EVT_RDS_GRP_MASK_REQ:
		{
			_log("🟡 RDS group mask request")
			tuner.setRdsOptionsMask(0)
		}

	default:
		{
			_log("🔴 Unknown event: %d", eventCode)
		}
	}

}
