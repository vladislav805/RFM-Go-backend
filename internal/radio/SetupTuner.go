package radio

import (
	"fm-go-bin/internal/radio/tavarua"
	"fm-go-bin/internal/system"
	"fm-go-bin/internal/v4l2"
	"fmt"
	"time"
)

func (tuner RadioTuner) SetupTuner() error {
	cap, err := tuner.QueryCapabilities()

	if err != nil {
		fmt.Printf("Failed to query capabilities: %v\n", err)
		return err
	}

	fmt.Printf("Version: %v\n", cap.Version)

	if err := system.SetProp("hw.fm.version", fmt.Sprint(cap.Version)); err != nil {
		return fmt.Errorf("failed to set prop 'hw.fm.version': %v", err)
	}

	if system.IsSMDTransportLayer() {
		if err := system.SetProp("ctl.start", "fm_dl"); err != nil {
			return fmt.Errorf("failed to set prop 'ctl.start': %v", err)
		}
	}

	if err := tuner.SetTunerState(v4l2.TUNER_STATE_RX); err != nil {
		return fmt.Errorf("failed to set state: %v", err)
	}

	if err := tuner.SetEmphasis(tavarua.FM_RX_EMP75); err != nil {
		return fmt.Errorf("failed to set emphasis: %v", err)
	}

	if err := tuner.SetSpacing(tavarua.FM_RX_SPACE_100KHZ); err != nil {
		return fmt.Errorf("failed to set spacing: %v", err)
	}

	if err := tuner.SetAntenna(0); err != nil {
		return fmt.Errorf("failed to set antenna: %v", err)
	}

	if err := tuner.SetBand(87500, 108000); err != nil {
		return fmt.Errorf("failed to set band: %v", err)
	}

	if err := tuner.SetRdsState(true); err != nil {
		return fmt.Errorf("failed to set rds state: %v", err)
	}

	if err := tuner.SetRdsSystem(tavarua.FM_RX_RDS_SYSTEM); err != nil {
		return fmt.Errorf("failed to set rds system: %v", err)
	}

	if err := tuner.SetRdsGroupOptions(); err != nil {
		return fmt.Errorf("failed to set rds group options: %v", err)
	}

	time.Sleep(time.Second)

	rdsOptions, err := tuner.GetRdsGroupOptions()

	if err != nil {
		return fmt.Errorf("failed to get rds options: %v", err)
	}

	fmt.Printf("rdsOptions: %v\n", rdsOptions)

	return nil
}
