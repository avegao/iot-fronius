package froniusCurrentDataMeter

import (
	"time"
	"fmt"
	"github.com/avegao/gocondi"
	"database/sql"
	"github.com/avegao/iot-api/entity/fronius"
	"github.com/pkg/errors"
)

type CurrentDataMeter struct {
	Body CurrentDataMeterBody   `json:"Body"`
	Head fronius.ResponseHeader `json:"Head"`
}

type CurrentDataMeterBody map[string]CurrentDataMeterBodyElement

type CurrentDataMeterBodyElement struct {
	Id              uint64
	CurrentACPhase1 float64  `json:"Current_AC_Phase_1"`
	CurrentACPhase2 *float64 `json:"Current_AC_Phase_2"`
	CurrentACPhase3 *float64 `json:"Current_AC_Phase_3"`
	CurrentACSum    float64  `json:"Current_AC_Sum"`
	Details struct {
		Manufacturer string `json:"Manufacturer"`
		Model        string `json:"Model"`
		Serial       string `json:"Serial"`
	} `json:"Details"`
	Enable                            int             `json:"Enable"`
	EnergyReactiveVArACPhase1Consumed int             `json:"EnergyReactive_VArAC_Phase_1_Consumed"`
	EnergyReactiveVArACPhase1Produced int             `json:"EnergyReactive_VArAC_Phase_1_Produced"`
	EnergyReactiveVArACPhase2Consumed sql.NullInt64   `json:"EnergyReactive_VArAC_Phase_2_Consumed"`
	EnergyReactiveVArACPhase2Produced sql.NullInt64   `json:"EnergyReactive_VArAC_Phase_2_Produced"`
	EnergyReactiveVArACPhase3Consumed sql.NullInt64   `json:"EnergyReactive_VArAC_Phase_3_Consumed"`
	EnergyReactiveVArACPhase3Produced sql.NullInt64   `json:"EnergyReactive_VArAC_Phase_3_Produced"`
	EnergyReactiveVArACSumConsumed    int             `json:"EnergyReactive_VArAC_Sum_Consumed"`
	EnergyReactiveVArACSumProduced    int             `json:"EnergyReactive_VArAC_Sum_Produced"`
	EnergyRealWACMinusAbsolute        int             `json:"EnergyReal_WAC_Minus_Absolute"`
	EnergyRealWACPhase1Consumed       int             `json:"EnergyReal_WAC_Phase_1_Consumed"`
	EnergyRealWACPhase1Produced       int             `json:"EnergyReal_WAC_Phase_1_Produced"`
	EnergyRealWACPhase2Consumed       sql.NullInt64   `json:"EnergyReal_WAC_Phase_2_Consumed"`
	EnergyRealWACPhase2Produced       sql.NullInt64   `json:"EnergyReal_WAC_Phase_2_Produced"`
	EnergyRealWACPhase3Consumed       sql.NullInt64   `json:"EnergyReal_WAC_Phase_3_Consumed"`
	EnergyRealWACPhase3Produced       sql.NullInt64   `json:"EnergyReal_WAC_Phase_3_Produced"`
	EnergyRealWACPlusAbsolute         int             `json:"EnergyReal_WAC_Plus_Absolute"`
	EnergyRealWACSumConsumed          int             `json:"EnergyReal_WAC_Sum_Consumed"`
	EnergyRealWACSumProduced          int             `json:"EnergyReal_WAC_Sum_Produced"`
	FrequencyPhaseAverage             float64         `json:"Frequency_Phase_Average"`
	MeterLocationCurrent              int             `json:"Meter_Location_Current"`
	PowerApparentSPhase1              float64         `json:"PowerApparent_S_Phase_1"`
	PowerApparentSPhase2              sql.NullFloat64 `json:"PowerApparent_S_Phase_2"`
	PowerApparentSPhase3              sql.NullFloat64 `json:"PowerApparent_S_Phase_3"`
	PowerApparentSSum                 float64         `json:"PowerApparent_S_Sum"`
	PowerFactorPhase1                 float64         `json:"PowerFactor_Phase_1"`
	PowerFactorPhase2                 sql.NullFloat64 `json:"PowerFactor_Phase_2"`
	PowerFactorPhase3                 sql.NullFloat64 `json:"PowerFactor_Phase_3"`
	PowerFactorSum                    float64         `json:"PowerFactor_Sum"`
	PowerReactiveQPhase1              float64         `json:"PowerReactive_Q_Phase_1"`
	PowerReactiveQPhase2              sql.NullFloat64 `json:"PowerReactive_Q_Phase_2"`
	PowerReactiveQPhase3              sql.NullFloat64 `json:"PowerReactive_Q_Phase_3"`
	PowerReactiveQSum                 float64         `json:"PowerReactive_Q_Sum"`
	PowerRealPPhase1                  float64         `json:"PowerReal_P_Phase_1"`
	PowerRealPPhase2                  sql.NullFloat64 `json:"PowerReal_P_Phase_2"`
	PowerRealPPhase3                  sql.NullFloat64 `json:"PowerReal_P_Phase_3"`
	PowerRealPSum                     float64         `json:"PowerReal_P_Sum"`
	TimeStamp                         int             `json:"TimeStamp"`
	Visible                           int             `json:"Visible"`
	VoltageACPhase1                   float64         `json:"Voltage_AC_Phase_1"`
	VoltageACPhase2                   sql.NullFloat64 `json:"Voltage_AC_Phase_2"`
	VoltageACPhase3                   sql.NullFloat64 `json:"Voltage_AC_Phase_3"`
}

func (currentData CurrentDataMeterBodyElement) getTableName() string {
	return "\"fronius\".\"current_data_meter\""
}

func (currentData CurrentDataMeterBodyElement) Persist() (err error) {
	if currentData.Id == 0 {
		err = currentData.insert()
	} else {
		err = errors.New("update not supported yet")
	}

	return
}

func (currentData CurrentDataMeterBodyElement) insert() (err error) {
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
