package observer

// Observer interface
type IObserver interface {
	Update() (bool, error)
}
