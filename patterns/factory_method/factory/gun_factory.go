package factory

import (
	"fmt"

	factory_method "github.com/amr-ahmed/Design-Patterns/patterns/factory_method"
)

func GetGun(gunType string) (factory_method.IGun, error) {

	if gunType == "ak47" {
		return factory_method.GetNewAk47(), nil
	}

	if gunType == "musket" {
		return factory_method.GetNewMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
