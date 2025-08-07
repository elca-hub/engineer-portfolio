package dto

import "devport/domain/model"

type WorkUrlDTO struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func NewWorkUrlDTO(url *model.WorkUrl) *WorkUrlDTO {
	return &WorkUrlDTO{
		Url:   url.Url(),
		Title: url.Title(),
	}
}
