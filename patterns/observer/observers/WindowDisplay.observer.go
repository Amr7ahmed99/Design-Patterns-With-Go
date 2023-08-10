package observers

import (
	"errors"
	"fmt"

	"github.com/amr-ahmed/Design-Patterns/patterns/observer/subjects"
)

type WindowDisplay struct {
	Station     *subjects.WeatherStation
	Temperature string
}

func GetNewWindowDisplay(station *subjects.WeatherStation) *WindowDisplay {
	return &WindowDisplay{
		Station:     station,
		Temperature: station.GetTemperature(),
	}
}

func (wd *WindowDisplay) Update() (bool, error) {
	if wd.Temperature == wd.Station.GetTemperature() {
		return false, errors.New("the temperature has not been changed")
	}
	wd.Temperature = wd.Station.GetTemperature()
	message := fmt.Sprintf("window has been notified: Current Temperature Is %s", wd.Temperature)
	fmt.Println(message)
	return true, nil
}
