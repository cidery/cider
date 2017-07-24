package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/cidery/cider/src/infrastructure/payload/response"
	"github.com/cidery/cider/src/infrastructure/payload/request"
	"github.com/cidery/cider/src/domain/service"
	"encoding/json"
	"net/http"
)

type WatcherController struct {
	registry *service.WatcherRegistry
}

func NewWatcherController(registry *service.WatcherRegistry) *WatcherController {
	return &WatcherController{registry: registry}
}

func (w *WatcherController) Bind(engine *gin.Engine) {
	engine.POST("/watchers/register", w.watchersRegister)
	engine.GET("/watchers/list", w.watchersList)
}

func (w *WatcherController) watchersRegister(c *gin.Context) {
	r := request.WatcherRegisterRequest{}
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&r); nil != err {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(err))
		return
	}

	if err := w.registry.RegisterWatcher(r.Id, r.Class, r.Scope); nil != err {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, response.NewEmptyResponse())
}

func (w *WatcherController) watchersList(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewWatcherListResponse(w.registry.Watchers()))
}
