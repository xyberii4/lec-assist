package tlservice

import (
	"context"
	"time"

	"github.com/xyberii4/lec-assist/backend/internal/twelvelabs"
)

type Service interface {
	CreateIndex(ctx context.Context, indexName string, models []string) (string, error)
	DeleteIndex(ctx context.Context, indexId string) error
	ListIndexes(ctx context.Context, req *ListIndexesQuery) ([]*IndexDetails, error)
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

func (s *service) validateCommonListQuery(q *CommonListQuery) error {
	// Validate Page
	if q.Page != nil {
		if *q.Page < 1 {
			return NewServiceError(ErrCodeInvalidArg, "page must be greater than 0", nil)
		}
	}

	// Validate PageLimit
	if q.PageLimit != nil {
		if *q.PageLimit < 1 || *q.PageLimit > 50 {
			return NewServiceError(ErrCodeInvalidArg, "page limit must be between 1 and 50", nil)
		}
	}

	// Validate SortBy
	if q.SortBy != nil {
		val := *q.SortBy
		if val != "created_at" && val != "updated_at" {
			return NewServiceError(ErrCodeInvalidArg, "sort_by must be either 'created_at' or 'updated_at'", nil)
		}
	}

	// Validate SortOption
	if q.SortOption != nil {
		val := *q.SortOption
		if val != "asc" && val != "desc" {
			return NewServiceError(ErrCodeInvalidArg, "sort_option must be either 'asc' or 'desc'", nil)
		}
	}

	// Validate CreatedAt
	if q.CreatedAt != nil {
		if q.CreatedAt.After(time.Now().UTC()) {
			return NewServiceError(ErrCodeInvalidArg, "created_at cannot be in the future", nil)
		}
	}

	// Validate UpdatedAt
	if q.UpdatedAt != nil {
		if q.UpdatedAt.After(time.Now().UTC()) {
			return NewServiceError(ErrCodeInvalidArg, "updated_at cannot be in the future", nil)
		}
	}

	return nil
}
