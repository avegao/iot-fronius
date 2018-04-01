package fronius

import "time"

type ResponseHeader struct {
	RequestArguments struct {
		Query string `json:"Query"`
		Scope string `json:"Scope"`
	} `json:"RequestArguments"`
	Status struct {
		Code        int    `json:"Code"`
		Reason      string `json:"Reason"`
		UserMessage string `json:"UserMessage"`
	} `json:"Status"`
	Timestamp time.Time `json:"Timestamp"`
}
