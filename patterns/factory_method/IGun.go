package factory_method

type IGun interface {
	SetName(string)
	GetName() string
	SetPower(int)
	GetPower() int
}
