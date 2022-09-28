package main

import (
	"github.com/digicon-hack-07/ez-toon/server/config"
	"github.com/digicon-hack-07/ez-toon/server/router"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	r, err := router.InjectRouter(c)
	if err != nil {
		panic(err)
	}

	r.Start()
}
