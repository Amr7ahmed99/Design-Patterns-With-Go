package factory_method

type Musket struct {
	Gun
}

func GetNewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
