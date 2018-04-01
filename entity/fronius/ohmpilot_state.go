package fronius

const (
	OhmpilotStateNormal               OhmpilotState = "normal"
	OhmpilotStateMinTemperature       OhmpilotState = "min-temperature"
	OhmpilotStateLegionallaProtection OhmpilotState = "legionella-protection"
	OhmpilotStateFault                OhmpilotState = "fault"
	OhmpilotStateWarning              OhmpilotState = "warning"
	OhmpilotStateBoost                OhmpilotState = "boost"
)

type OhmpilotState string

func (state OhmpilotState) String() string {
	return string(state)
}
