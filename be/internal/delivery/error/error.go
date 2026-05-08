package errors

import "errors"

var ErrDuplicateUser = errors.New("duplicate user")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredential = errors.New("invalid credential")