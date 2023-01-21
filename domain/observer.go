package domain

import "context"

type Observee interface {
	Register(observer Observer)
	Notify(ctx context.Context)
}
type Observer interface {
	Update(ctx context.Context)
}
