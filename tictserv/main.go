package main

import (
	"distributed.org/tictsrv/src"
	"distributed.org/tictsrv/src/requests"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	ServerPort = 3333
)

func main() {
	address := ":" + strconv.Itoa(ServerPort)
	wg := &sync.WaitGroup{}
	service := &src.TicketService{}
	rqch := make(chan src.RequestImp)
	rsch := make(chan src.ResponseImp)

	go src.HandleRequest(service, rqch, rsch)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/event", func(ctx *gin.Context) {
		var createReq requests.CreateEventRequest

		if err := ctx.ShouldBind(&createReq); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			rqch <- &createReq
		}
	})

	r.GET("/event", func(ctx *gin.Context) {
		var listReq requests.ListEventRequest

		if err := ctx.ShouldBind(&listReq); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			rqch <- &listReq
		}
	})

	r.POST("/ticket", func(ctx *gin.Context) {
		var bookReq requests.BookTicketRequest

		if err := ctx.ShouldBind(&bookReq); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			rqch <- &bookReq
		}
	})

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := r.Run(address); err != nil {
			return
		}
	}()

	log.Printf("Server is starting on port %d", ServerPort)
	wg.Wait()

}
