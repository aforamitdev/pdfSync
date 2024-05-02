package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/a-h/templ"
	"github.com/aforamitdev/pdfsync/server/app/cmd/ui"
	logger "github.com/aforamitdev/pdfsync/server/foundation"
)

func main() {
	var log *logger.Logger

	ctx := context.Background()

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "SEND ALERT")
		},
	}
	traceIDFn := func(ctx context.Context) string {
		return "test"
	}
	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES", traceIDFn, events)
	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg")
		os.Exit(1)
	}

	// select {
	// case err := <-serverErrors:
	// 	return fmt.Errorf("server error: %w", err)

	// 	return nil
	// }

}

func run(ctx context.Context, log *logger.Logger) error {

	fs := http.FileServer(http.Dir("./build"))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/dashboard", templ.Handler(ui.Dashboard()))
	serverErrors := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println(log)
	log.Info(ctx, "startup")

	server := &http.Server{
		Addr:         "localhost:9000",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      mux,
	}

	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		fmt.Println("shutdown", sig)
		ctx, cancel := context.WithTimeout(ctx, time.Second*20)
		cancel()
		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("could not stop gracefully: %w", err)
		}

	}

	return nil
	// select {
	// case err := <-serverErrors:
	// 	return fmt.Errorf("server error: %w", err)

	// 	return nil
	// }
}
