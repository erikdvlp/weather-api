package nws

type Points struct {
	Context  []any  `json:"-"`
	ID       string `json:"-"`
	Type     string `json:"-"`
	Geometry struct {
		Type        string    `json:"-"`
		Coordinates []float64 `json:"-"`
	} `json:"-"`
	Properties struct {
		ID                  string `json:"-"`
		Type                string `json:"-"`
		Cwa                 string `json:"-"`
		ForecastOffice      string `json:"-"`
		GridID              string `json:"gridId"`
		GridX               int    `json:"gridX"`
		GridY               int    `json:"gridY"`
		Forecast            string `json:"-"`
		ForecastHourly      string `json:"-"`
		ForecastGridData    string `json:"-"`
		ObservationStations string `json:"-"`
		RelativeLocation    struct {
			Type     string `json:"-"`
			Geometry struct {
				Type        string    `json:"-"`
				Coordinates []float64 `json:"-"`
			} `json:"-"`
			Properties struct {
				City     string `json:"-"`
				State    string `json:"-"`
				Distance struct {
					UnitCode string  `json:"-"`
					Value    float64 `json:"-"`
				} `json:"-"`
				Bearing struct {
					UnitCode string `json:"-"`
					Value    int    `json:"-"`
				} `json:"-"`
			} `json:"-"`
		} `json:"-"`
		ForecastZone    string `json:"-"`
		County          string `json:"-"`
		FireWeatherZone string `json:"-"`
		TimeZone        string `json:"-"`
		RadarStation    string `json:"-"`
	} `json:"properties"`
}
