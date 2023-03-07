package models

// SearchResponse ...
type SearchResponse struct {
	VideoID      string `json:"videoId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PublishedAt  string `json:"publishedAt"`
	ThumbnailURL string `json:"thumbnailURL"`
}
