package nestedmap

import (
	"github.com/vektra/errors"
)

type NestedMap struct {
	nestedMap map[string]map[string]string
}

func Init() *NestedMap {
	return &NestedMap{nestedMap: make(map[string]map[string]string)}
}

type NestedMapResult struct {
	Value string
	Error error
	Success bool
}

var (
	ErrEmptyInput = errors.New("One or more inputs are empty")
	ErrAlreadyExists = errors.New("Already exists")
)

func (m *NestedMap) Reserve(outerKey string, innerKey string, value string) (result NestedMapResult) {
	if outerKey == "" || innerKey == "" || value == "" {
		result.Error = ErrEmptyInput
		return
	}
	var innerMap map[string]string
	if _, ok := m.nestedMap[outerKey]; !ok {
		innerMap = make(map[string]string)
		m.nestedMap[outerKey] = innerMap
	}

	if _, ok := m.nestedMap[outerKey][innerKey]; ok {
		result.Error = ErrAlreadyExists
		return
	}
	m.nestedMap[outerKey][innerKey] = value
	result.Success = true
	result.Value = value
	return
}

/*
func (m *NestedMap) Set(outerKey string, innerKey string, value string) (result NestedMapResult) {
	if outerKey == "" || innerKey == "" || value == "" {
		result.Error = ErrEmptyInput
		return
	}
	var innerMap map[string]string
	if _, ok := m.nestedMap[outerKey]; !ok {
		innerMap = make(map[string]string)
		m.nestedMap[outerKey] = innerMap
	}
	innerMap[innerKey] = value
	result.Success = true
	result.Value = value
	return
}

func (m *NestedMap) Get(outerKey string, innerKey string) (result NestedMapResult) {
	if outerKey == "" || innerKey == "" {
		result.Error = ErrEmptyInput
		return
	}
	var innerMap map[string]string
	var ok bool
	if innerMap, ok = m.nestedMap[outerKey]; !ok {
		return
	}
	var value string
	if value, ok = innerMap[innerKey]; !ok {
		return
	}
	result.Success = true
	result.Value = value

	return
}
*/