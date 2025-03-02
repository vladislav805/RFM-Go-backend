package v4l2

import "unsafe"

// v4l2_capability
type Capability struct {
	Driver       [16]uint8
	Card         [32]uint8
	Bus_info     [32]uint8
	Version      uint32
	Capabilities uint32
	Device_caps  uint32
	Reserved     [3]uint32
}

// v4l2_control
type Control struct {
	Id    uint32
	Value uint32
}

// v4l2_frequency
type Frequency struct {
	Tuner     uint32
	Type      uint32
	Frequency uint32
	Reserved  [8]uint32
}

// v4l2_tuner
type Tuner struct {
	Index      uint32
	Name       [32]byte
	Type       uint32
	Capability uint32
	RangeLow   uint32
	RangeHigh  uint32
	RXSubChans uint32
	AudMode    uint32
	Signal     uint32
	AFC        int32
	Reserved   [4]uint32
}

// v4l2_buffer
type Buffer struct {
	Index     uint32
	Type      uint32
	BytesUsed uint32
	Flags     uint32
	Field     uint32
	Timeval   struct {
		TvSec  uint32
		TvNsec uint32
	}
	Timecode  uint32
	Sequence  uint32
	Memory    uint32
	UserPtr   unsafe.Pointer
	Length    uint32
	Reserved2 uint32
	Reserved  uint32
}

// v4l2_hw_freq_seek
type HardwareFrequencySeek struct {
	Tuner       uint32
	Type        uint32
	Seek_upward uint32
	Wrap_around uint32
	Spacing     uint32
	Rangelow    uint32
	Rangehigh   uint32
	Reserved    [5]uint32
}
