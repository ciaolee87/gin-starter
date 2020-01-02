package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

var router *gin.Engine

func init() {
	router = gin.New()
}

func main() {

	println(os.Getenv("GIN_MODE"))
}
