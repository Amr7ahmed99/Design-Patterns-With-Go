package strategy

type WildDuck struct {
}

func GetNewWildDuck() *WildDuck {
	return &WildDuck{}
}

func (duck WildDuck) Fly() string {
	return "Wild Duck Is Flying :)"
}

func (duck WildDuck) Eat() string {
	return "Wild Duck Is Eating :)"
}

func (duck WildDuck) Quack() string {
	return "Wild Duck Is Quacking :)"

}

func (duck WildDuck) Move() string {
	return "Wild Duck Is Moving :)"
}
