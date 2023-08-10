package main

import (
	"fmt"

	addon "github.com/amr-ahmed/Design-Patterns/patterns/decorator/add_on"
	"github.com/amr-ahmed/Design-Patterns/patterns/decorator/beverage"
)

func main() {
	nescafeCaramel := addon.AddMilk(addon.AddCaramel(beverage.GetNewNescafe()))
	fmt.Printf("Nescafe with Milk and Caramel: %d \n", nescafeCaramel.Cost())
	espressoCaramel := addon.AddCaramel(beverage.GetNewEspresso())
	fmt.Printf("Espresso wit Caramel: %d \n", espressoCaramel.Cost())

}
