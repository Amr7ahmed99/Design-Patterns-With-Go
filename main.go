package main

import (
	"fmt"

	"github.com/amr-ahmed/Design-Patterns/patterns/observer/observers"
	"github.com/amr-ahmed/Design-Patterns/patterns/observer/subjects"
)

func main() {

	station := subjects.GenerateNewWeatherStation()
	phone := observers.GetNewPhoneDisplay(station)
	window := observers.GetNewWindowDisplay(station)

	if _, err := station.Attach(phone); err != nil {
		fmt.Println(err)
	}
	if _, err := station.Attach(window); err != nil {
		fmt.Println(err)
	}
	if err := station.SetTemperature("20"); err != nil {
		fmt.Println(err)
	}
	if _, err := station.Attach(phone); err != nil {
		fmt.Println(err)
	}
	if _, err := station.Detach(window); err != nil {
		fmt.Println(err)
	}

	if err := station.SetTemperature("20"); err != nil {
		fmt.Println(err)
	}

}
