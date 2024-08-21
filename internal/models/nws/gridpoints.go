package nws

type GridPoints struct {
	Context  []any  `json:"-"`
	Type     string `json:"-"`
	Geometry struct {
		Type        string        `json:"-"`
		Coordinates [][][]float64 `json:"-"`
	} `json:"-"`
	Properties struct {
		Units             string `json:"-"`
		ForecastGenerator string `json:"-"`
		GeneratedAt       string `json:"-"`
		UpdateTime        string `json:"-"`
		ValidTimes        string `json:"-"`
		Elevation         struct {
			UnitCode string  `json:"-"`
			Value    float64 `json:"-"`
		} `json:"-"`
		Periods []struct {
			Number                     int    `json:"-"`
			Name                       string `json:"-"`
			StartTime                  string `json:"-"`
			EndTime                    string `json:"-"`
			IsDaytime                  bool   `json:"-"`
			Temperature                int    `json:"temperature"`
			TemperatureUnit            string `json:"-"`
			TemperatureTrend           string `json:"-"`
			ProbabilityOfPrecipitation struct {
				UnitCode string `json:"-"`
				Value    any    `json:"-"`
			} `json:"-"`
			WindSpeed        string `json:"-"`
			WindDirection    string `json:"-"`
			Icon             string `json:"-"`
			ShortForecast    string `json:"shortForecast"`
			DetailedForecast string `json:"-"`
		} `json:"periods"`
	} `json:"properties"`
}
