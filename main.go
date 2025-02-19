package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type (
	Wheather struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`
		Current struct {
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		} `json:"current"`
		Forecast struct {
			Forecastday []struct {
				Date string
				Hour []struct {
					TimeEpopch int64   `json:"time_epoch"`
					TempC      float64 `json:"temp_c"`
					Condition  struct {
						Text string `json:"text"`
					} `json:"condition"`
					ChanceOfRain float64 `json:"chance_of_rain"`
				} `json:"hour"`
			} `json:"forecastday"`
		} `json:"forecast"`
	}
)

func main() {

	q := "florianopolis"
	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=c1c5f0d3c5b2468bb3010800251302&q=" + q + "&days=1&aqi=no&alert=no&lang=pt")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Wheater API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var wheather Wheather
	err = json.Unmarshal(body, &wheather)
	if err != nil {
		panic(err)
	}

	location, current :=
		wheather.Location,
		wheather.Current

	fmt.Printf(
		"%s, %s: %0.fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, forecastDay := range wheather.Forecast.Forecastday {
		fmt.Printf(
			"---- %s -----\n", forecastDay.Date,
		)

		for _, hour := range forecastDay.Hour {
			date := time.Unix(hour.TimeEpopch, 0)

			message := fmt.Sprintf(
				"%s - %.0f°C, %.0f%%, %s\n",
				date.Format("15:04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text,
			)

			if hour.ChanceOfRain < 40 {
				fmt.Print(message)
			} else {
				color.Red(message)
			}
		}
	}

}
