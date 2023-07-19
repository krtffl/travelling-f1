package ergast

type Response struct {
	MRData MRData `json:"MRData"`
}

type MRData struct {
	Series    string    `json:"series"`
	URL       string    `json:"url"`
	Limit     string    `json:"limit"`
	Offset    string    `json:"offset"`
	Total     string    `json:"total"`
	RaceTable RaceTable `json:"raceTable"`
}

type RaceTable struct {
	Season string `json:"season"`
	Races  []Race `json:"races"`
}

type Race struct {
	Season   string  `json:"season"`
	Round    string  `json:"round"`
	URL      string  `json:"url"`
	RaceName string  `json:"raceName"`
	Circuit  Circuit `json:"circuit"`
}

type Circuit struct {
	CircuitId   string   `json:"circuitId"`
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"location"`
}

type Location struct {
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Locality string `json:"locality"`
	Country  string `json:"country"`
}
