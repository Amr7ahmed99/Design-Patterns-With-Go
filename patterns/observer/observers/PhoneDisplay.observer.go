package observers

import (
	"errors"
	"fmt"

	"github.com/amr-ahmed/Design-Patterns/patterns/observer/subjects"
)

type PhoneDisplay struct {
	Station     *subjects.WeatherStation
	Temperature string
}

func GetNewPhoneDisplay(station *subjects.WeatherStation) *PhoneDisplay {
	return &PhoneDisplay{
		Station:     station,
		Temperature: station.GetTemperature(),
	}
}

func (pd *PhoneDisplay) Update() (bool, error) {
	if pd.Temperature == pd.Station.GetTemperature() {
		return false, errors.New("the temperature has not been changed")
	}
	pd.Temperature = pd.Station.GetTemperature()
	message := fmt.Sprintf("phone has been notified: Current Temperature Is %s", pd.Temperature)
	fmt.Println(message)
	return true, nil
}
