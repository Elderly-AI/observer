package crawler

import (
	"context"

	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"
	pb "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
)

type Facade interface {
	GetVkUsersPhotosHandler(ctx context.Context, users []uint64) (usersPhotos []model.UserPhotos, err error)
}

type Implementation struct {
	pb.UnimplementedCrawlerServer
	CrawlerFacade Facade
}

func New(crawlerFacade Facade) Implementation {
	return Implementation{
		CrawlerFacade: crawlerFacade,
	}
}
