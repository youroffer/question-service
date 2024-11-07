package repoerr

import "errors"

// DB
var (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

// Category
var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrCategoryExists   = errors.New("category exists")
)
