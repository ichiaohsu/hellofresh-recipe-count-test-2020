package recipe

import (
	"log"
	"strconv"
)

type service struct {
	store Counter
}

func NewService(store Counter) service {
	return service{store: store}
}

// GetJSONOutput handles the final return object format
func (s service) GetJSONOutput(postcode string, from int, to int, names []string) (JSONOutput, error) {

	if len(names) == 0 {
		names = []string{"Potato", "Veggie", "Mushroom"}
	}
	var (
		result JSONOutput
		err    error
	)
	result.UniqueRecipeCount, err = s.store.CountDistinctRecipe()
	if err != nil {
		log.Println(err)
	}
	result.CountPerRecipe, err = s.store.CountGroupByRecipe()
	if err != nil {
		log.Println(err)
	}
	result.BusiestPostcode, err = s.store.CountBusiestPostcode()
	if err != nil {
		log.Println(err)
	}
	result.CountPerPostcodeAndTime.Postcode = postcode
	result.CountPerPostcodeAndTime.From = strconv.FormatInt(int64(from), 10) + "AM"
	result.CountPerPostcodeAndTime.To = strconv.FormatInt(int64(to), 10) + "PM"
	if from == 12 {
		from = 0
	}
	result.CountPerPostcodeAndTime.DeliveryCount, err = s.store.CountPerPostcodeAndTime(postcode, from, to)
	if err != nil {
		log.Println(err)
	}

	result.MatchByName, err = s.store.MatchRecipeByName(names...)
	if err != nil {
		log.Println(err)
	}
	return result, nil
}
