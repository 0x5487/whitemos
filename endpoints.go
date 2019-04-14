package main

import (
	"errors"
	"fmt"
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

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

func panicEndpoint(c *napnap.Context) {
	panic(errors.New("oops..."))
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

func authEndpoint(c *napnap.Context) {
	token := c.RequestHeader("Authorization")
	if len(token) == 0 {
		c.SetStatus(401)
		return
	}

	c.RespHeader("X-Secret", "abcd-key")
	c.SetStatus(200)
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
		println("err")
		c.String(500, err.Error())
		return
	}

	println("good")
	c.String(200, resp.String())
}

func timeoutEndpoint(c *napnap.Context) {
	c.Writer.WriteHeader(200)

	for i := 0; i < 50; i++ {
		fmt.Println("sleeping", i)
		//c.Writer.Write([]byte("s"))
		time.Sleep(1 * time.Second)
	}

	//c.String(200, "aaa")
	//time.Sleep(3 * time.Minute)
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

func hub1Endpoint(c *napnap.Context) {
	consumer := NewConsumer()
	msg := <-consumer.Queue

	if msg.OP == "done" {
		c.String(200, "done")
		return
	}

	c.String(200, "end")
}

func hub2Endpoint(c *napnap.Context) {
	consumer := NewConsumer()
	msg := <-consumer.Queue

	if msg.OP == "done" {
		c.String(200, "done")
		return
	}

	c.String(200, "end")
}
