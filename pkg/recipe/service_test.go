package recipe_test

import (
	"errors"
	"example.com/recipecount/internal/mocks"
	"example.com/recipecount/pkg/recipe"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetJSONOutputError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCounter(ctrl)
	s := recipe.NewService(mockStore)

	mockStore.EXPECT().CountDistinctRecipe().Return(0, errors.New("no sql rows in result")).Times(1)
	_, err := s.GetJSONOutput("10120", 11, 3, []string{})
	if err == nil {
		t.Fatal("should return error")
	}
}

func TestGetJSONOutputOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCounter(ctrl)
	s := recipe.NewService(mockStore)

	mockStore.EXPECT().CountDistinctRecipe().Return(300, nil)
	mockStore.EXPECT().CountGroupByRecipe().Return([]recipe.PerRecipe{
		{"Pulled Pork", 10},
	}, nil).Times(1)
	mockStore.EXPECT().CountBusiestPostcode().Return(recipe.BusiestPostcode{
		Postcode:      "10120",
		DeliveryCount: 50,
	}, nil).Times(1)
	mockStore.EXPECT().CountPerPostcodeAndTime("10000", 0, 6).Return(3, nil).Times(1)
	mockStore.EXPECT().MatchRecipeByName("Mango", "Strawberry").Return(nil, nil).Times(1)
	res, err := s.GetJSONOutput("10000", 12, 6, []string{"Mango", "Strawberry"})
	if err != nil {
		t.Fatal("should not return err")
	}
	if res.CountPerPostcodeAndTime.From != "12AM" || res.CountPerPostcodeAndTime.To != "6PM" || res.CountPerPostcodeAndTime.Postcode != "10000" {
		t.Fatal("count per postcode and time result error")
	}
}
