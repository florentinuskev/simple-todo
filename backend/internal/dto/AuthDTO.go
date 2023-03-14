package dto

import "github.com/florentinuskev/simple-todo/internal/dao"

type UserRegisterReq struct {
	Username string `json:"username" validate:"min=5, max=12"`
	Password string `json:"password" validate:"min=6, max=12"`
}

type UserRegisterRes struct {
	Status uint32    `json:"status"`
	User   *dao.User `json:"user"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

type UserLoginRes struct {
	Status uint32 `json:"status"`
	Token  string `json:"token"`
}

type GetProfileReq struct {
	UID string `json:"uid"`
}

type GetProfileRes struct {
	Status uint32    `json:"status"`
	User   *dao.User `json:"user"`
}
