package main

import (
	"encoding/json"
	"errors"
	"example.com/recipecount/internal/storage/sqlite"
	"example.com/recipecount/pkg/recipe"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
	"time"
)

func ReadJSON(db *sqlx.DB) error {
	start := time.Now()
	file, err := os.Open("hf_test_calculation_fixtures.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var jsonContent []recipe.Recipe
	if err := json.NewDecoder(file).Decode(&jsonContent); err != nil {
		return err
	}
	count := 0
	fmt.Printf("time usage after read json:%v\n", time.Now().Sub(start))
	for _, r := range jsonContent {
		_, err = db.Exec(`INSERT INTO recipe (recipe, postcode, delivery_from, delivery_to) VALUES (?, ?, ?, ?);`, r.Recipe, r.Postcode, r.Delivery.From, r.Delivery.To)
		if err != nil {
			return err
		}
		count++
		fmt.Printf("count:%d recipe:%s\n", count, r.Recipe)
	}

	// read open bracket
	//dec := json.NewDecoder(file)
	//t, err := dec.Token()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%T: %v\n", t, t)
	//count := 0
	//// while the array contains values
	//for dec.More() {
	//	var r recipe.Recipe
	//	// decode an array value (Message)
	//	if err := dec.Decode(&r); err != nil {
	//		log.Fatalf("cannot unmarshal json file: %s\n", err.Error())
	//	}
	//	_, err = db.Exec(`INSERT INTO recipe (recipe, postcode, delivery_from, delivery_to) VALUES (?, ?, ?, ?);`, r.Recipe, r.Postcode, r.Delivery.From, r.Delivery.To)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	count++
	//	fmt.Printf("count:%d recipe:%s\n", count, r.Recipe)
	//}
	//// finish token
	//t, err = dec.Token()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%T: %v\n", t, t)
	diff := time.Now().Sub(start)
	fmt.Printf("time usage:%v\n", diff)
	return nil
}

type recipeName []string

func (r *recipeName) String() string {
	return fmt.Sprint(*r)
}

func (r *recipeName) Set(value string) error {
	if len(*r) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dn := range strings.Split(value, ",") {
		*r = append(*r, dn)
	}
	return nil
}

func main() {
	// parse arguments
	var namesFlag recipeName
	flag.Var(&namesFlag, "match-names", "comma-separated list of recipe names to search for")
	deliverFrom := flag.Int("delivery-from", 11, "integer indicates start AM hour to search for")
	deliverTo := flag.Int("delivery-to", 3, "integer indicate end PM hour to search for")
	postcode := flag.String("postcode", "10120", "string to search for postcode")

	flag.Parse()

	db, err := sqlite.NewDB()
	if err != nil {
		log.Printf("creating table error: %s\n", err.Error())
	}
	defer db.Close()

	s := sqlite.NewStorage(db)
	// create service
	svc := recipe.NewService(s)

	//s.DB.Prepare("INSERT INTO ")
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
