package main

import (
	"fmt"

	strategy "github.com/amr-ahmed/Design-Patterns/patterns/strategy/ducks"
)

func main() {

	cityDuck := strategy.GetNewDuck(strategy.GetNewCityDuck(), strategy.GetNewCityDuck(), strategy.GetNewCityDuck(), strategy.GetNewCityDuck())
	fmt.Println("City Duck...")
	fmt.Println(cityDuck.Fly())
	fmt.Println(cityDuck.Eat())
	fmt.Println(cityDuck.Move())
	fmt.Println(cityDuck.Quack())
	fmt.Println()

	rubberDuck := strategy.GetNewDuck(strategy.GetNewRubberDuck(), strategy.GetNewRubberDuck(), strategy.GetNewRubberDuck(), strategy.GetNewRubberDuck())
	fmt.Println("Rubber Duck...")
	fmt.Println(rubberDuck.Fly())
	fmt.Println(rubberDuck.Eat())
	fmt.Println(rubberDuck.Move())
	fmt.Println(rubberDuck.Quack())
	fmt.Println()

	wildDuck := strategy.GetNewDuck(strategy.GetNewWildDuck(), strategy.GetNewWildDuck(), strategy.GetNewWildDuck(), strategy.GetNewWildDuck())
	fmt.Println("Wild Duck...")
	fmt.Println(wildDuck.Fly())
	fmt.Println(wildDuck.Eat())
	fmt.Println(wildDuck.Move())
	fmt.Println(wildDuck.Quack())
}
