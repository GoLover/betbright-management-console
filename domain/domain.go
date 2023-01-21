package domain

import "context"

type Operator interface {
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	Deactivate(ctx context.Context)
	Activate(ctx context.Context)
	Search(ctx context.Context)
	SearchAll(ctx context.Context)
}
