package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mujuiew/training/test/internal/core/login"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", login.Login)
	log.Fatal(router.Run(":8020"))
}
