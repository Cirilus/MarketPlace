package product

import "errors"

var (
	ErrCantCreateProdut = errors.New("there is error when created user")
	ErrProductNotExist  = errors.New("product doesn't exist")
	ErrGetProduts       = errors.New("there is error when getting products")
	ErrUserNoeExist     = errors.New("User doesn't exist")
)
