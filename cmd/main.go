package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/himmel520/question-service/config"
	"github.com/himmel520/question-service/internal/controller/ogen"
	"github.com/himmel520/question-service/internal/infrastructure/repository"
	"github.com/himmel520/question-service/internal/infrastructure/repository/postgres"
	categoryRepo "github.com/himmel520/question-service/internal/infrastructure/repository/postgres/category"
	categoryUC "github.com/himmel520/question-service/internal/usecase/category"

	authHandler "github.com/himmel520/question-service/internal/controller/ogen/handler/auth"
	categoryHandler "github.com/himmel520/question-service/internal/controller/ogen/handler/category"
	errHandler "github.com/himmel520/question-service/internal/controller/ogen/handler/error"
)

func main() {
	// TODO: добавить логгер
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := postgres.NewPG(cfg.DB.DBConn)
	if err != nil {
		log.Fatalf("unable to connect to pool: %v", err)
	}
	defer pool.Close()
	dbtx := repository.NewDBTX(pool)

	categoryRepo := categoryRepo.New()
	categoryUC := categoryUC.New(dbtx, categoryRepo)

	handler := ogen.NewHandler(ogen.HandlerParams{
		Auth:     authHandler.New(),
		Error:    errHandler.New(),
		Category: categoryHandler.New(categoryUC),
	})

	app, err := ogen.NewServer(handler, cfg.Srv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Printf("the server is starting on %v", cfg.Srv.Addr)

		if err := app.Run(); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := app.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err)
	}

	log.Println("the server is shut down")
}
