package usecase

import (
	"github.com/kuma0328/circle_board/domain/entity"
	"github.com/kuma0328/circle_board/domain/repositry"
)

var _ICircleUsecase = &CircleUsecase{}

type CircleUsecase struct {
	repo repositry.ICircleRepositry
}

type ICircleUsecase interface {
	CreateCircle(circle *entity.Circle) error
	GetAllCircle() ([]*entity.Circle, error)
	GetCircleById(id int64) (*entity.Circle, error)
}

func NewCircleUsecase(repo repositry.ICircleRepositry) ICircleUsecase {
	return &CircleUsecase {
		repo: repo,
	}
}

func (uc *CircleUsecase) CreateCircle(circle *entity.Circle) error {
	return uc.repo.CreateCircle(circle)
}

func (uc *CircleUsecase) GetAllCircle() ([]*entity.Circle, error) {
	return uc.repo.GetAllCircle()
}

func (uc *CircleUsecase) GetCircleById(id int64) (*entity.Circle, error) {
	return uc.repo.GetCircleById(id)
}