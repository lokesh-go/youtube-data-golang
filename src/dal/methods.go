package dal

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	ytModels "github.com/lokesh-go/youtube-data-golang/src/pkg/youtube/models"
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
}

// PushData ...
func (d *dal) PushData(response []ytModels.SearchResponse) (err error) {
	// Checks
	if len(response) == 0 {
		return nil
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
