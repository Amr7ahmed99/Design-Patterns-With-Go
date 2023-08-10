package addon

import "github.com/amr-ahmed/Design-Patterns/patterns/decorator/beverage"

type Milk struct {
	beverage beverage.Beverage
}

func AddMilk(b beverage.Beverage) *Milk {
	return &Milk{
		beverage: b,
	}
}

func (milk Milk) Cost() int {
	return milk.beverage.Cost() + 8
}
