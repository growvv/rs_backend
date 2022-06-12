package response

import "github.com/growvv/rs_demo/model"

type UserLoginResponse struct {
	Response
	UserId uint64 `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}
