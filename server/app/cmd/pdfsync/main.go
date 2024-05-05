package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	logger "github.com/aforamitdev/pdfsync/server/foundation"
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

	f, err := os.ReadDir("../pdfsync/server/dist")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range f {
		fmt.Println(file.Name() + "\n")
	}

	httpFS := http.FS(frontend)
	fileServer := http.FileServer(httpFS)

	serveIndex := serveFileContents("index.html", httpFS)

	mux := http.NewServeMux()

	mux.Handle("/", intercept404(fileServer, serveIndex))

	// mux.Handle("/static/", http.StripPrefix("/static/", fs))

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

func serveFileContents(file string, files http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Restrict only to instances where the browser is looking for an HTML file
		if !strings.Contains(r.Header.Get("Accept"), "text/html") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found")

			return
		}

		fmt.Println(files, file, "index")
		// Open the file and return its contents using http.ServeContent
		index, err := files.Open(file)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)

			return
		}

		fi, err := index.Stat()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)

			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(w, r, fi.Name(), fi.ModTime(), index)
	}
}

type hookedResponseWriter struct {
	http.ResponseWriter
	got404 bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == http.StatusNotFound {
		// Don't actually write the 404 header, just set a flag.
		hrw.got404 = true
	} else {
		hrw.ResponseWriter.WriteHeader(status)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.got404 {
		// No-op, but pretend that we wrote len(p) bytes to the writer.
		return len(p), nil
	}

	return hrw.ResponseWriter.Write(p)
}

func intercept404(handler, on404 http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hookedWriter := &hookedResponseWriter{ResponseWriter: w}
		handler.ServeHTTP(hookedWriter, r)

		if hookedWriter.got404 {
			on404.ServeHTTP(w, r)
		}
	})
}
