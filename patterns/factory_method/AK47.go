package factory_method

type AK47 struct {
	Gun
}

func GetNewAk47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
