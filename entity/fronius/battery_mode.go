package fronius

const (
	BatteryModeDisabled               BatteryMode = "disabled"
	BatteryModeNormal                 BatteryMode = "normal"
	BatteryModeService                BatteryMode = "service"
	BatteryModeChargeBoost            BatteryMode = "charge boost"
	BatteryModeNearylyDepleted        BatteryMode = "nearly depleted"
	BatteryModeSuspended              BatteryMode = "suspended"
	BatteryModeCalibrate              BatteryMode = "calibrate"
	BatteryModeGridSupport            BatteryMode = "grid support"
	BatteryModeDepleteRecovery        BatteryMode = "deplete recovery"
	BatteryModeNonOperableVoltage     BatteryMode = "non operable (voltage)"
	BatteryModeNonOperableTemperature BatteryMode = "non operable (temperature)"
	BatteryModePreheating             BatteryMode = "preheating"
	BatteryModeStartup                BatteryMode = "startup"
)

type BatteryMode string

func (mode BatteryMode) String() string {
	return string(mode)
}
