package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/asxcandrew/secret-server"
	"github.com/asxcandrew/secret-server/storage"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var logger log.Logger

func main() {
	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	var httpAddr = ":8000"

	errs := make(chan error, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "server", log.DefaultTimestampUTC)

	appConfig, err := server.ResolveConfig()

	if err != nil {
		errs <- err
	}

	db := storage.InitPGConnection(
		appConfig.DB.Host,
		appConfig.DB.Port,
		appConfig.DB.User,
		appConfig.DB.Password,
		appConfig.DB.Name,
	)

	defer db.Close()

	st := storage.NewPGStorage(db)
	s := server.NewSecretService(st)

	s = server.NewSecretLoggingService(logger, s)

	httpLogger := log.With(logger, "component", "http")

	routes := mux.NewRouter()

	routes.PathPrefix("/secret").Handler(server.MakeSecretHandler(s, httpLogger))

	srv := &http.Server{
		Addr:         httpAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}

	go func() {
		logger.Log("transport", "http", "address", httpAddr, "msg", "listening")

		errs <- srv.ListenAndServe()
	}()

	select {
	case <-c:
		shutdown(srv, wait)
	case <-errs:
		shutdown(srv, wait)
	}
}

func shutdown(srv *http.Server, wait time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logger.Log("transport", "shutting down...")
	os.Exit(0)
}
