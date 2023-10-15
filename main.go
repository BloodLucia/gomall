package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "HelloWorld")
	})

	err := r.Run(":3000")
	if err != nil {
		log.Fatalf("failed to run server: %s", err)
	}
}
