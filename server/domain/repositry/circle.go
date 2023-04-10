package repositry

import "github.com/kuma0328/circle_board/domain/entity"

type CircleRepositry interface {
	CreateCircle(circle *entity.Circle) error
	GetAllCircle() ([]*entity.Circle, error)
	GetCircleById(id int64) (*entity.Circle, error)
}