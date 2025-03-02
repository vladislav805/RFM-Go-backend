package main

import (
	"errors"
	"fm-go-bin/internal/radio"
	"fm-go-bin/internal/radio/tavarua"
	"fm-go-bin/internal/system"
	"fmt"
	"time"
)

func main() {
	err := prepareTuner()

	if err != nil {
		fmt.Printf("Failed to prepare tuner: %v\n", err)
		return
	}

	tuner, err := radio.New("/dev/radio0")

	if err != nil {
		fmt.Printf("Failed to initialize tuner: %v\n", err)
		return
	}

	defer tuner.Close()

	err = tuner.SetupTuner()

	fmt.Printf("SetupTuner/err: %v\n", err)

	if err := tuner.SetFrequency(90600); err != nil {
		fmt.Printf("setV4l2Ctrl(freq): %v\n", err)
	}

	time.Sleep(500 * time.Millisecond)

	kHz, err := tuner.GetFrequency()

	if err != nil {
		fmt.Printf("getFreq(err) = %v\n", err)
	} else {
		fmt.Printf("getFreq(res) = %d\n", kHz)
	}

	time.Sleep(time.Second)

	go func() {
		time.Sleep(10 * time.Second)

		fmt.Println("SET")
		tuner.SetFrequency(100500)

		time.Sleep(5 * time.Second)

		tuner.Seek(1)
	}()

	background(tuner)
}

func prepareTuner() error {
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

func background(tuner *radio.RadioTuner) {
	for {
		eventCodes := tuner.WaitEvents()

		for _, eventCode := range eventCodes {
			handleRdsEvent(eventCode, tuner)
		}

		time.Sleep(time.Millisecond * 10)
	}
}

func handleRdsEvent(eventCode byte, tuner *radio.RadioTuner) {
	switch eventCode {
	case tavarua.EVT_RADIO_READY:
		{
			fmt.Println("🟢 Radio ready")
		}

	case tavarua.EVT_RADIO_DISABLED:
		{
			fmt.Println("🟢 Radio disabled")
		}

	case tavarua.EVT_TUNE_SUCCESS:
		{
			kHz, _ := tuner.GetFrequency()
			fmt.Printf("🟢 Tune to frequency %d kHz\n", kHz)
		}

	case tavarua.EVT_SEEK_COMPLETE:
		{
			kHz, _ := tuner.GetFrequency()
			fmt.Printf("🟢 Seek completed at frequency %d kHz\n", kHz)
		}

	case tavarua.EVT_SCAN_NEXT:
		{
			kHz, _ := tuner.GetFrequency()
			fmt.Printf("🟢 Event scan next, frequency %d\n", kHz)
		}

	case tavarua.EVT_NEW_RAW_RDS:
		{
			fmt.Println("🟢 New raw RDS")
		}

	case tavarua.EVT_NEW_RT_RDS:
		{
			rt, pi, pty, err := tuner.ReadRT()

			fmt.Printf("🟢 New RT: '%v', PI: %d, PTY: %d, err: '%v'\n", rt, pi, pty, err)
		}

	case tavarua.EVT_NEW_PS_RDS:
		{
			ps, pi, pty, err := tuner.ReadPS()

			fmt.Printf("🟢 New PS: '%v', PI: %d, PTY: %d, err: '%v'\n", ps, pi, pty, err)
		}

	case tavarua.EVT_ERROR:
		{
			fmt.Println("🟢 Error")
		}

	case tavarua.EVT_BELOW_TH:
		{
			fmt.Println("🟢 Event below")
		}

	case tavarua.EVT_ABOVE_TH:
		{
			fmt.Println("🟢 Event above")
		}

	case tavarua.EVT_STEREO:
		{
			fmt.Println("🟢 Stereo = true")
		}

	case tavarua.EVT_MONO:
		{
			fmt.Println("🟢 Stereo = false")
		}

	case tavarua.EVT_RDS_AVAIL:
		{
			fmt.Println("🟢 RDS available")
		}

	case tavarua.EVT_RDS_NOT_AVAIL:
		{
			fmt.Println("🟢 RDS not available")
		}

	case tavarua.EVT_NEW_SEARCH_LIST:
		{
			fmt.Println("🟢 New search list")
		}

	case tavarua.EVT_NEW_AF_LIST:
		{
			af, err := tuner.ReadAF()

			fmt.Printf("🟢 New alternative frequencies list: %v, %v\n", af, err)
		}
	}

}
