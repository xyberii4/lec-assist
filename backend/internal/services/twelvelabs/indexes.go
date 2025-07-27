package tlservice

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/xyberii4/lec-assist/backend/internal/twelvelabs"
	sdk "github.com/xyberii4/lec-assist/backend/pkg/twelvelabs"
	"go.uber.org/zap"
)

func (s *service) CreateIndex(ctx context.Context, indexName string, models []string) (string, error) {
	zap.L().Debug("Service: Attempting to create index...", zap.Any("user_id", ctx.Value("userId")))
	if indexName == "" {
		return "", NewServiceError(ErrCodeInvalidArg, "index name cannot be empty", nil)
	}

	if len(models) == 0 {
		return "", NewServiceError(ErrCodeInvalidArg, "models cannot be empty", nil)
	}

	// TODO: check if user already has an associated index

	// Check if valid model
	modelsArr := make([]sdk.CreateIndexRequestModelsInner, len(models))
	for i, m := range models {
		if m != "marengo2.7" && m != "pegasus1.2" {
			return "", NewServiceError(ErrCodeInvalidArg, fmt.Sprintf("invalid model: %s", m), nil)
		}

		modelsArr[i] = *sdk.NewCreateIndexRequestModelsInner(m, []string{"visual", "audio"})

	}

	// Call the TwelveLabs API
	req := sdk.NewCreateIndexRequest(indexName, modelsArr)
	indexId, err := s.client.CreateIndex(ctx, req)
	if err != nil {
		zap.L().Error("Service: Failed to create index",
			zap.Any("user_id", ctx.Value("userId")),
			zap.Error(err))

		// Check if the error is an API error
		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			if apiErr.StatusCode == 409 {
				zap.L().Info("Service: Index already exists",
					zap.String("index_name", indexName),
					zap.Any("user_id", ctx.Value("userId")))

				return "", NewServiceError(ErrCodeConflict, "index already exists", nil)
			} else {
				zap.L().Warn("Service: API error while creating index",
					zap.Any("user_id", ctx.Value("userId")),
					zap.Error(err))

				return "", NewServiceError(ErrCodeExternalAPIError, fmt.Sprintf("failed to create index: %s", apiErr.Message), err)
			}
		}
		zap.L().Warn("Service: Unexpected error while creating index",
			zap.Any("user_id", ctx.Value("userId")),
			zap.Error(err))

		return "", NewServiceError(ErrCodeExternalAPIError, "unexpected error while creating index", err)
	}

	return indexId, nil
}

func (s *service) DeleteIndex(ctx context.Context, indexId string) error {
	zap.L().Debug("Service: Attempting to delete index...", zap.Any("user_id", ctx.Value("userId")), zap.String("index_id", indexId))
	if indexId == "" {
		return NewServiceError(ErrCodeInvalidArg, "index ID cannot be empty", nil)
	}

	// TODO: Verify if the index belongs to the user

	req := twelvelabs.DeleteIndexRequest{IndexId: indexId}
	if err := s.client.DeleteIndex(ctx, &req); err != nil {
		zap.L().Error("Service: Failed to delete index",
			zap.Any("user_id", ctx.Value("userId")),
			zap.String("index_id", indexId),
			zap.Error(err))

		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			zap.L().Warn("Service: API error while deleting index",
				zap.Any("user_id", ctx.Value("userId")),
				zap.String("index_id", indexId))

			return NewServiceError(ErrCodeExternalAPIError, apiErr.Message, nil)

		}
		zap.L().Warn("Service: Unexpected error while deleting index",
			zap.Any("user_id", ctx.Value("userId")),
			zap.String("index_id", indexId),
			zap.Error(err))

		return NewServiceError(ErrCodeExternalAPIError, "unexpected error while deleting index", err)
	}

	return nil
}

func (s *service) ListIndexes(ctx context.Context, req *ListIndexesQuery) ([]*IndexDetails, error) {
	zap.L().Debug("Service: Attempting to list indexes...", zap.Any("user_id", ctx.Value("userId")))
	// req must not be nil
	if req == nil || req.CommonListQuery == nil {
		return nil, NewServiceError(ErrCodeInvalidArg, "request cannot be nil", nil)
	}

	if err := s.validateCommonListQuery(req.CommonListQuery); err != nil {
		return nil, err
	}

	queryOpts := []interface{}{} // Query function options

	// Add default query options
	// If query parameters are not provided, default values will be used
	if req.Page != nil {
		queryOpts = append(queryOpts, twelvelabs.WithPage(*req.Page))
	}

	if req.PageLimit != nil {
		queryOpts = append(queryOpts, twelvelabs.WithPageLimit(*req.PageLimit))
	}

	if req.SortBy != nil {
		queryOpts = append(queryOpts, twelvelabs.WithSortBy(*req.SortBy))
	}

	if req.SortOption != nil {
		queryOpts = append(queryOpts, twelvelabs.WithSortOption(*req.SortOption))
	}

	if req.CreatedAt != nil {
		queryOpts = append(queryOpts, twelvelabs.WithCreatedAt(req.CreatedAt.Format(time.RFC3339)))
	}

	if req.UpdatedAt != nil {
		queryOpts = append(queryOpts, twelvelabs.WithUpdatedAt(req.UpdatedAt.Format(time.RFC3339)))
	}

	// optional query parameters
	// TODO: validate index name
	if req.IndexName != "" {
		queryOpts = append(queryOpts, twelvelabs.WithIndexName(req.IndexName))
	}

	// Model options can only be audio or visual, and should be serparated by commas
	if req.ModelOptions != "" {
		options := strings.Split(req.ModelOptions, ",")
		for _, o := range options {
			if o != "audio" && o != "visual" {
				return nil, NewServiceError(ErrCodeInvalidArg, fmt.Sprintf("invalid model option: %s", o), nil)
			}
		}
		queryOpts = append(queryOpts, twelvelabs.WithModelOptions(req.ModelOptions))
	}

	if req.ModelFamily != "" {
		if req.ModelFamily != "marengo" && req.ModelFamily != "pegasus" {
			return nil, NewServiceError(ErrCodeInvalidArg, "model family must be either 'marengo' or 'pegasus'", nil)
		}
		queryOpts = append(queryOpts, twelvelabs.WithModelFamily(req.ModelFamily))
	}

	// Create the query with the provided options
	query := twelvelabs.NewListIndexesQuery(queryOpts...)

	rawIndexDetails, err := s.client.ListIndexes(ctx, query)
	if err != nil {
		zap.L().Error("Service: Failed to list indexes")

		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			zap.L().Warn("Service: API error while listing indexes")

			return nil, NewServiceError(ErrCodeExternalAPIError, apiErr.Message, nil)
		}
		zap.L().Warn("Service: Unexpected error while listing indexes", zap.Error(err))

		return nil, NewServiceError(ErrCodeExternalAPIError, "unexpected error while listing indexes", err)
	}

	// Parse the index details
	indexDetails := make([]*IndexDetails, len(rawIndexDetails))
	for i, idx := range rawIndexDetails {
		indexDetails[i] = s.parseIndexDetails(idx)
	}

	return indexDetails, nil
}

func (s *service) RetrieveIndex(ctx context.Context, indexId string) (*IndexDetails, error) {
	zap.L().Debug("Service: Attempting to retrieve index details...", zap.Any("user_id", ctx.Value("userId")), zap.String("index_id", indexId))
	// TODO: check index exists and belongs to user
	if indexId == "" {
		return nil, NewServiceError(ErrCodeInvalidArg, "index id cannot be empty", nil)
	}

	rawIndexDetails, err := s.client.RetrieveIndex(ctx, &twelvelabs.RetrieveIndexRequest{IndexId: indexId})
	if err != nil {
		zap.L().Error("Service: Failed to retrieve index",
			zap.Any("user_id", ctx.Value("userId")),
			zap.String("index_id", indexId))

		var apiErr *twelvelabs.APIError
		if errors.As(err, &apiErr) {
			zap.L().Warn("Service: API error while retrieving index",
				zap.Any("user_id", ctx.Value("userId")),
				zap.String("index_id", indexId))

			return nil, NewServiceError(ErrCodeExternalAPIError, apiErr.Message, nil)
		}
		zap.L().Warn("Service: Unexpected error while retrieving index",
			zap.Any("user_id", ctx.Value("userId")),
			zap.String("index_id", indexId),
			zap.Error(err))

		return nil, NewServiceError(ErrCodeExternalAPIError, "unexpected error while retrieving index", err)
	}

	return s.parseIndexDetails(rawIndexDetails), nil
}

// converts the index details from the TwelveLabs API into a structured IndexDetails object
func (s *service) parseIndexDetails(idxDtls *twelvelabs.IndexDetails) *IndexDetails {
	createdAt, _ := time.Parse(time.RFC3339, idxDtls.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, idxDtls.UpdatedAt)
	expiresAt, _ := time.Parse(time.RFC3339, idxDtls.ExpiresAt)
	return &IndexDetails{
		IndexId:    idxDtls.IndexId,
		IndexName:  idxDtls.IndexName,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		ExpiresAt:  expiresAt,
		VideoCount: idxDtls.VideoCount,
	}
}
