package service

import "github.com/rostis232/psycho/internal/app/repository"

type ActivityService struct {
	repo repository.Activity
}

func NewActivityService (repo repository.Activity) *ActivityService {
	return &ActivityService{
		repo: repo,
	}
}