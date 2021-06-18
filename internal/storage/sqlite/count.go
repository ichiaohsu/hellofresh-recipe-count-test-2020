package sqlite

import (
	"example.com/recipecount/pkg/recipe"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

const tableNameRecipe = "recipe"

type Storage struct {
	DB *sqlx.DB
}

func (s Storage) CountDistinctRecipe() (int, error) {
	var count int
	if err := s.DB.Get(&count, "SELECT COUNT(DISTINCT recipe) FROM "+tableNameRecipe); err != nil {
		return 0, err
	}
	return count, nil
}

func (s Storage) CountGroupByRecipe() ([]recipe.PerRecipe, error) {
	var result []recipe.PerRecipe
	if err := s.DB.Select(&result, "SELECT recipe, COUNT(*) AS count FROM "+tableNameRecipe+" GROUP BY recipe ORDER BY recipe;"); err != nil {
		return nil, err
	}
	return result, nil
}

func (s Storage) CountBusiestPostcode() (recipe.BusiestPostcode, error) {
	var result recipe.BusiestPostcode
	if err := s.DB.Get(&result, "SELECT postcode, COUNT(*) AS delivery_count FROM "+tableNameRecipe+" GROUP BY postcode ORDER BY delivery_count DESC LIMIT 1;"); err != nil {
		log.Println(err)
		return recipe.BusiestPostcode{}, err
	}
	return result, nil
}

func (s Storage) CountPerPostcodeAndTime(postcode string, from int, to int) (int, error) {
	var count int
	if err := s.DB.Get(&count, "SELECT COUNT(recipe_id) AS delivery_count FROM "+tableNameRecipe+" WHERE postcode = ? AND delivery_from >= ? AND delivery_to <= ?", postcode, from, to); err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func (s Storage) MatchRecipeByName(names ...string) ([]string, error) {
	var result []string
	var conditions = make([]string, 0, len(names))
	for _, name := range names {
		conditions = append(conditions, "recipe LIKE '%"+name+"%'")
	}

	query := "SELECT DISTINCT recipe FROM " + tableNameRecipe
	if len(names) > 0 {
		query = query + " WHERE " + strings.Join(conditions, " OR ")
	}
	if err := s.DB.Select(&result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func NewStorage(DB *sqlx.DB) Storage {
	return Storage{DB: DB}
}
