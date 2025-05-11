package utils

import (
	"database/sql"
	"errors"
)

// ValidationErrors holds the list of validation error messages
type ValidationErrors struct {
	errors []map[string]interface{}
}

// NewValidationErrors initializes a new ValidationErrors instance
func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		errors: []map[string]interface{}{},
	}
}

// Add adds a new field + message pair
func (ve *ValidationErrors) Add(field, message string) {
	ve.errors = append(ve.errors, map[string]interface{}{
		"field":   field,
		"message": message,
	})
}

// All returns all collected errors
func (ve *ValidationErrors) All() []map[string]interface{} {
	return ve.errors
}

// IsNotFoundError returns true if the error is due to no rows found in the database
func IsNotFoundError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
