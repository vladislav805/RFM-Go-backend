package tavarua

import (
	"fm-go-bin/internal/env"
	"fm-go-bin/internal/system"
	"fm-go-bin/internal/v4l2"
	"fmt"
	"time"
)

// https://github.com/rmottola/Arctic-Fox/blob/ad84bb61f5054d6280f432949bcb3315839a5ae4/hal/gonk/GonkFMRadio.cpp#L122
// https://github.com/sonyxperiadev/vendor-qcom-opensource-fm/blob/aabbad189c592aa67facec37bb4d52eb5f991738/jni/android_hardware_fm.cpp#L121
func (tuner TavaruaRadio) StartTuner() error {
	cap, err := tuner.queryCapabilities()

	if err != nil {
		fmt.Printf("Failed to query capabilities: %v\n", err)
		return err
	}

	if env.IsVerbose {
		fmt.Printf("Version: %v\n", cap.Version)
	}

	if err := system.SetProp("hw.fm.version", fmt.Sprint(cap.Version)); err != nil {
		return fmt.Errorf("failed to set prop 'hw.fm.version': %v", err)
	}

	if system.IsSMDTransportLayer() {
		if err := system.SetProp("ctl.start", "fm_dl"); err != nil {
			return fmt.Errorf("failed to set prop 'ctl.start': %v", err)
		}
	}

	if err := tuner.setTunerState(v4l2.TUNER_STATE_RX); err != nil {
		return fmt.Errorf("failed to set state: %v", err)
	}

	if err := tuner.setPowerMode(FM_RX_POWER_MODE_NORMAL); err != nil {
		return fmt.Errorf("failed to set power mode: %v", err)
	}

	if err := tuner.setEmphasis(FM_RX_EMP75); err != nil {
		return fmt.Errorf("failed to set emphasis: %v", err)
	}

	if err := tuner.SetSpacing(FM_RX_SPACE_100KHZ); err != nil {
		return fmt.Errorf("failed to set spacing: %v", err)
	}

	if err := tuner.SetAntenna(0); err != nil {
		return fmt.Errorf("failed to set antenna: %v", err)
	}

	if err := tuner.SetBand(87_500, 108_000); err != nil {
		return fmt.Errorf("failed to set band: %v", err)
	}

	if err := tuner.SetRdsState(true); err != nil {
		return fmt.Errorf("failed to set rds state: %v", err)
	}

	if err := tuner.SetRdsSystem(FM_RX_RDS_SYSTEM); err != nil {
		return fmt.Errorf("failed to set rds system: %v", err)
	}

	if err := tuner.setRdsOptionsMask(0xffffffff); err != nil {
		return fmt.Errorf("failed to set rds options mask: %v", err)
	}

	if err := tuner.SetupRdsOptions(); err != nil {
		return fmt.Errorf("failed to set rds group options: %v", err)
	}

	time.Sleep(time.Millisecond * 50)

	return nil
}
