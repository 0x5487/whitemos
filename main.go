package main

import (
	"os"
	"strings"

	"github.com/jasonsoft/napnap"
)

var (
	_env string
)

func main() {
	nap := napnap.New()
	nap.ForwardRemoteIpAddress = true
	nap.SetRender("views/*")

	server := napnap.NewHttpEngine(":80")
	_env = strings.ToLower(os.Getenv("WHITEMOS_ENV"))
	println("Env:", _env)
	if _env == "development" {
		server.SetKeepAlivesEnabled(false)
		nap.UseFunc(dumpMiddleware())
	}

	// add health check
	nap.Use(napnap.NewHealth())

	// use static middle
	static := napnap.NewStatic("./public")
	nap.Use(static)

	router := napnap.NewRouter()
	router.Get("/", displayIndexEndpoint)
	router.Get("/hostname", getHostnameEndpoint)
	router.All("/api/hello-world", getHelloWorldEndpoint)
	router.All("/hello-world", getHelloWorldEndpoint)
	router.Get("/timeout", timeoutEndpoint)
	router.Get("/proxy", proxyEndpoint)
	router.Get("/500", throwInternalErrorEndpoint)
	router.Get("/400", throwBadRequestEndpoint)
	router.Get("/health", healthEndpoint)
	router.Get("/health/start", startHealthEndpoint)
	router.Get("/health/stop", stopHealthEndpoint)
	nap.Use(router)

	nap.Use(napnap.NewNotfoundMiddleware())

	nap.Run(server)
}
