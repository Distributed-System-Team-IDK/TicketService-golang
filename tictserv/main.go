package main

import (
	"distributed.org/tictsrv/utils"
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
	service := &utils.TicketService{}
	rqch := make(chan utils.RequestImp)
	rsch := make(chan utils.ResponseImp)

	go utils.HandleRequest(service, rqch, rsch)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/event", func(ctx *gin.Context) {
		var createReq utils.CreateEventRequest

		if err := ctx.ShouldBind(&createReq); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			rqch <- &createReq
		}
	})

	r.POST("/ticket", func(ctx *gin.Context) {
		// todo
	})

	r.GET("/ticket", func(ctx *gin.Context) {
		// todo
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
