package subjects

import (
	"errors"

	"github.com/amr-ahmed/Design-Patterns/patterns/observer"
)

type WeatherStation struct {
	listeners   []observer.IObserver
	temperature string
}

func GenerateNewWeatherStation() *WeatherStation {
	return &WeatherStation{}
}

func (ws *WeatherStation) Attach(o observer.IObserver) (bool, error) {
	for _, listener := range ws.listeners {
		if listener == o {
			return false, errors.New("observer already exists")
		}
	}
	ws.listeners = append(ws.listeners, o)
	return true, nil
}

func (ws *WeatherStation) Detach(o observer.IObserver) (bool, error) {
	for i, listener := range ws.listeners {
		if listener == o {
			ws.listeners = append(ws.listeners[:i], ws.listeners[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("observer not found")
}

func (ws *WeatherStation) Notify() (bool, error) {
	for _, listener := range ws.listeners {
		if _, err := listener.Update(); err != nil {
			return false, errors.New("error has been occurred while notifying listeners")
		}
	}
	return true, nil
}

func (ws WeatherStation) GetTemperature() string {
	return ws.temperature
}

func (ws *WeatherStation) SetTemperature(newTemperature string) error {
	ws.temperature = newTemperature
	_, err := ws.Notify()
	return err
}
