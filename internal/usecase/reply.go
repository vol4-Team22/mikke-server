package usecase

import (
	"context"
	"fmt"
	"mikke-server/internal/database"
	domain2 "mikke-server/internal/domain"
)

type SendReplyUsecase struct {
	Repo domain2.ReplyAdder
	DB   database.Execer
}

func (p *SendReplyUsecase) SendReply(ctx context.Context, reply *domain2.Reply) (*domain2.Reply, error) {
	err := p.Repo.SendReply(ctx, p.DB, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (p *ListRepliesUsecase) ListReplies(ctx context.Context, postID domain2.PostID) (domain2.Replies, error) {
	replies, err := p.Repo.ListReplies(ctx, p.DB, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return replies, nil
}

type ListRepliesUsecase struct {
	Repo domain2.ReplyLister
	DB   database.Queryer
}