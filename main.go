package main

import "github.com/gin-gonic/gin"

func main() {

	e := gin.Default()

	e.Run(":8000")
}
