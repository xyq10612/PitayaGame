package main

import (
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2/client"
	"time"
)

func main() {
	c := client.New(logrus.DebugLevel, 100*time.Millisecond)

	err := c.ConnectTo(":40000") // connect to proxy
	if err != nil {
		panic(err)
	}

	go func(c *client.Client) {

	}(c)

	select {}
}
