package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()

	// use static middle
	static := napnap.NewStatic("./public")
	nap.Use(static)

	// debug mode
	// nap.UseFunc(dumpMiddleware())

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

	nap.Run(":80")
}
