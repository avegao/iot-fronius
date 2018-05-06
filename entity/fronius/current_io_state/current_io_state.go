package froniusCurrentIoState

import (
	"time"
	"github.com/avegao/gocondi"
	"fmt"
	"errors"
)

type CurrentIoState struct {
	Id        uint64
	PinNumber int32
	Function  string
	Type      string
	Direction string
	Set       bool
	CreatedAt time.Time
}

func (state *CurrentIoState) getTableName() string {
	return "\"fronius\".\"current_io_state\""
}

func (state *CurrentIoState) Persist() (error) {
	if state.Id == 0 {
		id, err := state.insert()

		state.Id = id

		if err != nil {
			return err
		}
	} else {
		return errors.New("update not supported yet")
	}

	return nil
}

func (state *CurrentIoState) insert() (id uint64, err error) {
	const logTag = "CurrentIoState.Site.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("CurrentIoState", *state).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"pin_number",
			"function",
			"type",
			"direction",
			"set"
		) VALUES (
			$1,$2,$3,$4,$5
		) RETURNING id::int;`,
		state.getTableName(),
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", *state).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)

		return 0, err
	}

	if err := db.QueryRow(insertQuery,
		state.PinNumber,
		state.Function,
		state.Type,
		state.Direction,
		state.Set,
	).Scan(&id); err != nil {
		logger.WithError(err).Errorf("%s -> STOP", logTag)

		return 0, err
	}

	logger.
		WithField("CurrentIoState", *state).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		WithField("id", id).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
