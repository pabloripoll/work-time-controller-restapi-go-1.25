package errors

import "fmt"

// Base domain error
type DomainError struct {
	Message string
	Code    string
}

func (e DomainError) Error() string {
	return e.Message
}

// Specific error types
type EntityNotFoundError struct {
	DomainError
	EntityType string
	EntityID   string
}

func NewEntityNotFoundError(entityType, entityID string) *EntityNotFoundError {
	return &EntityNotFoundError{
		DomainError: DomainError{
			Message: fmt.Sprintf("%s with ID %s not found", entityType, entityID),
			Code:    "ENTITY_NOT_FOUND",
		},
		EntityType: entityType,
		EntityID:   entityID,
	}
}

type ValidationError struct {
	DomainError
	Field string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		DomainError: DomainError{
			Message: message,
			Code:    "VALIDATION_ERROR",
		},
	}
}

func NewFieldValidationError(field, message string) *ValidationError {
	return &ValidationError{
		DomainError: DomainError{
			Message: message,
			Code:    "VALIDATION_ERROR",
		},
		Field: field,
	}
}

type InvalidEmailError struct {
	DomainError
	Email string
}

func NewInvalidEmailError(email string) *InvalidEmailError {
	return &InvalidEmailError{
		DomainError: DomainError{
			Message: fmt.Sprintf("invalid email format: %s", email),
			Code:    "INVALID_EMAIL",
		},
		Email: email,
	}
}

type InvalidUUIDError struct {
	DomainError
	Value string
}

func NewInvalidUUIDError(value string) *InvalidUUIDError {
	return &InvalidUUIDError{
		DomainError: DomainError{
			Message: fmt.Sprintf("invalid UUID format: %s", value),
			Code:    "INVALID_UUID",
		},
		Value: value,
	}
}

type UnauthorizedError struct {
	DomainError
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		DomainError: DomainError{
			Message: message,
			Code:    "UNAUTHORIZED",
		},
	}
}

type ForbiddenError struct {
	DomainError
}

func NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{
		DomainError: DomainError{
			Message: message,
			Code:    "FORBIDDEN",
		},
	}
}

type DuplicateError struct {
	DomainError
	Field string
	Value string
}

func NewDuplicateError(field, value string) *DuplicateError {
	return &DuplicateError{
		DomainError: DomainError{
			Message: fmt.Sprintf("duplicate %s: %s", field, value),
			Code:    "DUPLICATE_ENTRY",
		},
		Field: field,
		Value: value,
	}
}

type InvalidCredentialsError struct {
	DomainError
}

func NewInvalidCredentialsError() *InvalidCredentialsError {
	return &InvalidCredentialsError{
		DomainError: DomainError{
			Message: "invalid email or password",
			Code:    "INVALID_CREDENTIALS",
		},
	}
}
