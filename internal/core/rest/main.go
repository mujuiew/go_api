package rest

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	Handle()
}
