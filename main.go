package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go_api/internal/core/login"
	"github.com/go_api/internal/core/pg"
	"github.com/go_api/internal/core/rest"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login.Login)
	log.Fatal(router.Run(":8020"))

	pg.InitDB()
	rest.Handle()
}
