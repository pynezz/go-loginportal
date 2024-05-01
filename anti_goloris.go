package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Client struct {
	IP                string
	ActiveConnections int
	UserAgent         string
	BodyBytesSent     int // default by goloris: 1000000
}

var clients map[string]*Client

func detectAnomalousBehaviour(c *fiber.Ctx) int {

	if clients[c.IP()] == nil {
		clients[c.IP()] = &Client{
			IP:                c.IP(),
			ActiveConnections: 1,
			UserAgent:         c.Get("User-Agent"),
			BodyBytesSent:     len(c.Request().Body()),
		}

	} else {
		clients[c.IP()].ActiveConnections++
		clients[c.IP()].BodyBytesSent += len(c.Request().Body())
	}

	if clients[c.IP()].ActiveConnections > 10 && clients[c.IP()].BodyBytesSent > 1000000 {
		log.Printf(
			"Anomalous behaviour detected from IP: " + c.IP() +
				" with User-Agent: " + c.Get("User-Agent") +
				" and BodyBytesSent: " + fmt.Sprintf("%d", len(c.Request().Body())) +
				" bytes. Closing connection." +
				"")

		return 444
	}

	return 200
}
