// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	recipe "example.com/recipecount/pkg/recipe"
	gomock "github.com/golang/mock/gomock"
)

// MockCounter is a mock of Counter interface.
type MockCounter struct {
	ctrl     *gomock.Controller
	recorder *MockCounterMockRecorder
}

// MockCounterMockRecorder is the mock recorder for MockCounter.
type MockCounterMockRecorder struct {
	mock *MockCounter
}

// NewMockCounter creates a new mock instance.
func NewMockCounter(ctrl *gomock.Controller) *MockCounter {
	mock := &MockCounter{ctrl: ctrl}
	mock.recorder = &MockCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCounter) EXPECT() *MockCounterMockRecorder {
	return m.recorder
}

// CountBusiestPostcode mocks base method.
func (m *MockCounter) CountBusiestPostcode() (recipe.BusiestPostcode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountBusiestPostcode")
	ret0, _ := ret[0].(recipe.BusiestPostcode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountBusiestPostcode indicates an expected call of CountBusiestPostcode.
func (mr *MockCounterMockRecorder) CountBusiestPostcode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountBusiestPostcode", reflect.TypeOf((*MockCounter)(nil).CountBusiestPostcode))
}

// CountDistinctRecipe mocks base method.
func (m *MockCounter) CountDistinctRecipe() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDistinctRecipe")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDistinctRecipe indicates an expected call of CountDistinctRecipe.
func (mr *MockCounterMockRecorder) CountDistinctRecipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDistinctRecipe", reflect.TypeOf((*MockCounter)(nil).CountDistinctRecipe))
}

// CountGroupByRecipe mocks base method.
func (m *MockCounter) CountGroupByRecipe() ([]recipe.PerRecipe, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountGroupByRecipe")
	ret0, _ := ret[0].([]recipe.PerRecipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountGroupByRecipe indicates an expected call of CountGroupByRecipe.
func (mr *MockCounterMockRecorder) CountGroupByRecipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountGroupByRecipe", reflect.TypeOf((*MockCounter)(nil).CountGroupByRecipe))
}

// CountPerPostcodeAndTime mocks base method.
func (m *MockCounter) CountPerPostcodeAndTime(postcode string, from, to int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountPerPostcodeAndTime", postcode, from, to)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountPerPostcodeAndTime indicates an expected call of CountPerPostcodeAndTime.
func (mr *MockCounterMockRecorder) CountPerPostcodeAndTime(postcode, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountPerPostcodeAndTime", reflect.TypeOf((*MockCounter)(nil).CountPerPostcodeAndTime), postcode, from, to)
}

// MatchRecipeByName mocks base method.
func (m *MockCounter) MatchRecipeByName(names ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range names {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MatchRecipeByName", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MatchRecipeByName indicates an expected call of MatchRecipeByName.
func (mr *MockCounterMockRecorder) MatchRecipeByName(names ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchRecipeByName", reflect.TypeOf((*MockCounter)(nil).MatchRecipeByName), names...)
}
