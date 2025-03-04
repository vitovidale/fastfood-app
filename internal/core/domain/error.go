package domain

import (
	"errors"
)

var (
	// general errors
	ErrorInternal        = errors.New("internal error")
	ErrorDataNotFound    = errors.New("record not found")
	ErrorConflictingData = errors.New("conflicting data")

	// category errors
	ErrorCategoryAlreadyActive   = errors.New("category already active")
	ErrorCategoryAlreadyInactive = errors.New("category already inactive")
	ErrCategoryAlreadyExists     = errors.New("category already exists")
	ErrorCategoryNotFound        = errors.New("category not found")

	// product errors
	ErrorProductAlreadyActive   = errors.New("product already active")
	ErrorProductAlreadyInactive = errors.New("product already inactive")
	ErrorProductAlreadyExists   = errors.New("product already exists")
	ErrorProductNotFound        = errors.New("product not found")

	// customer errors
	ErrorCustomerAlreadyExists   = errors.New("customer already exists")
	ErrorCustomerAlreadyInactive = errors.New("customer already inactive")
	ErrorCustomerAlreadyActive   = errors.New("customer already inactive")
	ErrorCustomerNotFound        = errors.New("customer not found")
	ErrorCustomerWrongPassword   = errors.New("wrong password")

	// order errors
	ErrorOrderAlreadyStarted    = errors.New("order already started")
	ErrorOrderAlreadyDone       = errors.New("order already done")
	ErrorOrderAlreadyProcessing = errors.New("order already processing")
	ErrorOrderAlreadyCancelled  = errors.New("order already cancelled")

	// healthcheck errors
	ErrorAppNotReady   = errors.New("app not ready")
	ErrorAppNotStarted = errors.New("app not started")
)
