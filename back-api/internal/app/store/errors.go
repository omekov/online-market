package store

import "errors"

var (
	// ErrRecordNotFound - если в базе нету сопостовления возвращается эту ошибку
	ErrRecordNotFound = errors.New("record not found")
)
