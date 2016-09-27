package main

import (
	"os"
	"strings"

	"github.com/jasonsoft/napnap"
)

func main() {
	nap := napnap.New()

	// use static middle
	static := napnap.NewStatic("./public")
	nap.Use(static)

	env := strings.ToLower(os.Getenv("WHITEMOS_ENV"))
	println("Env:", env)
	if env == "development" {
		nap.UseFunc(dumpMiddleware())
	}

	router := napnap.NewRouter()
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
	nap.UseFunc(notFoundMiddleware())

	server := napnap.NewHttpEngine(":80")
	nap.Run(server)
}
