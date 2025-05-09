package salesforce

import (
	"errors"
	"reflect"
)

type transformer interface {
	toProvider(input interface{}) (interface{}, error)
	fromProvider(input interface{}) (interface{}, error)
}

type SFTransformer struct {
	registry map[string]transformer
}

func NewSFTransformer() *SFTransformer {
	return &SFTransformer{
		registry: map[string]transformer{
			"Contact": contactTf{},
			// Future: "Deal": DealTransformer{},
		},
	}
}

func (t SFTransformer) ToProvider(input interface{}) (interface{}, error) {
	typeName := getTypeName(input)
	transformer, ok := t.registry[typeName]
	if !ok {
		return nil, errors.New("no transformer registered for type " + typeName)
	}
	return transformer.toProvider(input)
}

func (t SFTransformer) FromProvider(input interface{}) (interface{}, error) {
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
