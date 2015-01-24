package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	messagePtr := flag.String("message", "hello world", "display the message.")
	portPtr := flag.Int("port", 8080, "port number")
	flag.Parse()

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(200, *messagePtr)
	})

	// Listen and server on 0.0.0.0:8080
	port := fmt.Sprintf(":%d", *portPtr)
	router.Run(port)
}
