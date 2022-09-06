package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"test_tech/workers"
)

func Ticket(ctx *gin.Context) {
	body := ctx.Request.Body

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		log.Printf("read body error: %s", err.Error())
		ctx.Status(http.StatusBadRequest)
		return
	}

	workers.InputChannel <- bodyBytes
	ctx.Status(http.StatusAccepted)
}
