package postgres

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr/v2"
	"github.com/victoorraphael/simple-ms/adapters/database"
)

type Provider struct {
	sess *dbr.Session
}

func (p *Provider) Exec() *dbr.Session {
	return p.sess
}

func (p *Provider) Connect(cfg database.Config) error {
	dbcn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := dbr.Open("postgres", dbcn, nil)
	if err != nil {
		log.Println("failed to connect database")
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	p.sess = db.NewSession(nil)

	return nil
}

func (p *Provider) Close() error {
	return p.sess.Close()
}
