package radio

import (
	"errors"
	"fm-go-bin/internal/radio/helium"
	"fm-go-bin/internal/radio/tavarua"
)

func GetInstance() (RadioInterface, error) {
	if tavarua.Test() {
		return tavarua.New()
	}

	if helium.Test() {
		return helium.New()
	}

	return nil, errors.New("not implemented")
}
