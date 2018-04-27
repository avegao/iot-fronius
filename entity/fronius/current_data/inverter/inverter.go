package froniusCurrentDataInverter

import (
	"time"
	"github.com/avegao/gocondi"
	"fmt"
	"errors"
)

type CurrentDataInverter struct {
	Id          uint64
	DayEnergy   int32
	Pac         int32
	TotalEnergy int32
	YearEnergy  int32
	Timestamp   time.Time
	CreatedAt   time.Time
}

func (inverter *CurrentDataInverter) getTableName() string {
	return "\"fronius\".\"current_data_inverter\""
}

func (inverter *CurrentDataInverter) Persist() (error) {
	if inverter.Id == 0 {
		id, err := inverter.insert()

		inverter.Id = id

		if err != nil {
			return err
		}
	} else {
		return errors.New("update not supported yet")
	}

	return nil
}

func (inverter *CurrentDataInverter) insert() (id uint64, err error) {
	const logTag = "CurrentDataInverter.Site.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("CurrentDataInverter", *inverter).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"day_energy",
			"pac",
			"total_energy",
			"year_energy",
			"timestamp",
			"created_at"
		) VALUES (
			$1,$2,$3,$4,$5,$6
		) RETURNING id::int;`,
		inverter.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", *inverter).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if err := db.QueryRow(insertQuery,
		inverter.DayEnergy,
		inverter.Pac,
		inverter.TotalEnergy,
		inverter.YearEnergy,
		inverter.Timestamp,
		inverter.CreatedAt,
	).Scan(&id); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("CurrentDataInverter", *inverter).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		WithField("id", id).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
