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

type ListUploadsQuery struct {
	*CommonListQuery

	IndexId  string   `json:"index_id,omitempty"`
	Status   []string `json:"status,omitempty"`
	Filename string   `json:"filename,omitempty"`
	Duration float32  `json:"duration,omitempty"` // seconds
	Width    int32    `json:"width,omitempty"`    // pixels
	Height   int32    `json:"height,omitempty"`   // pixels
}

type UploadDetails struct {
	TaskId    string
	VideoId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
	IndexId   string
}

type ListVideosQuery struct {
	*CommonListQuery

	IndexId  string  `json:"index_id",omitempty`
	Filename string  `json:"filename,omitempty"`
	Duration float32 `json:"duration,omitempty"` // seconds
	Fps      float32 `json:"fps,omitempty"`      // frames per second
	Width    float32 `json:"width,omitempty"`    // pixels
	Height   int32   `json:"height,omitempty"`   // pixels
	Size     float32 `json:"size,omitempty"`     // bytes
}

type VideoDetails struct {
	VideoId   string
	CreatedAt time.Time
	UpdatedAt time.Time
	IndexId   string
	Filename  string
	Duration  float32
	Fps       float32
	Width     int32
	Height    int32
	Size      float32
}
