package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func printWeather(w http.ResponseWriter, current *Weather) {
	fmt.Fprintf(w,
		"\nNow - %s, %s, %.0f⁰C, %s\n\n",
		current.Location.Name,
		current.Location.Country,
		current.Current.TempC,
		current.Current.Condition.Text,
	)

	for _, hour := range current.Forecast.Forecastday[0].Hour {
		rain := hour.ChanceOfRain
		snow := hour.ChanceOfSnow
		temp := hour.TempC
		description := hour.Condition.Text

		// Convert time from Unix epoch to readable format and skip hours in the past
		date := time.Unix(hour.TimeEpoch, 0)
		dateformat := date.Format(time.Kitchen)
		if date.Before(time.Now()) {
			continue
		}

		format := ""

		switch {
		case snow > 0:
			format = "%07s - %3.0f⁰C,  | snow  %3.0f%%, | %s\n"
		case rain >= 0:
			format = "%07s - %3.0f⁰C,  | rain  %3.0f%%, | %s\n"

		}
		if strings.Contains(format, "snow %.0f%%") {
			fmt.Fprintf(w, format, dateformat, temp, snow, description)
		} else {
			fmt.Fprintf(w, format, dateformat, temp, rain, description)
		}

	}
}

//html template funcs

func unixTimeFormat(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("15:04")
}

func unixTimeIsPast(unixTime int64) bool {
	return time.Unix(unixTime, 0).Before(time.Now())
}

func intTemp(temp float64) int64 {
    return int64(temp)
}