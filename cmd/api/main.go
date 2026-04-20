package main

import (
	"github.com/gin-gonic/gin"
	"go-recipes/internal/routes"
)

func main() {
	r := gin.Default()
	routes.Setup(r)
	r.Run()
}
