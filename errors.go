package edance

import (
	"errors"
	"github.com/candbright/util/xgin"
	"net/http"
)

var (
	ErrNilSshClient      = errors.New("ssh client is nil")
	ErrUnsupportedDBType = errors.New("db type is not supported")
	ErrNilDB             = errors.New("db is nil")
)

var (
	ErrLoginFailed = xgin.StatusErr(errors.New("login failed"), http.StatusUnauthorized)

	ErrDBOpenFailed = func(err error) xgin.ResultErr {
		return xgin.CodeErr(3001, err)
	}
	ErrDBInitTablesFailed = func(err error) xgin.ResultErr {
		return xgin.CodeErr(3002, err)
	}
	ErrRandomUuid = func(err error) xgin.ResultErr {
		return xgin.CodeErr(3003, err)
	}
)
