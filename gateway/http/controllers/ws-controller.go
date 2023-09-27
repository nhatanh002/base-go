package controllers

import (
	"base-go/common/logger"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

type WsController struct {
}

func NewWsController() *WsController {
	return &WsController{}
}

func (controller *WsController) Mount(e *echo.Echo) {
	g := e.Group("/ws")
	g.GET("/ping", controller.Ping)
}

func (controller *WsController) Ping(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		logger.Error("Upgrade error: %s", err)
		return nil
	}
	defer ws.Close()
	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			logger.Error("Read error: %s", err)
			return nil
		}
		logger.Info("Client ping: %s", msg)
		// Write

		err = ws.WriteMessage(websocket.TextMessage, []byte("pong"))
		if err != nil {
			logger.Error("Write error: %s", err)
			return nil
		}

	}
}
