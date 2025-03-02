package tavarua

import (
	"errors"
	"fm-go-bin/internal/system"
	"fm-go-bin/internal/v4l2"
	"os"
	"time"
)

type TavaruaRadio struct {
	Control *v4l2.V4L2Ctl
}

const DEV_PATH = "/dev/radio0"

func New() (*TavaruaRadio, error) {
	err := startup()

	if err != nil {
		return nil, err
	}

	ctl, err := v4l2.New(DEV_PATH)

	if err != nil {
		return nil, err
	}

	return &TavaruaRadio{
		Control: ctl,
	}, nil
}

func Test() bool {
	_, err := os.Stat(DEV_PATH)

	return err == nil
}

func startup() error {
	system.SetProp("hw.fm.mode", "normal")
	system.SetProp("hw.fm.version", "0")
	system.SetProp("ctl.start", "fm_dl")

	successInit := false

	for range 600 {
		value, _ := system.GetProp("hw.fm.init")

		if len(value) > 0 && value[0] == '1' {
			successInit = true
			break
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}

	if !successInit {
		return errors.New("failed to prepare")
	}

	time.Sleep(500 * time.Millisecond)

	return nil
}

func (tuner TavaruaRadio) Close() {
	tuner.Control.Close()
}
