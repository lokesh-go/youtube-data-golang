package dal

// GetAllResponse ...
type GetAllResponse struct {
	VideoID      string `json:"videoId" bson:"_id"`
	Title        string `json:"title" bson:"title"`
	Description  string `json:"description" bson:"description"`
	PublishedAt  string `json:"publishedAt" bson:"publishedAt"`
	ThumbnailURL string `json:"thumbnailURL" bson:"thumbnailURL"`
}
