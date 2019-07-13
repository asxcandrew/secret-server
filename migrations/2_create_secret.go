package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table secrets...")
		_, err := db.Exec(`CREATE TABLE secrets (
												id serial PRIMARY KEY,
												hash varchar UNIQUE,
												body varchar,
												views_limit int,
												expires_at timestamptz,
												created_at timestamptz
											)
		`)
		return err
	},
		func(db migrations.DB) error {
			fmt.Println("dropping table secrets...")
			_, err := db.Exec(`DROP TABLE secrets`)
			return err
		})
}
