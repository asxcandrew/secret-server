package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table secret_views...")
		_, err := db.Exec(`CREATE TABLE secret_views (
												id serial PRIMARY KEY,
												secret_id int REFERENCES secrets(id),
												created_at timestamptz
											)
		`)
		return err
	},
		func(db migrations.DB) error {
			fmt.Println("dropping table secret_views...")
			_, err := db.Exec(`DROP TABLE secret_views`)
			return err
		})
}
