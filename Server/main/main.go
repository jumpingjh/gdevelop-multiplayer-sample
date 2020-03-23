package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"time"
)

const lag = 150

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alive",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		go func() {
			time.Sleep(lag * time.Millisecond)
			m.Broadcast(msg)
		}()
	})

	r.Run(":5000")
}
