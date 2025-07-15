package twelvelabs

type GistResult struct {
	VideoId  string   `json:"video_id"`
	Title    string   `json:"title,omitempty"`
	Topics   []string `json:"topics,omitempty"`
	Hashtags []string `json:"hashtags,omitempty"`
}

type SummariseResult struct {
	VideoId    string              `json:"video_id"`
	Summary    string              `json:"summary,omitempty"`
	Chapter    []*ChapterDetails   `json:"chapter,omitempty"`
	Highlights []*HighlightDetails `json:"highlights,omitempty"`
}

// chronological list of all the chapters in a video, providing a granular breakdown of its content
type ChapterDetails struct {
	ChapterNumber int32  `json:"chapter_number"`
	Start         int32  `json:"start"`                 // seconds from start of video
	End           int32  `json:"end"`                   // seconds from start of video
	Title         string `json:"title"`                 // chapter title
	Description   string `json:"description,omitempty"` // summary of chapter
}

// chronologically ordered list of the most important events within a video.
type HighlightDetails struct {
	Start       int32  `json:"start"`                 // seconds from start of video
	End         int32  `json:"end"`                   // seconds from start of video
	Title       string `json:"title"`                 // highlight title
	Description string `json:"description,omitempty"` // summary of highlight
}
