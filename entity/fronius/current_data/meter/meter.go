package froniusCurrentDataMeter

import (
	"time"
	"fmt"
	"github.com/avegao/gocondi"
	"github.com/pkg/errors"
)

type CurrentDataMeter struct {
	Id              uint64
	CurrentACPhase1 float64
	CurrentACPhase2 float64
	CurrentACPhase3 float64
	CurrentACSum    float64
	Details struct {
		Manufacturer string
		Model        string
		Serial       string
	}
	Enable                            bool
	EnergyReactiveVArACPhase1Consumed uint32
	EnergyReactiveVArACPhase1Produced uint32
	EnergyReactiveVArACPhase2Consumed uint32
	EnergyReactiveVArACPhase2Produced uint32
	EnergyReactiveVArACPhase3Consumed uint32
	EnergyReactiveVArACPhase3Produced uint32
	EnergyReactiveVArACSumConsumed    uint32
	EnergyReactiveVArACSumProduced    uint32
	EnergyRealWACMinusAbsolute        uint32
	EnergyRealWACPhase1Consumed       uint32
	EnergyRealWACPhase1Produced       uint32
	EnergyRealWACPhase2Consumed       uint32
	EnergyRealWACPhase2Produced       uint32
	EnergyRealWACPhase3Consumed       uint32
	EnergyRealWACPhase3Produced       uint32
	EnergyRealWACPlusAbsolute         uint32
	EnergyRealWACSumConsumed          uint32
	EnergyRealWACSumProduced          uint32
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
	TimeStamp                         uint32
	Visible                           bool
	VoltageACPhase1                   float64
	VoltageACPhase2                   float64
	VoltageACPhase3                   float64
}

func (currentData CurrentDataMeter) getTableName() string {
	return "\"fronius\".\"current_data_meter\""
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
	const logTag = "CurrentDataMeterBody.insert()"
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
		currentData.getTableName(),
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
		currentData.CurrentACPhase1,
		currentData.CurrentACPhase2,
		currentData.CurrentACPhase3,
		currentData.CurrentACSum,
		currentData.Enable,
		currentData.EnergyReactiveVArACPhase1Consumed,
		currentData.EnergyReactiveVArACPhase1Produced,
		currentData.EnergyReactiveVArACPhase2Consumed,
		currentData.EnergyReactiveVArACPhase2Produced,
		currentData.EnergyReactiveVArACPhase3Consumed,
		currentData.EnergyReactiveVArACPhase3Produced,
		currentData.EnergyReactiveVArACSumConsumed,
		currentData.EnergyReactiveVArACSumProduced,
		currentData.EnergyRealWACMinusAbsolute,
		currentData.EnergyRealWACPhase1Consumed,
		currentData.EnergyRealWACPhase1Produced,
		currentData.EnergyRealWACPhase2Consumed,
		currentData.EnergyRealWACPhase2Produced,
		currentData.EnergyRealWACPhase3Consumed,
		currentData.EnergyRealWACPhase3Produced,
		currentData.EnergyRealWACPlusAbsolute,
		currentData.EnergyRealWACSumConsumed,
		currentData.EnergyRealWACSumProduced,
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
		time.Unix(int64(currentData.TimeStamp), 0),
		currentData.Visible,
		currentData.VoltageACPhase1,
		currentData.VoltageACPhase2,
		currentData.VoltageACPhase3,
	); err != nil {
		logger.WithError(err).Panicf("%s -> STOP", logTag)
	}

	logger.
		WithField("currentData", currentData).
		WithField("time_spent", time.Since(startTimeLog).Nanoseconds()).
		Debugf(fmt.Sprintf("%s -> END", logTag))

	return
}
