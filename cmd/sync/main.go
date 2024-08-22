package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aforamitdev/pdfsync/docs"

	"github.com/aforamitdev/pdfsync/cmd/sync/handlers"
	logger "github.com/aforamitdev/pdfsync/zero"
	_ "github.com/mattn/go-sqlite3"
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
	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SYNC", traceIDFn, events)
	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup error", "msg")
		os.Exit(1)
	}

	// select {
	// case err := <-serverErrors:
	// 	return fmt.Errorf("server error: %w", err)

	// 	return nil
	// }

}

func run(ctx context.Context, log *logger.Logger) error {

	docs.SwaggerInfo.Title = "PDF sync app api "
	docs.SwaggerInfo.Description = "PDFs sync, share, read and mark pds"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:9000"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// mux.Handle("/", Intercept404(fileServer, serveIndex))

	serverErrors := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	db, err := sql.Open("sqlite3", "file:app.db")
	if err != nil {
		log.Error(ctx, "error connecting datanase1")
		os.Exit(1)
	}

	if err != nil {
		log.Error(ctx, "error connecting datanase")
		return err
	}

	server := &http.Server{
		Addr:         "localhost:9000",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		Handler:      handlers.APIMux("build", shutdown, log, db),
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
