package crawler

import (
	"context"

	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"
	pb "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
)

type CrawlerFacade interface {
	GetVkUsersPhotosHandler(ctx context.Context, users []uint64) (usersPhotos []model.UserPhotos, err error)
}

type Implementation struct {
	pb.UnimplementedCrawlerServer
	CrawlerFacade CrawlerFacade
}

func New(crawlerFacade CrawlerFacade) Implementation {
	return Implementation{
		CrawlerFacade: crawlerFacade,
	}
}
