package usecase

import (
	"github.com/kuma0328/circle_board/domain/entity"
	"github.com/kuma0328/circle_board/domain/repositry"
)

type circleUsecase struct {
	repo repositry.CircleRepositry
}

type CircleUsecase interface {
	CreateCircle(circle *entity.Circle) error
	GetAllCircle() ([]*entity.Circle, error)
	GetCircleById(id int64) (*entity.Circle, error)
}

func NewCircleUsecase(repo repositry.CircleRepositry) CircleUsecase {
	return circleUsecase {
		repo: repo,
	}
}

func (uc circleUsecase) CreateCircle(circle *entity.Circle) error {
	return uc.repo.CreateCircle(circle)
}

func (uc circleUsecase) GetAllCircle() ([]*entity.Circle, error) {
	return uc.repo.GetAllCircle()
}

func (uc circleUsecase) GetCircleById(id int64) (*entity.Circle, error) {
	return uc.repo.GetCircleById(id)
}