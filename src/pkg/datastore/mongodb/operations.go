package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Methods ...
type Methods interface {
	BulkWrite(ctx context.Context, collection string, models []mongo.WriteModel) error
	FindAll(ctx context.Context, collection string, skip, limit *int64, sort map[string]interface{}) ([]interface{}, error)
	CreateIndex(ctx context.Context, collection string, indexKeys []string, unique bool) (res string, err error)
	ListCollectionNames(ctx context.Context) (res []string, err error)
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

// CreateIndex ...
func (c *clients) CreateIndex(ctx context.Context, collection string, indexKeys []string, unique bool) (res string, err error) {
	// Forms keys
	var mapKeys bson.D
	for _, key := range indexKeys {
		k := primitive.E{Key: key, Value: 1}
		mapKeys = append(mapKeys, k)
	}

	// Forms index model
	indexModel := mongo.IndexModel{
		Keys:    mapKeys,
		Options: options.Index().SetUnique(unique),
	}

	// CreateIndex
	res, err = c.database.Collection(collection).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return res, err
	}

	// Returns
	return res, nil
}

// ListCollectionNames ...
func (c *clients) ListCollectionNames(ctx context.Context) (res []string, err error) {
	// Form query
	filter := bson.D{}

	// Gets
	res, err = c.database.ListCollectionNames(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Returns
	return res, nil
}
