package twelvelabs

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
