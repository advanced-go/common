package httpx

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/advanced-go/common/core"
	io2 "github.com/advanced-go/common/io"
	"net/http"
	"strings"
)

const (
	fileExistsError = "The system cannot find the file specified"
)

// NewResponseFromUri - read a Http response given a URL
func NewResponseFromUri(uri any) (*http.Response, *core.Status) {
	serverErr := &http.Response{StatusCode: http.StatusInternalServerError, Status: internalError, Header: make(http.Header)}
	if uri == nil {
		return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: URL is nil"))
	}
	//if u.Scheme != fileScheme {
	//	return serverErr, core.NewStatusError(core.StatusInvalidArgument, errors.New(fmt.Sprintf("error: Invalid URL scheme : %v", u.Scheme)))
	//}
	buf, status := io2.ReadFile(uri)
	if !status.OK() {
		if strings.Contains(status.Err.Error(), fileExistsError) {
			return &http.Response{StatusCode: http.StatusNotFound, Status: "Not Found", Header: make(http.Header)}, core.NewStatusError(core.StatusInvalidArgument, status.Err)
		}
		return serverErr, core.NewStatusError(core.StatusIOError, status.Err)
	}
	resp1, err2 := http.ReadResponse(bufio.NewReader(bytes.NewReader(buf)), nil)
	if err2 != nil {
		return serverErr, core.NewStatusError(core.StatusIOError, err2)
	}
	return resp1, core.StatusOK()

}
