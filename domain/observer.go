package domain

type Observee interface {
	Register(observer Observer)
	Notify()
}
type Observer interface {
	Update()
}
