package database

import (
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	wrapErrors "github.com/pkg/errors"

	model2 "crawlerdatacovid/internal/app/model"
	"crawlerdatacovid/internal/pkg/config"
	"crawlerdatacovid/internal/pkg/utils"
)

//Get connection
func GetConnect() *pg.DB {
	databaseURI := utils.Getenv("POSTGRES_CONN", "postgresql://postgres:123qwe@localhost:5432/covid?sslmode=disable")
	if config.ZookeeperValue != nil && config.ZookeeperValue.DatabaseURL != "" {
		databaseURI = config.ZookeeperValue.DatabaseURL
	}
	opts, err := pg.ParseURL(databaseURI)
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opts)
	// Wait until connection is healthy
	for {
		_, err := db.Exec("SELECT 1")
		if err == nil {
			err := InitTables(db)
			if err != nil {
				print(err)
			}
			break
		} else {
			print("Error: PostgreSQL is down")
			time.Sleep(1000 * time.Millisecond)
		}
	}
	return db
}

func InitTables(db *pg.DB) error {
	for _, model := range []interface{}{
		(*model2.DataCovid)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{IfNotExists: true, FKConstraints: true})
		if err != nil {
			return wrapErrors.Wrap(err, "Error initializing db")
		}
	}
	return nil
}
