package recipe

import (
	"errors"
	"regexp"
	"strconv"
)

var deliveryTimeMatch = regexp.MustCompile(`^"[A-Za-z]+ (\d+)AM - (\d+)PM"$`)

type DeliveryTime struct {
	From int `db:"delivery_from"`
	To   int `db:"delivery_to"`
}

// UnmarshalJSON handles the string format "Monday 8AM - 9PM" and turn into simple int fields
// 12AM means midnight so it's turned into 0
func (d *DeliveryTime) UnmarshalJSON(s []byte) error {
	matched := deliveryTimeMatch.FindStringSubmatch(string(s))
	if len(matched) != 3 {
		return errors.New("unmatched time group")
	}
	from, err := strconv.ParseInt(matched[1], 10, 32)
	if err != nil {
		return err
	}
	if from == 12 {
		from = 0
	}
	d.From = int(from)
	to, err := strconv.ParseInt(matched[2], 10, 32)
	if err != nil {
		return err
	}
	d.To = int(to)
	return nil
}

type Recipe struct {
	Postcode string       `json:"postcode" db:"postcode"`
	Recipe   string       `json:"recipe" db:"recipe"`
	Delivery DeliveryTime `json:"delivery"`
}

type PerRecipe struct {
	Recipe string `json:"recipe" db:"recipe"`
	Count  int    `json:"count" db:"count"`
}

type BusiestPostcode struct {
	Postcode      string `json:"postcode" db:"postcode"`
	DeliveryCount int    `json:"delivery_count" db:"delivery_count"`
}

type PerPostcodeAndTime struct {
	Postcode      string `json:"postcode"`
	From          string `json:"from"`
	To            string `json:"to"`
	DeliveryCount int    `json:"delivery_count" db:"delivery_count"`
}

// JSONOutput is the final JSON result template
type JSONOutput struct {
	UniqueRecipeCount       int                `json:"unique_recipe_count"`
	CountPerRecipe          []PerRecipe        `json:"count_per_recipe"`
	BusiestPostcode         BusiestPostcode    `json:"busiest_postcode"`
	CountPerPostcodeAndTime PerPostcodeAndTime `json:"count_per_postcode_and_time"`
	MatchByName             []string           `json:"match_by_name"`
}
