package service

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/kelein/cookbook/govoyage/pbgen"
	"github.com/kelein/cookbook/govoyage/pkg/bapi"
)

// TagService provides a Tag service
type TagService struct{}

// NewTagService creates a new TagService instance
func NewTagService() *TagService {
	return &TagService{}
}

// GetTags return list of tags
func (t *TagService) GetTags(ctx context.Context, req *pbgen.GetTagsRequest) (*pbgen.GetTagsReply, error) {
	api := bapi.NewAPI("http://localhost:8000")
	body, err := api.GetTags(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	tags := &pbgen.GetTagsReply{}
	if err := json.Unmarshal(body, tags); err != nil {
		slog.Error("unmarshal tags failed", "error", err)
		return nil, err
	}
	return tags, nil
}
