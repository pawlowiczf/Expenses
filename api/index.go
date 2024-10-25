package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Index(ctx *gin.Context) {
	data := struct{Name string}{Name: "Filip"}
	ctx.HTML(http.StatusOK, "index.html", data)
}
