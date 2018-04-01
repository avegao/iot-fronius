package froniusCurrentDataInverter

import "github.com/avegao/iot-api/entity/fronius"

type CurrentDataInverter struct {
	Body struct {
		// DayEnergy Energy generated on current day
		DayEnergy struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"DAY_ENERGY"`
		// Pac AC power
		Pac struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"PAC"`
		// TotalEnergy Energy generated overall
		TotalEnergy struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"TOTAL_ENERGY"`
		// YearEnergy Energy generated in current year
		YearEnergy struct {
			Unit   string         `json:"Unit"`
			Values map[string]int `json:"Values"`
		} `json:"YEAR_ENERGY"`
	} `json:"Body"`
	Head fronius.ResponseHeader `json:"Head"`
}
