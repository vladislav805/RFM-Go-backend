package radio

type Region uint32

const (
	REGION_US         Region = iota // = 0
	REGION_EU                       // = 1
	REGION_JAPAN                    // = 2
	REGION_JAPAN_WIDE               // = 3
	REGION_OTHER                    // = 4
)

var RegionBands map[Region][2]uint32

func init() {
	RegionBands = map[Region][2]uint32{
		REGION_US:         {87_500, 108_000},
		REGION_EU:         {87_500, 108_000},
		REGION_JAPAN:      {76_000, 90000},
		REGION_JAPAN_WIDE: {76_000, 108_000},
		REGION_OTHER:      {76_000, 108_000},
	}
}
