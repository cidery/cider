package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/cidery/cider/src/infrastructure/payload/response"
	"github.com/cidery/cider/src/infrastructure/payload/request"
	"encoding/json"
	"net/http"
	"github.com/cidery/cider/src/domain/errors"
	"github.com/cidery/cider/src/domain/service"
	"github.com/cidery/cider/src/domain/model"
)

type ProjectController struct {
	projectRegistry *service.ProjectRegistry
}

func NewProjectController(projectRegistry *service.ProjectRegistry) *ProjectController {
	return &ProjectController{projectRegistry}
}

func (p *ProjectController) Bind(engine *gin.Engine) {
	engine.POST("/projects/register", p.projectRegister)
}

func (p *ProjectController) projectRegister(c *gin.Context) {

	r := request.ProjectRegisterRequest{}
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&r); nil != err {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(err))
		return
	}

	if 0 == len(r.Target) {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(errors.NewProjectHasNoTargetsError()))
		return
	}

	buildTargets := make([]model.BuildTarget, 0)
	for _, t := range r.Target {
		buildTargets = append(buildTargets, model.NewBuildTarget(t.Action, t.Location))
	}

	err := p.projectRegistry.RegisterProject(r.Name, r.Locator, r.Watcher, buildTargets)
	if nil != err {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, response.NewEmptyResponse())
}
