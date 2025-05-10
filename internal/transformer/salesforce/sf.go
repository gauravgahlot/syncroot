package salesforce

import (
	"errors"
	"reflect"

	"github.com/gauravgahlot/syncroot/internal/types"
)

type transformer interface {
	toProvider(input types.Object) (interface{}, error)
	fromProvider(input interface{}) (types.Object, error)
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

func (t SFTransformer) ToProvider(input types.Object) (interface{}, error) {
	typeName := input.GetType()
	transformer, ok := t.registry[typeName]
	if !ok {
		return nil, errors.New("no transformer registered for type " + typeName)
	}

	return transformer.toProvider(input)
}

func (t SFTransformer) FromProvider(input interface{}) (types.Object, error) {
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
