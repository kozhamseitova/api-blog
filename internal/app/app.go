package app

import (
	"github.com/kozhamseitova/api-blog/internal/config"
	"github.com/kozhamseitova/api-blog/internal/handler"
	"github.com/kozhamseitova/api-blog/internal/repository/pgrepo"
	"github.com/kozhamseitova/api-blog/internal/service"
	"github.com/kozhamseitova/api-blog/pkg/client/postgres"
	"github.com/kozhamseitova/api-blog/pkg/httpserver"
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	conn, err := postgres.New(
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithUsername(cfg.DB.Username),
		postgres.WithPassword(cfg.DB.Password),
	)

	db := pgrepo.New(conn.Pool)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")

	token := jwttoken.New(cfg.Token.SecretKey)
	srvs := service.New(db, cfg, token)
	hndlr := handler.New(srvs)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
