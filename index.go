package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	messagePtr := flag.String("message", "hello world", "display the message.")
	portPtr := flag.Int("port", 8080, "port number")
	flag.Parse()

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, *messagePtr)
	})

	r.Use(static.Serve("/files", static.LocalFile("files", true)))

	// Listen and server on 0.0.0.0:8080
	port := fmt.Sprintf(":%d", *portPtr)
	r.Run(port)
}
