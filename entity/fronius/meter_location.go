package fronius

const (
	MeterLocationLoad   MeterLocation = "load"
	MeterLocationGrid   MeterLocation = "grid"
	MeterLocationUnkown MeterLocation = "unknown"
)

var (
	MeterLocationMap = map[string]MeterLocation{
		"load":    MeterLocationLoad,
		"grid":    MeterLocationGrid,
		"unknown": MeterLocationUnkown,
	}
)

type MeterLocation string

func (location MeterLocation) String() string {
	return string(location)
}

func NewMeterLocationFromString(meterLocationString string) MeterLocation {
	return MeterLocationMap[meterLocationString]
}

