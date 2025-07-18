package twelvelabs

import "fmt"

// -- API Error Handling --
type APIError struct {
	StatusCode int
	Code       string
	Operation  string
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("TwelveLabs API error for %s (%d): %s - %s", e.Operation, e.StatusCode, e.Code, e.Message)
}

// -- Rate Limits --
type RateLimit struct {
	Limit     int   `json:"limit"`     // Total number of requests allowed for endpoint
	Remaining int   `json:"remaining"` // Number of requests remaining for endpoint
	Used      int   `json:"used"`      // Number of requests made for enpoint
	Reset     int64 `json:"reset"`     // Time which used rates resets, in UTC epoch seconds
}

// -- Common Query Parameters for List Operations --

type commonListQueryParameters struct {
	Page       int32  `json:"page"`
	PageLimit  int32  `json:"page_limit"`           // Max: 50
	SortBy     string `json:"sort_by"`              // created_at/updated_at
	SortOption string `json:"sort_option"`          // asc/desc
	CreatedAt  string `json:"created_at,omitempty"` // optional, RFC3339 format
	UpdatedAt  string `json:"updated_at,omitempty"` // optional, RFC3339 format
}

type CommonListQueryOption func(*commonListQueryParameters)

func defaultCommonQueryParameters() commonListQueryParameters {
	return commonListQueryParameters{
		Page:       1,
		PageLimit:  10,
		SortBy:     "created_at",
		SortOption: "desc",
	}
}

func WithPage(page int32) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.Page = page
	}
}

func WithPageLimit(pageLimit int32) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.PageLimit = pageLimit
	}
}

func WithSortBy(sortBy string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.SortBy = sortBy
	}
}

func WithSortOption(sortOption string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.SortOption = sortOption
	}
}

func WithCreatedAt(createdAt string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt string) func(*commonListQueryParameters) {
	return func(q *commonListQueryParameters) {
		q.UpdatedAt = updatedAt
	}
}
