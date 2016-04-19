package main

import "github.com/jasonsoft/napnap"

func main() {
	nap := napnap.New()
	nap.UseFunc(dumpMiddleware())

	router := napnap.NewRouter()
	router.All("/api/hello-world", getHelloWorld)
	router.All("/hello-world", getHelloWorld)

	nap.Use(router)
	nap.Run(":8000")
}
