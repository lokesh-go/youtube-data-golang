package youtube

import (
	"google.golang.org/api/youtube/v3"

	ytPkgModels "github.com/lokesh-go/youtube-data-golang/src/pkg/youtube/models"
)

// Methods ...
type Methods interface {
	Search(searchText string) ([]ytPkgModels.SearchResponse, error)
}

// Search ...
func (s *service) Search(searchText string) (response []ytPkgModels.SearchResponse, err error) {
	// Initialises search results
	var pageToken string
	items := []*youtube.SearchResult{}

	// Travers all pages
	for {
		// Searches on youtube
		res, err := s.youtube.Search.List([]string{s.config.Youtube.Search.Part}).Q(searchText).Order(s.config.Youtube.Search.Order).Type(s.config.Youtube.Search.Type).MaxResults(s.config.Youtube.Search.MaxResults).PageToken(pageToken).Do()
		if err != nil {
			return nil, err
		}

		// Appends response
		items = append(items, res.Items...)

		// Checks for pagination
		if !s.config.Youtube.Search.Pagination.Enabled {
			break
		}

		// Assigns next page token
		pageToken = res.NextPageToken

		// If next page doesn't exists
		// Means all resutls are traversed
		if pageToken == "" {
			break
		}
	}

	// Forms response
	response = []ytPkgModels.SearchResponse{}
	for _, item := range items {
		if item == nil {
			continue
		}

		// Forms video id
		videoId := ""
		if item.Id != nil {
			videoId = item.Id.VideoId
		}

		// Assigns video snippet details
		title := ""
		description := ""
		publishedAt := ""
		thumbnailURL := ""
		if item.Snippet != nil {
			title = item.Snippet.Title
			description = item.Snippet.Description
			publishedAt = item.Snippet.PublishedAt
			if item.Snippet.Thumbnails != nil {
				if item.Snippet.Thumbnails.Default != nil {
					thumbnailURL = item.Snippet.Thumbnails.Default.Url
				}
			}
		}

		// Forms response
		res := ytPkgModels.SearchResponse{
			VideoID:      videoId,
			Title:        title,
			Description:  description,
			PublishedAt:  publishedAt,
			ThumbnailURL: thumbnailURL,
		}

		// Append response
		response = append(response, res)
	}

	// Returns
	return response, nil
}
