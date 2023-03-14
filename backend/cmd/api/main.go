package main

import (
	"github.com/florentinuskev/simple-todo/internal/server"
	"github.com/florentinuskev/simple-todo/public/db"
	"github.com/florentinuskev/simple-todo/public/utils"
)

func main() {

	// Load Environment Config
	c := &utils.Config{}
	c.LoadEnv()

	// DB Connection
	db := db.ConnectDB(c)

	s := server.NewServer(c, db)
	s.RunServer()

}
