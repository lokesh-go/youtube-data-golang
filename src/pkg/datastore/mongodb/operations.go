package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Methods ...
type Methods interface {
	BulkWrite(ctx context.Context, collection string, models []mongo.WriteModel) error
}

// Find ...
func (c *clients) BulkWrite(ctx context.Context, collection string, models []mongo.WriteModel) (err error) {
	// Set options
	opts := options.BulkWrite().SetOrdered(false)

	// Bulk write
	_, err = c.database.Collection(collection).BulkWrite(ctx, models, opts)
	if err != nil {
		return err
	}

	return nil
}
