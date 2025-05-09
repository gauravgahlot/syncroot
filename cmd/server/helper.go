package server

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/legacy"
)

func readSpecAndCreateRouter(ctx context.Context) (*routers.Router, error) {
	loader := openapi3.NewLoader()
	openapiPath := filepath.Join("api", "openapi.yaml")

	doc, err := loader.LoadFromFile(openapiPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load OpenAPI spec: %w", err)
	}

	if err := doc.Validate(ctx); err != nil {
		return nil, fmt.Errorf("failed to validate OpenAPI spec: %w", err)
	}

	router, err := legacy.NewRouter(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to create router: %w", err)
	}

	return &router, nil
}
