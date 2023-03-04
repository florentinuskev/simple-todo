package dto

import "github.com/florentinuskev/simple-todo/internal/dao"

type UserRegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
