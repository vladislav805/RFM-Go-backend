package radio

import (
	"errors"
	"fm-go-bin/internal/radio/tavarua"
)

func GetInstance() (RadioInterface, error) {
	if tavarua.Test() {
		ctl, err := tavarua.New()

		if err != nil {
			return nil, err
		}

		return ctl, nil
	}

	return nil, errors.New("not implemented")
}
