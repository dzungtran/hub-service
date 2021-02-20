package domain

import "errors"

var (
	ErrorNotFound   = errors.New("object not found")
	ErrorDuplicated = errors.New("object duplicated")
)
