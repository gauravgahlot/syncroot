package hubspot

import (
	"errors"
	"reflect"
)

type transformer interface {
	toProvider(input interface{}) (interface{}, error)
	fromProvider(input interface{}) (interface{}, error)
}

type HubSpotTransformer struct {
	registry map[string]transformer
}

func NewHubSpotTransformer() *HubSpotTransformer {
	return &HubSpotTransformer{
		registry: map[string]transformer{
			"Contact": contactTf{},
			// Future: "Deal": DealTransformer{},
		},
	}
}

func (t HubSpotTransformer) ToProvider(input interface{}) (interface{}, error) {
	typeName := getTypeName(input)
	transformer, ok := t.registry[typeName]
	if !ok {
		return nil, errors.New("no transformer registered for type " + typeName)
	}
	return transformer.toProvider(input)
}

func (t HubSpotTransformer) FromProvider(input interface{}) (interface{}, error) {
	typeName := getTypeName(input)
	transformer, ok := t.registry[typeName]
	if !ok {
		return nil, errors.New("no transformer registered for type " + typeName)
	}
	return transformer.fromProvider(input)
}

func getTypeName(v interface{}) string {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		return typ.Elem().Name()
	}

	return typ.Name()
}
