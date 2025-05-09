package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"go.uber.org/zap"
)

type openAPIValidator struct {
	router routers.Router
	logger *zap.Logger
	env    string
}

func NewOpenAPIValidator(router routers.Router, logger *zap.Logger, env string) *openAPIValidator {
	return &openAPIValidator{
		router: router,
		logger: logger,
		env:    env,
	}
}

func (v openAPIValidator) Validate(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/healthz":
			next.ServeHTTP(w, r)

			return
		}

		route, pathParams, err := v.router.FindRoute(r)
		if err != nil {
			v.logger.Warn("Route matching failed", zap.Error(err))
			http.Error(w, "Invalid route", http.StatusBadRequest)

			return
		}

		input := &openapi3filter.RequestValidationInput{
			Request:    r,
			PathParams: pathParams,
			Route:      route,
			Options: &openapi3filter.Options{
				MultiError:         true,
				AuthenticationFunc: validateAuthentication,
			},
		}

		if v.env != "production" {
			input.Options.AuthenticationFunc = openapi3filter.NoopAuthenticationFunc
		}

		if err := openapi3filter.ValidateRequest(context.Background(), input); err != nil {
			v.logger.Warn("Request validation failed", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		next.ServeHTTP(w, r)
	}
}

func validateAuthentication(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	token := input.RequestValidationInput.Request.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return errors.New("missing or invalid bearer token")
	}

	// TODO: implement JWT token validation

	return nil
}
