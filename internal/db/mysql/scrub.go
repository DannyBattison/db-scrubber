package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"danny.gg/db-scrubber/internal/cfg"
)

func Scrub(config *cfg.ScrubConfig) error {
	fmt.Println(config.Connection.Dsn)
	db, err := sql.Open("mysql", "root:root@/test")

	if err != nil {
		return err
	}

	defer db.Close()

	context := context.Background()

	tx, err := db.BeginTx(context, nil)

	if err != nil {
		return err
	}

	fmt.Println(config)

	for tableName, table := range config.Scrub.Tables {
		fmt.Println(tableName)
		for columnName, column := range table.Columns {
			if column.Generate != "" {
				_, err := db.Exec("UPDATE " + tableName + " SET " + columnName + " = '" + column.Generate + "'")

				if err != nil {
					tx.Rollback()

					return err
				}
			}
		}
	}

	tx.Commit()

	if err != nil {
		return err
	}

	return nil
}
