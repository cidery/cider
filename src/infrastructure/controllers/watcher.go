package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/cidery/cider/src/infrastructure/payload/response"
)

type WatcherController struct {
}

func NewWatcherController() *WatcherController {
	return &WatcherController{}
}

func (w *WatcherController) Bind(engine *gin.Engine) {
	engine.GET("/hello", w.test)
}

func (w *WatcherController) test(c *gin.Context) {
	c.JSON(200, response.NewTestResponse("World"))
}
