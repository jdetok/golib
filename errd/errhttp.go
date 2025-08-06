package errd

import (
	"net/http"
)

/*
used in http server
pass writer, msg, error. BuildErrStr will build it like any other error
*/
func (e *Err) HTTPErr(w http.ResponseWriter, msg string, err error) {
	e.Msg = msg
	http.Error(w, e.BuildErrStr(err), http.StatusInternalServerError)
}
