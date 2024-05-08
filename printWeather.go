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

		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		format := "%s - %.0f⁰C, rain %.0f%%, %s\n"

		switch {
		case snow > 0 && temp < 0 && temp >= -9 || snow > 0 && temp >= 0 && temp <= 9:
			format = "%s -  %.0f⁰C,   1| rain %.0f%%,   snow %.0f%%, | %s\n"
		case rain >= 0 && temp < 0 && temp >= -9 || rain > 0 && rain < 10 && temp >= 0 && temp <= 9:
			format = "%s -  %.0f⁰C,  2| rain %.0f%%,  %s\n"
		case rain >= 0 && rain < 10 && temp >= 0 && temp <= 9:
			format = "%s -  %.0f⁰C,  3| rain %.0f%%,   | %s\n"
		case rain >= 0 && rain < 10:
			format = "%s -  %.0f⁰C,  4| rain %.0f%%,   | %s\n"
		case rain > 9 && rain < 99:
			format = "%s -  %.0f⁰C,  5| rain %.0f%%,  | %s\n"
		case rain >= 100 && temp >= 10:
			format = "%s -  %.0f⁰C,  7| rain %.0f%%, | %s\n"
		case rain >= 100 && temp >= 0 && temp <= 9:
			format = "%s -  %.0f⁰C,   6| rain %.0f%%, | %s\n"
		default:
			format = "%s -  %.0f⁰C,   d| rain %.0f%%, | %s\n"
		}

		if strings.Contains(format, "snow %.0f%%") {
			fmt.Fprintf(w, format,
				date.Format("15:04"),
				temp,
				rain,
				snow,
				description,
			)
		} else {
			fmt.Fprintf(w, format,
				date.Format("15:04"),
				temp,
				rain,
				description,
			)
		}
	}
}
