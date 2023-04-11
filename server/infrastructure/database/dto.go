package database

import "github.com/kuma0328/circle_board/domain/entity"

type CircleDto struct {
	ID 				int64  `db:"id"`
	Name 			string `db:"name"`
	ImageURL  string `db:"image_url"`
	Genre   	string `db:"genre"`
	Location  string `db:"location"`
	Member  	int64  `db:"member"`
	Twitter 	string `db:"twitter"`
	Instagram string `db:"instagram"`
}

func CircleDtoToEntity(dto *CircleDto) *entity.Circle {
	return &entity.Circle{
		ID: 				dto.ID,
		Name: 			dto.Name,
		ImageURL: 	dto.ImageURL,
		Genre: 			dto.Genre,
		Location: 	dto.Location,
		Member: 		dto.Member,
		Twitter: 		dto.Twitter,
		Instagram: 	dto.Instagram,
	}
}

func CircleEntityToDto(entity *entity.Circle) *CircleDto {
	return &CircleDto{
		ID: 				entity.ID,
		Name: 			entity.Name,
		ImageURL: 	entity.ImageURL,
		Genre: 			entity.Genre,
		Location: 	entity.Location,
		Member: 		entity.Member,
		Twitter: 		entity.Twitter,
		Instagram: 	entity.Instagram,
	}
}
