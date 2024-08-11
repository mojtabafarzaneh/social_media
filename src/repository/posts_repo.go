package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mojtabafarzaneh/social_media/src/db"
	"github.com/mojtabafarzaneh/social_media/src/types"
	"gorm.io/gorm"
)

type Post interface {
	GetAllPosts(ctx context.Context) ([]*types.Post, error)
	InsertPost(ctx context.Context, post types.Post) ([]*types.Post, error)
	UpdatePost(ctx context.Context, id string) ([]*types.Post, error)
	DeletePost(ctx context.Context, id string) error
}

type PostgresPostRepo struct {
	DB *gorm.DB
}

func NewPostgresPostRepo() *PostgresPostRepo {
	return &PostgresPostRepo{
		DB: db.ConnectToDB(),
	}
}

func (pr *PostgresPostRepo) GetAllPosts(ctx context.Context) ([]*types.Post, error) {
	var posts []*types.Post

	if err := pr.DB.WithContext(ctx).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *PostgresPostRepo) InsertPost(ctx context.Context, post types.Post) ([]*types.Post, error) {
	var createPost = []*types.Post{{Content: post.Content, Author: post.Author}}

	if err := pr.DB.WithContext(ctx).Create(createPost).Error; err != nil {
		return nil, err
	}

	return createPost, nil
}

func (pr *PostgresPostRepo) UpdatePost(ctx context.Context, content string, id uuid.UUID) (*types.Post, error) {
	var posts *types.Post

	if err := pr.DB.WithContext(ctx).Model(&posts).Where("id = ?", id).Update("Content", content).Error; err != nil {
		return nil, err
	}
	if err := pr.DB.WithContext(ctx).First(&posts, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return posts, nil
}

func (pr *PostgresPostRepo) DeletePost(ctx context.Context, id uuid.UUID) error {
	var posts []*types.Post

	if err := pr.DB.WithContext(ctx).Delete(posts, id).Error; err != nil {
		return err
	}

	return nil
}

func (pr *PostgresPostRepo) GetPost(ctx context.Context, id uuid.UUID) ([]*types.Post, error) {
	var posts []*types.Post

	if err := pr.DB.WithContext(ctx).First(&posts, id).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *PostgresPostRepo) FindPost(ctx context.Context, content string) ([]*types.Post, error) {
	var posts []*types.Post

	if err := pr.DB.WithContext(ctx).Model(&posts).
		Where("content LIKE ?", "%"+content+"%").
		Find(&posts).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return posts, nil
}
