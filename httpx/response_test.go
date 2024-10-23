package httpx

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"io"
	"net/url"
)

const (
	testResponse = "file://[cwd]/resource/test-response.txt"
)

func readAll(body io.ReadCloser) ([]byte, *core.Status) {
	if body == nil {
		return nil, core.StatusOK()
	}
	defer body.Close()
	buf, err := io.ReadAll(body)
	if err != nil {
		return nil, core.NewStatusError(core.StatusIOError, err)
	}
	return buf, core.StatusOK()
}

func Example_NewResponseFromUri() {
	s := testResponse
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [status:%v] [statusCode:%v]\n", s, status0, resp.StatusCode)

	buf, status := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [status:%v] [content-length:%v]\n", status, len(buf)) //string(buf))

	//Output:
	//test: NewResponseFromUri(file://[cwd]/resource/test-response.txt) -> [status:OK] [statusCode:200]
	//test: readAll() -> [status:OK] [content-length:56]

}

func Example_NewResponseFromUri_URL_Nil() {
	resp, status0 := NewResponseFromUri(nil)
	fmt.Printf("test: NewResponseFromUri(nil) -> [error:[%v]] [statusCode:%v]\n", status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(nil) -> [error:[error: URL is nil]] [statusCode:500]

}

func _Example_NewResponseFromUri_Invalid_Scheme() {
	s := "https://www.google.com/search?q=golang"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%vl) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(https://www.google.com/search?q=golangl) -> [error:[error: Invalid URL scheme : https]] [statusCode:500]

}

func Example_NewResponseFromUri_HTTP_Error() {
	s := "file://[cwd]/resource/message.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/resource/message.txt) -> [error:[malformed HTTP status code "text"]] [statusCode:500]

}

func Example_NewResponseFromUri_504() {
	s := "file://[cwd]/resource/http-504.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	buf, status := readAll(resp.Body)
	fmt.Printf("test: readAll() -> [status:%v] [content-length:%v]\n", status, len(buf)) //string(buf))

	//Output:
	//test: NewResponseFromUri(file://[cwd]/resource/http-504.txt) -> [error:[<nil>]] [statusCode:504]
	//test: readAll() -> [status:OK] [content-length:0]

}

func Example_NewResponseFromUri_EOF_Error() {
	s := "file://[cwd]/resource/http-503-error.txt"
	u, _ := url.Parse(s)

	resp, status0 := NewResponseFromUri(u)
	fmt.Printf("test: NewResponseFromUri(%v) -> [error:[%v]] [statusCode:%v]\n", s, status0.Err, resp.StatusCode)

	//Output:
	//test: NewResponseFromUri(file://[cwd]/resource/http-503-error.txt) -> [error:[unexpected EOF]] [statusCode:500]

}

/*
func ExampleNewError() {
	status := core.StatusOK()
	//var resp *http.Response

	err := NewError(nil, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, nil, err)

	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", core.StatusOK(), nil, err)

	status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: invalid content"))
	err = NewError(status, nil)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", status, nil, err)

	resp, _ := NewResponse(http.StatusTeapot, nil, nil)
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [err:%v]\n", nil, resp != nil, err)

	resp, _ = NewResponse(http.StatusTeapot, nil, "error: response content")
	err = NewError(nil, resp)
	fmt.Printf("test: NewError() -> [status:%v] [resp:%v] [%v]\n", nil, resp != nil, err)

	//Output:
	//test: NewError() -> [status:<nil>] [resp:<nil>] [err:]
	//test: NewError() -> [status:OK] [resp:<nil>] [err:]
	//test: NewError() -> [status:Invalid Content [error: invalid content]] [resp:<nil>] [error: invalid content]
	//test: NewError() -> [status:<nil>] [resp:true] [err:]
	//test: NewError() -> [status:<nil>] [resp:true] [error: response content]

}


*/
