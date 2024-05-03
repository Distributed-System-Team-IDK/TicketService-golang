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

var rqch chan tictserv.RequestImp

func main() {
	address := ":" + strconv.Itoa(ServerPort)
	service := &tictserv.TicketService{}
	rqch = make(chan tictserv.RequestImp)

	go tictserv.HandleRequest(service, rqch)

	r := gin.Default()

	r.GET("/", handleGetRoot)
	r.POST("/event", handlePostEvent)
	r.GET("/event", handleGetEvent)
	r.POST("/ticket", handlePostTicket)

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
	close(rqch)

}

func handleGetRoot(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func handlePostEvent(ctx *gin.Context) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	createReq := requests.CreateEventRequest{
		ContextHolder: requests.ContextHolder{
			Context:   ctx,
			WaitGroup: wg,
		},
	}

	if err := ctx.ShouldBind(&createReq); err != nil {
		if err2 := ctx.Error(err); err2 != nil {
			log.Printf("Bind error could not be sent back to the client")
		}
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		wg.Add(1)
		rqch <- &createReq
	}
}

func handleGetEvent(ctx *gin.Context) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	listReq := requests.ListEventsRequest{
		ContextHolder: requests.ContextHolder{
			Context:   ctx,
			WaitGroup: wg,
		},
	}
	wg.Add(1)
	rqch <- &listReq
}

func handlePostTicket(ctx *gin.Context) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	bookReq := requests.BookTicketRequest{
		ContextHolder: requests.ContextHolder{
			Context:   ctx,
			WaitGroup: wg,
		},
	}

	if err := ctx.ShouldBind(&bookReq); err != nil {
		if err2 := ctx.Error(err); err2 != nil {
			log.Printf("Bind error could not be sent back to the client")
		}
		ctx.AbortWithStatus(http.StatusBadRequest)
	} else {
		wg.Add(1)
		rqch <- &bookReq
	}
}
