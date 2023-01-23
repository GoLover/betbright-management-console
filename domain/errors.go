package domain

import "errors"

var (
	ErrRepoRecordAlreadyExist          = errors.New(`this record already exist`)
	ErrRepoRecordNotFound              = errors.New(`record not found`)
	ErrDeliveryIncorrectInput          = errors.New(`incorrect input`)
	ErrUseCaseEnteredSportSlugNotFound = errors.New(`entered sport slug not found`)
	ErrSearchHasNoResult               = errors.New(`search has no result`)
)
