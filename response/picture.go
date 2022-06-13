package response

import "github.com/growvv/rs_demo/model"

type PictureResponse struct {
	Response
	Picture model.Picture `json:"picture,omitempty"`
}

type PictureListResponse struct {
	Response
	PictureList []model.Picture `json:"picture_list"`
}
