package helium

import (
	"fm-go-bin/internal/radio_state"
	"os"
	"plugin"
)

type HeliumRadio struct {
}

const HELIUM_NAME = "fm_helium.so"
const HELIUM_PATH = "/system_ext/lib64/fm_helium.so"

func New() (*HeliumRadio, error) {
	// handle, err := purego.Dlopen(HELIUM_NAME, purego.RTLD_LAZY)

	// fmt.Printf("dlopen: %v, %v\n", handle, err)

	// if err != nil {
	// 	return nil, err
	// }

	// defer purego.Dlclose(handle)

	// ptr, err := purego.Dlsym(handle, "hal_init")

	// fmt.Printf("dlsym: %v, %v\n", ptr, err)

	// if err != nil {
	// 	return nil, err
	// }

	// r1, r2, r3 := purego.SyscallN(ptr)

	// fmt.Printf("syscall: %v, %v, %v\n", r1, r2, r3)

	plugin, err := plugin.Open(HELIUM_PATH)

	if err != nil {
		return nil, err
	}

	hal_init, err := plugin.Lookup("hal_init")

	if err != nil {
		return nil, err
	}

	hal_init.(func())()

	return &HeliumRadio{}, nil
}

func Test() bool {
	_, err := os.Stat(HELIUM_PATH)

	return err == nil
}

func (tuner HeliumRadio) StartTuner() error {
	return nil
}

func (tuner HeliumRadio) Close() {

}

func (tuner HeliumRadio) GetFrequency() (uint32, error) {
	return 0, nil
}

func (tuner HeliumRadio) SetFrequency(kHz uint32) error {
	return nil
}

func (tuner HeliumRadio) GetSignalStrength() (uint32, error) {
	return 0, nil
}

func (tuner HeliumRadio) Seek(direction int8) error {
	return nil
}

func (tuner HeliumRadio) SetSpacing(spacing uint32) error {
	return nil
}

func (tuner HeliumRadio) SetAntenna(antenna uint32) error {
	return nil
}

func (tuner HeliumRadio) SetBand(lower uint32, upper uint32) error {
	return nil
}

func (tuner HeliumRadio) SetRdsState(state bool) error {
	return nil
}

func (tuner HeliumRadio) PollEvents(state *radio_state.RadioGlobalState, onEvent func()) {

}
