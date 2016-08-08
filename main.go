package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()

	// use static middle
	static := napnap.NewStatic("./public")
	nap.Use(static)

	// debug mode
	nap.UseFunc(dumpMiddleware())
	nap.Use(napnap.NewHealth())

	router := napnap.NewRouter()
	router.Get("/hostname", getHostnameEndpoint)
	router.All("/api/hello-world", getHelloWorldEndpoint)
	router.All("/hello-world", getHelloWorldEndpoint)
	router.Get("/timeout", timeoutEndpoint)
	router.Get("/500", throwInternalErrorEndpoint)
	router.Get("/400", throwBadRequestEndpoint)
	nap.Use(router)
	nap.UseFunc(notFoundMiddleware())

	nap.Run(":10081")
}
