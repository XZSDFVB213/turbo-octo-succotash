package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/XZSDFVB213/golang-todoapp/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("Failed to decode r.body: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	var (
		err error
	)
	v, ok := dest.(validatable)
	if ok {
		err = v.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}
	if err != nil {
		return fmt.Errorf("Failed to validate dto: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return nil
}
