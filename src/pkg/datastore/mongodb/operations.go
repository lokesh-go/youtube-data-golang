package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Methods ...
type Methods interface {
	BulkWrite(ctx context.Context, collection string, models []mongo.WriteModel) error
	FindAll(ctx context.Context, collection string, skip, limit *int64, sort map[string]interface{}) ([]interface{}, error)
}

// BulkWrite ...
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

// Find ...
func (c *clients) FindAll(ctx context.Context, collection string, skip, limit *int64, sort map[string]interface{}) (res []interface{}, err error) {
	// Find options
	options := &options.FindOptions{
		Skip:  skip,
		Limit: limit,
		Sort:  sort,
	}

	// Form query
	filter := bson.D{}

	// Finds
	cursor, err := c.database.Collection(collection).Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	// Close connection at the last
	defer cursor.Close(ctx)

	// Binds cursor response
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	// Returns
	return res, nil
}
