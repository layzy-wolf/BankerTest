package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/layzy-wolf/BankerTest/config"
	"github.com/layzy-wolf/BankerTest/internal/transport/http"
	"log"
	net "net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	log.Println("INFO config load")
	conf := config.Load()

	r := http.Handler(&ctx)

	srv := &net.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: r,
	}

	go func(srv *net.Server) {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, net.ErrServerClosed) {
			log.Fatalf("FATAL %v", err)
		}
	}(srv)

	<-ctx.Done()
	stop()

	log.Println("INFO server shutting down gracefully, press Ctrl+C to force")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("FATAL server forced to shutdown: ", err)
	}

	log.Println("INFO server exiting")
}
