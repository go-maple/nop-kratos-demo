package service

import (
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/go-maple/nop-kratos-demo/blog/api/blog/v1"
	"github.com/go-maple/nop-kratos-demo/blog/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

type BlogService struct {
	pb.UnimplementedBlogServiceServer

	log *log.Helper

	article *biz.ArticleUsecase
}
