package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ZeusPerez/go-crud-skeleton/internal/errors"
	"github.com/ZeusPerez/go-crud-skeleton/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kisielk/sqlstruct"
	log "github.com/sirupsen/logrus"
)

type MySQLConfig struct {
	URL string `envconfig:"DB_URL" default:"root@tcp(127.0.0.1:3306)/test_devs_crud"`
}

type MySQL interface {
	Close()
	Get(ctx context.Context, email string) (models.Dev, error)
	Create(ctx context.Context, dev models.Dev) error
	Update(ctx context.Context, dev models.Dev) (models.Dev, error)
	Delete(ctx context.Context, email string) error
}

func NewMySQLDev(cfg MySQLConfig) (MySQL, error) {
	db, err := sql.Open("mysql", cfg.URL)
	if err != nil {
		return mysqlDev{}, err
	}
	return mysqlDev{cfg: cfg, db: db}, nil

}

type mysqlDev struct {
	cfg MySQLConfig
	db  *sql.DB
}

func (m mysqlDev) Close() {
	m.db.Close()
}

func (m mysqlDev) Get(ctx context.Context, email string) (models.Dev, error) {
	var dev models.Dev
	query := fmt.Sprintf("SELECT * FROM devs WHERE email = '%s'", email)
	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		log.Error(err)
		return dev, err
	}

	defer rows.Close()

	if rows.Next() {
		err = sqlstruct.ScanAliased(&dev, rows, "d")
		if err != nil {
			log.Error(err)
			return dev, errors.Internal{Err: err}
		}
	} else {
		return dev, errors.NotFound{}
	}

	return dev, nil
}

func (m mysqlDev) Create(ctx context.Context, dev models.Dev) error {
	jsonLanguages, err := json.Marshal(dev.Languages)
	if err != nil {
		log.Error(err)
		return err
	}

	query := fmt.Sprintf("INSERT INTO devs(email, languages, expertise) VALUES('%s', '%s', '%d')", dev.Email, jsonLanguages, dev.Expertise)
	_, err = m.db.ExecContext(ctx, query)
	if err != nil {
		log.Error(err)
		return err
	}

	return err
}

func (m mysqlDev) Update(ctx context.Context, dev models.Dev) (models.Dev, error) {
	jsonLanguages, err := json.Marshal(dev.Languages)
	if err != nil {
		log.Error(err)
		return dev, err
	}

	query := fmt.Sprintf("UPDATE devs SET languages = '%s', expertise = '%d' WHERE email = '%s'", jsonLanguages, dev.Expertise, dev.Email)
	sqlResult, err := m.db.ExecContext(ctx, query)
	if err != nil {
		log.Error(err)
		return dev, err
	}

	rows, err := sqlResult.RowsAffected()
	if err != nil {
		log.Error(err)
		return dev, err
	}

	if rows == 0 {
		return dev, errors.NotFound{}
	}

	return dev, err
}

func (m mysqlDev) Delete(ctx context.Context, email string) error {
	query := fmt.Sprintf("DELETE FROM devs WHERE email = '%s'", email)
	_, err := m.db.ExecContext(ctx, query)
	if err != nil {
		log.Error(err)
		return err
	}

	return err
}
