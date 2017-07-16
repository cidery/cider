package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/cidery/cider/src/infrastructure/payload/response"
	"github.com/cidery/cider/src/infrastructure/payload/request"
	"github.com/cidery/cider/src/domain/service"
	"encoding/json"
)

type WatcherController struct {
	registry *service.WatcherRegistry
}

func NewWatcherController(registry *service.WatcherRegistry) *WatcherController {
	return &WatcherController{registry: registry}
}

func (w *WatcherController) Bind(engine *gin.Engine) {
	engine.POST("/watchers/register", w.watchersRegister)
	//engine.GET("/watchers/list", w.watchersList)
}

func (w *WatcherController) watchersRegister(c *gin.Context) {
	r := request.WatcherRegisterRequest{}
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&r); nil != err {
		c.JSON(500, response.NewErrorResponse(err))
		return
	}

	if err := w.registry.RegisterWatcher(r.Id, r.Class, r.Scope); nil != err {
		c.JSON(500, response.NewErrorResponse(err))
		return
	}

	c.JSON(200, response.NewEmptyResponse())
}
