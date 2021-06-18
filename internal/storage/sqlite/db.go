package sqlite

import (
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	//os.Remove("./recipe.db")
	db := sqlx.MustConnect("sqlite3", "./recipe.db")

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS recipe (
    	"recipe_id" INTEGER PRIMARY KEY AUTOINCREMENT, 
    	"recipe" VARCHAR(100) NOT NULL, 
    	"postcode" VARCHAR(10) NOT NULL, 
    	"delivery_from" integer, 
    	"delivery_to" integer);`)
	if err != nil {
		return nil, err
	}
	return db, nil
}
