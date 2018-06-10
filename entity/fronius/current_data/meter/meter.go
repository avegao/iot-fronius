package froniusCurrentDataMeter

import (
	"time"
	"fmt"
	"github.com/avegao/gocondi"
	"github.com/pkg/errors"
	pb "github.com/avegao/iot-fronius/resource/grpc"
)

const tableName = "\"fronius\".\"current_data_meter\""

type CurrentDataMeter struct {
	Id              uint64
	CurrentAcPhase1 float64
	CurrentAcPhase2 float64
	CurrentAcPhase3 float64
	CurrentAcSum    float64
	Details struct {
		Manufacturer string
		Model        string
		Serial       string
	}
	Enable                            bool
	EnergyReactiveVArAcPhase1Consumed uint32
	EnergyReactiveVArAcPhase1Produced uint32
	EnergyReactiveVArAcPhase2Consumed uint32
	EnergyReactiveVArAcPhase2Produced uint32
	EnergyReactiveVArAcPhase3Consumed uint32
	EnergyReactiveVArAcPhase3Produced uint32
	EnergyReactiveVArAcSumConsumed    uint32
	EnergyReactiveVArAcSumProduced    uint32
	EnergyRealWAcMinusAbsolute        uint32
	EnergyRealWAcPhase1Consumed       uint32
	EnergyRealWAcPhase1Produced       uint32
	EnergyRealWAcPhase2Consumed       uint32
	EnergyRealWAcPhase2Produced       uint32
	EnergyRealWAcPhase3Consumed       uint32
	EnergyRealWAcPhase3Produced       uint32
	EnergyRealWAcPlusAbsolute         uint32
	EnergyRealWAcSumConsumed          uint32
	EnergyRealWAcSumProduced          uint32
	FrequencyPhaseAverage             float64
	MeterLocationCurrent              uint32
	PowerApparentSPhase1              float64
	PowerApparentSPhase2              float64
	PowerApparentSPhase3              float64
	PowerApparentSSum                 float64
	PowerFactorPhase1                 float64
	PowerFactorPhase2                 float64
	PowerFactorPhase3                 float64
	PowerFactorSum                    float64
	PowerReactiveQPhase1              float64
	PowerReactiveQPhase2              float64
	PowerReactiveQPhase3              float64
	PowerReactiveQSum                 float64
	PowerRealPPhase1                  float64
	PowerRealPPhase2                  float64
	PowerRealPPhase3                  float64
	PowerRealPSum                     float64
	Timestamp                         time.Time
	VoltageAcPhase1                   float64
	VoltageAcPhase2                   float64
	VoltageAcPhase3                   float64
}

func (currentData CurrentDataMeter) Persist() (err error) {
	if currentData.Id == 0 {
		err = currentData.insert()
	} else {
		err = errors.New("update not supported yet")
	}

	return
}

func (currentData CurrentDataMeter) insert() (err error) {
	const logTag = "CurrentDataMeter.insert()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("currentData", currentData).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (
			"current_ac_phase_1",
			"current_ac_phase_2",
			"current_ac_phase_3",
			"current_ac_sum",
			"enable",
			"energy_reactive_v_ar_ac_phase_1_consumed",
			"energy_reactive_v_ar_ac_phase_1_produced",
			"energy_reactive_v_ar_ac_phase_2_consumed",
			"energy_reactive_v_ar_ac_phase_2_produced",
			"energy_reactive_v_ar_ac_phase_3_consumed",
			"energy_reactive_v_ar_ac_phase_3_produced",
			"energy_reactive_v_ar_ac_sum_consumed",
			"energy_reactive_v_ar_ac_sum_produced",
			"energy_real_w_ac_minus_absolute",
			"energy_real_w_ac_phase_1_consumed",
			"energy_real_w_ac_phase_1_produced",
			"energy_real_w_ac_phase_2_consumed",
			"energy_real_w_ac_phase_2_produced",
			"energy_real_w_ac_phase_3_consumed",
			"energy_real_w_ac_phase_3_produced",
			"energy_real_w_ac_plus_absolute",
			"energy_real_w_ac_sum_consumed",
			"energy_real_w_ac_sum_produced",
			"frequency_phase_average",
			"meter_location_current",
			"power_apparent_s_phase_1",
			"power_apparent_s_phase_2",
			"power_apparent_s_phase_3",
			"power_apparent_s_sum",
			"power_factor_phase_1",
			"power_factor_phase_2",
			"power_factor_phase_3",
			"power_factor_sum",
			"power_reactive_q_phase_1",
			"power_reactive_q_phase_2",
			"power_reactive_q_phase_3",
			"power_reactive_q_sum",
			"power_real_p_phase_1",
			"power_real_p_phase_2",
			"power_real_p_phase_3",
			"power_real_p_sum",
			"timestamp",
			"visible",
			"voltage_ac_phase_1",
			"voltage_ac_phase_2",
			"voltage_ac_phase_3"
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,
			$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,
			$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,
			$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,
			$40,$41,$42,$43,$44,$45,$46
		)`,
		tableName,
	)

	logger.
		WithField("query", insertQuery).
		WithField("parameters", currentData).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> Query to execute", logTag))

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.WithError(err).Debugf("%s -> STOP", logTag)
		return
	}

	if _, err := db.Exec(insertQuery,
		currentData.CurrentAcPhase1,
		currentData.CurrentAcPhase2,
		currentData.CurrentAcPhase3,
		currentData.CurrentAcSum,
		currentData.Enable,
		currentData.EnergyReactiveVArAcPhase1Consumed,
		currentData.EnergyReactiveVArAcPhase1Produced,
		currentData.EnergyReactiveVArAcPhase2Consumed,
		currentData.EnergyReactiveVArAcPhase2Produced,
		currentData.EnergyReactiveVArAcPhase3Consumed,
		currentData.EnergyReactiveVArAcPhase3Produced,
		currentData.EnergyReactiveVArAcSumConsumed,
		currentData.EnergyReactiveVArAcSumProduced,
		currentData.EnergyRealWAcMinusAbsolute,
		currentData.EnergyRealWAcPhase1Consumed,
		currentData.EnergyRealWAcPhase1Produced,
		currentData.EnergyRealWAcPhase2Consumed,
		currentData.EnergyRealWAcPhase2Produced,
		currentData.EnergyRealWAcPhase3Consumed,
		currentData.EnergyRealWAcPhase3Produced,
		currentData.EnergyRealWAcPlusAbsolute,
		currentData.EnergyRealWAcSumConsumed,
		currentData.EnergyRealWAcSumProduced,
		currentData.FrequencyPhaseAverage,
		currentData.MeterLocationCurrent,
		currentData.PowerApparentSPhase1,
		currentData.PowerApparentSPhase2,
		currentData.PowerApparentSPhase3,
		currentData.PowerApparentSSum,
		currentData.PowerFactorPhase1,
		currentData.PowerFactorPhase2,
		currentData.PowerFactorPhase3,
		currentData.PowerFactorSum,
		currentData.PowerReactiveQPhase1,
		currentData.PowerReactiveQPhase2,
		currentData.PowerReactiveQPhase3,
		currentData.PowerReactiveQSum,
		currentData.PowerRealPPhase1,
		currentData.PowerRealPPhase2,
		currentData.PowerRealPPhase3,
		currentData.PowerRealPSum,
		currentData.Timestamp,
		currentData.VoltageAcPhase1,
		currentData.VoltageAcPhase2,
		currentData.VoltageAcPhase3,
	); err != nil {
		logger.WithError(err).Errorf("%s -> STOP", logTag)
	}

	logger.
		WithField("currentData", currentData).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}

func (currentData CurrentDataMeter) ToGrpc() *pb.GetDataMeterResponse {
	return &pb.GetDataMeterResponse{
		CurrentAcPhase1:                   currentData.CurrentAcPhase1,
		CurrentAcPhase2:                   currentData.CurrentAcPhase2,
		CurrentAcPhase3:                   currentData.CurrentAcPhase3,
		CurrentAcSum:                      currentData.CurrentAcSum,
		EnergyReactiveVArAcPhase1Consumed: currentData.EnergyReactiveVArAcPhase1Consumed,
		EnergyReactiveVArAcPhase1Produced: currentData.EnergyReactiveVArAcPhase1Produced,
		EnergyReactiveVArAcPhase2Consumed: currentData.EnergyReactiveVArAcPhase2Consumed,
		EnergyReactiveVArAcPhase2Produced: currentData.EnergyReactiveVArAcPhase2Produced,
		EnergyReactiveVArAcPhase3Consumed: currentData.EnergyReactiveVArAcPhase3Consumed,
		EnergyReactiveVArAcPhase3Produced: currentData.EnergyReactiveVArAcPhase3Produced,
		EnergyReactiveVArAcSumConsumed:    currentData.EnergyReactiveVArAcSumConsumed,
		EnergyReactiveVArAcSumProduced:    currentData.EnergyReactiveVArAcSumProduced,
		EnergyRealWAcMinusAbsolute:        currentData.EnergyRealWAcMinusAbsolute,
		EnergyRealWAcPhase1Consumed:       currentData.EnergyRealWAcPhase1Consumed,
		EnergyRealWAcPhase1Produced:       currentData.EnergyRealWAcPhase1Produced,
		EnergyRealWAcPhase2Consumed:       currentData.EnergyRealWAcPhase2Consumed,
		EnergyRealWAcPhase2Produced:       currentData.EnergyRealWAcPhase2Produced,
		EnergyRealWAcPhase3Consumed:       currentData.EnergyRealWAcPhase3Consumed,
		EnergyRealWAcPhase3Produced:       currentData.EnergyRealWAcPhase3Produced,
		EnergyRealWAcPlusAbsolute:         currentData.EnergyRealWAcPlusAbsolute,
		EnergyRealWAcSumConsumed:          currentData.EnergyRealWAcSumConsumed,
		EnergyRealWAcSumProduced:          currentData.EnergyRealWAcSumProduced,
		FrequencyPhaseAverage:             currentData.FrequencyPhaseAverage,
		PowerApparentSPhase1:              currentData.PowerApparentSPhase1,
		PowerApparentSPhase2:              currentData.PowerApparentSPhase2,
		PowerApparentSPhase3:              currentData.PowerApparentSPhase3,
		PowerApparentSSum:                 currentData.PowerApparentSSum,
		PowerFactorPhase1:                 currentData.PowerFactorPhase1,
		PowerFactorPhase2:                 currentData.PowerFactorPhase2,
		PowerFactorPhase3:                 currentData.PowerFactorPhase3,
		PowerFactorSum:                    currentData.PowerFactorSum,
		PowerReactiveQPhase1:              currentData.PowerReactiveQPhase1,
		PowerReactiveQPhase2:              currentData.PowerReactiveQPhase2,
		PowerReactiveQPhase3:              currentData.PowerReactiveQPhase3,
		PowerReactiveQSum:                 currentData.PowerReactiveQSum,
		PowerRealPPhase1:                  currentData.PowerRealPPhase1,
		PowerRealPPhase2:                  currentData.PowerRealPPhase2,
		PowerRealPPhase3:                  currentData.PowerRealPPhase3,
		PowerRealPSum:                     currentData.PowerRealPSum,
		Timestamp:                         currentData.Timestamp.Unix(),
		VoltageAcPhase1:                   currentData.VoltageAcPhase1,
		VoltageAcPhase2:                   currentData.VoltageAcPhase2,
		VoltageAcPhase3:                   currentData.VoltageAcPhase3,
	}
}

func FindByDate(startDate time.Time, endDate time.Time) ([]CurrentDataMeter, error) {
	const logTag = "froniusCurrentDataMeter.FindByDate()"
	startTimeLog := time.Now()
	container := gocondi.GetContainer()

	logger := container.GetLogger()
	logger.
		WithField("startDate", startDate).
		WithField("endDate", endDate).
		Debugf(fmt.Sprintf("%s -> START", logTag))

	query := "SELECT id," +
		"current_ac_phase_1," +
		"current_ac_phase_2," +
		"current_ac_phase_3," +
		"current_ac_sum," +
		"energy_reactive_v_ar_ac_phase_1_consumed," +
		"energy_reactive_v_ar_ac_phase_1_produced," +
		"energy_reactive_v_ar_ac_phase_2_consumed," +
		"energy_reactive_v_ar_ac_phase_2_produced," +
		"energy_reactive_v_ar_ac_phase_3_consumed," +
		"energy_reactive_v_ar_ac_phase_3_produced," +
		"energy_reactive_v_ar_ac_sum_consumed," +
		"energy_reactive_v_ar_ac_sum_produced," +
		"energy_real_w_ac_minus_absolute," +
		"energy_real_w_ac_phase_1_consumed," +
		"energy_real_w_ac_phase_1_produced," +
		"energy_real_w_ac_phase_2_consumed," +
		"energy_real_w_ac_phase_2_produced," +
		"energy_real_w_ac_phase_3_consumed," +
		"energy_real_w_ac_phase_3_produced," +
		"energy_real_w_ac_sum_consumed," +
		"energy_real_w_ac_sum_produced," +
		"frequency_phase_average," +
		"power_apparent_s_phase_1," +
		"power_apparent_s_phase_2," +
		"power_apparent_s_phase_3," +
		"power_apparent_s_sum," +
		"power_factor_phase_1," +
		"power_factor_phase_2," +
		"power_factor_phase_3," +
		"power_factor_sum," +
		"power_reactive_q_phase_1," +
		"power_reactive_q_phase_2," +
		"power_reactive_q_phase_3," +
		"power_reactive_q_sum," +
		"power_real_p_phase_1," +
		"power_real_p_phase_2," +
		"power_real_p_phase_3," +
		"power_real_p_sum," +
		"\"timestamp\"," +
		"voltage_ac_phase_1," +
		"voltage_ac_phase_2," +
		"voltage_ac_phase_3 "

	query += "FROM " + tableName + " " +
		"WHERE \"timestamp\" >= ($1)::timestamptz " +
		"AND \"timestamp\" < ($2)::timestamptz"

	db, err := container.GetDefaultDatabase()
	if err != nil {
		logger.
			WithField("startDate", startDate).
			WithField("endDate", endDate).
			WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
			WithError(err).Debugf("%s -> STOP", logTag)
		return nil, err
	}

	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		logger.
			WithField("startDate", startDate).
			WithField("endDate", endDate).
			WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
			WithError(err).Debugf("%s -> STOP", logTag)
		return nil, err
	}

	defer rows.Close()

	results := make([]CurrentDataMeter, 0)

	for rows.Next() {
		var result CurrentDataMeter

		if err := rows.Scan(
			&result.Id,
			&result.CurrentAcPhase1,
			&result.CurrentAcPhase2,
			&result.CurrentAcPhase3,
			&result.CurrentAcSum,
			&result.EnergyReactiveVArAcPhase1Consumed,
			&result.EnergyReactiveVArAcPhase1Produced,
			&result.EnergyReactiveVArAcPhase2Consumed,
			&result.EnergyReactiveVArAcPhase2Produced,
			&result.EnergyReactiveVArAcPhase3Consumed,
			&result.EnergyReactiveVArAcPhase3Produced,
			&result.EnergyReactiveVArAcSumConsumed,
			&result.EnergyReactiveVArAcSumProduced,
			&result.EnergyRealWAcMinusAbsolute,
			&result.EnergyRealWAcPhase1Consumed,
			&result.EnergyRealWAcPhase1Produced,
			&result.EnergyRealWAcPhase2Consumed,
			&result.EnergyRealWAcPhase2Produced,
			&result.EnergyRealWAcPhase3Consumed,
			&result.EnergyRealWAcPhase3Produced,
			&result.EnergyRealWAcSumConsumed,
			&result.EnergyRealWAcSumProduced,
			&result.FrequencyPhaseAverage,
			&result.PowerApparentSPhase1,
			&result.PowerApparentSPhase2,
			&result.PowerApparentSPhase3,
			&result.PowerApparentSSum,
			&result.PowerFactorPhase1,
			&result.PowerFactorPhase2,
			&result.PowerFactorPhase3,
			&result.PowerFactorSum,
			&result.PowerReactiveQPhase1,
			&result.PowerReactiveQPhase2,
			&result.PowerReactiveQPhase3,
			&result.PowerReactiveQSum,
			&result.PowerRealPPhase1,
			&result.PowerRealPPhase2,
			&result.PowerRealPPhase3,
			&result.PowerRealPSum,
			&result.Timestamp,
			&result.VoltageAcPhase1,
			&result.VoltageAcPhase2,
			&result.VoltageAcPhase3,
		); err != nil {
			logger.
				WithField("startDate", startDate).
				WithField("endDate", endDate).
				WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
				WithError(err).Debugf("%s -> STOP", logTag)
			return nil, err
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		logger.
			WithField("startDate", startDate).
			WithField("endDate", endDate).
			WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
			WithError(err).Debugf("%s -> STOP", logTag)
		return nil, err
	}

	logger.
		WithField("startDate", startDate).
		WithField("endDate", endDate).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return results, err
}
