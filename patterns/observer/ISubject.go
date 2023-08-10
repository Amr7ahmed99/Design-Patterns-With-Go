package observer

// Subject interface
type ISubject interface {
	Attach(o IObserver) (bool, error)
	Detach(o IObserver) (bool, error)
	Notify() (bool, error)
}
