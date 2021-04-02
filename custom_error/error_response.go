package custom_error

import "fmt"

type Info struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Error   Info   `json:"error"`
	ID      int    `json:"id"`
}

func (err *Response) Generate() *Response {
	err.Jsonrpc = "2.0"
	err.ID = 1
	return err
}

func InvalidMethod() *Response {
	errRes := new(Response).Generate()
	errRes.Error = Info{Code: -32601, Message: "Method not found"}
	return errRes
}

func InvalidParams(msg string) *Response{
	errRes := new(Response).Generate()
	errRes.Error = Info{Code: -32601, Message: fmt.Sprintf("Params error: %s", msg)}
	return errRes
}

func DatabaseError(msg string) *Response{
	errRes := new(Response).Generate()
	errRes.Error = Info{Code: -32601, Message: fmt.Sprintf("Database error: %s", msg)}
	return errRes
}
