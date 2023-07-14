package app

import (
	"api-blog/internal/config"
	"api-blog/internal/repository/pgrepo"
	"api-blog/pkg/httpserver"
	"log"
)

func Run(cfg *config.Config) error {
	conn, err := pgrepo.New(
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}

	log.Println("connection success")

	server := httpserver.New()
}
