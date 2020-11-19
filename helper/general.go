package helper

import (
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc/status"
)

func CatchError(e *error) {
	if err := recover(); err != nil {
		*e = fmt.Errorf("%v", err)
	}
}

func CatchErrorGeneral() {
	if err := recover(); err != nil {
		log.Println("Error =>", err)
	}
}

// RPCErrDesc get decription error of rpc
func RPCErrDesc(err error) error {
	desc := RPCErr(err).Message()

	return errors.New(desc)
}

// RPCErrCode get code error of rpc
func RPCErrCode(err error) error {
	code := RPCErr(err).Code()

	return errors.New(code.String())
}

// RPCErr get rpc error struct method
func RPCErr(err error) *status.Status {
	r, _ := status.FromError(err)
	return r
}
