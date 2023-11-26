package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/danielalejandrorosero/to_do_list/app/utils/consts"
	"github.com/danielalejandrorosero/to_do_list/infra/logger"
	"github.com/go-chi/chi"
)

func Serve() {
	r := chi.NewRouter()

	Start(r)
	srv := &http.Server{
		Addr:         consts.Port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port ", consts.Port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	GracefulShutdown(srv)
}

// server will gracefully shutdown within 5 sec
func GracefulShutdown(srv *http.Server) {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan
	logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	logger.Info("Server gracefully stopped!")
}
