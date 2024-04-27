package usecase

import (
	"context"
	"fmt"
	"mikke-server/database"
	"mikke-server/domain"
)

type PostUsecase struct {
	Repo PostAdder
	DB   database.Execer
}

func (p *PostUsecase) PostQuestion(ctx context.Context, user_id int, title string, comment string) (*domain.Post, error) {
	post := &domain.Post{
		UserID:  domain.UserID(user_id),
		Title:   title,
		Comment: comment,
	}
	err := p.Repo.PostQuestion(ctx, p.DB, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

type ListPostsUsecase struct {
	Repo PostLister
	DB   database.Queryer
}

func (p *ListPostsUsecase) ListPosts(ctx context.Context) (domain.Posts, error) {
	posts, err := p.Repo.ListPosts(ctx, p.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return posts, nil
}
