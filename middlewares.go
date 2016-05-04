package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/jasonsoft/napnap"
)

func dumpMiddleware() napnap.MiddlewareFunc {
	return func(c *napnap.Context, next napnap.HandlerFunc) {
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
		next(c)
	}
}

func notFoundMiddleware() napnap.MiddlewareFunc {
	return func(c *napnap.Context, next napnap.HandlerFunc) {
		c.SetStatus(404)
	}
}
