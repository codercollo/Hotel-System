// Package validator wraps go-playground/validator with JSON-aware field names,
// structured apierror output, and a one-call DecodeAndValidate helper for handlers.
package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	// Use the json struct tag as the field name in validation errors.
	// RegisterTagNameFunc requires func(reflect.StructField) string — NOT an interface.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Decode decodes the JSON body of r into dst.
// Returns a structured apierror on malformed JSON.
func Decode(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalErr *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxErr):
			return apierror.New(http.StatusBadRequest, apierror.CodeBadRequest,
				fmt.Sprintf("malformed JSON at position %d", syntaxErr.Offset))
		case errors.As(err, &unmarshalErr):
			return apierror.New(http.StatusBadRequest, apierror.CodeBadRequest,
				fmt.Sprintf("invalid value for field %q", unmarshalErr.Field))
		default:
			return apierror.New(http.StatusBadRequest, apierror.CodeBadRequest, "invalid request body")
		}
	}
	return nil
}

// Validate runs struct-tag validation on v and returns a structured apierror
// with per-field messages on failure.
func Validate(v any) error {
	if err := validate.Struct(v); err != nil {
		var verrs validator.ValidationErrors
		if errors.As(err, &verrs) {
			fields := make(map[string]string, len(verrs))
			for _, fe := range verrs {
				fields[fe.Field()] = fieldMessage(fe)
			}
			return apierror.NewValidation(fields)
		}
		return apierror.New(http.StatusBadRequest, apierror.CodeBadRequest, err.Error())
	}
	return nil
}

// DecodeAndValidate decodes r.Body into dst then validates struct tags.
// This is the primary helper used in HTTP handlers.
func DecodeAndValidate(r *http.Request, dst any) error {
	if err := Decode(r, dst); err != nil {
		return err
	}
	return Validate(dst)
}

// ValidateVar validates a single value against a tag expression.
// e.g. ValidateVar(email, "required,email")
func ValidateVar(field any, tag string) error {
	if err := validate.Var(field, tag); err != nil {
		return apierror.New(http.StatusBadRequest, apierror.CodeBadRequest, err.Error())
	}
	return nil
}

func fieldMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email address"
	case "min":
		return fmt.Sprintf("must be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", fe.Param())
	case "len":
		return fmt.Sprintf("must be exactly %s characters", fe.Param())
	case "oneof":
		return fmt.Sprintf("must be one of: %s", fe.Param())
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("must be less than or equal to %s", fe.Param())
	case "url":
		return "must be a valid URL"
	case "uuid":
		return "must be a valid UUID"
	default:
		return fmt.Sprintf("failed validation: %s", fe.Tag())
	}
}
