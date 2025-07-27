package tlservice

import "time"

// pointers to check if field has been set
type CommonListQuery struct {
	Page       *int32
	PageLimit  *int32
	SortBy     *string
	SortOption *string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type ListIndexesQuery struct {
	*CommonListQuery

	IndexName    string
	ModelOptions string // audio/visual, multiple options must be comma-separated
	ModelFamily  string // marengo/pegasus
}

type IndexDetails struct {
	IndexId    string    `json:"index_id"`
	IndexName  string    `json:"index_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ExpiresAt  time.Time `json:"expires_at"`
	VideoCount int32     `json:"video_count"`
}
