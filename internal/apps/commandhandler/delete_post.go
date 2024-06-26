package commandhandler

import (
	"context"
	"github.com/qmstar0/nightsky-api/internal/domain/aggregates"
	"github.com/qmstar0/nightsky-api/internal/domain/commands"
)

type DeletePostHandler struct {
	postRepo aggregates.PostRepository
}

func NewDeletePostHandler(repository aggregates.PostRepository) DeletePostHandler {
	if repository == nil {
		panic("missing PostRepository")
	}
	return DeletePostHandler{postRepo: repository}
}

func (d DeletePostHandler) Handle(ctx context.Context, cmd commands.DeletePost) error {
	return d.postRepo.Delete(ctx, cmd.ID)
}
