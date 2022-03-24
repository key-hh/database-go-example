package internal

import (
	"context"
	"fmt"
	"go-database/internal/repository"
	"go-database/internal/service"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go-database/internal/handler"
)

type App struct {
	conf *Config
}

func NewApp(conf *Config) *App {
	return &App{
		conf: conf,
	}
}

func (a *App) Run() {
	ctx := context.Background()

	sqlxRepo := repository.NewSQLXRepository()
	if err := sqlxRepo.Init(); err != nil {
		log.Fatalf("sqlx repository error %v", err)
	}
	defer sqlxRepo.Close()

	ormRepo := repository.NewORMRepository()
	if err := ormRepo.Init(ctx); err != nil {
		log.Fatalf("orm repository error %v", err)
	}
	defer ormRepo.Close()

	sqlxSrv := service.NewSampleService(sqlxRepo)
	sqlxHandler := handler.NewSampleHandler(sqlxSrv)

	ormSrv := service.NewSampleService(ormRepo)
	ormHandler := handler.NewSampleHandler(ormSrv)

	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Route("/sqlx", func(r chi.Router) {
			r.Get("/samples", sqlxHandler.List)
			r.Get("/samples/{id}", sqlxHandler.Get)
			r.Post("/", sqlxHandler.Create)
		})

		r.Route("/orm", func(r chi.Router) {
			r.Get("/samples", ormHandler.List)
			r.Get("/samples/{id}", ormHandler.Get)
			r.Post("/", ormHandler.Create)
		})
	})

	addr := fmt.Sprintf("%s:%d", a.conf.Host, a.conf.Port)
	log.Printf("Serving HTTP at: %s", addr)
	server := http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       time.Duration(a.conf.ReadTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(a.conf.ReadHeaderTimeout) * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("ListenAndServe: %v", err)
		} else {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}
}
