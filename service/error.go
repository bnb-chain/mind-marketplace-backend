package service

import (
	"fmt"
)

// Verify Interface Compliance
var _ error = (*Err)(nil)

// Err defines service errors.
type Err struct {
	Code    int64  `json:"code"`
	Message string `json:"error"`
}

func (e Err) Enrich(message string) Err {
	return Err{
		Code:    e.Code,
		Message: fmt.Sprintf("%s: %s", e.Message, message),
	}
}

func (e Err) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

var (
	// NoErr - success
	NoErr = Err{Code: 2000}

	InvalidAddressErr   = Err{Code: 3000, Message: "invalid address"}
	EmptyUserNameErr    = Err{Code: 3010, Message: "user name is empty"}
	ExistedUserNameErr  = Err{Code: 3020, Message: "user name is existed"}
	InvalidSigErr       = Err{3030, "invalid signature"}
	InvalidTimestampErr = Err{3040, "invalid timestamp"}

	InvalidKeywordErr = Err{Code: 3500, Message: "invalid keyword, too short"}
	TooBigLimitErr    = Err{Code: 3600, Message: "limit is too big"}

	//NotFoundErr - not found error
	NotFoundErr = Err{Code: 4000, Message: "cannot find"}

	//InternalErr -internal error
	InternalErr = Err{Code: 5000, Message: "internal error"}
)
