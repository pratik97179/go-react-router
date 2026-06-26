package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("Starting server at 8000...")
	r.Run(":8000")
}
