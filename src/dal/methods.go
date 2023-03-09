package dal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	ytModels "github.com/lokesh-go/youtube-data-golang/src/pkg/youtube/models"
	utils "github.com/lokesh-go/youtube-data-golang/src/utils"
)

const (
	IdIdentifier           = "_id"
	SetIdentifier          = "$set"
	TitleIdentifier        = "title"
	DescriptionIdentifier  = "description"
	PublishedAtIdentifier  = "publishedAt"
	ThumbnailURLIdentifier = "thumbnailURL"
)

// Methods ...
type Methods interface {
	PushData(response []ytModels.SearchResponse) error
	GetAllData(currentPage int64) ([]GetAllResponse, error)
}

// PushData ...
func (d *dal) PushData(response []ytModels.SearchResponse) (err error) {
	// Checks
	if len(response) == 0 {
		return nil
	}

	// Get all collections
	collections, err := d.dbServices.ListCollectionNames(context.Background())
	if err != nil {
		return err
	}

	// Check
	if !utils.Contains(collections, d.config.Datastores.Youtube.Collections.Youtube) {
		// Create index if collection is not there
		// Note: If Collection is already there thne it will not create the index.
		d.dbServices.CreateIndex(context.Background(), d.config.Datastores.Youtube.Collections.Youtube, []string{TitleIdentifier}, true)
		d.dbServices.CreateIndex(context.Background(), d.config.Datastores.Youtube.Collections.Youtube, []string{PublishedAtIdentifier}, true)
	}

	// Forms query
	models := []mongo.WriteModel{}
	for _, data := range response {
		// Forms filter for each record
		filter := bson.D{
			primitive.E{Key: IdIdentifier, Value: data.VideoID},
		}

		// Forms update value
		update := bson.D{
			primitive.E{Key: SetIdentifier, Value: bson.D{
				primitive.E{Key: TitleIdentifier, Value: data.Title},
				primitive.E{Key: DescriptionIdentifier, Value: data.Description},
				primitive.E{Key: PublishedAtIdentifier, Value: data.PublishedAt},
				primitive.E{Key: ThumbnailURLIdentifier, Value: data.ThumbnailURL},
			}},
		}
		models = append(models, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	// Bulk operation
	err = d.dbServices.BulkWrite(context.Background(), d.config.Datastores.Youtube.Collections.Youtube, models)
	if err != nil {
		return err
	}

	// Returns
	return nil
}

// GetData ...
func (d *dal) GetAllData(currentPage int64) (response []GetAllResponse, err error) {
	// Gets config
	limit := int64(d.config.Datastores.Youtube.Pagination.ResponsePerPage)
	skip := (currentPage - 1) * limit
	sort := map[string]interface{}{
		PublishedAtIdentifier: -1,
	}

	// Get all data
	data, err := d.dbServices.FindAll(context.Background(), d.config.Datastores.Youtube.Collections.Youtube, &skip, &limit, sort)
	if err != nil {
		return nil, err
	}

	// Ranges
	response = []GetAllResponse{}
	for _, d := range data {
		res := GetAllResponse{}
		bytes, _ := utils.BSONMarshal(d)
		utils.BSONUnmarshal(bytes, &res)
		response = append(response, res)
	}

	// Returns
	return response, nil
}
