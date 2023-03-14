package dto

import "github.com/florentinuskev/simple-todo/internal/dao"

type GetTodosReq struct {
	UID string `json:"uid"`
}

type GetTodosRes struct {
	Status uint32      `json:"status"`
	Todos  []*dao.Todo `json:"todos"`
}

type GetTodoReq struct {
	ID string `json:"id"`
}

type GetTodoRes struct {
	Status uint32    `json:"status"`
	Todo   *dao.Todo `json:"todo"`
}

type NewTodoReq struct {
	UID  string `json:"uid"`
	Todo string `json:"todo"`
}

type NewTodoRes struct {
	Status uint32    `json:"status"`
	Todo   *dao.Todo `json:"todo"`
}

type EditTodoReq struct {
	ID     string `json:"id"`
	Todo   string `json:"todo"`
	IsDone bool   `json:"is_done"`
}

type EditTodoRes struct {
	Status uint32    `json:"status"`
	Todo   *dao.Todo `json:"todo"`
}

type DeleteTodoReq struct {
	ID string `json:"id"`
}

type DeleteTodoRes struct {
	Status uint32 `json:"status"`
	Msg    string `json:"msg"`
}
