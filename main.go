package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5065"
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		for {
			websocket.Message.Send(ws, "Hello, Client!")
			msg := ""
			log.WithFields(log.Fields{
				"animal": "walrus",
			}).Info("A walrus appears")
			
			err:=websocket.Message.Receive(ws, &msg)
			
			if err!= nil{
				fmt.Println(err)
				break
			}
			
			fmt.Println(msg)
			

		}
	})))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	log.WithFields(log.Fields{
		"port": port,
	}).Info("Connected")
	e.Run(standard.New(":" + port))

}
