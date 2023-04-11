package jsonutil

import "github.com/kuma0328/circle_board/domain/entity"

type CircleJson struct {
	ID 				int64  `json:"id"`
	Name 			string `json:"name"`
	ImageURL  string `json:"image_url"`
	Genre   	string `json:"genre"`
	Location  string `json:"location"`
	Member  	int64  `json:"member"`
	Twitter 	string `json:"twitter"`
	Instagram string `json:"instagram"`
}

func CircleJsonToEntity(dto *CircleJson) *entity.Circle {
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

func CircleEntityToJson(entity *entity.Circle) *CircleJson {
	return &CircleJson{
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