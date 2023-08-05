package strategy

type CityDuck struct {
}

func GetNewCityDuck() *CityDuck {
	return &CityDuck{}
}

func (duck CityDuck) Fly() string {
	return "City Duck Is Flying :)"
}

func (duck CityDuck) Eat() string {
	return "City Duck Is Eating :)"
}

func (duck CityDuck) Quack() string {
	return "City Duck Is Quacking :)"

}

func (duck CityDuck) Move() string {
	return "City Duck Is Moving :)"
}
