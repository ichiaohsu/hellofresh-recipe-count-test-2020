//go:generate mockgen -package=mocks -source=storage.go -destination=../../internal/mocks/mock_storage.go
package recipe

type Counter interface {
	CountDistinctRecipe() (int, error)
	CountGroupByRecipe() ([]PerRecipe, error)
	CountBusiestPostcode() (BusiestPostcode, error)
	CountPerPostcodeAndTime(postcode string, from int, to int) (int, error)
	MatchRecipeByName(names ...string) ([]string, error)
}
