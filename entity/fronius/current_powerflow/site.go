package froniusCurrentPowerflow

import (
	"time"
	"github.com/avegao/gocondi"
	"fmt"
	"errors"
	"github.com/avegao/iot-fronius/entity/fronius"
)

type Site struct {
	Id   int
	Mode fronius.SiteMode

	// BatteryStandby True when battery is in standby
	BatteryStandby bool

	// BackupMode Field is available if configured (false) or active (true)
	// if not available, mandatory config is not set.
	BackupMode bool

	// PowerFromGrid This value is null if no meter is enabled (+ from grid, - to grid)
	PowerFromGrid float64

	// PowerLoad This value is null if no meter is enabled (+ generator, - consumer)
	PowerLoad float64

	// PowerAkku This value is null if no battery is active (+ charge, - discharge)
	PowerAkku float64

	// PowerFromPV This value is null if Inverter is not running (+ production (default))
	PowerFromPV float64

	// RelativeSelfConsumption Current relative self consumption in %, null if no smart meter is connected
	RelativeSelfConsumption uint8

	// RelativeAutonomy Current relative autonomy in %, null if no smart meter is connected
	RelativeAutonomy uint8

	MeterLocation fronius.MeterLocation

	// EnergyDay Energy [Wh] this day, null if no Inverter is connected
	EnergyDay float64

	// EnergyYear Energy [Wh] this year, null if no Inverter is connected
	EnergyYear float64

	// EnergyTotal Energy [Wh] ever since, null if no Inverter is connected
	EnergyTotal float64
}

func (site Site) getTableName() string {
	return "\"fronius\".\"current_powerflow_site\""
}

func (site Site) Persist() (Site, error) {
	if site.Id == 0 {
		id, err := site.insert()

		site.Id = id

		if err != nil {
			return site, err
		}
	} else {
		return site, errors.New("update not supported yet")
	}

	return site, nil
}

func (site *Site) insert() (id int, err error) {
	const logTag = "CurrentPowerflow.Site.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("Site", *site).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"battery_standby",
			"backup_mode",
			"power_from_grid",
			"power_load",
			"power_akku",
			"power_from_pv",
			"relative_self_consumption",
			"relative_autonomy",
			"meter_location",
			"energy_day",
			"energy_year",
			"energy_total"
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,
			$10,$11,$12
		) RETURNING id::int;`,
		site.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", *site).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if err := db.QueryRow(insertQuery,
		site.BatteryStandby,
		site.BackupMode,
		site.PowerFromGrid,
		site.PowerLoad,
		site.PowerAkku,
		site.PowerFromPV,
		site.RelativeSelfConsumption,
		site.RelativeAutonomy,
		site.MeterLocation,
		site.EnergyDay,
		site.EnergyYear,
		site.EnergyTotal,
	).Scan(&id); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("Site", *site).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		WithField("id", id).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
