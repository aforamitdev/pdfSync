package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go:embed build
var UI embed.FS

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

	var frontend fs.FS = os.DirFS("../pdfsync/server/dist")

	httpFS := http.FS(frontend)
	fileServer := http.FileServer(httpFS)

	serveIndex := ServeFileContents("index.html", httpFS)

	mux := http.NewServeMux()

	mux.Handle("/", Intercept404(fileServer, serveIndex))

	serverErrors := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println(log)
	log.Info(ctx, "startup message ")

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
