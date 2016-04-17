package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()
	nap.UseFunc(dumpMiddleware())

	router := napnap.NewRouter()
	router.Get("/api/hello-world", getHelloWorld)
	router.Get("/hello-world", getHelloWorld)

	nap.Use(router)
	nap.Run(":8000")
}
