package http

import (
	"github.com/gin-gonic/gin"

	template "janusapi/pkg/template"
)

//Handler is the instance od the templare's handler
type Handler struct {
	usecase template.Usecase
}

//InitHandler initialises template's http handler
func InitHandler(router gin.IRouter, uc template.Usecase) {
	handler := &Handler{uc}
	route := router.Group("/template")
	route.POST("/template", handler.Create)
}

//Create inititialises template's handlers with gin
func (h *Handler) Create(c *gin.Context) {

}
