package main

import (
	"encoding/json"
	"example.com/recipecount/internal/storage/sqlite"
	"example.com/recipecount/pkg/recipe"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// parse arguments
	var namesFlag recipeName
	flag.Var(&namesFlag, "match-names", "comma-separated list of recipe names to search for")
	deliverFrom := flag.Int("delivery-from", 11, "integer indicates start AM hour to search for")
	deliverTo := flag.Int("delivery-to", 3, "integer indicate end PM hour to search for")
	postcode := flag.String("postcode", "10120", "string to search for postcode")

	flag.Parse()

	db := sqlx.MustConnect("sqlite3", "./recipe.db")
	defer db.Close()

	s := sqlite.NewStorage(db)
	// create service
	svc := recipe.NewService(s)

	o, err := svc.GetJSONOutput(*postcode, *deliverFrom, *deliverTo, namesFlag)
	if err != nil {
		log.Println(err)
	}
	// format output
	j, err := json.MarshalIndent(o, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
