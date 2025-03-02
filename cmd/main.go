package main

import (
	"encoding/json"
	"fm-go-bin/internal/radio"
	"fm-go-bin/internal/radio_state"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("======= RFM Go backend =======")

	tuner, err := radio.GetInstance()

	if err != nil {
		fmt.Printf("Failed to initialize tuner: %v\n", err)
		return
	}

	defer tuner.Close()

	state := radio_state.RadioGlobalState{
		Enabled:        false,
		Frequency:      1,
		Stereo:         false,
		RdsAvailable:   false,
		RdsPS:          "",
		RdsRT:          "",
		RdsPI:          0,
		RdsPTY:         0,
		RdsAF:          nil,
		SignalStrength: 0,
	}

	if err := tuner.StartTuner(); err != nil {
		fmt.Printf("Failed to start tuner: %v\n", err)
		return
	}

	if err := tuner.SetFrequency(90_600); err != nil {
		fmt.Printf("SetFrequency(freq): %v\n", err)
		return
	}

	channel := make(chan interface{})

	go tuner.PollEvents(&state, func() {
		result, err := json.Marshal(state)

		if err == nil {
			fmt.Println(string(result))
		}
	})

	srv(&state, tuner)

	<-channel
}

func srv(state *radio_state.RadioGlobalState, tuner radio.RadioInterface) {
	http.Handle("/get", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonContent, err := json.Marshal(state)

		if err == nil {
			w.Header().Add("content-type", "application/json; charset=utf-8")
			w.Write(jsonContent)
		}
	}))

	http.Handle("/set/{kHz}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		kHz, err := strconv.Atoi(r.PathValue("kHz"))

		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}

		err = tuner.SetFrequency(uint32(kHz))

		result := err == nil

		jsonContent, _ := json.Marshal(result)

		w.Header().Add("content-type", "application/json; charset=utf-8")
		w.Write(jsonContent)
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
