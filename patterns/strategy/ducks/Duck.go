package strategy

import (
	"github.com/amr-ahmed/Design-Patterns/patterns/strategy/ducks/behaviors"
)

// Duck embeds IFly, IEat, IQuack and IMove interface
// Client for all behaviors
type Duck struct {
	iFly   behaviors.IFly
	iEat   behaviors.IEat
	iQuack behaviors.IQuack
	iMove  behaviors.IMove
}

func GetNewDuck(iFly behaviors.IFly, iEat behaviors.IEat, iQuack behaviors.IQuack, iMove behaviors.IMove) *Duck {
	return &Duck{
		iFly:   iFly,
		iEat:   iEat,
		iQuack: iQuack,
		iMove:  iMove,
	}
}

func (duck Duck) Fly() string {
	return duck.iFly.Fly()
}

func (duck Duck) Eat() string {
	return duck.iEat.Eat()
}

func (duck Duck) Quack() string {
	return duck.iQuack.Quack()
}

func (duck Duck) Move() string {
	return duck.iMove.Move()
}
