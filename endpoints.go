package main

import (
	"os"
	"time"

	"github.com/jasonsoft/napnap"
	"github.com/jasonsoft/request"
	uuid "github.com/satori/go.uuid"
)

var (
	_isHealth = true
)

type Payload struct {
	Message string `json:"message"`
}

func newPayLoad() Payload {
	return Payload{
		Message: "Hello, World!",
	}
}

type IndexViewModel struct {
	Hostname  string
	ClientIP  string
	RequestID string
	Date      string
	Env       string
}

func displayIndexEndpoint(c *napnap.Context) {
	hostname, _ := os.Hostname()
	u1 := uuid.NewV4()
	nowUTC := time.Now().UTC()

	data := IndexViewModel{
		Hostname:  hostname,
		ClientIP:  c.RemoteIPAddress(),
		RequestID: u1.String(),
		Date:      nowUTC.String(),
		Env:       _env,
	}

	c.Render(200, "index.html", data)
}

func getHelloWorldEndpoint(c *napnap.Context) {
	payLoad := newPayLoad()

	if c.Request.Header.Get("Accept") == "application/json" {
		c.JSON(200, payLoad)
		return
	}

	c.String(200, payLoad.Message)
}

func getHostnameEndpoint(c *napnap.Context) {
	hostname, _ := os.Hostname()
	c.String(200, hostname)
}

func throwInternalErrorEndpoint(c *napnap.Context) {
	c.String(500, "internal error/發生未知錯誤")
}

func throwBadRequestEndpoint(c *napnap.Context) {
	c.String(400, "bad request/找不到")
}

func proxyEndpoint(c *napnap.Context) {
	clientIP := c.RemoteIPAddress()
	url := c.Query("url")
	resp, err := request.
		GET(url).
		Set("X-Forwarded-For", clientIP).
		End()

	if err != nil {
		c.String(500, err.Error())
	}

	c.String(200, resp.String())
}

func timeoutEndpoint(c *napnap.Context) {
	time.Sleep(3 * time.Minute)
}

func healthEndpoint(c *napnap.Context) {
	if _isHealth {
		c.String(200, "OK")
		return
	}

	c.String(500, "Oops..")
}

func startHealthEndpoint(c *napnap.Context) {
	_isHealth = true
	c.SetStatus(200)
}

func stopHealthEndpoint(c *napnap.Context) {
	_isHealth = false
	c.SetStatus(200)
}
