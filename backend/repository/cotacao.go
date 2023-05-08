package repository

import (
	"context"
	"database/sql"
	"jorge/env"

	_ "github.com/mattn/go-sqlite3"
)

var CotacaoRepo *CotacaoRepository

type CotacaoRepository struct {
	db *sql.DB
}

func init() {
	db, err := sql.Open("sqlite3", "./db.sql")

	if err != nil {
		panic(err)
	}

	CotacaoRepo = &CotacaoRepository{
		db: db,
	}
}

func (cr *CotacaoRepository) Create(cotacao string) error {
	stmt, err := cr.db.Prepare("INSERT INTO cotacoes VALUES ($1, $2)")

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), env.DATABASE_UPDATE_TIMEOUT)

	defer cancel()

	_, err = stmt.ExecContext(ctx, nil, cotacao)

	if err != nil {
		return err
	}

	return nil
}

func (cr *CotacaoRepository) Up() error {
	_, err := cr.db.Exec("CREATE TABLE cotacoes (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, value TEXT NOT NULL)")

	if err != nil {
		return err
	}

	return nil
}
