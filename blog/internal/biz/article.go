package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Like      int64
}

type ArticleRepo interface {
	// db
	ListArticle(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	DeleteArticle(ctx context.Context, id int64) error

	// redis
	GetArticleLike(ctx context.Context, id int64) (rv int64, err error)
	IncArticleLike(ctx context.Context, id int64) error
}

type HelloMessage struct {
	Message string
}

type ViteRepo interface {
	SayHello(ctx context.Context, content string) (*HelloMessage, error)
}

type ArticleUsecase struct {
	repo   ArticleRepo
	vite   ViteRepo
	logger log.Logger
}

func NewArticleUsecase(
	repo ArticleRepo,
	vite ViteRepo,
	logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{
		repo:   repo,
		vite:   vite,
		logger: logger,
	}
}

func (uc *ArticleUsecase) List(ctx context.Context) (ps []*Article, err error) {
	ps, err = uc.repo.ListArticle(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (p *Article, err error) {
	p, err = uc.repo.GetArticle(ctx, id)
	if err != nil {
		return
	}
	err = uc.repo.IncArticleLike(ctx, id)
	if err != nil {
		return
	}
	p.Like, err = uc.repo.GetArticleLike(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) (string, error) {
	msg, err := uc.vite.SayHello(ctx, article.Content)
	if err != nil {
		return "", err
	}

	uc.logger.Log(log.LevelWarn, "msg:", msg)
	uc.repo.CreateArticle(ctx, article)
	return msg.Message, nil
}

func (uc *ArticleUsecase) Update(ctx context.Context, id int64, article *Article) error {
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteArticle(ctx, id)
}
