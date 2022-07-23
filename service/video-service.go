package service

import "ginframe/entity"


type VideoService interface{
	SaveVideo(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService{
	return &videoService{}
}

func (service *videoService)SaveVideo(video entity.Video) entity.Video{
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService)FindAll() []entity.Video{
	return service.videos
}