package froniusCurrentPowerflow

import (
	"time"
	"fmt"
	"github.com/avegao/gocondi"
	"github.com/avegao/iot-fronius/entity/fronius"
)

type Inverter struct {
	BatteryMode *fronius.BatteryMode

	// DeviceType Device type of Inverter
	DeviceType uint16

	// EnergyDay Energy in Wh this day, null if no Inverter is connected
	EnergyDay float64

	// EnergyDay Energy in Wh ever since, null if no Inverter is connected
	EnergyTotal float64

	// EnergyDay Energy in Wh this year, null if no Inverter is connected
	EnergyYear float64

	// CurrentPower current power in Watt, null if not running
	CurrentPower float64

	// Soc Current state of charge in % ( 0 - 100% )
	Soc uint8
}

func (inverter Inverter) getTableName() string {
	return "\"fronius\".\"current_powerflow_inverter\""
}

func (inverter Inverter) Persist(siteId int) (error) {
	return inverter.insert(siteId)
}

func (inverter Inverter) insert(siteId int) (err error) {
	const logTag = "CurrentPowerflow.Inverter.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("Inverter", inverter).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"id_site",
			"battery_mode",
			"device_type",
			"energy_day",
			"energy_year",
			"energy_total",
			"current_power",
			"soc"
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8
		);`,
		inverter.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", inverter).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if _, err = db.Exec(insertQuery,
		siteId,
		inverter.BatteryMode,
		inverter.DeviceType,
		inverter.EnergyDay,
		inverter.EnergyYear,
		inverter.EnergyTotal,
		inverter.CurrentPower,
		inverter.Soc,
	); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("Inverter", inverter).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
