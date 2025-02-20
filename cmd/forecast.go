package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
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

var forecastCommand = &cobra.Command{
	Use:   "forecast",
	Short: "Print the city forecast",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		q := "florianopolis"
		d := 2
		if len(os.Args) > 2 {
			q = os.Args[2]

			if len(os.Args) > 3 {
				var err error
				d, err = strconv.Atoi(os.Args[3])
				if err != nil {
					fmt.Println("Invalid number of days, using default value 1")
					d = 1
				}
			}
		}

		res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=c1c5f0d3c5b2468bb3010800251302&q=%s&days=%v&aqi=no&alert=no&lang=pt", q, d))
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
	},
}

func init() {
	rootCmd.AddCommand(forecastCommand)
}
