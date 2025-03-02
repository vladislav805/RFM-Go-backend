package radio

import "fm-go-bin/internal/v4l2"

type RadioTuner struct {
	Control *v4l2.V4L2Ctl
}

func New(path string) (*RadioTuner, error) {
	ctl, err := v4l2.New(path)

	if err != nil {
		return nil, err
	}

	return &RadioTuner{
		Control: ctl,
	}, nil
}

func (tuner RadioTuner) Close() {
	tuner.Control.Close()
}
