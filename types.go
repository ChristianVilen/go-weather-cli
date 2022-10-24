package main

type CurrentConditions struct {
	Datetime       string
	DatetimeEpoch  int
	Temp           float64
	Feelslike      float64
	Humidity       float64
	Dew            float64
	Precip         float64
	Precipprob     interface{}
	Snow           float64
	Snowdepth      float64
	Preciptype     interface{}
	Windgust       interface{}
	Windspeed      float64
	Winddir        float64
	Pressure       float64
	Visibility     float64
	Cloudcover     float64
	Solarradiation float64
	Solarenergy    float64
	Uvindex        float64
	Conditions     string
	Icon           string
	Stations       []string
	Sunrise        string
	SunriseEpoch   int
	Sunset         string
	SunsetEpoch    int
	Moonphase      float64
}

type Response struct {
	QueryCost         int
	Latitude          float64
	Longitude         float64
	ResolvedAddress   string
	Address           string
	Timezone          string
	Tzoffset          float64
	Description       string
	CurrentConditions CurrentConditions
}
