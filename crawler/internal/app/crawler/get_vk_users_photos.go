package crawler

import (
	"context"
	"github.com/Elderly-AI/observer/crawler/internal/pkg/model"

	pb "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
)

func (i *Implementation) GetVkUsersPhotosHandler(ctx context.Context, req *pb.GetVkUsersPhotosHandlerRequest) (*pb.GetVkUsersPhotosHandlerResponse, error) {
	usersPhotos, err := i.CrawlerFacade.GetVkUsersPhotosHandler(ctx, req.Users)
	if err != nil {
		return nil, err
	}
	return i.convertResponse(usersPhotos), nil
}

func (i *Implementation) convertResponse(usersPhotos []model.UserPhotos) *pb.GetVkUsersPhotosHandlerResponse {
	return &pb.GetVkUsersPhotosHandlerResponse{
		Photos: i.convertUsersPhotosResponse(usersPhotos),
	}
}

func (i *Implementation) convertUsersPhotosResponse(usersPhotos []model.UserPhotos) []*pb.GetVkUsersPhotosHandlerResponse_UserPhotos {
	vkPhotos := make([]*pb.GetVkUsersPhotosHandlerResponse_UserPhotos, 0, len(usersPhotos))
	for _, userPhotos := range usersPhotos {
		vkPhotos = append(vkPhotos, i.convertUserPhotosResponse(userPhotos))
	}
	return vkPhotos
}

func (i *Implementation) convertUserPhotosResponse(photos model.UserPhotos) *pb.GetVkUsersPhotosHandlerResponse_UserPhotos {
	return &pb.GetVkUsersPhotosHandlerResponse_UserPhotos{
		User:   photos.UserID,
		Photos: photos.Photos,
	}
}
