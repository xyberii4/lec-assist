package tlservice

import (
	"context"
	"time"

	"github.com/xyberii4/lec-assist/backend/internal/twelvelabs"
)

type Service interface {
	CreateIndex(ctx context.Context, indexName string, models []string) (string, error)
	DeleteIndex(ctx context.Context, indexId string) error
	ListIndexes(ctx context.Context, q *ListIndexesQuery) ([]*IndexDetails, error)
	RetrieveIndex(ctx context.Context, indexId string) (*IndexDetails, error)
}

type service struct {
	client twelvelabs.Client
}

func NewService(client twelvelabs.Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) parseCommonListQuery(q *CommonListQuery) ([]interface{}, error) {
	queryOpts := []interface{}{}
	// Validate Page
	if q.Page != nil {
		if *q.Page < 1 {
			return nil, NewServiceError(ErrCodeInvalidArg, "page must be greater than 0", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithPage(*q.Page))
	}

	// Validate PageLimit
	if q.PageLimit != nil {
		if *q.PageLimit < 1 || *q.PageLimit > 50 {
			return nil, NewServiceError(ErrCodeInvalidArg, "page limit must be between 1 and 50", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithPageLimit(*q.PageLimit))
	}

	// Validate SortBy
	if q.SortBy != nil {
		val := *q.SortBy
		if val != "created_at" && val != "updated_at" {
			return nil, NewServiceError(ErrCodeInvalidArg, "sort_by must be either 'created_at' or 'updated_at'", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithSortBy(val))
	}

	// Validate SortOption
	if q.SortOption != nil {
		val := *q.SortOption
		if val != "asc" && val != "desc" {
			return nil, NewServiceError(ErrCodeInvalidArg, "sort_option must be either 'asc' or 'desc'", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithSortOption(val))
	}

	// Validate CreatedAt
	if q.CreatedAt != nil {
		if q.CreatedAt.After(time.Now().UTC()) {
			return nil, NewServiceError(ErrCodeInvalidArg, "created_at cannot be in the future", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithCreatedAt(q.CreatedAt.Format(time.RFC3339)))
	}

	// Validate UpdatedAt
	if q.UpdatedAt != nil {
		if q.UpdatedAt.After(time.Now().UTC()) {
			return nil, NewServiceError(ErrCodeInvalidArg, "updated_at cannot be in the future", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithUpdatedAt(q.UpdatedAt.Format(time.RFC3339)))
	}

	return queryOpts, nil
}
