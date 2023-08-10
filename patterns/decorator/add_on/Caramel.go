package addon

import "github.com/amr-ahmed/Design-Patterns/patterns/decorator/beverage"

type Caramel struct {
	beverage beverage.Beverage
}

func AddCaramel(b beverage.Beverage) *Caramel {
	return &Caramel{
		beverage: b,
	}
}

func (caramel Caramel) Cost() int {
	return caramel.beverage.Cost() + 12
}
