package main

import (
	"encoding/json"
	"example.com/recipecount/pkg/recipe"
	"flag"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"sync"
)

const poolSize = 10

func ReadJSON(path string, rchan chan recipe.Recipe) {
	defer close(rchan)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if _, err = dec.Token(); err != nil {
		panic(err)
	}
	// while the array contains values
	for dec.More() {
		var r recipe.Recipe
		// decode an array value (Message)
		if err := dec.Decode(&r); err != nil {
			log.Fatalf("cannot unmarshal json file: %s\n", err.Error())
		}
		rchan <- r
	}

	// finish token
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	os.Remove("./recipe.db")
	filepath := flag.String("filepath", "hf_test_calculation_fixtures.json", "filepath for fixture JSON file")

	db := sqlx.MustConnect("sqlite3", "./recipe.db")
	defer db.Close()

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS recipe (
    	"recipe_id" INTEGER PRIMARY KEY AUTOINCREMENT, 
    	"recipe" VARCHAR(100) NOT NULL, 
    	"postcode" VARCHAR(10) NOT NULL, 
    	"delivery_from" integer, 
    	"delivery_to" integer);`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE INDEX idx_recipe ON recipe (recipe)`)
	if err != nil {
		log.Printf("error creating index: %v\n", err)
	}
	stmt, err := db.Prepare(`INSERT INTO recipe (recipe, postcode, delivery_from, delivery_to) VALUES (?, ?, ?, ?);`)
	if err != nil {
		panic(err)
	}

	rchan := make(chan recipe.Recipe, poolSize)
	go ReadJSON(*filepath, rchan)
	wg := sync.WaitGroup{}
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func(rchan chan recipe.Recipe) {
			for r := range rchan {
				_, err := stmt.Exec(r.Recipe, r.Postcode, r.Delivery.From, r.Delivery.To)
				if err != nil {
					break
				}
			}
			wg.Done()
		}(rchan)
	}
	wg.Wait()
}
