package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/meteedev/go_choreography/config"
	"github.com/meteedev/go_choreography/order/internal/adapter/db"
	"github.com/meteedev/go_choreography/order/internal/adapter/handler"
	"github.com/meteedev/go_choreography/order/internal/application/core/service"
	"github.com/meteedev/go_choreography/pkg/messenger"
)

func main() {

	dbConn, err := db.NewDb()
	if err != nil {
		log.Fatalf("Failed to initialize db connection : %v", err)
	}

	store := db.NewStore(dbConn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := messenger.NewRabbitMQ(ctx, config.GetAmqpURL())
	if err != nil {
		log.Fatalf("Failed to initialize Messenger: %v", err)
	}

	ms := messenger.NewMessengerService(r)

	s := service.NewOrderService(ms, store)

	h := handler.NewHandler(s)
	e := echo.New()

	registerRoute(e, h)

	go func() {
		startServer(e)
	}()

	gracefulShutdownServer(e)

}

func startServer(e *echo.Echo) {

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}

}

func registerRoute(e *echo.Echo, h *handler.OrderHandler) {
	e.POST("/orders", h.CreateOrder)
}

func gracefulShutdownServer(e *echo.Echo) {

	sig, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-sig.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("shutting down the server:", err.Error())
	}
	log.Println("server shutdown gracefully")

}
