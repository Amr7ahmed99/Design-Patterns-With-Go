package beverage

type Nescafe struct {
}

func GetNewNescafe() *Nescafe {
	return &Nescafe{}
}

func (nescafe *Nescafe) Cost() int {
	return 25
}
