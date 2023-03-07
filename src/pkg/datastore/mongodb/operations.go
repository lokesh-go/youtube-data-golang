package mongodb

import "context"

// Methods ...
type Methods interface {
	Find(ctx context.Context) error
}

// Find ...
func (c *clients) Find(ctx context.Context) (err error) {
	return nil
}
