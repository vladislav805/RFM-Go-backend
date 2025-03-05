package radio

import (
	"errors"
	"fm-go-bin/internal/radio/tavarua"
)

func GetInstance() (RadioInterface, error) {
	if tavarua.Test() {
		return tavarua.New()
	}

		return ctl, nil
	}

	return nil, errors.New("not implemented")
}
