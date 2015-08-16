package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"runtime"
)

type JsonHello struct {
	Message string `json:"message"`
}

func main() {
	runtime.GOMAXPROCS(4)

	messagePtr := flag.String("message", "Hello, World!", "display the message.")
	portPtr := flag.Int("port", 8080, "port number")
	flag.Parse()

	r := gin.New()

	r.GET("/json", func(c *gin.Context) {
		myJson := JsonHello{Message: "Hello, World!"}
		c.JSON(200, myJson)

	})

	r.GET("/plaintext", func(c *gin.Context) {
		c.String(200, *messagePtr)
	})

	r.Use(static.Serve("/files", static.LocalFile("files", true)))

	// Listen and server on 0.0.0.0:8080
	port := fmt.Sprintf(":%d", *portPtr)
	r.Run(port)
}
