package strategy

type RubberDuck struct {
}

func GetNewRubberDuck() *RubberDuck {
	return &RubberDuck{}
}

func (duck RubberDuck) Fly() string {
	return "Rubber Duck Can Not Fly :("
}

func (duck RubberDuck) Eat() string {
	return "Rubber Duck Can Not Eat :("
}

func (duck RubberDuck) Quack() string {
	return "Rubber Duck Can Not Quack :("

}

func (duck RubberDuck) Move() string {
	return "Rubber Duck Can Not Move :("
}
