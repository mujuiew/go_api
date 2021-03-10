package main

import (
	"fmt"

	"github.com/go_api/internal/core/rest"
)

func main() {
	// router.POST("/login", login.Login)
	// log.Fatal(router.Run(":8020"))

	fmt.Println("hello")
	rest.Handle()
}
