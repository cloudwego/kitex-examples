package errno

import (
	"errors"
	"fmt"
)

const (
	// System Code
	SuccessCode    = 0
	ServiceErrCode = 10001
	ParamErrCode   = 10002

	// User ErrCode
	LoginErrCode            = 11001
	UserNotExistErrCode     = 11002
	UserAlreadyExistErrCode = 11003

	// Shop Errcode
	ShopAlreadyExistErrCode = 12001
	ShopNotExistErrCode     = 12002

	// Product Errcode
	BrandNotExistErrCode = 12001
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success             = NewErrNo(SuccessCode, "Success")
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	ShopAlreadyExistErr = NewErrNo(ShopAlreadyExistErrCode, "Shop has been settled")
	ShopNotExistErr     = NewErrNo(ShopNotExistErrCode, "User has not settle a shop yet")
	BrandNotExistErr    = NewErrNo(BrandNotExistErrCode, "Brand is not exist")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
