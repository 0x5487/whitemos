package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()
	nap.UseFunc(dumpMiddleware())

	router := napnap.NewRouter()
	router.All("/api/hello-world", getHelloWorld)
	router.All("/hello-world", getHelloWorld)
	router.Get("/500", throwInternalError)
	router.Get("/400", throwBadRequest)
	nap.Use(router)

	nap.UseFunc(notFoundMiddleware())
	nap.Run(":8000")
}
