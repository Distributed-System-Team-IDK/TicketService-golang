package main

import (
	"distributed.org/tictserv"
	"distributed.org/tictserv/requests"
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
	service := &tictserv.TicketService{}
	rqch := make(chan tictserv.RequestImp)

	go tictserv.HandleRequest(service, rqch)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/event", func(ctx *gin.Context) {
		wg := &sync.WaitGroup{}
		defer wg.Wait()
		createReq := requests.CreateEventRequest{
			ContextHolder: requests.ContextHolder{
				Context:   ctx,
				WaitGroup: wg,
			},
		}

		if err := ctx.ShouldBind(&createReq); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			wg.Add(1)
			rqch <- &createReq
		}
	})

	r.GET("/event", func(ctx *gin.Context) {
		wg := &sync.WaitGroup{}
		defer wg.Wait()
		listReq := requests.ListEventRequest{
			ContextHolder: requests.ContextHolder{
				Context:   ctx,
				WaitGroup: wg,
			},
		}
		wg.Add(1)
		rqch <- &listReq
	})

	r.POST("/ticket", func(ctx *gin.Context) {
		wg := &sync.WaitGroup{}
		defer wg.Wait()
		bookReq := requests.BookTicketRequest{
			ContextHolder: requests.ContextHolder{
				Context:   ctx,
				WaitGroup: wg,
			},
		}

		if err := ctx.ShouldBind(&bookReq); err != nil {
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			wg.Add(1)
			rqch <- &bookReq
		}
	})

	wg := &sync.WaitGroup{}
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
