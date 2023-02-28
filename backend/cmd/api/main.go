package main

import (
	"github.com/florentinuskev/simple-todo/internal/server"
	"github.com/florentinuskev/simple-todo/public/utils"
)

func main() {

	// Load Environment Config
	c := &utils.Config{}
	c.LoadEnv()

	s := server.NewServer(c)
	s.RunServer()

}
