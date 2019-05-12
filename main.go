package main

import (
	"os"
	"strings"
	"time"

	"github.com/jasonsoft/napnap"
)

var (
	_env string
)

func main() {
	nap := napnap.New()
	nap.SetRender("./templates")

	_env = strings.ToLower(os.Getenv("WHITEMOS_ENV"))
	println("Env:", _env)
	if _env == "development" {
		nap.UseFunc(dumpMiddleware())
	}

	// add health check
	nap.Use(napnap.NewHealth())

	// use static middle
	static := napnap.NewStatic("./public")
	nap.Use(static)

	router := napnap.NewRouter()
	router.Get("/", displayIndexEndpoint)
	router.Get("/panic", panicEndpoint)
	router.Get("/hostname", getHostnameEndpoint)
	router.Get("/auth", authEndpoint)
	router.All("/api/hello-world", getHelloWorldEndpoint)
	router.All("/hello-world", getHelloWorldEndpoint)
	router.Get("/timeout", timeoutEndpoint)
	router.Get("/proxy", proxyEndpoint)
	router.Get("/500", throwInternalErrorEndpoint)
	router.Get("/400", throwBadRequestEndpoint)
	router.Get("/health", healthEndpoint)
	router.Get("/health/start", startHealthEndpoint)
	router.Get("/health/stop", stopHealthEndpoint)

	//router.Get("/hubs/1", )
	nap.Use(router)
	nap.Use(napnap.NewNotfoundMiddleware())

	server := napnap.NewHttpEngine(":10080")
	server.ReadTimeout = 30 * time.Second
	server.WriteTimeout = 30 * time.Second
	server.IdleTimeout = 30 * time.Second
	//server.SetKeepAlivesEnabled(false)

	nap.Run(server)
}
