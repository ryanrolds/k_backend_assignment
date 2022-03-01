package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func StartServer(ctx context.Context, wg *sync.WaitGroup, port int64, conn *sql.DB) *http.Server {
	a := &API{
		db: conn,
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", a.health).Methods("GET")
	router.HandleFunc("/services", a.listServices).Methods("GET")
	router.HandleFunc("/services/{id}", a.getService).Methods("GET")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: requestLogger(router),
	}

	wg.Add(1)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Errorf("error starting server: %w", err))
			return
		}

		log.Info("HTTP server stopped")
		wg.Done()
	}()

	go func() {
		<-ctx.Done()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Error(fmt.Errorf("Problem shutting down server: %w", err))
		}
	}()

	return server
}
