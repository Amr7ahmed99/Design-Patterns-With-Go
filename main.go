package main

import (
	"fmt"

	"github.com/amr-ahmed/Design-Patterns/patterns/factory_method"
	"github.com/amr-ahmed/Design-Patterns/patterns/factory_method/factory"
)

func main() {
	ak47, _ := factory.GetGun("ak47")
	printGunDetails(ak47)

	musket, _ := factory.GetGun("musket")
	printGunDetails(musket)

}

func printGunDetails(gun factory_method.IGun) {
	fmt.Printf("Gun: %s \t", gun.GetName())
	fmt.Printf("Power: %d \n", gun.GetPower())
}
