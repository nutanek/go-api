package main

import (
	"github.com/gin-gonic/gin"
	// "net/http"		
	"fmt"
	"github.com/nutanek/go-api/controllers"
)

func main() {

	// init Routes
	r := gin.Default()

	v1 := r.Group("/v1.0")
	{
		v1.GET("/user", controllers.TestApi)
	}

	fmt.Println("server is starting...")

	r.Run() // listen and serve on 0.0.0.0:8080

}

