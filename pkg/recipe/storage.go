package recipe

type Counter interface {
	CountDistinctRecipe() (int, error)
	CountGroupByRecipe() ([]PerRecipe, error)
	CountBusiestPostcode() (BusiestPostcode, error)
	CountPerPostcodeAndTime(postcode string, from int, to int) (int, error)
	MatchRecipeByName(names ...string) ([]string, error)
}
